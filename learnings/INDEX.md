# Vibe-Learn: Central Learning Hub Index

Welcome to your portable **Vibe-Learn Catalog**. Every time you solve a ticket using `vibe-learn` or spawn the `shadow-tutor`, your deep-dive guide, interactive sandbox, and PR Defense flashcards are indexed here using relative links.

> **Interactive Browser Dashboard:** Open [`./index.html`](./index.html) in your browser (or preview it in your IDE) for a searchable, tag-filtered dashboard of all your learnings!

---

## Learning Sessions Catalog

| Date | Ticket / Topic | Protocol / Domain | Concurrency & Architecture Concepts | Links |
| :--- | :--- | :--- | :--- | :--- |
| **2026-11-24** | **[TELCO-402: MagicSock Direct-UDP to DERP Fallback](./2026-11-24-telco-402-magicsock/LEARNING_GUIDE.md)** | Tailscale `MagicSock`, `Disco` NAT Hole-Punching (`ping`/`pong`), `DERP` Encrypted TCP Relay Fallback | Go `sync.RWMutex` (`RLock` fast-path vs `Lock`), Channel-Close 1-to-N Broadcast (`close(probeCh)`), Lock release before blocking `select`, `time.NewTimer` cleanup | [Learning Guide](./2026-11-24-telco-402-magicsock/LEARNING_GUIDE.md) · [Interactive Sandbox](./2026-11-24-telco-402-magicsock/sandbox.html) · [Example Code](../examples/telco-magicsock/conn.go) |

---

## Quick Concept Lookup

### Network & Telecom Protocols
- **Tailscale MagicSock & NAT Traversal (`Disco` + `DERP`)**: [TELCO-402 Guide — Layer 1](./2026-11-24-telco-402-magicsock/LEARNING_GUIDE.md)

### Concurrency & Systems Primitives (Go)
- **`sync.RWMutex` Reader Fast-Path vs. Writer Exclusivity**: [TELCO-402 Guide — Layer 3.1](./2026-11-24-telco-402-magicsock/LEARNING_GUIDE.md)
- **1-to-N Goroutine Broadcast via `close(chan struct{})`**: [TELCO-402 Guide — Layer 3.2](./2026-11-24-telco-402-magicsock/LEARNING_GUIDE.md)
- **Avoiding Lock-Wait Deadlocks Across `select` + `time.NewTimer`**: [TELCO-402 Guide — Layer 3.3](./2026-11-24-telco-402-magicsock/LEARNING_GUIDE.md)
