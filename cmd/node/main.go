package main

import (
    "context"
    "flag"
    "log"
    "os"
    "os/signal"
    "strings"
    "syscall"

    "merklecoin/internal/p2p"
)

// main wires flags and hands off to the P2P node; everything else hangs off that node.
func main() {
    rpcAddr := flag.String("rpc", ":8080", "RPC listen address")
    p2pAddr := flag.String("p2p", ":30333", "P2P listen address")
    dataDir := flag.String("data", "./data", "Data directory for chain state and wallets")
    mine := flag.Bool("mine", false, "Enable CPU miner")
    bootstrap := flag.String("bootstrap", "", "Comma-separated list of peer addresses to connect to on startup")

    flag.Parse()

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigs
        cancel()
    }()

    nodeCfg := p2p.Config{
        RPCListen:     *rpcAddr,
        P2PListen:     *p2pAddr,
        DataDir:       *dataDir,
        EnableMining:  *mine,
        BootstrapPeers: splitPeers(*bootstrap),
    }

    node, err := p2p.NewNode(nodeCfg)
    if err != nil {
        log.Fatalf("failed to create node: %v", err)
    }

    if err := node.Run(ctx); err != nil {
        log.Fatalf("node terminated: %v", err)
    }
}

func splitPeers(raw string) []string {
    if raw == "" {
        return nil
    }
    var peers []string
    for _, p := range strings.Split(raw, ",") {
        if trimmed := strings.TrimSpace(p); trimmed != "" {
            peers = append(peers, trimmed)
        }
    }
    return peers
}
