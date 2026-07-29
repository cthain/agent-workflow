---
task-id: CON-005
roadmap-id: CON-005
status: delivered
archived: 2026-07-29
review: review-03.md
capability-impact:
  type: add
  ids:
    - CAP-005
---

# Summary

## Task

Implement the Concoct CLI foundation in Go, delivering the first executable
commands from the CON-003 command and state-machine contract.

## Delivered outcome

Concoct now provides a Linux-buildable Go CLI with `init` and read-only
`status` commands. The CLI embeds the complete checked-in template tree,
initializes projects independently of the caller's working directory, stages
generated files without creating a commit, discovers projects from nested
directories, validates workflow metadata and review evidence, and reports
deterministic workflow states, diagnostics, and next actions.

The legacy shell entry point remains executable as a thin source-checkout
wrapper around the Go implementation.

## Key decisions

- Embed `all:templates` in the binary so installed initialization does not
  depend on repository-relative runtime paths.
- Preserve the established bootstrap behavior of staging generated files while
  creating no initial commit.
- Treat malformed or contradictory durable artifacts as `invalid` rather than
  guessing a workflow state.
- Keep recovery-disposition matching textual because the accepted notes schema
  does not define structured finding identifiers.

## Files and areas changed

- Added the Go module, embedded-template boundary, CLI entry point, and internal
  CLI, project, and workflow packages.
- Added unit and integration tests for initialization, discovery, metadata,
  state selection, recovery evidence, diagnostics, Git behavior, and
  non-mutation.
- Replaced the broken shell initializer implementation with a delegation
  wrapper and updated README usage and development guidance.

## Verification

- `go test ./...` — passed.
- `go vet ./...` — passed.
- `bash -n cmd/concoct/concoct` and executable-mode check — passed.
- `git diff --check` — passed before archival.
- External temporary-directory initialization — passed, including root files,
  dotfiles, nested templates, personas, prompts, current/archive directories,
  bootstrap guidance, staged Git files, and no commit.
- Nested-directory `status` — reported `ready` and left Git status unchanged.

## Review outcome

`review-03.md` approved the implementation after the findings in
`review-01.md` and `review-02.md` were remediated. No material findings
remain.

## Capability changes

- Added `CAP-005` for executable project initialization and workflow-status
  reporting.
- Updated `CAP-003` to remove its obsolete broken-initializer limitation and
  replace the failed bootstrap evidence with the accepted end-to-end result.
- Unrelated CAP-003 template limitations remain recorded.

## Follow-up work

- CON-006 can build deterministic prompt rendering on the delivered CLI and
  state infrastructure.
- Structured review-finding identifiers may eventually replace textual
  remediation-disposition detection.

## References

- Roadmap: `.concoct/roadmap.md#con-005--implement-the-concoct-cli-foundation-in-go`
- Capability: `.concoct/capabilities.md#cap-005--executable-cli-initialization-and-workflow-status`
- Approval: `review-03.md`
- Normative contracts: `doc/command-reference.md`, `doc/state-machine.md`
