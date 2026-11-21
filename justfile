set shell := ["bash", "-euo", "pipefail", "-c"]
set positional-arguments

# List available recipes.
default:
    @just --list

# Enter the Nix development shell (also loaded by .envrc via direnv).
dev:
    nix develop

# Build the local node binary, using Go's incremental build cache.
build:
    mkdir -p bin
    go build -o bin/node ./cmd/node

# Build and run the node; pass CLI flags after the recipe name.
run *args: build
    ./bin/node "$@"

# Run all unit/integration tests; pass Go test flags after the recipe name.
test *args:
    go test "$@" ./...

# Run all tests with Go's race detector.
test-race *args:
    go test -race "$@" ./...

# Format Go source in place.
fmt:
    go fmt ./...

# Run Go's static analyzer.
vet:
    go vet ./...

# Build, vet, and test the project without modifying source files.
check: build vet test

# Re-run checks whenever Go source or module files change (Ctrl-C to stop).
watch:
    watchexec --restart --exts go,mod,sum -- just check

# Build the reproducible Nix package; output is result/bin/node.
nix-build:
    nix build .#merklecoin

# Build and test the package through the flake's checks.
nix-check:
    nix flake check

# Run the packaged node; pass CLI flags after the recipe name.
nix-run *args:
    nix run .#merklecoin -- "$@"

# Remove local build output only; leave chain/wallet data and Go caches intact.
clean:
    rm -rf bin
