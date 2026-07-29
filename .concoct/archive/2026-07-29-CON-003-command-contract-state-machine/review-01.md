---
task-id: CON-003
review: 1
status: changes-requested
created: 2026-07-29
persona: reviewer
---

# Review 01

## Outcome

`changes-requested`

## Summary

The two documentation deliverables are present in the repository's canonical `doc/` directory, clearly distinguish planned command behavior from the legacy initializer, and provide complete per-command sections, artifact-backed states, remediation cycles, failure handling, and transactional archive ordering. One major contract defect remains: the blocked-review recovery path cannot be derived from the normative detection and command rules. Later CLI work would have to invent the durable evidence and transition that CON-003 is intended to settle.

## Acceptance criteria assessment

- Passed: `doc/command-reference.md` covers all seven initial commands and all ten required contract fields for each command.
- Passed: supported ordinary states are defined from repository artifacts, with explicit invalid-state and state-preserving behavior.
- Passed: the happy path requires no separate `handoff` command.
- Passed: one or more `changes-requested → code → review` cycles preserve sequential reviews through `remediates-review` evidence.
- Changes required: blocked-review routing names responsible roles, but does not define an observable recovery transition that is consistent with review precedence and valid command starting states.
- Passed: archival requires approval evidence, preserves review history, reconciles roadmap and capability records, and delays clearing current artifacts until validation succeeds.
- Passed: terminology, canonical paths, persona names, planned-versus-implemented language, and optional-command boundaries are otherwise consistent across the deliverables.

## Findings

### Finding 1 — Blocked-review recovery has no reachable artifact-backed transition

- Severity: `major`
- Status: `unresolved`
- Evidence: `doc/state-machine.md` defines the latest `blocked` review as `review-blocked` and states that an approved or blocked review cannot be bypassed with remediation metadata (lines 43 and 72–75). It later says durable evidence can return the task to an implementation or review phase, but does not identify that evidence (lines 118–127 and 206–214). `doc/command-reference.md` excludes `review-blocked` from `code` until owned artifacts already record a valid return to implementation (lines 281–287), while `review` accepts only `implementation-complete` when no existing review determines the actionable state (lines 342–344 and 375–381). No initial command or detection rule can establish the prerequisite return state.
- Impact: A repository with a resolved blocked review remains deterministically `review-blocked` under the stated precedence rules. Implementers of state detection and prompt rendering must invent how a Product Owner, Task Planner, Developer, or human records resolution, whether another review is appended, and which command becomes valid. This violates the goal that later CLI work not invent transition rules and the acceptance requirement for clear, actionable recovery from blocked outcomes.
- Required outcome: Define the durable evidence and state-detection rule for each supported blocked-review resolution route, and align the relevant command starting states and failure guidance with it. The corrected contract must preserve the blocked review append-only, prevent a blocked outcome from being treated as approval, identify the responsible artifact owner, and make at least one explicit trace from `review-blocked` to a valid `code` or `review` invocation without depending on an optional command.

## Prior finding disposition

No prior `.concoct/current/review-NN.md` files existed.

## Verification performed

- Read the complete active plan, notes, capability ledger, reviewer persona, both deliverables, repository guidance, relevant roadmap entry, canonical skill, personas, and handoff prompts.
- Inspected the complete tracked diff and all untracked task deliverables; noted that the ordinary Git diff does not include untracked document contents.
- Confirmed `doc/command-reference.md` and `doc/state-machine.md` exist and `.concoct/docs/` is absent.
- Counted seven command contracts and verified each has one section for purpose, required inputs, valid starting states, files read, selected persona, files created or updated, prompt produced, resulting state, failure conditions, and recommended next commands.
- Traced the happy path, repeated changes-requested remediation, blocked-review routing, invalid evidence, and archive ordering manually.
- Searched the deliverables for obsolete persona names and stale `.planning/` or `.concoct/docs/` paths; no stale match was found.
- Ran `git diff --check`; it passed.
- No automated test suite or Markdown linter is present. This documentation-only contract was therefore verified with structural searches and manual transition tracing.

## Scope assessment

The task implementation is documentation-only and the deliverables are confined to the approved command and state-machine contract. Active task artifacts were updated as expected. The working tree also contains a modification to `.concoct/roadmap.md`; notes identify it as pre-existing and outside both developer passes, so it is not attributed to this implementation or made a requested task change. No source, tests, templates, prompts, personas, capabilities, or archives were changed by the documented implementation.

## Documentation and compatibility assessment

Placement under `doc/` follows repository conventions and local cross-links resolve. The documents correctly avoid claiming the future command surface is implemented and preserve compatibility context for the legacy shell initializer. Apart from Finding 1, the contracts are internally usable and keep tool-specific behavior out of the durable workflow design.

## Capability impact assessment

The declared impact is accurate: after approval and archival, CAP-001 should be updated to reference the normative command and state-machine documents while continuing to state that command enforcement is not yet implemented. This task does not add a runtime CLI capability and does not change CAP-002.

## Handoff

Responsible role: Developer.

Resolve Finding 1 within the existing documentation-only scope, update the developer evidence and handoff, and request a fresh independent review.

Next command:

```text
concoct code
```
