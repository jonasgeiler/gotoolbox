import type { AsyncFunctionArguments } from "@actions/github-script";
import type { Api } from "@octokit/plugin-rest-endpoint-methods";
import type { PaginateInterface } from "@octokit/plugin-paginate-rest";
import { readFile } from "node:fs/promises";
import { basename, dirname, join as joinPaths } from "node:path";
import { RequestError } from "@octokit/request-error";

type ScriptFunctionArgs = AsyncFunctionArguments & {
	github: Api & {
		paginate: PaginateInterface;
	};
	workspace: string;
};

const RENOVATE_CONFIG_FILE_PATH = joinPaths(
	import.meta.dirname,
	"..",
	"renovate.jsonc",
);
const RENOVATE_CONFIG: {
	customManagers: {
		customType: "regex" | "jsonata";
		managerFilePatterns: string[];
		matchStrings: string[];
	}[];
} = Function(
	`"use strict"; return (${await readFile(RENOVATE_CONFIG_FILE_PATH)});`,
)();

function renovatePatternToRegExp(renovatePattern: string): RegExp {
	let negated = false;
	if (renovatePattern.startsWith("/")) {
		renovatePattern = renovatePattern.substring(1);
		/* node:coverage disable */
	} else if (renovatePattern.startsWith("!/")) {
		renovatePattern = renovatePattern.substring(2);
		negated = true;
	} else {
		const error = new Error(
			`Invalid renovate config pattern: ${renovatePattern}`,
		);
		error.name = "InvalidRenovatePatternError";
		throw error;
	}
	/* node:coverage enable */

	let patternFlags = "";
	if (renovatePattern.endsWith("/")) {
		renovatePattern = renovatePattern.slice(0, -1);
		/* node:coverage disable */
	} else if (renovatePattern.endsWith("/i")) {
		renovatePattern = renovatePattern.slice(0, -2);
		patternFlags += "i";
	} else {
		const error = new Error(
			`Invalid renovate config pattern: ${renovatePattern}`,
		);
		error.name = "InvalidRenovatePatternError";
		throw error;
	}
	/* node:coverage enable */

	/* node:coverage disable */
	if (negated) {
		// Turn the pattern into a negative look-ahead.
		renovatePattern = `(?!${renovatePattern})`;
	}
	/* node:coverage enable */

	return new RegExp(renovatePattern, patternFlags);
}

function isRenovateCustomManagerForVersionGoFiles(
	customManager: (typeof RENOVATE_CONFIG)["customManagers"][number],
	versionGoFilePath: string,
) {
	for (let managerFilePattern of customManager.managerFilePatterns) {
		const managerFileRegex = renovatePatternToRegExp(managerFilePattern);
		const managerFileMatches = versionGoFilePath.match(managerFileRegex);
		if (managerFileMatches) {
			return true;
		}
	}
	return false;
}

interface RenovateManagerFields {
	datasource: string;
	depName: string;
	currentValue: string;
}

async function getRenovateManagerFields(
	versionGoFilePath: string,
): Promise<RenovateManagerFields> {
	const versionGoFileContents = await readFile(versionGoFilePath, "utf8");

	for (const customManager of RENOVATE_CONFIG.customManagers) {
		// Skip custom managers that don't use regex or don't manage version.go files.
		if (
			customManager.customType !== "regex" ||
			!isRenovateCustomManagerForVersionGoFiles(
				customManager,
				versionGoFilePath,
			)
		) {
			continue;
		}

		// Match custom manager data in file.
		for (const matchString of customManager.matchStrings) {
			const matches = versionGoFileContents.match(matchString);
			if (
				matches &&
				matches.groups &&
				matches.groups.datasource &&
				matches.groups.depName &&
				matches.groups.currentValue
			) {
				return matches.groups as unknown as RenovateManagerFields;
			}
		}
	}

	const error = new Error(
		`Failed to match renovate manager fields in: ${versionGoFilePath}`,
	);
	error.name = "NoMatchingRenovateManagerFieldsError";
	throw error;
}

