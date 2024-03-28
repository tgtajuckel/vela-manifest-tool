package main

import (
	"bytes"
	"testing"
)

func TestManifestSpec_New_Validate(t *testing.T) {
	man := defaultFixture(t)

	assertImageMatch(t, "index.docker.io/octocat/hello-world:latest", man.Image)
	assertImageMatch(t, "index.docker.io/octocat/hello-world:latest-linux-amd64",
		man.Manifests[0].Image)
	assertImageMatch(t, "index.docker.io/octocat/hello-world:latest-linux-arm64-v8",
		man.Manifests[1].Image)
	var data bytes.Buffer
	man.Render(&data)
	expected := "image: index.docker.io/octocat/hello-world:latest\n" +
		"manifests:\n" +
		"- image: index.docker.io/octocat/hello-world:latest-linux-amd64\n" +
		"  platform:\n" +
		"    os: linux\n" +
		"    architecture: amd64\n" +
		"- image: index.docker.io/octocat/hello-world:latest-linux-arm64-v8\n" +
		"  platform:\n" +
		"    os: linux\n" +
		"    architecture: arm64\n" +
		"    variant: v8\n"
	if data.String() != expected {
		t.Errorf("failed yaml rendering.\nexpected:\n%sactual:\n%s", expected, data.String())
	}
}

func defaultFixture(t *testing.T) ManifestSpec {
	registry := Registry{
		Name:      "index.docker.io",
		Username:  "test",
		Password:  "pass",
		PushRetry: 1,
		DryRun:    true,
	}
	repo := Repo{
		Name:              "/octocat/hello-world",
		Tags:              []string{"latest"},
		Platforms:         []string{"linux/amd64", "linux/arm64/v8"},
		ComponentTemplate: "{{.Repo}}:{{.Tag}}-{{.Os}}-{{.Arch}}{{if .Variant}}-{{.Variant}}{{end}}",
	}
	ms, err := NewManifestSpec(&registry, &repo)
	if err != nil {
		t.Fatalf("error encountered: %v", err)
	}
	if len(ms) != 1 {
		t.Fatalf("should only have returned a single manifest spec")
	}
	return (ms)[0]
}

func assertImageMatch(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("image mismatch\nexpected: %s !=\nactual:   %s", expected, actual)
	}
}
