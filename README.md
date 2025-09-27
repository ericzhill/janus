# Janus

Janus is a small Go package for managing database schema versions. Named after the Roman god of transitions and duality, Janus focuses on helping you move a database schema both forward and backward safely, predictably, and with a developer-friendly API.

## Why “Janus”?
Janus looks both to the past and the future. Likewise, this package is designed to move a schema to any point in its version history—upgrades and rollbacks—with equal care.

## What it does
- Tracks schema versions in a simple, explicit manner
- Applies migrations forward (up) and backward (down)
- Encourages safe, incremental change with clear ordering
- Aims to be easy to consume in applications and tooling

## Goals
- Effective: migrations apply deterministically, in order
- Reversible: every forward change has a corresponding backward path
- Simple: keep the API and operational model straightforward

## Core concept: directional, versioned transitions
Janus models every schema change as an explicit transition between two semantic versions. Each transition is stored in a file whose name encodes the "from" and "to" versions separated by a directional arrow.

- Format: v<from_semver> -> v<to_semver>.sql
- Example (forward): v1.0.0 -> v1.1.0.sql contains the DDL to move from 1.0.0 to 1.1.0
- Example (reverse): v1.1.0 -> v1.0.0.sql contains the DDL to rollback from 1.1.0 to 1.0.0

Notes
- Transitions can skip versions. You might have v1.1.0 -> v1.3.0.sql without requiring a v1.2.0 waypoint.
- Multiple paths can exist to reach the same target. Janus can choose an available path that satisfies the move you requested.

A few simple examples
- Direct neighbors:
  - v1.0.0 -> v1.1.0.sql
  - v1.1.0 -> v1.0.0.sql
- Skipping versions:
  - v1.1.0 -> v1.3.0.sql
  - v1.3.0 -> v1.1.0.sql
- Multiple paths to v2.0.0 from v1.0.0:
  - Path A: v1.0.0 -> v1.1.0.sql, v1.1.0 -> v1.2.0.sql, v1.2.0 -> v2.0.0.sql
  - Path B: v1.0.0 -> v1.3.0.sql, v1.3.0 -> v2.0.0.sql

We’ll keep the README focused on the concept. More detailed, end-to-end examples will live in TUTORIAL.md.

## Getting started
This repository currently focuses on the core concept and interface for versioned schema transitions. Typical usage looks like:

- Define your schema versions and migrations (up/down steps)
- Ask Janus to transition the database to a target version
- Janus determines the delta and executes steps in order

Implementation details (types, helpers, and storage format) may evolve. When ready, examples and API documentation will be added to show:
- How to declare migrations
- How to run “up to latest”
- How to roll back to an earlier version

## Status
Early-stage. README describes the intent and direction; API and examples will follow as the package evolves.

## Contributing
- Open issues for ideas, questions, or bugs
- Propose clear, minimal PRs that move the core forward

## License
This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
