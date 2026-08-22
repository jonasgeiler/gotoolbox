package gotoolbox

import (
	"os"
	"path/filepath"
)

const toolCacheDirName = "gotoolbox"
const toolCacheDirHiddenName = ".gotoolbox"

func ToolCacheDir() string {
	// Try to use user's cache dir (e.g. /home/username/.cache/gotoolbox).
	if userCacheDir, err := os.UserCacheDir(); err == nil {
		return filepath.Join(userCacheDir, toolCacheDirName)
	}

	// Otherwise try to use user's home dir (e.g. /home/username/.gotoolbox).
	if userHomeDir, err := os.UserHomeDir(); err == nil {
		return filepath.Join(userHomeDir, toolCacheDirHiddenName)
	}

	// Fall back to global temp dir (e.g. /tmp/gotoolbox).
	return filepath.Join(os.TempDir(), toolCacheDirName)
}
