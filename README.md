# Vibe-Learn: Portable Learning-While-Vibe-Coding Kit

Turn vibe coding into guided engineering mastery across **Antigravity** and **Cursor**.
Designed for engineers tackling tickets involving unfamiliar network/domain protocols, unfamiliar codebase architecture, or tricky concurrency primitives.

---

## Quick Start for Any Developer (1-Command Install)

Clone this repository anywhere on your machine and run `./install.sh`:

```bash
git clone <your-repo-url> vibe-learn
cd vibe-learn

# Install Antigravity plugin + initialize your personal Learning Hub:
./install.sh

# Or choose a custom Learning Hub directory & install rules into a Cursor repo:
./install.sh --hub-dir ~/my-learnings --cursor-repo ~/projects/my-service
```

---

## How the Portable Learning Hub Works (`<LEARNING_HUB_DIR>`)

No user paths are hardcoded anywhere in the rules or skills. The agent dynamically resolves `<LEARNING_HUB_DIR>` in this priority order:

1. **`$VIBE_LEARN_HUB`** environment variable (if set)
2. **`~/.config/vibe-learn/hub_path`** (automatically created by `./install.sh`, pointing to `<cloned-repo>/learnings` by default)
3. **`<workspace-root>/.vibe-learn/learnings`** (if your team creates a `.vibe-learn` folder inside a project repo to check shared learnings into git!)
4. **`~/.vibe-learn/learnings`** (zero-config fallback on any OS)

Every learning session is saved under `<LEARNING_HUB_DIR>/<YYYY-MM-DD>-<slug>/` and indexed using **100% relative paths** (`./<YYYY-MM-DD>-<slug>/...`) in:
- **`<LEARNING_HUB_DIR>/index.html`** — Searchable, tag-filtered browser dashboard
- **`<LEARNING_HUB_DIR>/INDEX.md`** — Markdown table & concept index

Because all links inside `learnings/` are relative, you can move the directory, share it as a zip, or publish it to GitHub Pages / GitLab Pages with zero broken links.

---

## What's Included

- **`rules/AGENTS.md`**: Always-On Glass-Box Coding rule (adds concise 2–4 bullet **Learning Anchors** and **"Delete-This-Line" failure modes** whenever protocol or concurrency code is written).
- **`skills/vibe-learn/`**: The 4-Phase workflow (Pre-Flight Brief + Interactive HTML Sandbox -> Glass-Box Coding -> Background `shadow-tutor` -> PR Defense Check) plus `scripts/resolve-hub.sh` and `resources/dashboard-template.html`.
- **`agents/shadow-tutor.md`**: Background technical mentor subagent that builds line-linked `LEARNING_GUIDE.md` guides and flashcards from your `git diff`.
- **`.cursor/rules/*.mdc`**: Drop-in Cursor IDE rules (`vibe-mentor.mdc` and `vibe-learn-workflow.mdc`).
- **`examples/telco-magicsock/`**: Reference Go implementation & concurrency test suite for Ticket `TELCO-402` (Tailscale MagicSock Direct-UDP to DERP Relay fallback).
