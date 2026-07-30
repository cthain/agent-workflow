# Notes

## Planning summary

- Readiness: `ready`.
- `CON-028` exists with status `planned`, no unresolved roadmap dependencies,
  and a coherent advisory command outcome.
- The deterministic task branch is checked out at the prompt-recorded base; the
  worktree was clean, no Git operation was active, and no current task artifacts
  conflicted with planning.

## Confirmed findings

- `internal/workflow.Detect` currently reports `concoct roadmap or concoct plan
  <roadmap-id>` for ready state and is the durable state authority.
- `internal/prompt` already centralizes role selection, stable input ordering,
  relevant archive discovery, canonical handoff inclusion, and deterministic
  prompt rendering.
- `internal/cli.runPrompt` already owns shared stdout and create-only file output
  behavior; the prompt command dispatch does not include `next`.
- Initialization output and bootstrap guidance explicitly recommend `concoct
  roadmap`; integration prints the workflow report after returning to ready.
- Ready-state guidance is repeated in `README.md`, `doc/command-reference.md`,
  `doc/state-machine.md`, the Concoct skill, personas, prompt assets, tests, and
  embedded templates.
- The accepted implementation history consistently uses shared workflow parsing
  as authority, appends canonical prompt assets, protects full output with golden
  tests, and treats malformed durable evidence as invalid.
- The current accepted task-origin surface contains roadmap planning and human
  product input/roadmap maintenance. `CON-019` and `CON-027` describe future
  non-roadmap origins and bug work, but both remain candidates and unavailable.

## Capability prerequisite compatibility

- `CAP-001`: compatible. Its limitation reserves semantic correctness to role
  judgment; `next` explicitly delegates the recommendation to the Product Owner
  and only validates/assembles durable evidence.
- `CAP-005`: compatible. Its status command remains read-only and structurally
  authoritative; changing the ready recommendation does not merge status with
  Product Owner judgment.
- `CAP-006`: compatible. `next` extends deterministic prompt rendering while
  preserving create-only output, conservative archive relevance, and the rule
  that rendering does not complete role work or mutate state.
- `CAP-008`: compatible. Existing planning eligibility remains the authority for
  whether a roadmap item may be presented as plannable; limitation compatibility
  and semantic selection remain Product Owner/Task Planner judgments.

## Decisions

- Plan one capability addition, provisionally `CAP-009`, rather than treating
  this as a non-observable internal change.
- Reuse and expose existing planning-eligibility rules for next-action evidence;
  do not create a separate eligibility definition.
- Treat supported Product Owner roadmap/product input as the only present
  alternative to planning. Future task-origin types must remain absent until
  their accepted contracts exist.
- Require explicit prompt outcomes rather than storing a recommendation artifact;
  `next` is advisory and leaves ready state unchanged.
- Cover source/template guidance as part of the product contract, not as optional
  documentation cleanup.

## Risks

- The broad set of ready-state strings creates drift risk across executable
  output, documentation, skills, personas, prompts, tests, and templates.
- Deterministic ordering may be mistaken for automatic prioritization unless the
  prompt and tests clearly separate structural evidence from Product Owner
  judgment.
- Independent eligibility parsing would diverge from `plan`; shared workflow
  helpers and agreement tests are important review targets.

## Relevant history

- CON-003 established durable artifacts as state authority and separated prompt
  rendering from completed role work.
- CON-005 delivered initialization and read-only status with invalid-state
  diagnostics.
- CON-006 delivered deterministic role-aware prompt rendering, create-only
  output, canonical handoff reuse, and full-output golden tests.
- CON-007/CAP-008 delivered accepted prerequisite validation and retained semantic
  limitation compatibility as planner judgment.
- CON-015/CAP-007 delivered the Git branch and integration lifecycle whose ready
  outputs must change without weakening its safety boundaries.
- The legacy archive records older initialization and template limitations; the
  accepted later archives supersede its broken-initializer finding.

## Handoff

- Current state: `planned` after the roadmap status and paired active artifacts
  validate.
