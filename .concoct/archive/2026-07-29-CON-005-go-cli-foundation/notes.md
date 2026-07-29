# Notes

## Planning summary

- Readiness result: `ready`.
- CON-005 exists in `.concoct/roadmap.md` with status `planned`, priority `critical`, and a coherent initial scope of `init`, `status`, and shared CLI infrastructure.
- Its sole dependency, CON-003, is delivered, approved in `review-02.md`, archived, and reflected in CAP-001.
- `.concoct/current/` contained only `.gitkeep` before planning; there was no conflicting active task or review.
- Repository reality supports the roadmap assumptions, and no unresolved product decision requires Product Owner action before implementation.

## Confirmed findings

- There is no `go.mod`, Go source, or automated test suite.
- The only executable implementation is `cmd/concoct/concoct`, an executable Bash script with the legacy form `./concoct <project> [destination-parent]`.
- The script copies dotfiles, creates current/archive directories, writes a bootstrap prompt, initializes Git, runs `git add .`, and creates no commit.
- From its checked-in location the script looks for `cmd/concoct/templates` and `cmd/concoct/personas`, which do not exist. Repository templates live at root `templates/`, and shipped personas already live in `templates/.concoct/personas/`.
- README currently documents a root-level `./concoct` command that is absent from the repository inventory, while the implementation is under `cmd/concoct/`.
- `doc/command-reference.md` explicitly reserves staging as a CLI-foundation decision. Current implementation, README, and legacy archive evidence consistently establish staging-without-commit as existing bootstrap behavior.
- `doc/state-machine.md` requires artifact-backed detection for ready, active implementation, all review outcomes, remediation, blocked-review recovery, invalid evidence, and partial archival consistency. Empty tracked placeholders are not active artifacts.
- Templates contain conventional root adapters and `.concoct/` content, including dotfiles and nested prompt/persona files. Three legacy writer persona files are empty and some stale references are known capability gaps; broad cleanup is not required for CON-005 unless it blocks correct bootstrap guidance or validation.

## Decisions

- Preserve staging of generated files and the no-commit behavior. This is treated as compatibility grounded in existing behavior, not invented product direction.
- Propose `CAP-005` for the new executable CLI capability. On accepted archival, also reconcile CAP-003's initializer limitation and evidence.
- Require `status` to implement the complete CON-003 detection contract now, including remediation and blocked-review evidence, even though commands that produce those artifacts are later roadmap work.
- Leave the exact Go package layout and standard-library-versus-small-YAML-library choice to the Developer, constrained by explicitness, tested schema support, and avoiding a framework-heavy CLI.
- Do not change `.concoct/roadmap.md` during planning; CON-005 is already `planned`, and the prompt did not authorize a status transition.

## Assumptions

- A template distribution mechanism can be implemented and tested without requiring a Product Owner decision; it is a technical design choice so long as installed use is reliable and the generated contract is unchanged.
- Human-readable status formatting may vary from the roadmap example while retaining every applicable field, deterministic state names, diagnostics, and next actions.
- The retained or replacement executable entry point can be documented without preserving the shell script's optional second positional destination argument, because the normative CON-003 contract requires exactly one project target name or path.

## Risks and open questions

- Installed template discovery is the largest architectural risk: a solution that works only in the source checkout would repeat the current failure.
- Metadata parsing must be strict enough for safe state detection without becoming an accidental general-purpose YAML or workflow framework.
- Initialization failure cleanup must distinguish a target created by the current invocation from pre-existing or ambiguous data.
- A compatibility wrapper can drift if it retains business logic; if kept, it should only locate and invoke the Go CLI.
- Product question threshold: return to the Product Owner only if implementation would materially change bootstrap contents, staging/no-commit semantics, the normative status fields/states, or command scope. No such question is currently known.

## Relevant history

