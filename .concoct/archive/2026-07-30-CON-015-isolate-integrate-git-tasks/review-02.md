---
task-id: CON-015
review: 2
status: changes-requested
created: 2026-07-30
persona: reviewer
---

# Review 02

## Outcome

`changes-requested`

## Summary

The remediation fixes prepared-phase recovery and prevents abort or continuation from crossing an unexpected committed trunk HEAD. One destructive case from Review 01 remains: before the integration commit, abort treats every staged change as transaction-owned and then hard-resets it. An unrelated staged file or modification can therefore still be discarded.

## Acceptance criteria assessment

- Prepared recovery on the task branch can now safely continue or abort, satisfying the interruption case from Review 01 Finding 2.
- Exact phase-to-HEAD validation prevents abort and later continuation from rewinding unrelated committed trunk work.
- The requirement that unsafe recovery never overwrite work remains unmet for unrelated staged changes present during a prepared, squashed, or conflicted phase.
- Existing automated checks pass, but the destructive-operation tests cover untracked and committed unrelated work, not unrelated staged work.

## Findings

### Finding 1 — Abort still discards unrelated staged work

- Severity: critical
- Status: unresolved
- Evidence: `validateAbortWorktree` in `internal/integration/integration.go:315-335` permits any status entry whose second column is blank, including `A ` for a newly staged file and `M ` for an unrelated staged modification. `abort` then calls `ResetHard(r.PreIntegrationHead)` at line 260. A newly created file that is staged during conflict recovery is therefore accepted by the guard and removed from both the index and worktree by the reset. The new `dirty abort` test at `internal/integration/integration_test.go:210-223` exercises only an untracked `??` file, so it does not cover this path.
- Impact: User work unrelated to Concoct can still be destroyed by `concoct integrate --abort`. This is the remaining destructive portion of Review 01 Finding 1 and violates the explicit constraints to never discard a dirty worktree or overwrite work during unsafe/interrupted recovery.
- Required outcome: Abort must distinguish transaction-owned index/worktree state from unrelated staged additions or modifications and refuse before reset when unrelated staged work exists. Preserve all refused work and recovery evidence. Add real-repository tests for at least a staged new file and a staged unrelated tracked-file modification during pre-commit recovery.

## Prior finding disposition

- Review 01 Finding 1 — partially fixed. Exact HEAD checks, unrelated-operation checks, clean guards for committed phases, and tests for untracked/committed drift are present. The staged-work destructive case above remains unresolved.
- Review 01 Finding 2 — fixed. `prepared` recovery now validates the task branch and unchanged trunk, then either completes checkout/squash on continue or removes only pre-mutation recovery on abort. Both paths have real-repository tests.

## Verification performed

- `go test ./...` — passed.
- `go vet ./...` — passed.
- `go build ./cmd/concoct` — passed.
- `git diff --check` — passed.
- `bash -n cmd/concoct/concoct.sh` — passed.
- Executable-mode check for `cmd/concoct/concoct.sh` — passed.
- Inspected the remediation code, real-repository recovery tests, task-plan remediation metadata, notes dispositions, complete implementation diff, and relevant workflow/documentation contracts.
- Compared changed template files with their repository source counterparts; no mismatches were found.

## Capability impact assessment

The planned `add CAP-007` impact remains accurate, but capability truth must remain unchanged until destructive abort behavior is safe and the task is approved.

## Scope assessment

The remediation remains within CON-015. Resolving the remaining finding is required by the existing destructive-operation constraints and does not broaden scope.

## Documentation assessment

The documentation's promise of exact, safe abort remains appropriate as the target behavior. No additional documentation change is required if the implementation is corrected to match it.

## Handoff

- Current state: `changes-requested`.
- Work remaining: protect unrelated staged work during pre-commit abort, add staged-addition and staged-modification regression tests, record the Review 02 disposition, and rerun the complete verification suite.
- Expected next role: Developer in review-remediation mode.
- Recommended next command: `concoct code`.
