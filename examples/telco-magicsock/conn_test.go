package magicsock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDERPFallbackOnTimeout(t *testing.T) {
	conn := NewConn(40 * time.Millisecond)
	conn.RegisterPeer("node-athens-01", "derp-fra-1")

	start := time.Now()
	res, err := conn.Send(Packet{PeerID: "node-athens-01", Payload: []byte("wg-encrypted-frame")})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.PathUsed != PathDERP {
		t.Fatalf("expected fallback to %s, got %s", PathDERP, res.PathUsed)
	}
	if time.Since(start) > 250*time.Millisecond {
		t.Fatalf("Send took too long to fall back to DERP: %v", time.Since(start))
	}
}

func TestConcurrentSendAndDiscoPongUpgrade(t *testing.T) {
	conn := NewConn(80 * time.Millisecond)
	conn.RegisterPeer("node-salonika-02", "derp-ath-1")

	var wg sync.WaitGroup
	// Spawn 30 concurrent packet senders while Disco pongs arrive concurrently
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			res, err := conn.Send(Packet{
				PeerID:  "node-salonika-02",
				Payload: []byte(fmt.Sprintf("pkt-%d", id)),
			})
			if err != nil {
				t.Errorf("sender %d failed: %v", id, err)
				return
			}
			if res.PathUsed != PathDirectUDP && res.PathUsed != PathDERP {
				t.Errorf("unexpected path used: %s", res.PathUsed)
			}
		}(i)
	}

	// Simulate asynchronous Disco Pong arriving mid-flight from UDP socket
	time.Sleep(15 * time.Millisecond)
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func(attempt int) {
			defer wg.Done()
			conn.HandleDiscoPong(DiscoPong{
				PeerID:   "node-salonika-02",
				UDPAddr:  "198.51.100.42:41641",
				Received: time.Now(),
			})
		}(j)
	}

	wg.Wait()

	// Subsequent packet should now use Direct UDP fast-path
	res, err := conn.Send(Packet{PeerID: "node-salonika-02", Payload: []byte("post-upgrade")})
	if err != nil {
		t.Fatalf("post-upgrade send failed: %v", err)
	}
	if res.PathUsed != PathDirectUDP {
		t.Fatalf("expected upgraded path %s, got %s", PathDirectUDP, res.PathUsed)
	}
}
