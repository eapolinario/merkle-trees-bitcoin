{
  description = "Merklecoin scaffold with Go toolchain via Nix flakes";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.11";

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forEachSystem = f: nixpkgs.lib.genAttrs systems (system:
        let
          pkgs = import nixpkgs { inherit system; };
          go = pkgs.go_1_21;
        in
        f { inherit pkgs go system; });
    in
    {
      packages = forEachSystem ({ pkgs, go, system }: {
        default = pkgs.buildGoModule {
          pname = "merklecoin";
          version = "0.0.1";
          src = ./.;
          modRoot = ".";
          subPackages = [ "cmd/node" ];
          vendorHash = null; # No external deps yet; set after adding modules.
          ldflags = [ "-s" "-w" ];
          meta = with pkgs.lib; {
            description = "Learning scaffold for a Bitcoin-like chain using Merkle trees";
            homepage = "https://example.invalid/merklecoin";
            license = licenses.mit;
            platforms = systems;
          };
        };
      });

      devShells = forEachSystem ({ pkgs, go, system }: {
        default = pkgs.mkShell {
          packages = [
            go
            pkgs.gotools
            pkgs.gopls
          ];
          env = {
            GO111MODULE = "on";
          };
        };
      });
    };
}
