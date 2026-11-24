# Ticket TELCO-402: Race Condition & Hanging Packets in MagicSock Direct-UDP to DERP Fallback

**Reported by:** Core Network SRE Team  
**Severity:** P1 (Customer VPN Drops & Panics in Edge POPs)  
**Component:** `telco.internal/magicsock` (Tailscale / WireGuard NAT Traversal Engine)

## Background (For New Engineers)
Our mesh VPN uses a custom socket multiplexer called **`MagicSock`** (modeled after Tailscale's `magicsock`) underneath WireGuard:
1. Ideally, two peers talk over a **Direct UDP** path (`PathDirectUDP`) discovered via **Disco (`ping`/`pong`)** NAT hole-punching probes.
2. Behind harsh Symmetric NATs or carrier-grade firewalls, direct UDP fails. When that happens, packets must fall back to **DERP** (Detoured Encrypted Routing Protocol — an encrypted TCP/HTTPS relay server) while background Disco probes continue trying to upgrade the connection back to Direct UDP.

## Symptoms Observed in Production
1. **Crash (`fatal error: concurrent map read and map write`)**: When 50+ worker goroutines send packets through `Conn.Send()` at the same time a Disco `pong` probe response arrives in `Conn.HandleDiscoPong()`, the daemon crashes.
2. **Packet Hang / Blackhole**: When a peer has no verified direct UDP path (or its probe fails/times out), `Conn.Send()` fires a probe and blocks forever waiting on `peer.probeCh`, dropping customer voice/data traffic instead of falling back to the DERP relay after `ProbeTimeout` and waking all waiting callers.

## Acceptance Criteria
1. Eliminate all data races in `Conn` (`go test -race ./...` must pass cleanly under high concurrency).
2. In `Conn.Send()`: if a peer is not yet on `PathDirectUDP` (or its direct route has expired), probe for a direct path up to `c.ProbeTimeout`, and if no Disco `pong` arrives in time, **gracefully fall back to sending via `PathDERP`** without leaking goroutines or deadlocking multiple concurrent senders.
3. When a `HandleDiscoPong` arrives from a peer, safely upgrade the peer's state to `PathDirectUDP` and notify any waiting senders immediately.
