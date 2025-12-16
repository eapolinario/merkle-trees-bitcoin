# Merklecoin Plan

## Goals
- Build a Bitcoin-like prototype emphasizing Merkle tree usage, proof-of-work, and basic networking.
- Keep components small and testable; ship a CLI node for local dev networks.

## Tasks (incremental)
- **Serialization & hashing**: Deterministic header/tx encoding; compute prev-block hashes; replace placeholder PoW header serializer.
- **Merkle**: Add proof generation for leaves; verify roots during block validation.
- **Consensus/validation**: Implement PoW mining loop, difficulty targets/adjustment, and full block validation (coinbase rules, fee/amount checks, double-spend prevention).
- **State**: Build in-memory UTXO set; add disk persistence (e.g., boltdb/pebble); snapshot/restore on start.
- **P2P**: Define message types (handshake, inv, getdata, tx, block); peer manager with dialing/retry; gossip loop for tx/block propagation.
- **Wallet/tx creation**: Persist keys, derive addresses, build/sign transactions, and simple coin selection.
- **CLI/RPC**: Commands for wallet create/show balance, submit tx, inspect chain/peers; light JSON-RPC or gRPC server.
- **Testing**: Unit tests for merkle proofs, tx hashing/validation, PoW targets; integration test that mines a short chain with multiple nodes.
- **Tooling**: After adding dependencies, lock `vendorHash` in `flake.nix`; add lint/format helpers (gofmt/staticcheck) to dev shell.
