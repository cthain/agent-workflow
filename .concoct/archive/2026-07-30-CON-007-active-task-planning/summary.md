---
task-id: CON-007
roadmap-id: CON-007
status: archived
archived: 2026-07-30
review: review-02.md
delivery: pending-integration
capability-impact:
  type: add
  ids:
    - CAP-008
---

# Summary

## Task

Complete active task planning so `concoct plan <roadmap-id>` validates one
eligible roadmap item and its accepted capability prerequisites, then supplies
the deterministic context needed for a safe, implementation-ready Task Planner
transition.

## Delivered outcome

Planning eligibility now resolves every declared capability prerequisite
uniquely against active capability truth before branch creation or prompt
output. The Task Planner prompt includes deterministic limitation and archive
provenance context while leaving semantic compatibility and role-owned plan
creation to the planner. Existing paired-artifact validation, selected-only
roadmap activation, Git identity, no-overwrite behavior, and planned-state
detection remain the completion authority.

## Key decisions

- Extend the existing workflow parser and prompt renderer rather than create a
  second planning model or state authority.
- Treat capability identity and active status as structural eligibility, while
  keeping limitation compatibility as Task Planner judgment.
- Reject malformed, duplicate, missing, and inactive prerequisite evidence at
  the planning-eligibility boundary before Git or output mutation.
- Preserve prompt rendering as guidance; it does not author or activate task
  artifacts.

## Files and areas changed

- Extended `internal/workflow` with capability prerequisite parsing,
  validation, diagnostics, and tests.
- Extended `internal/prompt` with prerequisite, limitation, and archive context,
  including golden-output coverage.
- Added CLI/Git boundary tests for safe pre-branch failure.
- Reconciled command and state-machine documentation plus source/template Task
  Planner persona and handoff assets.

## Verification

- `go test -count=1 ./...` — passed immediately before archival.
- `go vet ./...` — passed.
- `bash -n cmd/concoct/concoct.sh` and executable-mode validation — passed.
- `git diff --check 302fb230813be6d08aa6b8a8cd942b3ed9245fa7..HEAD` — passed.
- Review 02 additionally passed focused workflow, CLI, prompt, and Git tests; a
  Go build; and source/template parity checks.

## Review outcome

`review-02.md` approved the implementation after confirming that Review 01's
single diagnostic-context finding was fixed. Both review files are preserved.

## Capability changes

Added CAP-008 for validated active task planning. CAP-001, CAP-005, CAP-006,
and CAP-007 remain prerequisite capabilities and are not replaced.

## Skipped and follow-up work

- Direct agent execution and automated role completion remain future work.
- Capability schema changes must update the intentionally narrow Markdown
  parser and its tests.
- Delivery is pending `concoct integrate`; CON-007 remains active and current
  task evidence remains intact until integration succeeds.

## References

- Roadmap item: `CON-007` in `.concoct/roadmap.md`
- Capability: `CAP-008` in `.concoct/capabilities.md`
- Approving review: `review-02.md`
- Prior review: `review-01.md`
- Normative contracts: `doc/command-reference.md`, `doc/state-machine.md`
