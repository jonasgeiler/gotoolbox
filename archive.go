package gotoolbox

import (
	"archive/tar"
	"archive/zip"
	"compress/bzip2"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func OpenFileInArchive(
	archiveFile *os.File,
	archiveFileSize int64,
	inArchiveFilePath string,
) (io.ReadCloser, error) {
	archiveFilePath := archiveFile.Name()
	archiveFilePathWithoutExt, archiveFileExt := splitArchiveFilePath(archiveFilePath)
	if archiveFileExt == ".gz" || archiveFileExt == ".bz2" {
		// Likely a .tar.gz or .tar.bz2 file, so try again and prepend.
		archiveFileExt = filepath.Ext(archiveFilePathWithoutExt) + archiveFileExt
	}

	switch archiveFileExt {
	case ".tar":
		if err := resetArchiveFilePosition(archiveFile); err != nil {
			return nil, err
		}
		return openFileInTarArchive(
			archiveFile,
			archiveFilePath,
			inArchiveFilePath,
			false,
		)

	case ".tar.gz", ".taz", ".tgz":
		if err := resetArchiveFilePosition(archiveFile); err != nil {
			return nil, err
		}
		gzipReader, err := gzip.NewReader(archiveFile)
		if err != nil {
			return nil, fmt.Errorf(
				"creating gzip reader for %q: %w",
				archiveFilePath, err,
			)
		}
		return openFileInTarArchive(
			gzipReader,
			archiveFilePath,
			inArchiveFilePath,
			true,
		)

	case ".tar.bz2", ".tb2", ".tbz", ".tbz2", ".tz2":
		if err := resetArchiveFilePosition(archiveFile); err != nil {
			return nil, err
		}
		return openFileInTarArchive(
			&ReadCloserWrapper{
				Reader: bzip2.NewReader(archiveFile),
			},
			archiveFilePath,
			inArchiveFilePath,
			false, // It is closable but a noop function.
		)

	case ".zip":
		return openFileInZipArchive(
			archiveFile,
			archiveFilePath,
			inArchiveFilePath,
			archiveFileSize,
		)

	default:
		return nil, fmt.Errorf(
			"unsupported archive file format %q for %q",
			archiveFileExt, archiveFilePath,
		)
	}
}

// splitArchiveFilePath splits path immediately before the final extension,
// separating it into a path without the extension and an extension.
// The extension is the suffix beginning at the final dot in the
// final element of path; it is empty if there is no dot.
// splitArchiveFilePath returns the same extension as [filepath.Ext].
// The returned values have the property that path = root+ext.
func splitArchiveFilePath(path string) (root, ext string) {
	// Note: This could be a generic "SplitExt" function and doesn't really have
	// anything to do with archives, but it's only used by OpenFileInArchive.
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[:i], path[i:]
		}
	}
	return path, ""
}

// resetArchiveFilePosition resets the current file position for some
// archive readers that don't use io.ReaderAt, only the current position
// of the incoming file reader.
func resetArchiveFilePosition(archiveFile *os.File) error {
	if _, err := archiveFile.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf(
			"seeking to start of archive file %q: %w",
			archiveFile.Name(), err,
		)
	}
	return nil
}

// ReadCloserWrapper is an io.ReadCloser which wraps an io.Reader with an
// optionally definable io.Closer.
type ReadCloserWrapper struct {
	io.Reader
	Closer io.Closer
}

var _ io.ReadCloser = (*ReadCloserWrapper)(nil)

func (r *ReadCloserWrapper) Close() error {
	if r.Closer == nil {
		return nil
	}
	return r.Closer.Close()
}

func openFileInTarArchive(
	archiveFile io.ReadCloser,
	archiveFilePath string,
	inArchiveFilePath string,
	archiveFileClosable bool,
) (io.ReadCloser, error) {
	var closer io.Closer
	if archiveFileClosable {
		closer = archiveFile
	}

	tarReader := tar.NewReader(archiveFile)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			if closer != nil {
				_ = closer.Close()
			}
			return nil, fmt.Errorf(
				"file %q not found in archive %q",
				inArchiveFilePath, archiveFilePath,
			)
		}
		if err != nil {
			if closer != nil {
				_ = closer.Close()
			}
			return nil, fmt.Errorf(
				"reading archive %q: %w",
				archiveFilePath, err,
			)
		}

		if header.Typeflag == tar.TypeReg && header.Name == inArchiveFilePath {
			return &ReadCloserWrapper{
				Reader: tarReader,
				Closer: closer,
			}, nil
		}
	}
}

func openFileInZipArchive(
	archiveFile io.ReaderAt,
	archiveFilePath string,
	inArchiveFilePath string,
	archiveFileSize int64,
) (io.ReadCloser, error) {
	zipReader, err := zip.NewReader(archiveFile, archiveFileSize)
	if err != nil {
		return nil, fmt.Errorf(
			"creating zip reader for %q: %w",
			archiveFilePath, err,
		)
	}

	var inArchiveFile *zip.File
	for _, file := range zipReader.File {
		if file.Name == inArchiveFilePath {
			inArchiveFile = file
		}
	}
	if inArchiveFile == nil {
		return nil, fmt.Errorf(
			"file %q not found in archive %q",
			inArchiveFilePath, archiveFilePath,
		)
	}

	inArchiveFileReader, err := inArchiveFile.Open()
	if err != nil {
		return nil, fmt.Errorf(
			"opening file %q in archive: %w",
			inArchiveFile.Name, err,
		)
	}
	return inArchiveFileReader, nil
}
