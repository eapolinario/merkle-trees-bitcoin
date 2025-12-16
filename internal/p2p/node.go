package p2p

import (
    "context"
    "errors"
    "fmt"
    "log"
    "net"
    "sync"
    "time"
)

// Config holds node-level configuration values.
type Config struct {
    RPCListen      string
    P2PListen      string
    DataDir        string
    EnableMining   bool
    BootstrapPeers []string
}

// Node ties together networking, consensus, and storage. Many parts are stubbed.
type Node struct {
    cfg    Config
    peers  *PeerManager
    gossip *Gossip
    stop   chan struct{}
    once   sync.Once
}

// NewNode constructs a Node with default subcomponents.
func NewNode(cfg Config) (*Node, error) {
    if cfg.P2PListen == "" {
        return nil, errors.New("P2P listen address required")
    }

    pm := NewPeerManager()
    g := NewGossip(pm)

    return &Node{
        cfg:    cfg,
        peers:  pm,
        gossip: g,
        stop:   make(chan struct{}),
    }, nil
}

// Run starts network listeners and the gossip loop; this is a stub for now.
func (n *Node) Run(ctx context.Context) error {
    log.Printf("node starting: rpc=%s p2p=%s mining=%v", n.cfg.RPCListen, n.cfg.P2PListen, n.cfg.EnableMining)
    ln, err := net.Listen("tcp", n.cfg.P2PListen)
    if err != nil {
        return fmt.Errorf("p2p listen failed: %w", err)
    }
    defer ln.Close()

    go n.acceptLoop(ln)
    go n.gossip.Run(ctx)
    go n.connectBootstrap()

    <-ctx.Done()
    n.once.Do(func() { close(n.stop) })
    log.Println("node stopping")
    return nil
}

func (n *Node) acceptLoop(ln net.Listener) {
    for {
        conn, err := ln.Accept()
        if err != nil {
            select {
            case <-n.stop:
                return
            default:
            }
            log.Printf("accept error: %v", err)
            continue
        }
        // TODO: handshake and upgrade connection; for now just close.
        conn.Close()
    }
}

func (n *Node) connectBootstrap() {
    for _, peer := range n.cfg.BootstrapPeers {
        if err := n.peers.AddStatic(peer); err != nil {
            log.Printf("bootstrap add failed for %s: %v", peer, err)
        }
    }
    // TODO: dial peers and start sessions.
    time.Sleep(10 * time.Millisecond)
}

// Peer represents a remote node; mostly a placeholder for now.
type Peer struct {
    Addr string
}

// PeerManager tracks known peers; not thread-safe yet.
type PeerManager struct {
    peers []Peer
}

// NewPeerManager constructs a manager.
func NewPeerManager() *PeerManager {
    return &PeerManager{peers: make([]Peer, 0)}
}

// AddStatic adds a peer to the known set.
func (pm *PeerManager) AddStatic(addr string) error {
    if addr == "" {
        return errors.New("empty peer addr")
    }
    pm.peers = append(pm.peers, Peer{Addr: addr})
    return nil
}

// Gossip handles message dissemination; a stub for now.
type Gossip struct {
    peers *PeerManager
}

// NewGossip returns a Gossip struct.
func NewGossip(pm *PeerManager) *Gossip {
    return &Gossip{peers: pm}
}

// Run pumps periodic gossip ticks; real implementation will multiplex messages.
func (g *Gossip) Run(ctx context.Context) {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            // TODO: broadcast inventory, transactions, and blocks.
        }
    }
}
