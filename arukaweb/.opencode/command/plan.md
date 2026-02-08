---
description: Starts a new feature from the roadmap
agent: planner
subtask: false
---

User has selected the $ARGUMENTS feature from the plan
Read the `product/plan.md` and read the item that the user selected, remember it as [feature] for future reference.
If use didn't select a feature, let the user know that and we finish.

Create a file in `product/features/` with the number of the [feature] and the name of the feature, e.g. `001-frontend-landing-page.md`, remember it as [feature-file]. If such file exists let the user know that and we finish. Otherwise we continue.

Read the `product/description.md` file, review existing code related to [feature] to get a better context of the project.

Inside the [feature-file] create two main sections:
`### Description` where you should detail the description of the [feature] so the developer can understand what he needs to do.
`### Tasks` where you should detail all the tasks needed to implement the [feature], the tasks should be as detailed as possible, assume the developer is an expert one so you don't need to explain many things.

IMPORTANT: If some aspect of the [feature] is not clear to you, ask the user for clarifications, do this as many times as needed.
When the [feature-file] is ready, let the user know that and we finish.
