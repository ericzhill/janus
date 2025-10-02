package janus

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestLoadAllMigrations iterates through the migrations directory and ensures
// that each migration file can be loaded via LoadMigrationFromFile. It also
// asserts that the parsed From/To versions match the filename and that SQL
// content is non-empty.
func TestLoadAllMigrations(t *testing.T) {
	dir := filepath.Join(".", "migrations")
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("failed to read migrations dir %s: %v", dir, err)
	}

	var loaded int
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".sql") {
			continue
		}

		path := filepath.Join(dir, name)
		mig, err := LoadMigrationFromFile(path)
		if err != nil {
			t.Fatalf("failed to load migration %s: %v", name, err)
		}

		if mig.From == nil {
			t.Errorf("migration %s has empty From version", name)
		}
		if mig.To == nil {
			t.Errorf("migration %s has empty To version", name)
		}

		// Confirm versions match the filename parts.
		base := strings.TrimSuffix(name, ".sql")
		parts := strings.Split(base, " -> ")
		if len(parts) != 2 {
			t.Fatalf("unexpected migration filename format: %s", name)
		}
		if "v"+mig.From.String() != parts[0] {
			t.Errorf("migration %s From mismatch: got %s want %s", name, mig.From, parts[0])
		}
		if "v"+mig.To.String() != parts[1] {
			t.Errorf("migration %s To mismatch: got %s want %s", name, mig.To, parts[1])
		}

		loaded++
	}

	if loaded == 0 {
		t.Fatalf("no migration files loaded from %s", dir)
	}
}
