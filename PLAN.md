# Merklecoin Plan

## Goals
- Build a Bitcoin-like prototype emphasizing Merkle tree usage, proof-of-work, and basic networking.
- Keep components small and testable; ship a CLI node for local dev networks.

## Current state
- The Go/Nix development and build tooling is in place.
- Three unit tests pass; `tests/local_mining_test.go` is still skipped.
- Transaction and header hashing truncate integers and omit fields.
- Merkle verification only handles siblings on the right; proof generation is missing.
- `Chain.Append` does not validate blocks, and `-mine` does not start a miner.

## Next milestone: validated local mining
Build a single node that mines a short, valid chain and supports transaction-inclusion proofs. Establish correctness before expanding networking or RPC.

Keep this milestone deliberately small:
- Coinbase-only blocks with a fixed reward.
- A fixed, easy difficulty suitable for local tests.
- In-memory chain state, with no persistence or external dependencies required.
- A documented prototype format, not Bitcoin wire compatibility or production security.

## Implementation sequence

### 1. Canonical serialization and hashing
Primary files: `internal/tx/tx.go`, `internal/blockchain/block.go`, `internal/consensus/pow.go`.

- [ ] Define a canonical binary encoding with full-width integers, explicit collection/byte-string lengths, fixed-size hashes, and a specified byte order and timestamp resolution.
- [ ] Include all transaction fields, including scripts and lock time, and all consensus-relevant header fields.
- [ ] Expose a header hash used consistently for block identity, previous-block linkage, and proof of work.
- [ ] Specify a deterministic genesis block, its expected hash, and its PoW treatment.
- [ ] Add golden encoding/hash vectors and regression tests for previously omitted fields and integer values above 255.

Acceptance: identical logical values produce identical bytes and hashes; changing any committed field changes the hash in the test vectors. No independent placeholder header serializer remains.

### 2. Complete Merkle proofs
Primary files: `internal/merkle/tree.go`, `internal/merkle/tree_test.go`.

- [ ] Generate a proof for any valid leaf index.
- [ ] Represent sibling direction explicitly or derive it from a leaf index so verification handles either side.
- [ ] Document empty-tree, single-leaf, and duplicate-last behavior for odd levels.
- [ ] Avoid modifying input buffers when concatenating hashes.
- [ ] Test every leaf across several tree sizes, including right-hand and odd-tail leaves.
- [ ] Test invalid indices, tampered proofs/roots, and input immutability.

Acceptance: every leaf in the test trees has a verifiable inclusion proof; invalid proofs are rejected under the documented tree conventions.

### 3. Mining and mandatory block validation
Primary files: `internal/consensus/`, `internal/blockchain/chain.go`, `cmd/node/main.go`.

- [ ] Implement a cancellable nonce-search loop using the canonical header hash.
- [ ] Define target calculation and boundary behavior; reject zero/invalid difficulty without panicking.
- [ ] Validate previous-block linkage, the recomputed Merkle root, the configured difficulty, and proof of work.
- [ ] Enforce basic coinbase rules: one coinbase, the fixed reward, and a unique identity per block, for example by committing the block height.
- [ ] Reject ordinary spending transactions until UTXO/signature validation is implemented.
- [ ] Make validation mandatory before appending; rejection must leave chain state unchanged.
- [ ] Choose an acyclic validation boundary: `consensus` already imports `blockchain`, so do not introduce mutual imports.
- [ ] Wire `-mine` to actual mining and report accepted block heights/hashes; keep application orchestration separate from P2P transport responsibilities.

Acceptance: a node can mine and append valid coinbase-only blocks, reject malformed candidates, and stop mining on context cancellation.

### 4. End-to-end acceptance test
Primary file: `tests/local_mining_test.go`.

- [ ] Replace the skipped test with a bounded local test that mines five blocks after genesis.
- [ ] Verify height, previous-hash linkage, expected difficulty, proof of work, and transaction commitments.
- [ ] Verify transaction-inclusion proofs for mined blocks; exercise nontrivial multi-leaf paths in the Merkle unit tests.
- [ ] Verify wrong previous hashes, altered transactions/roots, and invalid work are rejected without advancing the chain.
- [ ] Verify cancellation stops the miner promptly.
- [ ] Use controlled inputs and an easy target to keep tests fast and reproducible.

Milestone completion checks: `just check` and `just test-race` pass, the local-mining test no longer skips, and `-mine` performs real mining.

## Suggested PR boundaries
1. Canonical serialization and hashing, with golden vectors.
2. Complete Merkle proof generation/verification and edge-case tests.
3. Mining, mandatory validation, CLI wiring, and the local-mining integration test.

## Follow-up milestones
1. **Spendable state**: In-memory UTXOs, signature verification, amount/fee checks, overflow protection, and double-spend rejection.
2. **Persistence**: Store chain/UTXO state and verify restart recovery; choose a storage backend once the state model is settled.
3. **Two-node networking**: Message framing and handshake, peer lifecycle, chain synchronization, and transaction/block propagation.
4. **Wallet and RPC usability**: Persist keys, derive addresses, select coins, build/sign transactions, query balances, and submit transactions.

Keep build/test tooling current as these land; set `vendorHash` in `flake.nix` if external Go dependencies are introduced.

## Deferred
Difficulty adjustment, forks/reorgs, full Bitcoin-style scripting, and Bitcoin protocol compatibility are outside the next milestone. This remains a learning prototype, not a production cryptocurrency.
