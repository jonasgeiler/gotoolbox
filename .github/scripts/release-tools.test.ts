import { describe, it } from "node:test";
import assert from "node:assert/strict";
import releaseToolVersions from "./release-tools.ts";
import { RequestError } from "@octokit/request-error";

// Helper to build a fake GitHub client for a given release.
function makeGithubStub(toolOwner: string, toolRepo: string, toolRelease: any) {
	const _createdReleases: any[] = [];
	return {
		_createdReleases,
		rest: {
			repos: {
				getReleaseByTag: async ({
					owner,
					repo,
					tag,
				}: {
					owner: string;
					repo: string;
					tag: string;
				}) => {
					if (
						owner !== toolOwner ||
						repo !== toolRepo ||
						tag !== toolRelease.tag_name
					) {
						throw new RequestError(
							`No stubbed release for ${owner}/${repo}@${tag}`,
							404,
							{
								request: {
									url: `stub://${owner}/${repo}@${tag}`,
									body: "",
									headers: {
										Accept: "application/json",
									},
									method: "GET",
								},
							},
						);
					}
					return { data: toolRelease };
				},
				createRelease: async (args: any) => {
					_createdReleases.push(args);
					return { data: {} };
				},
			},
		},
	} as any;
}

describe("release-tools github script", () => {
	it("should work for the dprint tool, which doesn't use a 'v' prefix", async () => {
		process.env.CHANGED_VERSION_GO_FILE_PATHS =
			".github/scripts/test-files/dprint/version.go";
		const github = makeGithubStub("dprint", "dprint", {
			tag_name: "0.55.2",
			name: "dprint 0.55.2",
			body: "dprint release notes",
			body_html: null,
			html_url: "https://example.com/dprint/dprint/0.55.2",
		});
		const context = {
			repo: {
				owner: "toolbox-owner1",
				repo: "toolbox-repo1",
			},
			sha: "deadbeef1",
		} as any;
		await releaseToolVersions({ github, context } as any);

		assert.equal(
			github._createdReleases.length,
			1,
			"Expected one release to be created",
		);
		const createdRelease = github._createdReleases[0];
		assert.equal(createdRelease.owner, "toolbox-owner1");
		assert.equal(createdRelease.repo, "toolbox-repo1");
		assert.equal(createdRelease.tag_name, "dprint/v0.55.2");
		assert.equal(createdRelease.target_commitish, "deadbeef1");
		assert.equal(createdRelease.tag_name, createdRelease.name);
		assert.match(createdRelease.body, /\bdprint\/dprint@0\.55\.2\b/);
		assert.match(createdRelease.body, /\bdprint 0\.55\.2\b/);
		assert.match(createdRelease.body, /\bdprint release notes\b/);
		assert.match(
			createdRelease.body,
			/\bhttps:\/\/example\.com\/dprint\/dprint\/0\.55\.2\b/,
		);
		assert.equal(createdRelease.make_latest, "false");
	});

	it("should work for the goreleaser tool, which uses a 'v' prefix", async () => {
		process.env.CHANGED_VERSION_GO_FILE_PATHS =
			".github/scripts/test-files/goreleaser/version.go";
		const github = makeGithubStub("goreleaser", "goreleaser", {
			tag_name: "v2.18.0",
			name: "goreleaser v2.18.0",
			body: "goreleaser release notes",
			body_html: null,
			html_url: "https://example.com/goreleaser/goreleaser/v2.18.0",
		});
		const context = {
			repo: {
				owner: "toolbox-owner2",
				repo: "toolbox-repo2",
			},
			sha: "deadbeef2",
		} as any;
		await releaseToolVersions({ github, context } as any);

		assert.equal(
			github._createdReleases.length,
			1,
			"Expected one release to be created",
		);
		const createdRelease = github._createdReleases[0];
		assert.equal(createdRelease.owner, "toolbox-owner2");
		assert.equal(createdRelease.repo, "toolbox-repo2");
		assert.equal(createdRelease.tag_name, "goreleaser/v2.18.0");
		assert.equal(createdRelease.target_commitish, "deadbeef2");
		assert.equal(createdRelease.tag_name, createdRelease.name);
		assert.match(createdRelease.body, /\bgoreleaser\/goreleaser@v2\.18\.0\b/);
		assert.match(createdRelease.body, /\bgoreleaser v2\.18\.0\b/);
		assert.match(createdRelease.body, /\bgoreleaser release notes\b/);
		assert.match(
			createdRelease.body,
			/\bhttps:\/\/example\.com\/goreleaser\/goreleaser\/v2\.18\.0\b/,
		);
		assert.equal(createdRelease.make_latest, "false");
	});

	it("should work for the golangci-lint tool, which uses a 'v' prefix and a html release notes body", async () => {
		process.env.CHANGED_VERSION_GO_FILE_PATHS =
			".github/scripts/test-files/golangci-lint/version.go";
		const github = makeGithubStub("golangci", "golangci-lint", {
			tag_name: "v2.12.2",
			name: "golangci-lint v2.12.2",
			body: "golangci-lint release notes",
			body_html: "<b>golangci-lint</b> release notes",
			html_url: "https://example.com/golangci/golangci-lint/v2.12.2",
		});
		const context = {
			repo: {
				owner: "toolbox-owner3",
				repo: "toolbox-repo3",
			},
			sha: "deadbeef3",
		} as any;
		await releaseToolVersions({ github, context } as any);

		assert.equal(
			github._createdReleases.length,
			1,
			"Expected one release to be created",
		);
		const createdRelease = github._createdReleases[0];
		assert.equal(createdRelease.owner, "toolbox-owner3");
		assert.equal(createdRelease.repo, "toolbox-repo3");
		assert.equal(createdRelease.tag_name, "golangci-lint/v2.12.2");
		assert.equal(createdRelease.target_commitish, "deadbeef3");
		assert.equal(createdRelease.tag_name, createdRelease.name);
		assert.match(createdRelease.body, /\bgolangci\/golangci-lint@v2\.12\.2\b/);
		assert.match(createdRelease.body, /\bgolangci-lint v2\.12\.2\b/);
		assert.match(createdRelease.body, /<b>golangci-lint<\/b> release notes\b/);
		assert.match(
			createdRelease.body,
			/\bhttps:\/\/example\.com\/golangci\/golangci-lint\/v2\.12\.2\b/,
		);
		assert.equal(createdRelease.make_latest, "false");
	});

	it("should be able to do multiple releases at once", async () => {
		process.env.CHANGED_VERSION_GO_FILE_PATHS = [
			".github/scripts/test-files/golangci-lint/version.go",
			".github/scripts/test-files/golangci-lint/version.go",
		].join("\n");
		const github = makeGithubStub("golangci", "golangci-lint", {
			tag_name: "v2.12.2",
			name: "golangci-lint v2.12.2",
			body: "golangci-lint release notes",
			body_html: "<b>golangci-lint</b> release notes",
			html_url: "https://example.com/golangci/golangci-lint/v2.12.2",
		});
		const context = {
			repo: {
				owner: "toolbox-owner",
				repo: "toolbox-repo",
			},
			sha: "deadbeef",
		} as any;
		await releaseToolVersions({ github, context } as any);

		assert.equal(
			github._createdReleases.length,
			2,
			"Expected two releases to be created",
		);
	});

	it("should reject no renovate tag being found in version.go", async () => {
		process.env.CHANGED_VERSION_GO_FILE_PATHS =
			".github/scripts/test-files/no-renovate/version.go";
		const github = makeGithubStub("golangci", "golangci-lint", {
			tag_name: "v2.12.2",
			name: "golangci-lint v2.12.2",
			body: "golangci-lint release notes",
			body_html: "<b>golangci-lint</b> release notes",
			html_url: "https://example.com/golangci/golangci-lint/v2.12.2",
		});
		const context = {
			repo: {
				owner: "toolbox-owner",
				repo: "toolbox-repo",
			},
			sha: "deadbeef",
		} as any;
		await assert.rejects(
			() => releaseToolVersions({ github, context } as any),
			(error: unknown) => {
				assert.ok(
					error instanceof AggregateError,
					"Expected an AggregateError",
				);
				assert.equal(error.errors.length, 1, "Expected one error to be thrown");
				const subError = error.errors[0];
				assert.equal(subError.name, "NoMatchingRenovateManagerFieldsError");
				return true;
			},
		);

		assert.equal(
			github._createdReleases.length,
			0,
			"Expected no releases to be created",
		);
	});

	it("should reject a release not being found", async () => {
		process.env.CHANGED_VERSION_GO_FILE_PATHS =
			".github/scripts/test-files/golangci-lint/version.go";
		const github = makeGithubStub("golangci", "golangci-lint", {
			tag_name: "v0.0.0",
			name: "golangci-lint v0.0.0",
			body: "golangci-lint release notes",
			body_html: "<b>golangci-lint</b> release notes",
			html_url: "https://example.com/golangci/golangci-lint/v0.0.0",
		});
		const context = {
			repo: {
				owner: "toolbox-owner",
				repo: "toolbox-repo",
			},
			sha: "deadbeef",
		} as any;
		await assert.rejects(
			() => releaseToolVersions({ github, context } as any),
			(error: unknown) => {
				assert.ok(
					error instanceof AggregateError,
					"Expected an AggregateError",
				);
				assert.equal(error.errors.length, 1, "Expected one error to be thrown");
				const subError = error.errors[0];
				assert.ok(subError instanceof RequestError, "Expected a RequestError");
				assert.match(subError.message, /\bgolangci\/golangci-lint@2\.12\.2\b/);
				return true;
			},
		);

		assert.equal(
			github._createdReleases.length,
			0,
			"Expected no releases to be created",
		);
	});

	it("should reject an unsupported datasource in version.go", async () => {
		process.env.CHANGED_VERSION_GO_FILE_PATHS =
			".github/scripts/test-files/unsupported-datasource/version.go";
		const github = makeGithubStub("golangci", "golangci-lint", {
			tag_name: "v2.12.2",
			name: "golangci-lint v2.12.2",
			body: "golangci-lint release notes",
			body_html: "<b>golangci-lint</b> release notes",
			html_url: "https://example.com/golangci/golangci-lint/v2.12.2",
		});
		const context = {
			repo: {
				owner: "toolbox-owner",
				repo: "toolbox-repo",
			},
			sha: "deadbeef3",
		} as any;
		await assert.rejects(
			() => releaseToolVersions({ github, context } as any),
			(error: unknown) => {
				assert.ok(
					error instanceof AggregateError,
					"Expected an AggregateError",
				);
				assert.equal(error.errors.length, 1, "Expected one error to be thrown");
				const subError = error.errors[0];
				assert.equal(subError.name, "UnsupportedDatasourceError");
				return true;
			},
		);

		assert.equal(
			github._createdReleases.length,
			0,
			"Expected no releases to be created",
		);
	});

	it("should reject an invalid depName in version.go", async () => {
		process.env.CHANGED_VERSION_GO_FILE_PATHS =
			".github/scripts/test-files/invalid-depname/version.go";
		const github = makeGithubStub("golangci", "golangci-lint", {
			tag_name: "v2.12.2",
			name: "golangci-lint v2.12.2",
			body: "golangci-lint release notes",
			body_html: "<b>golangci-lint</b> release notes",
			html_url: "https://example.com/golangci/golangci-lint/v2.12.2",
		});
		const context = {
			repo: {
				owner: "toolbox-owner",
				repo: "toolbox-repo",
			},
			sha: "deadbeef3",
		} as any;
		await assert.rejects(
			() => releaseToolVersions({ github, context } as any),
			(error: unknown) => {
				assert.ok(
					error instanceof AggregateError,
					"Expected an AggregateError",
				);
				assert.equal(error.errors.length, 1, "Expected one error to be thrown");
				const subError = error.errors[0];
				assert.equal(subError.name, "InvalidDepNameError");
				return true;
			},
		);

		assert.equal(
			github._createdReleases.length,
			0,
			"Expected no releases to be created",
		);
	});

	it("should reject an empty CHANGED_VERSION_GO_FILE_PATHS", async () => {
		process.env.CHANGED_VERSION_GO_FILE_PATHS = "";
		const github = makeGithubStub("golangci", "golangci-lint", {
			tag_name: "v2.12.2",
			name: "golangci-lint v2.12.2",
			body: "golangci-lint release notes",
			body_html: "<b>golangci-lint</b> release notes",
			html_url: "https://example.com/golangci/golangci-lint/v2.12.2",
		});
		const context = {
			repo: {
				owner: "toolbox-owner",
				repo: "toolbox-repo",
			},
			sha: "deadbeef3",
		} as any;
		await assert.rejects(
			() => releaseToolVersions({ github, context } as any),
			(error: any) => {
				assert.equal(error.name, "InvalidEnvError");
				return true;
			},
		);

		assert.equal(
			github._createdReleases.length,
			0,
			"Expected no releases to be created",
		);
	});
});
