//go:build exclude

package github_scripts_test

// Version that uses an invalid depName in it's Renovate tag.
//
// renovate: datasource=github-releases depName=invalid-depname
const Version = "2.12.2"
