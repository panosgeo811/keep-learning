# Vibe Learn Plugin (`vibe-learn`)

Turn vibe coding into guided engineering mastery. Designed for developers working on unfamiliar protocols, unfamiliar codebases, or tricky concurrency primitives.

## What's Inside the Bundle

1. **Always-On Rule (`rules/AGENTS.md`)**: Enforces "Glass-Box Coding" with concise **Learning Anchors** and **"Delete-This-Line" failure modes** whenever concurrency or protocol logic is written.
2. **Skill (`skills/vibe-learn/SKILL.md`)**: Implements the 4-Phase Learning Workflow:
   - **Phase 1**: Pre-Flight Concept Brief + Interactive Generative UI Protocol/Concurrency Sandbox
   - **Phase 2**: Glass-Box Vibe Coding with Learning Anchors
   - **Phase 3**: Background `shadow-tutor` Subagent spawning
   - **Phase 4**: PR Defense Check (interactive mock code review quiz)
3. **Subagent (`agents/shadow-tutor.md`)**: Background technical tutor that inspects your `git diff` and builds a line-linked `LEARNING_GUIDE.md` artifact while you continue coding.
