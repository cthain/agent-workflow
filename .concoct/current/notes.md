# Task notes

## Structural decisions

### Superseding layout decision

- `.concoct/` is reserved for generated client installs, not for organizing Concoct's own source repository.
- Source directories move back to the repository root.
- Client task state and personas move under `.concoct/`; conventional root adapters remain at the client root because their tools require those locations.

- The template skill and repository skill serve different installed contexts, so both are retained.
- `concoct` continues to stage generated files with `git add .`; this preserves existing behavior without creating a commit.

## Verification

### Client `.concoct/` layout

- `bash -n concoct` passed and executable mode remains set.
- End-to-end initialization passed in a system temporary directory.
- Verified `.concoct/current/`, `.concoct/archive/`, `.concoct/personas/`, the bootstrap prompt, all six personas, and staged Git state.
- Verified conventional adapters remain at the client root, including `AGENTS.md`, `CLAUDE.md`, `.codex/`, and `.github/`.
- Verified generated projects contain neither `.planning/` nor a root `personas/` directory.
- Verified this source repository contains no `.concoct/` directory.
- `git diff --check` passed.

- `bash -n concoct` passed.
- The previous end-to-end result covered the superseded client layout; verification for the new `.concoct/` client layout is recorded below.
- Verified generated files are staged but no commit is created.
- `git diff --check` passed.
- Stale product-name search passed; `multi-agent-workflow.md` remains intentionally named as a workflow artifact.

## Tooling note

- Ruby was unavailable for a scripted Markdown link scan. The README's two local documentation targets were checked directly and exist.
