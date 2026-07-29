# Concoct Transition Prompts

These prompts coordinate the human and agent transitions in the Concoct workflow.

```text
human input
  → product-owner
  → task-planner
  → developer
  → reviewer
      ├─ changes requested → developer
      ├─ blocked → responsible role or human
      └─ approved → archivist
                      → product-owner / next task
```

They can be used manually now and later rendered by the `concoct` CLI.