async function getReleaseByVersionTag(
	github: ScriptFunctionArgs["github"],
	owner: string,
	repo: string,
	tag: string,
) {
	try {
		return await github.rest.repos.getReleaseByTag({
			owner,
			repo,
			tag,
			headers: {
				Accept: "application/vnd.github.html+json",
			},
		});
	} catch (error: unknown) {
		if (error instanceof RequestError && error.status === 404) {
			// Try again with/without a "v" prefix.
			try {
				return await github.rest.repos.getReleaseByTag({
					owner,
					repo,
					tag: tag.startsWith("v") ? tag.substring(1) : `v${tag}`,
					headers: {
						Accept: "application/vnd.github.html+json",
					},
				});
			} catch {
				// Ignore this error and rethrow the original error below.
			}
		}

		throw error;
	}
}

async function releaseToolFromVersionGo(
	{ github, context }: ScriptFunctionArgs,
	versionGoFilePath: string,
) {
	const owner = context.repo.owner;
	const repo = context.repo.repo;
	const target_commitish = context.sha;

	const {
		datasource: toolVersionDatasource,
		depName: toolRepoFull, // full repo = "owner/repo"
		currentValue: toolVersion,
	} = await getRenovateManagerFields(versionGoFilePath);
	if (toolVersionDatasource !== "github-releases") {
		const error = new Error(
			`Unsupported version.go datasource: ${toolVersionDatasource}`,
		);
		error.name = "UnsupportedDatasourceError";
		throw error;
	}

	const [toolOwner, toolRepo] = toolRepoFull.split("/");
	if (!toolOwner || !toolRepo) {
		const error = new Error(`Invalid version.go depName: ${toolRepoFull}`);
		error.name = "InvalidDepNameError";
		throw error;
	}

	const toolRelease = (
		await getReleaseByVersionTag(github, toolOwner, toolRepo, toolVersion)
	).data;

	// Construct a tag name that Go module discovery likes (name and "v" prefix).
	let moduleVersion = toolVersion;
	if (!moduleVersion.startsWith("v")) {
		moduleVersion = "v" + moduleVersion;
	}
	const moduleName = basename(dirname(versionGoFilePath));
	const tag_name = `${moduleName}/${moduleVersion}`;

	const body = `###### Release notes from [\`${toolOwner}/${toolRepo}@${toolRelease.tag_name}\`](${toolRelease.html_url})
# ${toolRelease.name}

${toolRelease.body_html || toolRelease.body || "(No release notes)"}`;

	// Create release and it's git tag.
	await github.rest.repos.createRelease({
		owner,
		repo,
		tag_name,
		target_commitish,
		name: tag_name,
		body,

		// None of the releases should be "latest", since we have multiple
		// different releases for individual tools in one repo.
		make_latest: "false",
	});
}

export default async (args: ScriptFunctionArgs) => {
	if (!process.env.CHANGED_VERSION_GO_FILE_PATHS?.trim()) {
		const error = new Error(
			"Empty CHANGED_VERSION_GO_FILE_PATHS environment variable.",
		);
		error.name = "InvalidEnvError";
		throw error;
	}
	const results = await Promise.allSettled(
		process.env.CHANGED_VERSION_GO_FILE_PATHS.split("\n").map(
			(relativeVersionGoFilePath) =>
				releaseToolFromVersionGo(
					args,
					// Resolve version.go file path from repo root.
					joinPaths(import.meta.dirname, "..", "..", relativeVersionGoFilePath),
				),
		),
	);
	const errors = [];
	for (const result of results) {
		if (result.status === "rejected") {
			errors.push(result.reason);
		}
	}
	if (errors.length > 0) {
		throw new AggregateError(
			errors,
			"Some errors occurred while releasing tool versions",
		);
	}
};
