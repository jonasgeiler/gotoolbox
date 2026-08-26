package gotoolbox

import (
	"archive/tar"
	"archive/zip"
	"compress/bzip2"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

type ArchiveFormat int

const (
	NotAnArchive ArchiveFormat = iota
	ZipArchive
	TarArchive
	TarGzipArchive
	TarBzip2Archive
)

func OpenFileInArchive(
	archiveFile *os.File,
	archiveFileSize int64,
	archiveFormat ArchiveFormat,
	inArchiveFilePath string,
) (io.ReadCloser, error) {
	switch archiveFormat {
	case ZipArchive:
		return openFileInZipArchive(
			archiveFile,
			inArchiveFilePath,
			archiveFileSize,
		)

	case TarArchive:
		if err := resetArchiveFilePosition(archiveFile); err != nil {
			return nil, err
		}
		return openFileInTarArchive(
			archiveFile,
			archiveFile.Name(),
			inArchiveFilePath,
			false,
		)

	case TarGzipArchive:
		if err := resetArchiveFilePosition(archiveFile); err != nil {
			return nil, err
		}
		gzipReader, err := gzip.NewReader(archiveFile)
		if err != nil {
			return nil, fmt.Errorf(
				"creating gzip reader for %q: %w",
				archiveFile.Name(), err,
			)
		}
		return openFileInTarArchive(
			gzipReader,
			archiveFile.Name(),
			inArchiveFilePath,
			true,
		)

	case TarBzip2Archive:
		if err := resetArchiveFilePosition(archiveFile); err != nil {
			return nil, err
		}
		return openFileInTarArchive(
			&ReadCloserWrapper{
				Reader: bzip2.NewReader(archiveFile),
			},
			archiveFile.Name(),
			inArchiveFilePath,
			false, // It is closable but a noop function.
		)

	default:
		return nil, fmt.Errorf(
			"unsupported archive file format: %q",
			archiveFile.Name(),
		)
	}
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
	archiveFile *os.File,
	inArchiveFilePath string,
	archiveFileSize int64,
) (io.ReadCloser, error) {
	zipReader, err := zip.NewReader(archiveFile, archiveFileSize)
	if err != nil {
		return nil, fmt.Errorf(
			"creating zip reader for %q: %w",
			archiveFile.Name(), err,
		)
	}

	var inArchiveFile *zip.File
	for _, file := range zipReader.File {
		if file.Name == inArchiveFilePath {
			inArchiveFile = file
			break
		}
	}
	if inArchiveFile == nil {
		return nil, fmt.Errorf(
			"file %q not found in archive %q",
			inArchiveFilePath, archiveFile.Name(),
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
