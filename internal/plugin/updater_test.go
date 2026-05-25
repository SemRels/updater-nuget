package plugin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpdaterUpdateProjectFile(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "updater-nuget-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "demo.csproj")
	if err := os.WriteFile(file, []byte("<Project>\n  <Version>1.2.3</Version>\n</Project>\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := NewUpdater().Update(file, "1.3.0"); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	got, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), "<Version>1.3.0</Version>") {
		t.Fatalf("updated file = %s", got)
	}
}

func TestUpdaterMissingFile(t *testing.T) {
	t.Parallel()

	err := NewUpdater().Update(filepath.Join(t.TempDir(), "demo.csproj"), "1.3.0")
	if err == nil || !strings.Contains(err.Error(), "read") {
		t.Fatalf("expected read error, got %v", err)
	}
}

func TestUpdaterMissingVersion(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "updater-nuget-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "demo.csproj")
	if err := os.WriteFile(file, []byte("<Project></Project>\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	err = NewUpdater().Update(file, "1.3.0")
	if err == nil || !strings.Contains(err.Error(), "Version element not found") {
		t.Fatalf("expected version error, got %v", err)
	}
}
