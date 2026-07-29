# Separate source and client layouts

## Goal

Keep Concoct's source files in ordinary root-level directories while isolating Concoct-owned material in generated client projects under `.concoct/`.

## Decisions

- Concoct source directories live at this repository's root; `.concoct/` is reserved for client installs.
- Generated projects use `.concoct/current/`, `.concoct/archive/`, and `.concoct/personas/`.
- Conventional tool adapters such as `AGENTS.md`, `.codex/`, and `.github/` remain at the generated project root.
- The generated-project Codex skill remains in the root adapter template; the repository-facing skill lives under `skills/`.
- `git add .` remains because staging is existing behavior and the script makes no commit.

## Work

- [x] Inspect repository structure, references, and bootstrap command behavior.
- [x] Move source assets out of `.concoct/`.
- [x] Consolidate client-owned task state and personas under `.concoct/`.
- [x] Update documentation, prompts, adapters, and bootstrap output.
- [x] Verify paths, shell syntax, and end-to-end initialization.

## Completion criteria

- [x] This source repository has no `.concoct/` directory.
- [x] Generated projects keep removable Concoct material under `.concoct/`.
- [x] Root adapters reference the client `.concoct/` paths.
- [x] `concoct` passes syntax and end-to-end checks.
