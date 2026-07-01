# updater-nuget

[![Latest Release](https://img.shields.io/github/v/release/SemRels/updater-nuget?label=version\&color=blue)](https://github.com/SemRels/updater-nuget/releases/latest)

Updates the version property in a NuGet project file.

This plugin is distributed as the standalone Go binary `semrel-plugin-updater-nuget`. Semrel executes the binary as a subprocess, provides plugin configuration through `SEMREL_PLUGIN_*` environment variables, provides release context through `SEMREL_*` environment variables, reads standard output, and treats exit code `0` as success and any non-zero exit code as failure. Install the binary in `~/.semrel/plugins/` or anywhere on your `$PATH`.

## Installation

### Binary

```bash
go install github.com/SemRels/updater-nuget/cmd/plugin@latest
```

### Docker

Pre-built, multi-platform images (linux/amd64, linux/arm64) are published to the GitHub Container Registry on every release:

```bash
docker pull ghcr.io/semrels/updater-nuget:latest
```

Images are signed with [cosign](https://github.com/sigstore/cosign) and include a full SBOM attestation. Verify the signature:

```bash
cosign verify ghcr.io/semrels/updater-nuget:latest \
  --certificate-identity-regexp 'https://github.com/SemRels/updater-nuget/.github/workflows/release.yml.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```


## Configuration

```yaml
plugins:
  - name: updater-nuget
    path: ~/.semrel/plugins/semrel-plugin-updater-nuget
    env:
      SEMREL_PLUGIN_FILE: "src/MyApp/MyApp.csproj"
      SEMREL_PLUGIN_PROPERTY: "Version"
```

## `SEMREL_PLUGIN_*` variables

| Name | Required | Description | Default |
| --- | --- | --- | --- |
| `SEMREL_PLUGIN_FILE` | Optional | Path or glob that resolves to the `.csproj` file to update. | *.csproj |
| `SEMREL_PLUGIN_PROPERTY` | Optional | XML property name that stores the package version. | Version |

## `SEMREL_*` release context used

| Variable | Description |
| --- | --- |
| `SEMREL_VERSION` | Resolved release version for the current run. |
| `SEMREL_NEXT_VERSION` | Next version computed by semrel for the release. |
| `SEMREL_DRY_RUN` | Whether semrel is running in dry-run mode. |

## Example behavior

The plugin updates the configured XML property in the selected project file to the next version.

## License

Apache-2.0
