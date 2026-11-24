#!/usr/bin/env bash
# Resolves the portable Learning Hub directory for vibe-learn across any user's machine
# and initializes index.html + INDEX.md if they do not exist yet.
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SKILL_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
TEMPLATE_HTML="$SKILL_DIR/resources/dashboard-template.html"

# Priority 1: Explicit environment variable $VIBE_LEARN_HUB
# Priority 2: User config pointer ~/.config/vibe-learn/hub_path
# Priority 3: Repo-local .vibe-learn/learnings if .vibe-learn exists in current working directory
# Priority 4: Default portable fallback ~/.vibe-learn/learnings
resolve_hub_dir() {
  if [[ -n "${VIBE_LEARN_HUB:-}" ]]; then
    echo "$VIBE_LEARN_HUB"
  elif [[ -f "$HOME/.config/vibe-learn/hub_path" ]]; then
    cat "$HOME/.config/vibe-learn/hub_path"
  elif [[ -d ".vibe-learn" ]]; then
    echo "$(pwd)/.vibe-learn/learnings"
  else
    echo "$HOME/.vibe-learn/learnings"
  fi
}

HUB_DIR="$(resolve_hub_dir)"
mkdir -p "$HUB_DIR"

if [[ ! -f "$HUB_DIR/index.html" && -f "$TEMPLATE_HTML" ]]; then
  cp "$TEMPLATE_HTML" "$HUB_DIR/index.html"
fi

if [[ ! -f "$HUB_DIR/INDEX.md" ]]; then
  cat > "$HUB_DIR/INDEX.md" <<'EOF'
# Vibe-Learn: Central Learning Hub Index

Welcome to your portable **Vibe-Learn Catalog**. Every time you solve a ticket using `vibe-learn` or spawn the `shadow-tutor`, your deep-dive guide, interactive sandbox, and PR Defense flashcards are indexed here using relative links.

> **Interactive Browser Dashboard:** Open [`./index.html`](./index.html) in your browser for a searchable, tag-filtered dashboard of all your learnings!

---

## Learning Sessions Catalog

| Date | Ticket / Topic | Protocol / Domain | Concurrency & Architecture Concepts | Links |
| :--- | :--- | :--- | :--- | :--- |
EOF
fi

echo "$HUB_DIR"
