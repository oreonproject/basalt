{
  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs =
    inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [ ];
      systems = [ "x86_64-linux" ];
      perSystem =
        {
          config,
          self',
          inputs',
          pkgs,
          system,
          ...
        }:
        {
          packages.default = import ./default.nix {
            inherit (pkgs) buildGoModule;
            lib = inputs.nixpkgs.lib;
          };
          apps.default = {
            type = "app";
            program = "${self'.packages.default}/bin/basalt";
          };
          devShells = {
            default = pkgs.mkShell {
              buildInputs = [
                pkgs.gcc
                pkgs.go
                pkgs.git
              ];
            };
          };
        };
      flake = { };
    };
}
