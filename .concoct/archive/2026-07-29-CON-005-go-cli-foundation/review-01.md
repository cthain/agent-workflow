---
task-id: CON-005
review: 1
status: changes-requested
created: 2026-07-29
persona: reviewer
---

# Review 01

## Outcome

`changes-requested`

## Summary

The Go CLI builds, its initialization path is well covered, and the implementation is appropriately scoped. However, status detection does not yet implement two required parts of the CON-003 contract: later reviews must supersede historical remediation metadata, and canonical state-bearing metadata must be validated rather than merely decoded. These defects can make valid repositories invalid and malformed repositories appear valid.

The current repository's `CON-005` roadmap item is still `planned` while the current task is populated. That planner-owned artifact mismatch correctly causes this implementation's `status` command to report `invalid`; it did not prevent independent code review and is not assigned to the Developer as a code defect.

## Acceptance criteria assessment

- Go CLI, explicit command dispatch, Linux build, help, and error behavior: met.
- Embedded template initialization independent of the caller directory: met.
- Complete template, dotfile, nested-content, Git staging, and no-commit behavior: met by tests and inspection.
- Target refusal and partial-target diagnostics: met for the tested and inspected paths.
- Root/nested discovery and read-only status behavior: met.
- Normative state and invalid-evidence detection: not met; see Findings 1 and 2.
- Applicable status fields and next actions: met for states the detector classifies correctly.
- Compatibility wrapper and documentation: met.
- Required verification: passed.

## Findings

### Finding 1 — Later reviews do not supersede historical recovery metadata

- Severity: major
- Status: unresolved
- Evidence: `internal/workflow/workflow.go:206-226` validates any non-empty `remediates-review` or `blocked-review-resolution` against the latest review before allowing the latest review outcome to determine state. If remediation for `review-01.md` is followed by `review-02.md`, the retained `remediates-review: review-01.md` is therefore rejected because it does not name the latest review. The same failure occurs for a retained blocked-review resolution after a later review. `doc/state-machine.md:75` explicitly says a later review causes a stale remediation reference to be ignored for detection, and `doc/state-machine.md:97` gives the same supersession rule for blocked-review resolution. No test covers either later-review scenario.
- Impact: A normal remediation or blocker-resolution cycle can produce a valid `review-approved`, `review-changes-requested`, or `review-blocked` repository that `concoct status` incorrectly reports as `invalid`. In particular, an approved second review can become impossible to archive unless a role performs an otherwise optional metadata cleanup.
- Required outcome: Apply recovery metadata only when it refers to the actual latest review; otherwise let the later valid review supersede it as the normative contract requires. Add tests for at least a second review after changes-requested remediation and a second review after blocked-review resolution, including an approved outcome.

### Finding 2 — Required state-bearing metadata is not fully validated

- Severity: major
- Status: unresolved
- Evidence: `internal/workflow/workflow.go:341-371` accepts any roadmap status string other than using `active` for active-count checks, so a roadmap item with `Status: mystery` can coexist with an otherwise empty current directory and report `ready`. The task and review decoders also define required schema fields but do not validate several of them: task `title`, `created`, `updated`, and capability-impact `rationale`, and review `created` and `persona`. Indeed, the review fixtures in `internal/workflow/workflow_test.go:109-110` omit `created` and `persona` and are accepted. The task plan's Phase 2 requires validation of known status values and required metadata, while the acceptance contract requires incomplete active artifacts and malformed metadata to produce actionable `invalid` diagnostics.
- Impact: Partially populated or schema-invalid canonical artifacts can be presented as valid workflow states. That weakens status as a safe transition gate and contradicts the task's central artifact-validation requirement.
- Required outcome: Validate the canonical roadmap status vocabulary and all required task/review state metadata, with file-and-field-specific diagnostics. Add focused tests showing unknown roadmap statuses and missing required task/review metadata produce `invalid`.

## Prior finding disposition

No prior reviews exist for CON-005.

## Verification performed

- `go test ./...` — passed.
- `go vet ./...` — passed.
- `git diff --check` — passed.
- `bash -n cmd/concoct/concoct` — passed.
- Retained wrapper executable-mode check — passed.
- Inspected the complete diff, all new Go source and tests, README changes, the normative command/state documents, current plan and notes, capabilities, and repository instructions.

## Capability impact assessment

The declared `add` impact for proposed `CAP-005` is appropriate once the findings are resolved and the implementation is approved. Archival should also reconcile CAP-003's broken-initializer limitation and stale verification evidence. No capability ledger update is appropriate before approval.

## Scope, documentation, and compatibility assessment

The implementation remains within CON-005: it adds only `init`, `status`, their shared infrastructure, tests, a thin source-checkout wrapper, and directly affected README guidance. The embedded template strategy addresses installed-binary independence without introducing agent integration or later workflow commands. Documentation accurately describes the chosen invocation and staged/no-commit behavior. The retained wrapper remains executable and contains delegation only.

## Risks and follow-up

- The active roadmap/task mismatch is owned by the role authorized to update `.concoct/roadmap.md`; it should be reconciled before relying on this repository itself as a valid status fixture or before archival.
- The textual remediation-disposition heuristic is intentionally limited by the unstructured notes schema. It is not an additional blocking finding for this task, but future structured finding identifiers would make this validation more reliable.

## Handoff

- Current state: changes requested for CON-005.
- Work required: correct review supersession and complete required metadata/status validation; add regression tests for both findings.
- Expected next role: Developer in remediation mode.
- Recommended next command: `concoct code`.