- `.concoct/archive/2026-07-29-CON-003-command-contract-state-machine/summary.md` records the approved normative command and state-machine contract and calls out downstream implementation of remediation and blocked-review metadata.
- `.concoct/archive/2026-07-29-legacy-hitl-restructuring/summary.md` records the initializer path failure and lack of tests.
- The legacy archive notes and plan record that staging was intentionally retained while avoiding an automatic commit.

## Developer handoff

- Current state: `planned` for CON-005.
- Work completed: readiness validation, repository inspection, scope definition, acceptance criteria, verification plan, and capability-impact proposal.
- Work remaining: implement and verify the Go CLI foundation; update these artifacts with durable implementation evidence and a reviewer handoff.
- Review focus after implementation: installed template resolution, full state-machine fidelity, malformed/contradictory evidence, read-only status behavior, safe initialization failures, staged/no-commit Git behavior, and wrapper/migration correctness.
- Expected next role: Developer.
- Recommended next command: `concoct code`.

## Implementation findings and decisions

### Embedded template distribution

- The Go module embeds `all:templates`, so installed binaries carry root files,
  nested content, and dotfiles without runtime repository-path assumptions.
- Initialization substitutes the generated project directory name for the
  template's `<project-name>` placeholders, writes bootstrap guidance, runs Git
  initialization and staging, and deliberately creates no commit.
- Creation failures preserve any target created by the invocation and report
  its exact path plus a safe inspection/removal action. Existing targets and
  missing destination parents are never modified.

### CLI and compatibility boundary

- `cmd/concoct/main.go` is the installable Go entry point with explicit `init`,
  `status`, and help dispatch and non-zero operational-error exits.
- `cmd/concoct/concoct` is retained as an executable source-checkout wrapper.
  It delegates to the Go CLI and passes the caller directory explicitly so
  relative initialization and nested-directory status retain binary behavior.
- The implementation uses only the standard library plus `gopkg.in/yaml.v3`
  for strict decoding of state-bearing front matter.

### State detection

- Project discovery walks parents and requires the canonical `AGENTS.md`,
  roadmap, and capabilities contract.
- Detection validates project identity, active roadmap/task identity, task
  statuses, capability impact, contiguous zero-padded reviews, internal review
  numbering, task identity, exactly one documented outcome, remediation, and
  blocked-review-resolution evidence before selecting a state.
- Placeholder notes require special handling because notes have no YAML front
  matter. Detection recognizes the shipped placeholder markers; other non-empty
  notes are treated as populated durable context.
- Representative interrupted archival evidence is rejected when a current task
  points to a roadmap item already marked delivered.

### Owned-artifact inconsistency

- The active roadmap item remains `planned`, as established during planning.
  The normative contract requires it to be `active` while a populated current
  task exists, so `concoct status` correctly reports this repository as
  `invalid` with an actionable diagnostic.
- The Developer cannot update `.concoct/roadmap.md`. This does not affect the
  generated-project or fixture verification, but the next Reviewer should route
  this existing workflow-artifact correction to the owning role rather than
  requesting that state validation accept contradictory evidence.

## Test results

- `go test ./...` — passed, including table-driven normative-state,
  remediation, blocked-resolution, invalid-evidence, read-only detection,
  discovery, template-copy, and real-Git initialization coverage.
- `go vet ./...` — passed.
- `go build -o <temporary-path>/concoct ./cmd/concoct` — passed on Linux.
- Built CLI help, argument-error behavior, external-directory `init`, root and
  nested `status`, and the retained shell wrapper — passed.
- Generated-project checks confirmed root files, dotfiles, nested prompts,
  personas, current/archive directories, bootstrap guidance, staged files, and
  no Git commit.
- File metadata snapshots before and after nested `status` were identical.
- `bash -n cmd/concoct/concoct` and executable-mode check — passed.
- `git diff --check` — passed after final code, documentation, and artifact updates.
- Relevant stale-path/persona search found stale references only in unchanged
  legacy workflow documentation and the planning history recorded above; the
  directly affected README and wrapper use current paths and roles.

