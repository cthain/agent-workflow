---
task-id: CON-028
roadmap-id: CON-028
status: archived
archived: 2026-07-30
review: review-02.md
delivery: pending-integration
capability-impact:
  type: add
  ids:
    - CAP-009
---

# Summary

## Task

Add `concoct next` as the single ready-state recommendation so Product Owner
judgment can select one valid follow-up from deterministic repository evidence
without the CLI choosing work or changing workflow state.

## Delivered outcome

Concoct now permits `concoct next` only in valid ready state and renders a
byte-deterministic Product Owner prompt containing canonical roadmap,
capability, dependency, prerequisite, relevant archive, and supported-origin
evidence. The guidance requires one evidence-backed recommendation: plan an
eligible item, perform supported product or roadmap work, resolve a named
blocker or inconsistency, or report that no actionable work is recorded.

Ready-state status, initialization, bootstrap, and successful integration now
recommend exactly `concoct next`. Rendering remains advisory and read-only:
it does not rank or select work, create or activate a task, or mutate durable
workflow evidence.

## Key decisions

- Reuse shared planning eligibility as the structural authority instead of
  introducing a competing parser.
- Keep deterministic evidence presentation separate from semantic Product Owner
  judgment.
- Bound supported origins to roadmap planning and human product input or roadmap
  maintenance until future origin contracts are accepted.
- Preserve create-only output safety, invalid-evidence refusal, nested project
  discovery, and Git lifecycle boundaries.

## Files and areas changed

- Extended workflow inspection, prompt rendering, CLI dispatch, initialization,
  and integration guidance plus focused Go tests and golden fixtures.
- Added the canonical next-action prompt and updated Product Owner, Archivist,
  handoff, skill, documentation, and embedded-template guidance.
- Reconciled ready-state recommendations across source and installed assets.

## Verification

- `go test -count=1 ./...` — passed.
- `go vet ./...` — passed.
- `go build ./cmd/concoct` — passed.
- Formatting, source/template parity, executable-mode validation, stale-guidance
  searches, generated-project execution, and `git diff --check` — passed.
- Full-output golden coverage exercises eligible work, supported product input,
  roadmap reconciliation, a dependency blocker, and no actionable work.
- The legacy `bash -n cmd/concoct/concoct` instruction is inapplicable because
  the tracked compatibility path is an ELF executable; this predates CON-028.

## Review outcome

`review-02.md` approved the implementation after confirming that Review 01's
major finding—missing outcome and ready-transition regression coverage—was
fixed. Both reviews are preserved.

## Capability changes

Added `CAP-009` for deterministic, evidence-backed, read-only Product Owner
next-action recommendation. Existing CAP-001, CAP-005, CAP-006, and CAP-008
remain related prerequisite capabilities and retain their existing descriptions.

## Skipped and follow-up work

Unsupported task origins, direct agent execution, lifecycle mutation, and
automated prioritization remain outside CON-028. The compatibility executable
and stale shell-specific repository check should be reconciled separately.
Delivery remains pending `concoct integrate`; CON-028 stays active and current
task evidence remains intact until integration succeeds.

## References

- Roadmap item: `CON-028` in `.concoct/roadmap.md`
- Capability: `CAP-009` in `.concoct/capabilities.md`
- Approving review: `review-02.md`
- Prior review: `review-01.md`
- User documentation: `README.md`
- Normative contracts: `doc/command-reference.md`, `doc/state-machine.md`
