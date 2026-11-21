{
  description = "Merklecoin scaffold with Go toolchain via Nix flakes";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forEachSystem = f: nixpkgs.lib.genAttrs systems (system:
        let
          pkgs = import nixpkgs { inherit system; };
          go = pkgs.go_1_25;
        in
        f { inherit pkgs go system; });
    in
    {
      packages = forEachSystem ({ pkgs, go, ... }:
        let
          merklecoin = (pkgs.buildGoModule.override { inherit go; }) {
            pname = "merklecoin";
            version = "0.0.1";
            src = pkgs.lib.fileset.toSource {
              root = ./.;
              fileset = pkgs.lib.fileset.unions [
                ./go.mod
                (pkgs.lib.fileset.maybeMissing ./go.sum)
                ./cmd
                ./internal
                ./tests
              ];
            };
            modRoot = ".";
            subPackages = [ "cmd/node" ];
            vendorHash = null; # No external deps yet; set after adding modules.
            ldflags = [ "-s" "-w" ];
            # The default check only tests subPackages, missing the library tests.
            checkPhase = ''
              runHook preCheck
              go test -p "$NIX_BUILD_CORES" ./...
              runHook postCheck
            '';
            meta = with pkgs.lib; {
              description = "Learning scaffold for a Bitcoin-like chain using Merkle trees";
              homepage = "https://example.invalid/merklecoin";
              license = licenses.mit;
              platforms = systems;
              mainProgram = "node";
            };
          };
        in
        {
          inherit merklecoin;
          default = merklecoin;
        });

      checks = forEachSystem ({ system, ... }: {
        merklecoin = self.packages.${system}.merklecoin;
      });

      devShells = forEachSystem ({ pkgs, go, ... }: {
        default = pkgs.mkShell {
          packages = [
            go
            pkgs.gotools
            pkgs.gopls
            pkgs.just
            pkgs.watchexec
          ];
          env = {
            GO111MODULE = "on";
            GOTOOLCHAIN = "local";
            # Do not inherit a GOROOT from another Go installation (e.g. mise).
            GOROOT = "${go}/share/go";
          };
        };
      });
    };
}
