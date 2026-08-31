package gotoolbox

import (
	"os"
	"path/filepath"
)

const (
	cacheDirName       = "gotoolbox"
	cacheDirHiddenName = ".gotoolbox"
)

func CacheDirPath() string {
	// Try to use user's cache dir (e.g. /home/username/.cache/gotoolbox).
	if userCacheDirPath, err := os.UserCacheDir(); err == nil {
		return filepath.Join(userCacheDirPath, cacheDirName)
	}

	// Otherwise try to use user's home dir (e.g. /home/username/.gotoolbox).
	if userHomeDirPath, err := os.UserHomeDir(); err == nil {
		return filepath.Join(userHomeDirPath, cacheDirHiddenName)
	}

	// Fall back to global temp dir (e.g. /tmp/gotoolbox).
	return filepath.Join(os.TempDir(), cacheDirName)
}

// TODO: Better cache management. At the moment it would grow indefinitely.
