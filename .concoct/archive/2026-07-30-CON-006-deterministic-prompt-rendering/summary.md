---
task-id: CON-006
roadmap-id: CON-006
status: delivered
archived: 2026-07-30
review: review-02.md
capability-impact:
  type: add
  ids:
    - CAP-006
---

# Summary

## Task

Implement deterministic, inspectable prompt rendering for `concoct roadmap`,
`concoct plan <roadmap-id>`, `concoct code`, and `concoct review` without
launching an agent or treating rendered output as completed role work.

## Delivered outcome

Concoct now validates repository state and renders complete role-aware prompts
for roadmap intake, task planning, development, and independent review. The
renderer distinguishes initial and continuing implementation,
changes-requested remediation, post-remediation review, and both validated
blocked-review recovery routes. Output is deterministic, goes to stdout by
default, and can be written byte-identically to a create-only output path.
Rendering itself leaves workflow artifacts and state unchanged.

## Key decisions

- Reuse `internal/workflow` as the validation and state authority through a
  narrow read-only prompt context.
- Append the repository's canonical manual handoff assets rather than creating
  a divergent body of durable role rules.
- Sort discovered reviews and relevant archive summaries deterministically.
- Refuse existing output destinations and remove incomplete newly created
  output after write or close failure.
- Protect the complete observable prompt bytes with explicit, opt-in-updated
  golden fixtures.

## Files and areas changed

- Added `internal/prompt` rendering, state/mode selection, tests, and nine
  full-output golden fixtures.
- Extended `internal/cli` with the four prompt commands, argument validation,
  output handling, and integration tests.
- Extended `internal/workflow` with validated read-only prompt context and
  roadmap dependency validation.
- Updated `README.md` with prompt-only command and output guidance.

## Verification

- `go test ./...` — passed after remediation and before archival.
- `go vet ./...` — passed.
- Temporary Linux CLI build and `help` smoke test — passed during Review 02.
- `bash -n cmd/concoct/concoct.sh` and executable-mode check — passed.
- `git diff --check` — passed before archival.
- Nine checked-in golden cases cover all four commands and every materially
  distinct development and review mode in the accepted plan.

## Review outcome

`review-02.md` approved the implementation after confirming that Review 01's
sole major finding—missing full-output golden coverage—was fixed. Review 01 is
preserved as changes-requested history.

## Capability changes

- Added `CAP-006` for executable deterministic, role-aware prompt rendering.
- Updated CAP-002 to relate the manual prompt corpus to executable rendering.
- Updated CAP-005 to distinguish its initialization/status behavior from the
  prompt-rendering surface now recorded by CAP-006.
- Removed prompt rendering from the known capability gaps while retaining the
  absence of role-completion mutations, archival automation, and direct agent
  execution.

## Skipped and follow-up work

- Prompt commands do not perform role-completion workflow mutations or launch
  agents; those remain assigned to later roadmap items.
- The CLI does not implement `archive`.
- Archive relevance remains conservative until a future archive index or
  richer metadata contract is delivered.

## References

- Roadmap item: `CON-006` in `.concoct/roadmap.md`
- Capability: `CAP-006` in `.concoct/capabilities.md`
- Approving review: `review-02.md`
- Prior review: `review-01.md`
- User documentation: `README.md`
- Normative contracts: `doc/command-reference.md`, `doc/state-machine.md`
