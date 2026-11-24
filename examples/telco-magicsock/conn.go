package magicsock

import (
	"errors"
	"sync"
	"time"
)

// PeerState tracks the active routing path for a single remote WireGuard peer.
type PeerState struct {
	PeerID       string
	CurrentPath  PathState
	DirectAddr   string
	DERPRegion   string
	LastPongTime time.Time

	// probeCh is closed (never sent to) when a Disco pong upgrades this peer to PathDirectUDP.
	// Closing a channel in Go broadcasts a zero-value wakeup to ALL goroutines waiting on <-probeCh.
	probeCh     chan struct{}
	probeClosed bool
}

// Conn is a simplified Tailscale-style MagicSock multiplexer.
type Conn struct {
	ProbeTimeout time.Duration
	mu           sync.RWMutex
	peers        map[string]*PeerState
}

func NewConn(probeTimeout time.Duration) *Conn {
	return &Conn{
		ProbeTimeout: probeTimeout,
		peers:        make(map[string]*PeerState),
	}
}

// RegisterPeer adds a peer with its home DERP relay region before traffic starts.
func (c *Conn) RegisterPeer(peerID, derpRegion string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.peers[peerID] = &PeerState{
		PeerID:      peerID,
		CurrentPath: PathUnknown,
		DERPRegion:  derpRegion,
		probeCh:     make(chan struct{}),
	}
}

// HandleDiscoPong is called asynchronously by the UDP read loop whenever a NAT hole-punching
// "pong" packet arrives from a peer.
func (c *Conn) HandleDiscoPong(pong DiscoPong) {
	c.mu.Lock()
	defer c.mu.Unlock()

	peer, ok := c.peers[pong.PeerID]
	if !ok {
		return
	}

	peer.CurrentPath = PathDirectUDP
	peer.DirectAddr = pong.UDPAddr
	peer.LastPongTime = pong.Received

	// Broadcast wakeup: closing an open channel wakes EVERY waiting goroutine at once,
	// and never blocks even if 0 goroutines are currently waiting.
	// Guard with probeClosed because closing an already-closed channel panics in Go.
	if !peer.probeClosed {
		close(peer.probeCh)
		peer.probeClosed = true
	}
}

// Send transmits an encrypted WireGuard packet to a peer, preferring Direct UDP and
// falling back to DERP Relay if NAT traversal times out.
func (c *Conn) Send(pkt Packet) (SendResult, error) {
	c.mu.RLock()
	peer, ok := c.peers[pkt.PeerID]
	if !ok {
		c.mu.RUnlock()
		return SendResult{}, errors.New("unknown peer")
	}

	// Fast path: 99.9% of steady-state packets hit this read-locked check without blocking each other.
	if peer.CurrentPath == PathDirectUDP {
		addr := peer.DirectAddr
		c.mu.RUnlock()
		return SendResult{
			PeerID:   pkt.PeerID,
			PathUsed: PathDirectUDP,
			Endpoint: addr,
		}, nil
	}

	// Snapshot the broadcast channel and fallback DERP region while holding RLock,
	// then MUST release RLock before waiting so HandleDiscoPong can acquire mu.Lock()!
	probeCh := peer.probeCh
	derpRegion := peer.DERPRegion
	c.mu.RUnlock()

	// Slow path: Wait up to c.ProbeTimeout for a Disco Pong broadcast, or fall back to DERP.
	timer := time.NewTimer(c.ProbeTimeout)
	defer timer.Stop()

	select {
	case <-probeCh:
		// Disco Pong arrived and closed probeCh! Re-acquire RLock to safely read the new DirectAddr.
		c.mu.RLock()
		addr := peer.DirectAddr
		c.mu.RUnlock()
		return SendResult{
			PeerID:   pkt.PeerID,
			PathUsed: PathDirectUDP,
			Endpoint: addr,
		}, nil

	case <-timer.C:
		// Direct UDP NAT hole-punching timed out. Upgrade PathUnknown -> PathDERP if still unknown,
		// and route this packet via the encrypted TCP DERP relay so customer traffic never hangs.
		c.mu.Lock()
		if peer.CurrentPath == PathUnknown {
			peer.CurrentPath = PathDERP
		}
		c.mu.Unlock()

		return SendResult{
			PeerID:   pkt.PeerID,
			PathUsed: PathDERP,
			Endpoint: derpRegion,
		}, nil
	}
}
