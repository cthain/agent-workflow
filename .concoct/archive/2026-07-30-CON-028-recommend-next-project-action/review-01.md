---
task-id: CON-028
review: 1
status: changes-requested
created: 2026-07-30
persona: reviewer
---

# Review 01

## Outcome

`changes-requested`

## Summary

The implementation establishes the intended read-only `concoct next` command,
reuses planning eligibility, changes the ready-state recommendation, and keeps
the reviewed source/template assets aligned. The code compiles and the existing
test suite passes. Acceptance is withheld because the committed verification
does not demonstrate several explicitly required outcome and transition cases,
while the task plan marks that verification phase complete.

## Acceptance criteria assessment

- Passed by inspection and existing tests: ready-only dispatch, deterministic
  repeated rendering for the covered fixture, nested discovery, stdout/file
  byte equality, create-only output refusal, workflow-artifact non-mutation,
  shared plan eligibility, argument validation, and source/template parity.
- Passed by inspection: `status`, initialization, bootstrap, documentation,
  skills, personas, and handoffs use `concoct next` at the changed ready-state
  boundaries.
- Not demonstrated by committed tests: full-output coverage for every required
  recommendation outcome, invalid canonical evidence at the `next` boundary,
  initialization/bootstrap next-command output, and successful integration's
  sole ready-state recommendation.

## Findings

### Finding 1 — Required outcome and transition coverage is incomplete

- Severity: major
- Status: open
- Evidence: `.concoct/current/task-plan.md:190-199` requires focused invalid-input
  tests, full-output golden cases for an eligible item, supported non-planning
  Product Owner work, roadmap reconciliation, a specific blocker, and no work,
  plus initialization and integration ready-output exercises. The acceptance
  contract repeats these requirements at lines 205-232. In
  `internal/prompt/render_test.go:12-72`, only one `next` fixture is protected by
  a full-output golden. `internal/prompt/render_test.go:91-103` covers the empty
  roadmap only with two substring assertions. There are no separate full-output
  cases for supported non-planning work or reconciliation, and no direct
  `next` invalid-canonical-evidence test. `internal/project/project_test.go:14-63`
  does not assert initialization output or bootstrap content recommends exactly
  `concoct next`. `internal/integration/integration_test.go:12-30` verifies
  delivery cleanup but does not assert the resulting ready report's sole next
  command.
- Impact: regressions in required Product Owner outcome guidance, invalid-state
  refusal, and the two lifecycle entry points can pass the suite. This conflicts
  with the task's observable acceptance contract and makes the recorded claim
  that Phase 5 is complete unsupported.
- Required action: add the planned full-output golden cases for the distinct
  evidence/outcome scenarios; add focused invalid canonical evidence coverage
  for `next`; and assert initialization/bootstrap and successful integration
  recommend exactly `concoct next`. Keep non-mutation and deterministic-output
  checks for the added cases, then rerun the documented verification.

## Verification performed

- `go test -count=1 ./...`: passed.
- `go vet ./...`: passed.
- `go build ./cmd/concoct`: passed.
- `git diff --check 20321a3..HEAD`: passed.
- Executable-bit check for `cmd/concoct/concoct`: passed.
- Source/template byte parity for the changed skill, Product Owner and Archivist
  personas, next-action prompt, and archival handoffs: passed.
- Complete implementation diff from `20321a3` through `f9b60ed`: inspected.

The documented `bash -n cmd/concoct/concoct` check was not rerun because the
tracked path is an ELF executable rather than a shell source; the Developer's
notes accurately record this pre-existing repository-instruction mismatch.

## Capability impact assessment

The declared `add` impact for provisional `CAP-009` is appropriate if the work
is later approved. Existing capability truth must remain unchanged during
remediation.

## Scope and documentation assessment

The implementation remains within CON-028's product scope. The changed source
and installed template assets are aligned, and the reviewed documentation
correctly distinguishes `status`, `next`, `roadmap`, and `plan`. No unrelated
implementation change was identified.

## Handoff

- Current state: review changes requested; implementation remains on the
  recorded task branch.
- Work remaining: complete the missing verification coverage in Finding 1 and
  update Developer-owned plan/notes evidence honestly.
- Known risk: required outcome and ready-transition behavior is insufficiently
  regression-protected despite broad checks passing.
- Artifact created: `.concoct/current/review-01.md`.
- Expected next role: Developer in remediation mode.
- Recommended next command: `concoct code`.
