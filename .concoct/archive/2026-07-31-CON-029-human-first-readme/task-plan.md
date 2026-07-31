---
id: CON-029
title: Introduce Concoct through a human-first README
roadmap-id: CON-029
status: implementation-complete
created: 2026-07-30
updated: 2026-07-30
capability-impact:
  type: none
  ids: []
  rationale: Reframes existing accepted behavior for human onboarding without changing runtime behavior or capability truth.
git:
  enabled: true
  trunk: main
  task-branch: concoct/con-029-introduce-concoct-through-a-human-first-readme
  base: cd20a422961d1848ee51d521f3e29a72192a8dc4
  status: active
---

# Task Plan

## Goal

Rewrite the root `README.md` around a new user's journey so a developer can
quickly understand what Concoct is, decide whether it fits, and follow one
representative workflow from project initialization through accepted Git
integration without first learning the repository's internal layout.

## Context

The README currently opens with a short product description, then moves
immediately into installed files and repository structure. Executable usage is
split between a bootstrap section and a command list, while the distinction
between CLI effects and agent-performed role work is explained only after the
commands. This is accurate in pieces but asks a new user to understand Concoct's
internals before seeing the normal workflow and its value.

## Why this matters

Concoct's primary value is a repeatable, inspectable collaboration loop between
a developer and repository-aware coding agents. The README is the first product
surface most users encounter; it should make that value and the present
automation boundary legible before serving contributors or explaining layout.

## Current state

- The opening names Concoct and durable context, but does not clearly identify
  the intended developer or connect the workflow's roles and retained history
  to practical user value.
- Installed workflow structure and repository layout precede the first usable
  quick start.
- Initialization instructions use source-build paths before showing the normal
  end-to-end journey.
- Role commands are listed together with lifecycle operations, requiring nearby
  prose to clarify that prompt rendering does not execute the selected role.
- Detailed workflow, command, state-machine, and multi-agent documentation
  already exists and can be linked instead of reproduced.

## Target state

- The README opens in product language: what Concoct does, why durable workflow
  evidence matters, who it serves, and the maturity boundary a user should
  expect.
- A concise quick start begins with `concoct init hello-world`, explains the
  generated repository's initial commit responsibility, and walks through
  `next`, roadmap intake, planning, development, review/remediation, archival,
  and Git integration.
- Each role-prompt step explicitly tells the user that the command renders
  validated guidance for a human or agent to perform; only documented CLI
  lifecycle mutations are described as automatic.
- Internal layout, source-checkout development, naming conventions, and other
  contributor details remain discoverable after the user journey.
- Existing detailed documents carry normative depth through focused links.

## Design constraints

- Keep the root README useful both to prospective users and repository
  contributors, with the normal user journey taking precedence.
- Describe only behavior supported by CAP-001, CAP-005, CAP-006, and CAP-007 and
  the current CLI help/state-machine contracts.
- Preserve the distinction between durable role work and prompt rendering:
  `next`, `roadmap`, `plan`, `code`, `review`, and `archive` render guidance and
  do not autonomously complete their roles.
- State the bounded CLI mutations accurately: `init` creates and stages a Git
  project without committing; Git-backed `plan` establishes the task branch;
  `integrate` performs guarded local squash integration and may require human
  conflict resolution.
- Describe agent neutrality as a shared repository contract usable by capable
  tools, not as identical native integration or behavior across tools.
- Keep claims concise and link to `doc/workflow.md`,
  `doc/command-reference.md`, `doc/state-machine.md`, and development guidance
  for detail.
- Preserve the current command spelling and supported flags. Do not invent an
  installation or release channel that the repository does not provide.
- Maintain the repository's existing Markdown style and relative-link
  conventions.

## Non-goals

- No CLI, workflow, template, persona, prompt, test, or runtime behavior changes.
- No redesign of the state machine, role boundaries, Git lifecycle, or command
  surface.
- No broad rewrite of detailed documentation or generated-project guidance.
- No claim that Concoct launches agents, validates semantic correctness, makes
  product decisions, or resolves merge conflicts autonomously.
- No universal tool-specific setup guide, hosted-service workflow, package
  publishing, or unsupported installation instructions.
- No new diagrams or branded visual assets unless a concrete readability need
  emerges during implementation.

## Working assumptions

- `README.md` is the only product file that needs content changes; the linked
  detailed documents remain authoritative enough for this onboarding task.
- A representative Git-backed happy path is the clearest default journey, with
  review remediation and conflict handling mentioned as branches rather than
  expanded into full tutorials.
- The existing source build and wrapper commands are the only installation-like
  instructions that can be stated without inventing distribution behavior.
- Capability impact remains `none` because the task changes presentation of
  accepted behavior, not observable product behavior.

## Risks and open questions

- A short quick start can accidentally imply that prompt commands execute role
  work. Review every transition for an explicit actor and effect.