- Work completed: repository and archive inspection, prerequisite compatibility
  judgment, implementation-ready plan, risk analysis, acceptance criteria, and
  verification design.
- Work remaining: implement `CON-028`, verify all outcome classes and portable
  guidance, update these artifacts with results, and prepare independent review.
- Unresolved questions: none requiring Product Owner action; local evidence and
  prompt representation are Developer decisions within the plan constraints.
- Expected next role: Developer.
- Recommended next command: `concoct code`.

## Implementation findings and decisions

- Added `workflow.InspectNextActionEvidence` as the read-only evidence boundary.
  It requires valid `ready` state, presents roadmap and capability records in
  stable identifier order, and calls `InspectPlanEligibility` for every item so
  `next` and `plan` cannot drift on structural plannability.
- Evidence includes roadmap title, status, priority, dependencies, capability
  prerequisites, eligibility or exact blocker, archive provenance, accepted
  capability status/limitations/provenance, and only the two currently
  supported work origins.
- Added `next` to the shared prompt renderer and CLI prompt-output path. The
  command authorizes no writes, uses create-only output handling, and never
  enters the Git task-branch creation path.
- Added a canonical next-action handoff requiring exactly one outcome: plan an
  eligible item, perform supported product/roadmap work, resolve a named
  blocker, or report no actionable recorded work. Deterministic ordering is
  explicitly presentation rather than automated selection.
- Changed status, initialization, bootstrap, archival handoffs, normative docs,
  skills, personas, and embedded templates to make `concoct next` the sole
  ready-state recommendation while retaining `roadmap` and `plan` as explicit
  Product Owner follow-ups.

## Verification results

- `go test -count=1 ./...`: passed.
- `go vet ./...`: passed.
- `go build ./cmd/concoct`: passed.
- Generated-project exercise with a freshly built binary: passed initialization,
  staged complete template/dotfile copy, Git creation with no generated commit,
  bootstrap/status `concoct next` guidance, root/nested deterministic `status`
  and `next` output, and workflow-artifact non-mutation hashing.
- Source/template parity checks for the Concoct skill, Product Owner and
  Archivist personas, next-action prompt, and archival handoffs: passed.
- `git diff --check`: passed.
- Stale ready-state roadmap-or-plan and obsolete prompt-command searches: no
  stale matches in task-relevant source, templates, docs, skills, or personas.
- `cmd/concoct/concoct` executable-bit check: passed.

### Attempt: shell syntax check for compatibility command

- Tried: `bash -n cmd/concoct/concoct` as required by legacy repository guidance.
- Error/result: Bash reported `cannot execute binary file`.
- Why it failed: the checked-in compatibility command is currently an ELF Go
  executable, not a shell script; this task did not change that file.
- Next approach: verified its executable bit and verified the current Go command
  through build, tests, and generated-project execution. The stale shell-specific
  repository instruction remains a follow-up outside CON-028.

## Handoff to reviewer

### Implemented

Implemented `concoct next` as the sole valid ready-state recommendation command,
with deterministic authoritative evidence, Product Owner guidance, CLI/output
safety, portable templates, and updated ready transitions.

### Key decisions

- Reused `InspectPlanEligibility` rather than duplicating planning rules.
- Kept recommendation semantic and prompt-driven; the CLI exposes evidence but
  never ranks, selects, activates, or persists work.
- Modeled supported origins as an explicit bounded list so future accepted
  origins can extend the evidence boundary without changing its advisory role.

### Files changed

- Workflow, prompt, CLI, initialization, and focused Go tests/golden output under
  `internal/`.
- `README.md`, command/state-machine documentation, source skill/personas and
  handoffs, plus matching embedded template assets.
- Active task plan and notes only; roadmap, capabilities, reviews, and archives
  remain untouched.

### Verification

All Go tests, vet, build, generated-project execution, nested discovery,
determinism, non-mutation, parity, executable-bit, stale-string, and diff checks
passed. The inapplicable legacy `bash -n` check is documented above.

### Known risks

