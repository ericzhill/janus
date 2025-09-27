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
TBD