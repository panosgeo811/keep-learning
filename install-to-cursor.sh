#!/usr/bin/env bash
set -euo pipefail

TARGET_REPO="${1:-.}"
VIBE_LEARN_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

mkdir -p "$TARGET_REPO/.cursor/rules"
cp "$VIBE_LEARN_ROOT/.cursor/rules/vibe-mentor.mdc" "$TARGET_REPO/.cursor/rules/"
cp "$VIBE_LEARN_ROOT/.cursor/rules/vibe-learn-workflow.mdc" "$TARGET_REPO/.cursor/rules/"

echo "Installed Vibe-Learn Cursor rules into $TARGET_REPO/.cursor/rules/"
echo "  - vibe-mentor.mdc (Always-On Glass-Box Coding & Learning Anchors)"
echo "  - vibe-learn-workflow.mdc (On-Demand 4-Phase Workflow + Central Learning Hub Archival)"
