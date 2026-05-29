package gotoolbox

import (
	"os"
	"path/filepath"
)

func ToolCacheDir() string {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		cacheDir = os.TempDir()
	}
	toolCacheDir := filepath.Join(cacheDir, "gotoolbox")
	if err = os.MkdirAll(toolCacheDir, 0755); err != nil {
		// Fall back to not using a subdir.
		return cacheDir
	}
	return toolCacheDir
}

func ResolveToolPath(binName string) string {
	return filepath.Join(ToolCacheDir(), binName)
}
