---
task-id: CON-005
review: 2
status: changes-requested
created: 2026-07-29
persona: reviewer
---

# Review 02

## Outcome

`changes-requested`

## Summary

Review 01's required-field validation is complete, and its recovery-metadata
supersession fix works for the tested case where a later review is directly
authoritative. However, the fix does not compose the two recovery mechanisms:
a retained historical `remediates-review` prevents a valid resolution for the
latest blocked review from being evaluated. This can leave a correctly
resolved blocked-review workflow incorrectly reported as `review-blocked`.

The CLI otherwise builds and passes its automated checks. External
initialization, embedded template copying, Git staging without a commit, and
read-only nested-directory status also passed independent verification.

## Acceptance criteria assessment

- Go CLI, explicit dispatch, Linux build, help, and operational errors: met.
- Embedded initialization, template completeness, Git staging, and no-commit
  behavior: met.
- Project discovery and read-only status: met.
- Required metadata and roadmap-status validation from Review 01: met.
- Normative state and recovery precedence: not fully met; see Finding 1.
- Documentation, compatibility wrapper, scope, and capability declaration: met.
- Required verification: passed.

## Findings

### Finding 1 — Historical remediation bypasses the latest blocked-review resolution

- Severity: major
- Status: unresolved
- Evidence: `internal/workflow/workflow.go:222-226` jumps directly to
  `latestReview` whenever `remediates-review` names an earlier
  `changes-requested` review. That jump bypasses the
  `blocked-review-resolution` evaluation at lines 243-253. Therefore, after
  `review-01` requests changes, remediation completes, and `review-02` is
  blocked, a task plan that retains `remediates-review: review-01.md` and adds
  a valid `blocked-review-resolution` for `review-02.md` is reported from the
  latest review outcome instead of as `implementation-in-progress` or
  `implementation-complete`. `doc/state-machine.md` requires the stale
  remediation field to be ignored and valid resolution evidence naming the
  exact latest blocked review to select the implementation state. The new test
  at `internal/workflow/workflow_test.go:115-132` covers only a later approved
  review with one historical recovery field; it does not cover recovery fields
  composed across successive review outcomes.
- Impact: A valid repeated review cycle can remain stuck in `review-blocked`
  after the responsible role has persisted all required resolution evidence.
  The reported next action is wrong and the next review transition cannot be
  selected from durable state as the CON-003 contract requires.
- Required outcome: Treat a valid historical recovery reference as
  non-authoritative without skipping evaluation of any recovery mapping that
  applies to the latest review. Add a regression test with a historical
  `remediates-review` and a valid `blocked-review-resolution` naming a later
  blocked review, covering at least one supported route and its expected
  implementation state.

## Prior finding disposition

- Review 01 Finding 1 — partially fixed. Later approved reviews supersede a
  single retained recovery field, but historical remediation incorrectly
  suppresses a current blocked-review resolution as described above.
- Review 01 Finding 2 — fixed. Canonical roadmap statuses and the required task
  and review metadata are validated with focused diagnostics and tests.

## Verification performed

- `go test ./...` — passed.
- `go vet ./...` — passed.
- `go build -o <temporary-path>/concoct ./cmd/concoct` — passed.
- Built CLI help — passed.
- Built CLI initialization from an external temporary directory — passed.
- Generated root files, nested content, dotfiles, reviewer persona, and
  bootstrap prompt — present.
- Generated Git repository — files staged, no commit created.
- `concoct status` from a generated nested directory — reported `ready`; Git
  status was unchanged before and after.
- `bash -n cmd/concoct/concoct` and executable-mode check — passed.
- `git diff --check` — passed before creation of this review artifact.
- Inspected the complete implementation diff, relevant source and tests,
  README, command contract, state-machine contract, task plan, notes,
  capabilities, and Review 01.

## Capability impact assessment

The proposed `add` impact for CAP-005 remains accurate after the state defect
is fixed and the implementation is approved. Archival should also reconcile
CAP-003's initializer limitation and obsolete evidence. No capability ledger
change is appropriate before approval.

## Scope, documentation, and compatibility assessment

The changes remain within CON-005. The embedded template strategy supports an
installed binary independently of the checkout, and the shell entry point is a
thin executable wrapper. README guidance accurately covers build, init,
status, staging, and no-commit behavior. No material unrelated changes were
found.

## Risks and follow-up

- The current repository's roadmap item remains `planned` while its task is
  populated. This planner-owned mismatch correctly makes this checkout an
  invalid status target and remains outside Developer ownership.
- Recovery precedence should be tested as a sequence of accumulated durable
  fields, because task-plan cleanup is intentionally optional during later
  owned updates.

## Handoff

- Current state: changes requested for CON-005.
- Work required: make historical recovery metadata non-authoritative without
  bypassing valid recovery evidence for the latest review, and add the focused
  regression test.
- Expected next role: Developer in remediation mode.
- Recommended next command: `concoct code`.
