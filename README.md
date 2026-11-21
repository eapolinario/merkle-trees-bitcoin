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

## Development

Enter the Nix development shell, then use `just` for the edit/build/test loop:

```sh
nix develop
just check
just run -p2p :30333
```

The shell includes Go 1.25, gotools, gopls, just, and watchexec. It sets `GOROOT`
to match the Nix compiler, avoiding conflicts with other Go installations.
If you use direnv, `direnv allow` loads the same shell through `.envrc`.
Alternatively, install Go 1.25 or newer and just yourself (plus watchexec for
`just watch`); normally leave `GOROOT` unset so Go finds its own toolchain.

| Command | Purpose |
| --- | --- |
| `just` | List available recipes |
| `just build` | Build `bin/node` using Go's incremental cache |
| `just run -p2p :30334` | Build and run the node with CLI flags |
| `just test` | Run all tests |
| `just test -run TestVerifyProof -count=1` | Run selected tests without caching |
| `just test-race` | Run tests with the race detector |
| `just fmt` | Format Go source in place |
| `just vet` | Run Go's static analyzer |
| `just check` | Build, vet, and test without changing source files |
| `just watch` | Re-run checks on Go/module changes; Ctrl-C to stop |
| `just clean` | Remove `bin/`, preserving node data and Go caches |

The node is still a scaffold: mining, RPC, and persistence are not implemented,
and the local-mining integration test is currently skipped.

## Nix package

```sh
nix build .#merklecoin       # or: just nix-build / nix build
./result/bin/node -help
nix run .#merklecoin -- -p2p :30334  # or: just nix-run -p2p :30334
nix flake check              # or: just nix-check
```

Both the default and `merklecoin` package use the same Go 1.25 toolchain as the
dev shell. Package builds run `go test ./...`, including the library tests.
There are no external Go dependencies yet; after adding any, set `vendorHash`
in `flake.nix` to the hash of the vendored modules.

## Next steps
1) Implement deterministic serialization for headers and transactions; compute prev-block hashes.
2) Wire proof-of-work mining and difficulty adjustment; add block validation rules (merkle root check, coinbase, fees, double-spend prevention).
3) Build a minimal UTXO set in memory, then persist to disk.
4) Flesh out P2P message types (handshake, inv, getdata, tx, block) and actual networking flows.
5) Add RPC/CLI commands for wallet creation, balance queries, and transaction submission.
6) Write unit tests for Merkle proofs, transaction hashing, and PoW target calculation; expand integration tests to mine a few blocks.

## Notes
- All cryptography is placeholder and not suitable for production; this scaffold is for learning and experimentation.
