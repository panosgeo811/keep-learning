package magicsock

import "time"

// PathState represents how MagicSock currently routes encrypted WireGuard packets to a peer.
type PathState string

const (
	// PathUnknown means we haven't established a path yet.
	PathUnknown PathState = "UNKNOWN"
	// PathDirectUDP means NAT hole-punching succeeded via Disco ping/pong probes.
	PathDirectUDP PathState = "DIRECT_UDP"
	// PathDERP means direct UDP timed out or failed, so we route via the encrypted TCP DERP relay.
	PathDERP PathState = "DERP_RELAY"
)

// Packet represents an encrypted WireGuard frame ready to be transmitted.
type Packet struct {
	PeerID  string
	Payload []byte
}

// SendResult records which transport path actually carried the packet.
type SendResult struct {
	PeerID   string
	PathUsed PathState
	Endpoint string
}

//DiscoPong represents an incoming NAT traversal probe response from a peer's UDP endpoint.
type DiscoPong struct {
	PeerID   string
	UDPAddr  string
	Received time.Time
}
