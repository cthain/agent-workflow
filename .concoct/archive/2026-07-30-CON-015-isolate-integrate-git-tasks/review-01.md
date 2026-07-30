---
task-id: CON-015
review: 1
status: changes-requested
created: 2026-07-30
persona: reviewer
---

# Review 01

## Outcome

`changes-requested`

## Summary

The implementation establishes the intended Git-backed planning, archival, integration, documentation, and template contracts, and the full automated suite passes. However, two integration recovery paths can either overwrite work that is not part of the recorded transaction or strand valid recovery evidence. Those defects violate the acceptance criteria for destructive-operation safety and interruption recovery.

## Acceptance criteria assessment

- Git planning, deterministic branch naming, collision refusal, role checkout validation, archival/integration states, local squash integration, conflict continuation, upstream policy, non-Git compatibility, documentation, and source/template parity are implemented and covered by passing tests.
- The `--abort` safety criterion is not met because abort does not validate the current trunk HEAD or worktree before a hard reset.
- The interruption-safety criterion is not met at the prepared-to-trunk-checkout boundary because recovery is persisted before checkout, while both recovery commands require the trunk already to be checked out.

## Findings

### Finding 1 — Abort can discard unrelated trunk work

- Severity: critical
- Status: unresolved
- Evidence: `internal/integration/integration.go:200-211` checks only that the current branch name equals the recorded trunk, then unconditionally runs `git reset --hard <pre-integration-head>`. It does not require a clean worktree, reject an unrelated Git operation, or prove that current `HEAD` is one of the transaction's recorded integration/delivery commits. Likewise, `resume` at lines 108-133 validates only the branch name before proceeding from recorded later phases. A user can commit unrelated work on the trunk while recovery exists and then run `concoct integrate --abort`; the hard reset removes those commits from the branch. Uncommitted changes are discarded as well.
- Impact: The command can overwrite work outside Concoct's integration transaction. This directly violates the constraints to never discard a dirty worktree and to preserve recovery metadata, and the acceptance criterion that unsafe/interrupted scenarios never overwrite work.
- Required outcome: Before any destructive reset or phase continuation, validate the checkout, operation state, cleanliness/staged state appropriate to that phase, and exact recorded HEAD relationship. Abort must refuse rather than reset when the trunk has advanced or contains unrelated work. Add real-repository tests covering dirty abort, an unrelated trunk commit after recovery creation, and later-phase continuation with unexpected HEAD.

### Finding 2 — Prepared recovery can be stranded on the task branch

- Severity: major
- Status: unresolved
- Evidence: `internal/integration/integration.go:91-97` writes the exclusive `phase: prepared` recovery record before checking out the recorded trunk. If checkout fails or execution stops after the write and before checkout completes, the task branch remains checked out. Both `resume` at lines 108-112 and `abort` at lines 200-204 reject unless the trunk is already checked out, so neither documented recovery command can repair the state. A plain retry also refuses because recovery already exists (lines 62-64).
- Impact: A failure at a deliberately persisted checkpoint leaves an integration transaction that cannot be recovered using the advertised commands, contrary to the plan's retry/idempotency invariants and the acceptance criterion requiring interrupted scenarios to remain recoverable.
- Required outcome: Make the prepared phase recoverable regardless of whether checkout completed—for example, safely validate and complete the recorded checkout in `--continue`, or safely remove/rollback the pre-mutation recovery in `--abort`. Add a real-repository test simulating interruption immediately after the prepared record is written while the task branch remains checked out.

## Prior finding disposition

No prior reviews exist.

## Verification performed

- `go test ./...` — passed.
- `go vet ./...` — passed.
- `git diff --check` — passed.
- `bash -n cmd/concoct/concoct.sh` — passed.
- Executable-mode check for `cmd/concoct/concoct.sh` — passed.
- Inspected the complete diff and the relevant Git, integration, workflow, CLI, prompt, test, documentation, persona, and template contracts.
- Compared each changed template counterpart with its repository source counterpart; no mismatches were found.

## Capability impact assessment

The declared impact remains correct: successful completion should add `CAP-007` for optional Git-backed task isolation, archival, integration, and recovery while retaining the non-Git workflow. Capability truth must not be updated until these safety findings are resolved and the task is approved.

## Scope assessment

The implementation is within CON-015's product scope. The requested remediation is limited to the already-required integration safety and recovery behavior.

## Documentation assessment

The documentation consistently describes exact abort and interruption recovery, but the implementation does not yet fulfill those claims. Documentation changes are needed only if the chosen safe behavior changes the public recovery instructions.

## Handoff

- Current state: `changes-requested`.
- Work remaining: remediate both recovery findings, add the specified destructive-operation and interruption tests, rerun the complete verification suite, and record finding dispositions in notes.
- Expected next role: Developer in review-remediation mode.
- Recommended next command: `concoct code`.
