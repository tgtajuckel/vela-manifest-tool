## Description

This plugin enables you to build and publish [Docker Manifest List](https://www.docker.com/)
or [OCI Image Index](https://github.com/opencontainers/image-spec/blob/main/image-index.md)
in a Vela pipeline.

Source Code: https://github.com/go-vela/vela-manifest-tool

Registry: https://hub.docker.com/r/target/vela-manifest-tool

## Usage

> **NOTE:**
>
> Users should refrain from using latest as the tag for the Docker image.
>
> It is recommended to use a semantically versioned tag instead.

Sample of building and publishing an image:

```yaml
steps:
  - name: publish_hello-world
    image: target/vela-manifest-tool:latest
    pull: always
    parameters:
      registry: index.docker.io
      repo: index.docker.io/octocat/hello-world
      manifests:
        - image: /octocat/hello-world:latest-linux-amd64
          platform:
            os: linux
            arch: amd64
        - image: /octocat/hello-world:latest-linux-arm64-v8
          platform:
            os: linux
            arch: arm64
            variant: v8
```

Sample of building an image without publishing:

```yaml
steps:
  - name: publish_hello-world
    image: target/vela-manifest-tool:latest
    pull: always
    parameters:
+     dry_run: true
      registry: index.docker.io
      repo: index.docker.io/octocat/hello-world
      tags: [ "latest" ]
      platforms:
        - linux/amd64
        - linux/arm64/v8
      imageTemplate: /octocat/hello-world:latest-{{ .Os }}-{{ .Arch }}{{ if .Variant }}-{{ .Variant}}{{ end }}
```
