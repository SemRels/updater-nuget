package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunUpdatesProjectFile(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "updater-nuget-main-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Errorf("remove temporary directory: %v", err)
		}
	})

	file := filepath.Join(dir, "demo.csproj")
	if err := os.WriteFile(file, []byte("<Project><Version>1.0.0</Version></Project>"), 0o644); err != nil {
		t.Fatal(err)
	}

	env := map[string]string{"SEMREL_VERSION": "v1.1.0", "SEMREL_PLUGIN_FILE": file}
	var stdout, stderr bytes.Buffer
	if code := run(&stdout, &stderr, func(key string) string { return env[key] }); code != 0 {
		t.Fatalf("run() code = %d stderr = %s", code, stderr.String())
	}

	got, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), "<Version>1.1.0</Version>") {
		t.Fatalf("updated file = %s", got)
	}
}

func TestRunDryRun(t *testing.T) {
	t.Parallel()

	env := map[string]string{"SEMREL_VERSION": "1.1.0", "SEMREL_DRY_RUN": "true"}
	var stdout, stderr bytes.Buffer
	if code := run(&stdout, &stderr, func(key string) string { return env[key] }); code != 0 {
		t.Fatalf("run() code = %d", code)
	}
	if !strings.Contains(stdout.String(), "[dry-run]") {
		t.Fatalf("stdout = %q", stdout.String())
	}
}

func TestRunRequiresVersion(t *testing.T) {
	t.Parallel()

	var stdout, stderr bytes.Buffer
	if code := run(&stdout, &stderr, func(string) string { return "" }); code != 1 {
		t.Fatalf("run() code = %d", code)
	}
	if !strings.Contains(stderr.String(), "SEMREL_VERSION is required") {
		t.Fatalf("stderr = %q", stderr.String())
	}
}