## Handoff to reviewer

### Implemented

- Added the Go CLI foundation with embedded template distribution, `init`,
  project discovery, strict artifact parsing, full status state selection, and
  structured diagnostics.
- Added unit and integration tests covering the state matrix, invalid evidence,
  initialization, Git semantics, discovery, and non-mutation.
- Replaced the broken shell initializer logic with a thin, tested delegation
  wrapper and updated README invocation, distribution, status, and development
  guidance.

### Key decisions

- Embed the exact checked-in template tree in the binary rather than resolve
  templates beside the executable.
- Preserve staged generated files and no automatic commit.
- Report invalid durable evidence instead of guessing, including the current
  repository's planner-owned roadmap/task mismatch.

### Files changed

- `go.mod`, `go.sum`, `templates.go`
- `cmd/concoct/main.go`, `cmd/concoct/concoct`
- `internal/cli/cli.go`
- `internal/project/project.go`, `internal/project/project_test.go`
- `internal/workflow/workflow.go`, `internal/workflow/workflow_test.go`
- `README.md`
- `.concoct/current/task-plan.md`, `.concoct/current/notes.md`

### Checks

- All required Go tests, vetting, build, CLI, external init, nested status,
  read-only, Git, wrapper, template-layout, and diff checks passed as recorded
  above.

### Risks and skipped work

- The Markdown roadmap item parser intentionally implements the checked-in
  schema rather than a general Markdown data model; strict YAML is used for all
  front matter.
- Completed remediation disposition validation is deterministic but textual: it
  requires at least one recognized disposition term per review finding because
  the normative notes schema does not define structured finding IDs.
- No out-of-scope workflow commands, agent launching, archive mutation, template
  cleanup, or capability updates were implemented.
- The roadmap/task status mismatch described above is unresolved because the
  Developer does not own the roadmap.

### Capability impact

- Remains `add` for proposed CAP-005. After approval, archival should also
  reconcile CAP-003's initializer limitation. Capability truth was not changed
  during implementation.

### Suggested review focus

- State precedence and invalid diagnostics, especially remediation and blocked
  review recovery.
- Initialization failure safety, embedded-template completeness, staged/no-
  commit behavior, and installed-binary independence from the checkout.
- Strict read-only behavior of status and the wrapper's caller-directory
  preservation.
- Correct routing of the planner-owned roadmap status inconsistency.

Expected next role: Reviewer.

Recommended next command: `concoct review`.

## Review 01 remediation

### Finding dispositions

- Finding 1 — later reviews do not supersede historical recovery metadata:
  **fixed**. Detection now ignores retained `remediates-review` metadata when
  it identifies an earlier `changes-requested` review, and likewise ignores a
  retained `blocked-review-resolution` when it identifies an earlier `blocked`
  review. Arbitrary, missing, wrong-outcome, and current-review mismatches
  remain invalid. Regression tests cover an approved second review after both
  recovery paths.
- Finding 2 — required state-bearing metadata is not fully validated:
  **fixed**. Roadmap items now reject statuses outside `candidate`, `planned`,
  `active`, `blocked`, `delivered`, `deferred`, and `cancelled`. Task plans now
  require `title`, `created`, `updated`, and `capability-impact.rationale` in
  addition to the previously validated fields. Reviews now require `created`
  and `persona`, with `persona` constrained to `reviewer`. Focused invalid-state
  tests cover every newly enforced field and unknown roadmap status.

### Verification after remediation

- `go test ./...` — passed.
- `go vet ./...` — passed.
- `go build -o <temporary-path>/concoct ./cmd/concoct` — passed.
- Built CLI help, external-directory initialization, generated template and
  persona presence, staged Git state, no generated commit, and nested-directory
  status — passed in temporary directories outside this repository.
- `bash -n cmd/concoct/concoct` and executable-mode check — passed.
- `git diff --check` — passed.

