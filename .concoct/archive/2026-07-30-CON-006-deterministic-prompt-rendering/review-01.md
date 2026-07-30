---
task-id: CON-006
review: 1
status: changes-requested
created: 2026-07-30
persona: reviewer
---

# Review 01

## Outcome

`changes-requested`

## Summary

The implementation provides the four prompt-only commands, preserves the
existing workflow detector as the state authority, selects the reviewed modes,
renders stable ordered context, and uses create-only file output. Independent
project checks pass. Approval is withheld because the task's explicit golden
test requirement has not been implemented, leaving the complete rendered
prompt contract substantially less protected than the accepted plan requires.

## Acceptance criteria assessment

- The four commands, role selection, state validation, stdout output, and
  create-only `--output` behavior are implemented.
- Initial, continuation, remediation, post-remediation, and blocked-review
  recovery modes have focused test cases.
- Repeated rendering and stdout/file equality have automated coverage, and
  workflow artifacts are checked for non-mutation in the CLI success path.
- Existing `init`, `status`, wrapper behavior, Linux compilation, and project
  checks remain compatible in the verification performed.
- The golden-test criterion is not met. There are no golden fixtures or
  full-output comparisons for any role or mode.

## Findings

### Finding 1 — Required golden prompt coverage is absent

- Severity: `major`
- Status: `unresolved`
- Evidence: `.concoct/current/task-plan.md` Phase 4 requires golden tests for
  every supported command and materially distinct development and review
  modes, and the acceptance criteria repeat that requirement. The only new
  renderer test, `internal/prompt/render_test.go`, checks selected substrings,
  section-heading presence, and equality between two renders from the same
  implementation. The repository contains no prompt golden fixtures and no
  expected full-output comparison. The developer notes nevertheless report
  golden verification as complete.
- Impact: Deletion, reordering, duplication, or unintended wording changes in
  most of a rendered prompt can pass the current suite as long as the handful
  of asserted markers remain. This weakens verification of the deterministic,
  inspectable prompt bytes that are the task's principal observable output and
  leaves an explicit acceptance criterion unsatisfied.
- Required outcome: Add stable expected-output golden coverage for all four
  roles and the materially distinct code/review modes identified by the plan,
  then rerun the documented checks. Update the task status and notes honestly
  if the resulting fixtures expose additional discrepancies.

## Prior finding disposition

No prior review exists for this task.

## Verification performed

- `go test ./...` — passed.
- `go vet ./...` — passed.
- `git diff --check` — passed.
- `bash -n cmd/concoct/concoct.sh` — passed.
- `test -x cmd/concoct/concoct.sh` — passed.
- Inspected the complete tracked diff, all untracked Go source and tests,
  active task artifacts, manual prompt assets, and the normative command/state
  documentation.
- Searched source, tests, documentation, and active artifacts for golden or
  `testdata` coverage; none exists for rendered prompts.

## Capability impact assessment

The proposed `CAP-006` addition remains the correct capability impact if the
task is subsequently approved. CAP-002 and CAP-005 should retain their current
truth until archival reconciliation. No capability update is appropriate in
the current changes-requested state.

## Scope assessment

The reviewed implementation remains within CON-006's prompt-only boundary. It
does not launch agents, persist reviews, implement role-completion mutations,
archive work, or add Git task lifecycle behavior. The pre-existing untracked
`cmd/concoct/concoct` binary was not treated as implementation work.

## Documentation assessment

The README accurately describes the four prompt-rendering commands, default
stdout behavior, create-only output policy, and non-mutating boundary. No
material user-documentation defect was found.

## Risks and follow-up

The conservative archive-summary selection and append-verbatim prompt asset
strategy are documented design choices. They should be captured in the golden
cases so their observable ordering and content remain intentional.

## Handoff

- Current state: `changes-requested` after review 01.
- Required next role: Developer.
- Work remaining: resolve Finding 1, record the disposition and fresh
  verification in notes, return the task to `implementation-complete`, and add
  a new reviewer handoff without editing this review.
- Suggested next review focus: full rendered-byte golden coverage across all
  required roles and modes, plus continued determinism and non-mutation.
- Recommended next command: `concoct code`.
