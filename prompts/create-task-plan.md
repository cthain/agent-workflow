# Prompt: Create a task plan

Use this prompt when you have an idea and want an assistant to turn it into agent-ready planning artifacts.

```text
I want to turn the following idea into implementation-ready planning artifacts.

Please produce the contents for:

- `.planning/current/task-plan.md`
- `.planning/current/notes.md`

Use the lightweight agent planning workflow.

Follow these conventions:

- hyphens, not underscores, in file and directory names
- uppercase Markdown filenames for long-lived project artifacts
- lowercase hyphenated Markdown filenames for task artifacts
- keep the plan useful, not ceremonial
- include clear goals, context, constraints, non-goals, phases, and completion criteria
- include initial notes with decisions, assumptions, risks, and follow-up ideas
- keep the files agent-neutral so they can be used by Codex, Claude Code, Copilot, Aider, or another coding agent

Here is the idea:

[PASTE IDEA HERE]

Project context:

[PASTE CONTEXT OR SAY "use what you know"]

Preferred tech stack:

[PASTE OR SAY “recommend one”]

First implementation target:

[PASTE OR SAY “help me define it”]
```

## Optional addition

If the task reveals a durable project rule, add:

```text
Also suggest any changes that should be made to `AGENTS.md`, but keep them separate from the task-specific files.
```
