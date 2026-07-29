---
task-id: CON-003
roadmap-id: CON-003
status: delivered
archived: 2026-07-29
review: review-02.md
capability-impact:
  type: update
  ids:
    - CAP-001
---

# Summary

## Task

Define the normative contract for Concoct's seven initial commands and its artifact-backed workflow state machine without implementing the planned CLI.

## Delivered outcome

- Added `doc/command-reference.md`, covering purpose, inputs, starting states, artifact access, persona selection, effects, prompts, resulting states, failures, and next commands for all seven initial commands.
- Added `doc/state-machine.md`, defining observable workflow states, precedence, valid and invalid transitions, remediation loops, blocked-review recovery, and transactional archival.
- Kept prompt rendering distinct from completed role work and clearly labeled the command surface as normative design rather than current executable behavior.

## Key decisions

- Durable repository artifacts, rather than conversation history, determine workflow state.
- Review outcomes are explicit states, with sequential append-only review history.
- Changes-requested remediation uses `remediates-review`; blocked-review recovery uses validated `blocked-review-resolution` evidence owned by the Task Planner or Developer.
- Archival clears current artifacts only after archive, capability, roadmap, and cross-reference writes validate.

## Files and areas changed

- `doc/command-reference.md`
- `doc/state-machine.md`
- `.concoct/capabilities.md`
- `.concoct/roadmap.md`

The accepted implementation was documentation-only. No CLI, parser, validator, test, template, prompt, or persona implementation changed.

## Verification

- Confirmed seven command contracts contain all ten required contract sections.
- Manually traced the happy path, repeated changes-requested remediation, both blocked-review recovery routes, approved archival, and representative invalid states.
- Confirmed both documents are under `doc/`, their relative link resolves, and `.concoct/docs/` is absent.
- Searched for obsolete persona names and stale in-scope paths.
- `git diff --check` passed before approval.
- No Markdown linter or automated test suite was available; Review 02 judged structural searches and manual transition tracing proportionate for this documentation-only task.

## Review outcome

`review-02.md` approved the implementation after confirming that Review 01's blocked-review recovery finding was fixed. Review 01 remains preserved as prior review history.

## Capability changes

Updated `CAP-001` to include the accepted normative command reference and artifact-backed state machine as part of Concoct's durable file-based workflow contract. The capability continues to state that a state-aware CLI does not yet enforce these contracts.

## Follow-up work

- Later schema, state-detection, and command implementation must implement `remediates-review` and `blocked-review-resolution` consistently.
- CON-005 through CON-009 remain responsible for executable CLI foundations, prompt rendering, planning, implementation/review loops, and archival automation.

## References

- Roadmap item: `CON-003` in `.concoct/roadmap.md`
- Capability: `CAP-001` in `.concoct/capabilities.md`
- Approving review: `review-02.md`
- Product documentation: `doc/command-reference.md`, `doc/state-machine.md`
