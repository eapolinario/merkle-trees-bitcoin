# Merklecoin (scaffold)

This is a skeleton for a Bitcoin-like cryptocurrency focused on Merkle tree mechanics. It intentionally leaves room to fill in consensus, networking, persistence, and wallet features.

## Layout
- `cmd/node`: CLI entrypoint to start a node.
- `internal/blockchain`: Block/header definitions and chain management.
- `internal/consensus`: Proof-of-work and validation stubs.
- `internal/merkle`: Merkle tree builder and proof verification.
- `internal/p2p`: Placeholder networking node, gossip loop, and peer manager.
- `internal/tx`: Transaction/UTXO primitives.
- `internal/wallet`: Ephemeral keypair management and signing.
- `tests`: Integration harness placeholder.

## Nix
- `flake.nix` provides a Go toolchain and build recipe.
- Dev shell: `nix develop` (drops you into an env with Go 1.21, gotools, gopls).
- Build the node binary: `nix build .#merklecoin` (output in `./result/bin/node`).

## Next steps
1) Implement deterministic serialization for headers and transactions; compute prev-block hashes.
2) Wire proof-of-work mining and difficulty adjustment; add block validation rules (merkle root check, coinbase, fees, double-spend prevention).
3) Build a minimal UTXO set in memory, then persist to disk.
4) Flesh out P2P message types (handshake, inv, getdata, tx, block) and actual networking flows.
5) Add RPC/CLI commands for wallet creation, balance queries, and transaction submission.
6) Write unit tests for Merkle proofs, transaction hashing, and PoW target calculation; expand integration tests to mine a few blocks.

## Notes
- Go tooling is not installed in this environment; `go.mod` is provided but `go` commands will fail until installed.
- All cryptography is placeholder and not suitable for production; this scaffold is for learning and experimentation.
