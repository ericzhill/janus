// Package janus provides types for representing schema migrations.
package janus

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver"
)

// Migration represents a single schema transition from one version to another
// along with the SQL to perform that transition.
//
// Fields:
//   - From: source version (e.g., "v1.2.3" or "empty").
//   - To: target version (e.g., "v2.0.0" or "empty").
//   - SQL: the SQL statements to execute for this migration.
//
// This is intentionally minimal; validation and execution are handled elsewhere.
type Migration struct {
	From *semver.Version
	To   *semver.Version
	SQL  string
}

// NewMigration constructs a Migration with the provided from/to versions and SQL.
func NewMigration(from, to *semver.Version, sql string) Migration {
	return Migration{From: from, To: to, SQL: sql}
}

// LoadMigrationFromFile loads a Migration from a file path.
// The file name must be in the form: "<from> -> <to>.sql" where versions are
// either semantic versions with a leading 'v' (e.g., v1.2.3) or the sentinel
// "empty". The function reads the SQL content and validates versions.
func LoadMigrationFromFile(path string) (Migration, error) {
	base := filepath.Base(path)
	if !strings.HasSuffix(base, ".sql") {
		return Migration{}, fmt.Errorf("migration file must have .sql extension: %s", base)
	}
	name := strings.TrimSuffix(base, ".sql")
	parts := strings.Split(name, " -> ")
	if len(parts) != 2 {
		return Migration{}, fmt.Errorf("invalid migration file name, expected '<from> -> <to>.sql': %s", base)
	}
	fromV := semver.MustParse(parts[0])
	toV := semver.MustParse(parts[1])
	sqlBytes, err := os.ReadFile(path)
	if err != nil {
		return Migration{}, fmt.Errorf("read migration file %s: %w", path, err)
	}
	return Migration{From: fromV, To: toV, SQL: string(sqlBytes)}, nil
}
