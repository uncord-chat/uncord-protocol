# Contributing to Uncord Protocol

Thanks for your interest in contributing. This repository contains the shared protocol definitions used by both the server and client. Changes here affect both projects, so care is needed.

## Getting started

1. Fork the repository and clone your fork.
2. Run `npm install` to set up BiomeJS for TypeScript formatting.
3. Ensure you have Go installed for working with the Go definitions.

## How this repo works

Every constant is defined twice: once in Go and once in TypeScript. Both must stay in sync. If you add a permission bit, opcode, or error code, update both the `.go` and `.ts` files.

## Pull request guidelines

- One concern per pull request. Don't mix new constants with refactors.
- Keep Go and TypeScript definitions in sync within the same PR.
- Run `make fmt` and `make lint` before submitting.
- Write a clear PR description explaining what changed and why.
- Reference the issue number if one exists.

## Code style

- Run `make fmt` to format both Go and TypeScript files.
- Go follows `gofmt`. TypeScript follows the BiomeJS configuration in `biome.json`.
- Commit messages: use imperative mood ("Add permission bit" not "Added permission bit").

## Reporting issues

If you find a mismatch between Go and TypeScript definitions, or a missing constant, open a GitHub issue describing the discrepancy.

## Code of conduct

Be constructive, be patient, and be kind. Harassment and discrimination are not tolerated. See [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for details.
