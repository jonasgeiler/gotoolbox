package gotoolbox

import (
	"fmt"
	"io"
	"strings"
)

func GitHubReleaseAssetURL(repo, version, asset string) string {
	// Example: https://github.com/dprint/dprint/releases/download/0.54.0/dprint-x86_64-unknown-linux-gnu.zip
	var urlBuilder strings.Builder
	urlBuilder.WriteString("https://github.com/")
	urlBuilder.WriteString(repo)
	urlBuilder.WriteString("/releases/download/")
	urlBuilder.WriteString(version)
	urlBuilder.WriteString("/")
	urlBuilder.WriteString(asset)
	return urlBuilder.String()
}

func FetchGitHubReleaseAsset(
	repo, version, asset string,
) (io.ReadCloser, error) {
	url := GitHubReleaseAssetURL(repo, version, asset)
	resp, err := Fetch(url)
	if err != nil {
		return nil, fmt.Errorf(
			"fetching GitHub release asset from %q: %w",
			url, err,
		)
	}
	return resp.Body, nil
}
