---
name: vibe-learn
description: >-
  Structured 4-phase learning-while-vibe-coding workflow for developers tackling
  tickets involving unfamiliar protocols, unfamiliar codebase architecture, or
  tricky concurrency/systems concepts. Activate when the user wants to learn
  along while solving a ticket, requests a pre-flight protocol/codebase brief,
  asks for an interactive concurrency or protocol visualizer, wants to spawn the
  background shadow-tutor subagent, wants a PR Defense Check, or wants to view
  or update the central Learning Hub in ~/projects/vibe-learn/learnings/.
---

# Vibe-Learn: Guided Mastery While Vibe Coding

Use this skill to help a developer ship their ticket quickly **and** deeply understand every protocol concept, architectural boundary, and concurrency primitive the LLM writes—while indexing every lesson in the central **Learning Hub** (`~/projects/vibe-learn/learnings/`).

---

## The 4-Phase Agentic Workflow (+ Central Learning Hub Archival)

You can execute all 4 phases end-to-end when tackling a ticket, or jump directly to the phase the user asks for.

### Phase 1: The Pre-Flight Concept Brief (Before Coding)
Before writing the fix, orient the developer in under 90 seconds of reading:
1. **60-Second Protocol / Domain Primer**: Explain *only* the slice of the protocol (e.g., WireGuard/Tailscale NAT traversal, STUN/DERP fallback, keepalives) needed for this ticket. Use a Mermaid sequence or state diagram.
2. **Codebase Mini-Map**: Identify the 2–4 key files/structs involved with clickable `file://` links and explain who calls whom.
3. **Interactive Concept Sandbox (Generative UI)**:
   - Read the `generative_ui` skill (`/home/zabatistas/.gemini/antigravity/builtin/skills/generative_ui/SKILL.md`) if you haven't already in this session.
   - Build a compact (<480px tall) interactive HTML widget in the artifact directory and embed it inline with `<agent-embed src="file:///<artifact_path>/<widget>.html"></agent-embed>`.
   - Also copy the widget to `~/projects/vibe-learn/learnings/<YYYY-MM-DD>-<topic-slug>/sandbox.html` so it lives permanently in the user's Learning Hub.

---

### Phase 2: Glass-Box Vibe Coding (Writing the Code)
When implementing the ticket:
1. Write clean, well-tested code that solves the ticket and run tests to verify.
2. Add **Educational Micro-Comments** only where non-obvious concurrency or protocol invariants exist.
3. Include a **Learning Anchor** summary in chat:
   - **Why this concurrency pattern?** (e.g., `sync.RWMutex` vs `sync.Mutex`, channel-close broadcast vs point-to-point send, `select` with `time.NewTimer`).
   - **The "Delete-This-Line" Bomb**: Explicitly name one line in the diff and explain what race condition, deadlock, or protocol bug happens if a developer deletes it.

---

### Phase 3: The Background `shadow-tutor` Subagent & Learning Hub Update
To keep the main coding conversation fast while building a permanent library of lessons:
1. Invoke `shadow-tutor` via `invoke_subagent` in the background (or generate directly if the subagent is read-only) to produce a comprehensive `LEARNING_GUIDE.md` containing:
   - Layer 1: Protocol / Domain Mental Model
   - Layer 2: Codebase Call-Graph & State Transitions
   - Layer 3: Concurrency & Systems Mechanics anchored to exact code lines
   - 3 "What-If" mutation experiments
   - 3 Expandable `<details><summary>` PR review flashcards
2. **Persist & Index in the Central Learning Hub (`~/projects/vibe-learn/learnings/`)**:
   - Save the guide to `~/projects/vibe-learn/learnings/<YYYY-MM-DD>-<topic-slug>/LEARNING_GUIDE.md` (as well as the conversation artifact directory).
   - Register the new entry in both:
     - [`~/projects/vibe-learn/learnings/INDEX.md`](file:///home/zabatistas/projects/vibe-learn/learnings/INDEX.md)
     - [`~/projects/vibe-learn/learnings/index.html`](file:///home/zabatistas/projects/vibe-learn/learnings/index.html) (add the new entry to the `LEARNINGS` JavaScript array so it immediately appears in the searchable browser dashboard).

---

### Phase 4: The PR Defense Check (Mock Senior Reviewer)
Once the code passes tests, switch into **Supportive Senior Reviewer Mode**:
1. Use `ask_question` (or ask 2–3 targeted questions in chat) simulating a real PR review on the exact lines just written:
   - **Question 1 (Protocol Layer)**: Why the protocol behaves a specific way in an edge case.
   - **Question 2 (Concurrency Layer)**: What happens under concurrent load or different lock ordering.
   - **Question 3 (State / Architecture Layer)**: How state transitions or callers react under race conditions.
2. After the developer answers, confirm what they nailed, fill in any blind spots referencing the exact code lines, and link them to their updated Learning Hub page (`~/projects/vibe-learn/learnings/index.html`).