## Handoff to reviewer after Review 01

### Implemented

- Corrected later-review precedence for retained remediation and blocked-review
  recovery metadata.
- Added strict required-field validation for state-bearing roadmap, task, and
  review metadata.
- Added regression coverage for both Review 01 findings.

### Key decisions

- Historical recovery metadata is ignored only when it resolves to a real
  earlier review with the outcome required by that metadata. This preserves
  the normative supersession behavior without allowing malformed references to
  bypass validation.
- Roadmap status validation uses the canonical vocabulary documented by the
  Product Owner persona and already used by the checked-in roadmap.

### Files changed during remediation

- `internal/workflow/workflow.go`
- `internal/workflow/workflow_test.go`
- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`

### Risks and unresolved work

- No Review 01 findings remain unresolved.
- The existing planner-owned roadmap/task status mismatch remains outside the
  Developer's artifact ownership, as Review 01 explicitly noted.
- Capability impact remains `add` for proposed CAP-005; capabilities and
  roadmap were not changed during remediation.

### Suggested review focus

- Confirm that a later review supersedes only genuine earlier recovery
  metadata and that malformed references remain invalid.
- Confirm diagnostics name the affected file and missing or invalid field.
- Re-run the full state and initialization test suite.

Expected next role: Reviewer.

Recommended next command: `concoct review`.

## Review 02 remediation

### Finding dispositions

- Finding 1 — historical remediation bypasses the latest blocked-review
  resolution: **fixed**. A valid retained `remediates-review` for an earlier
  changes-requested review now falls through instead of jumping directly to
  the latest review outcome. Detection therefore evaluates a
  `blocked-review-resolution` that names the later blocked review. Invalid
  historical references still produce actionable diagnostics.
- Review 01 Finding 1 is now **fixed** rather than partially fixed: recovery
  metadata composes across successive changes-requested and blocked review
  outcomes.

### Verification after remediation

- `go test ./...` — passed, including the new composed-recovery regression.
- `go vet ./...` — passed.
- `go build -o <temporary-path>/concoct ./cmd/concoct` — passed.
- Built CLI help, external-directory initialization, generated root and nested
  template content, dotfiles, personas, bootstrap guidance, staged Git state,
  no generated commit, and nested-directory read-only status — passed in an
  isolated temporary directory outside this repository.
- `bash -n cmd/concoct/concoct` and executable-mode check — passed.
- `git diff --check` — passed before this artifact update.
- An initial manual-verification command was rejected before execution because
  its automatic cleanup used a prohibited destructive operation. The check was
  rerun successfully without automatic cleanup; no product or repository state
  was affected by the rejected attempt.

## Handoff to reviewer after Review 02

### Implemented

- Corrected recovery precedence so historical remediation does not suppress a
  valid blocked-review resolution for the latest review.
- Added a regression test for review-01 changes requested, retained
  remediation metadata, review-02 blocked, and a valid route-to-code
  resolution selecting `implementation-in-progress`.

### Key decisions

- Recovery fields targeting the latest review remain authoritative and return
  their implementation state immediately after validation.
- Genuine historical recovery fields are ignored individually, allowing other
  recovery evidence to be evaluated; references that do not identify an
  earlier review with the required outcome remain invalid.

### Files changed during remediation

- `internal/workflow/workflow.go`
- `internal/workflow/workflow_test.go`
- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`

### Risks and unresolved work

- No Review 02 findings remain unresolved.
- The existing planner-owned roadmap/task status mismatch remains outside the
  Developer's artifact ownership.
- Capability impact remains `add` for proposed CAP-005; capabilities and
  roadmap were not changed during remediation.

### Suggested review focus

- Confirm that composed recovery metadata selects the resolution applicable to
  the latest review without weakening validation of historical references.
- Re-run the state tests and inspect the focused regression sequence.

Expected next role: Reviewer.

Recommended next command: `concoct review`.