- Capability limitation text is intentionally preserved as authoritative
  evidence and can make prompts verbose; deterministic full-output coverage
  protects the format.
- The compatibility executable versus shell-check instruction mismatch predates
  this task and should be reconciled separately.

### Skipped or unresolved work

None within CON-028. No unsupported task origins, direct execution, lifecycle
mutation, or automated prioritization were added.

### Capability impact

Still `add` for provisional `CAP-009`: deterministic, read-only Product Owner
next-action recommendation. Accepted capability truth was not changed.

### Suggested review focus

Verify agreement between `next` and `plan` eligibility, invalid-evidence
handling, no-work/blocker outcome guidance, Git non-mutation, ready-state output
consistency, and source/template parity.

Current state: `implementation-complete`. Recommended next command: `concoct review`.

## Review 01 remediation

### Finding 1 — Required outcome and transition coverage is incomplete

- Disposition: `fixed`.
- Added full-output golden coverage for eligible roadmap work, supported product
  input, roadmap reconciliation, a specific dependency blocker, and no
  actionable recorded work.
- Each added `next` fixture verifies repeated byte-identical rendering and
  workflow-artifact non-mutation.
- Added invalid canonical roadmap evidence coverage that verifies `next` refuses
  the repository without mutation.
- Added assertions that initialization output and bootstrap guidance recommend
  exactly `concoct next`, and that the successful integration ready report has
  the same sole recommendation.

## Remediation verification

- `go test -count=1 ./...`: passed.
- `go vet ./...`: passed.
- `go build ./cmd/concoct`: passed.
- `git diff --check`: passed.
- `cmd/concoct/concoct` executable-bit check: passed.
- The pre-existing ELF executable versus `bash -n` instruction mismatch remains
  unchanged and outside CON-028, as documented in the original handoff.

## Post-remediation handoff to reviewer

- Current state: `implementation-complete`; Review 01 Finding 1 is fixed.
- Transition evidence: the task plan now records
  `remediates-review: review-01.md`, allowing the workflow detector to
  distinguish this completed remediation from the implementation originally
  reviewed. This marker was initially omitted and corrected before Review 02.
- Work completed: added the missing golden, invalid-evidence, initialization,
  bootstrap, integration, determinism, and non-mutation regression coverage.
- Decisions: kept remediation test-only because inspection and existing checks
  confirmed the product implementation already met the requested behavior.
- Files changed: focused tests and golden fixtures under `internal/prompt/`,
  `internal/project/`, and `internal/integration/`, plus Developer-owned notes.
- Known risks: none newly introduced; the documented compatibility executable
  mismatch remains an unrelated repository follow-up.
- Capability impact: unchanged `add` for provisional `CAP-009`; capability truth
  remains untouched.
- Suggested review focus: confirm every Review 01 scenario is now asserted and
  that the golden fixtures distinguish eligibility, reconciliation, blocker,
  supported-origin, and no-work evidence.
- Expected next role: Reviewer.
- Recommended next command: `concoct review`.

## Archival handoff

- Current state: `review-approved`; archival evidence is being committed on the
  recorded Git task branch.
- Accepted outcome: CON-028 adds a deterministic, read-only Product Owner
  recommendation command and makes `concoct next` the sole ready-state entry
  point without selecting work or mutating workflow evidence.
- Approving review: `review-02.md` (`approved`).
- Capability reconciliation: add CAP-009; CAP-001, CAP-005, CAP-006, and CAP-008
  remain related prerequisite capabilities and retain their existing truth.
- Verification: the approved review records passing full Go tests, vet, build,
  formatting, source/template parity, stale-guidance searches, executable-mode
  validation, and diff checks. The legacy ELF-versus-shell-check mismatch is a
  separate follow-up and does not affect acceptance.
- Archive destination:
  `.concoct/archive/2026-07-30-CON-028-recommend-next-project-action/`.
- Delivery remains pending integration; current task evidence must remain intact
  and CON-028 must remain active until `concoct integrate` succeeds.
- Expected next transition: `concoct integrate`.