- Presenting the complete loop can become verbose or duplicate normative docs.
  Keep the README at journey level and use links for edge cases and schemas.
- `concoct plan` has a bounded Git mutation even though it is also a prompt
  command; wording must preserve this exception.
- The README should be candid about current maturity without burying the value
  proposition or speculating about future delivery dates.
- No product decision remains unresolved. Heading order, prose length, and the
  exact quick-start presentation are Developer choices within these constraints.

## Implementation phases

### Phase 1 — Establish the human-first narrative

Status: `complete`

- Rework the opening to identify the intended developer, the problem Concoct
  solves, and the value of durable context, explicit roles and transitions,
  independent review, and retained history.
- Add a concise fit/maturity statement that sets accurate expectations about
  agent neutrality, prompt rendering, and human judgment.

### Phase 2 — Build the representative quick start

Status: `complete`

- Begin with `concoct init hello-world`, enter the generated project, explain
  review/commit of staged bootstrap files, and show state inspection.
- Walk through `next`, `roadmap`, `plan`, `code`, `review`, `archive`, and
  `integrate` in valid order, naming what the CLI does and what the developer or
  coding agent must do after each rendered prompt.
- Briefly cover changes-requested review loops and human-resolved integration
  conflicts without turning the quick start into a recovery manual.

### Phase 3 — Reorder supporting and contributor material

Status: `complete`

- Move installed workflow structure and repository layout after the product and
  user-journey sections.
- Preserve source-build, wrapper, development, naming, and manual repository
  rename information where still useful, consolidating duplicated setup prose.
- Add focused links to workflow, command-reference, state-machine, and
  multi-agent/development material rather than duplicating their contracts.

### Phase 4 — Verify accuracy and prepare review

Status: `complete`

- Compare every command and effect in the README with current CLI help,
  capability limitations, command reference, and state machine.
- Check heading flow, code-block continuity, relative links, terminology, and
  stale or contradictory onboarding language.
- Run repository-standard proportional checks, update active artifacts with
  results, and prepare an independent reviewer handoff.

## Acceptance criteria

- The opening tells a new user what Concoct does, why its durable workflow is
  useful, who it is for, and the current maturity boundary before describing
  files or repository internals.
- A user can follow a representative journey beginning with exactly `concoct
  init hello-world` through roadmap/product intake, planning, implementation,
  independent review, archival, and Git integration without first reading the
  repository layout.
- The journey explains the value of durable context, defined roles, explicit
  transitions, independent review, and retained archive history in practical
  terms.
- Every displayed executable command exists in current CLI help and appears in
  a valid state order consistent with the documented state machine.
- Prompt-rendering commands are never described as autonomously performing
  Product Owner, Task Planner, Developer, Reviewer, or Archivist work; the
  Git-backed planning branch mutation is called out accurately.
- `init`, `status`, prompt output behavior, archival, and integration claims
  agree with accepted capability truth, including staging without an initial
  commit, create-only prompt files where discussed, local squash integration,
  guarded human conflict resolution, and optional upstream behavior.
- Agent-neutral language promises a portable shared contract without claiming
  identical native features across all named tools.
- Workflow, command-reference, state-machine, multi-agent, and contributor
  information remains easy to find through working relative links and later
  sections.
- The README states limitations honestly, including that Concoct does not launch
  agents or establish semantic correctness on its own.
- The change remains documentation-only and does not alter runtime behavior or
  accepted capability records.

## Verification

- Run `go run ./cmd/concoct help` and compare the displayed command surface to
  every command used in the README.
- Manually trace the README quick start against `doc/state-machine.md` and the
  detailed contracts in `doc/command-reference.md`.
- Inspect CAP-001, CAP-005, CAP-006, and CAP-007 limitations against all product
  and automation claims.
- Check every README relative link resolves to an existing repository path.
- Search the README for language that implies prompt commands launch agents,
  complete role work, or guarantee semantic correctness.
- Run `go test ./...` and `go vet ./...` if implementation changes anything
  beyond Markdown; otherwise document why these unchanged-code checks were not
  needed.
- Run `git diff --check`.
- Inspect the final README diff for human-first section ordering, concision, and
  scope discipline.

## Capability impact

Expected impact is `none`. The README will make CAP-001, CAP-005, CAP-006, and
CAP-007 easier to understand and adopt, but it will not add, update, or remove
accepted product behavior. If implementation discovers that an accurate quick
start requires undocumented or unavailable behavior, stop and route that gap
to the Product Owner rather than changing capability truth in this task.

## Handoff expectations

The Developer should first set the task to `implementation-in-progress`, then
rewrite only the root README unless repository evidence proves a linked factual
correction is necessary. Before review, record the final narrative structure,
command-contract checks, links checked, files changed, skipped checks, known
risks, and capability-impact confirmation in `notes.md`; set status to
`implementation-complete`; and recommend `concoct review`.
