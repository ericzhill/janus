# Janus CLI Command Reference

The `janus` CLI provides a developer-friendly interface for managing database schema versions using explicit, versioned transitions. This document describes the available commands and options.

## Usage

```
janus <command> [options]
```

## Commands

### plan

Simulates a migration between two schema versions and shows the steps (DDL files) that would be executed, without making changes to the database.

**Usage:**
```
janus plan --from <version> --to <version> [--migrations <dir>] [--format <text|json>]
```

**Options:**
- `--from <version>`: The current schema version (e.g., `1.0.0`).
- `--to <version>`: The target schema version (e.g., `2.0.0`).
- `--migrations <dir>`: Directory containing migration files (default: `./migrations`).
- `--format <text|json>`: Output format for the plan (default: `text`).

**Example:**
```
janus plan --from 1.0.0 --to 2.0.0 --migrations ./db/migrations --format json
```

### apply

Executes the migration steps required to move the database from one version to another, applying the DDL files in order.

**Usage:**
```
janus apply --from <version> --to <version> [--migrations <dir>] [--dry-run]
```

**Options:**
- `--from <version>`: The current schema version.
- `--to <version>`: The target schema version.
- `--migrations <dir>`: Directory containing migration files (default: `./migrations`).
- `--dry-run`: Show the steps that would be executed, but do not apply changes.

**Example:**
```
janus apply --from 1.1.0 --to 1.3.0 --migrations ./db/migrations
```

## Migration File Format

Migration files must follow the naming convention:

```
v<from_semver> -> v<to_semver>.sql
```

Example:
- `v1.0.0 -> v1.1.0.sql`
- `v1.1.0 -> v1.0.0.sql`

## Notes
- The `plan` command is safe to run and does not modify the database.
- The `apply` command will execute DDL statements unless `--dry-run` is specified.
- Janus will determine the shortest available path between versions, even if multiple migration files exist.

For more details, see the README.md.

