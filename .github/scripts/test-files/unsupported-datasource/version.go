//go:build exclude

package github_scripts_test

// Version that uses an unsupported datasource in it's Renovate tag.
//
// renovate: datasource=unsupported depName=golangci/golangci-lint
const Version = "2.12.2"
