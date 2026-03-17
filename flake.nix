{
  description = "Description for the project";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = inputs@{ flake-parts, self, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
				./package.nix
      ];
      systems = [ "x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin" ];
      perSystem = {pkgs, self', ... }: {
				devShells.default = pkgs.mkShell {
					name = "go";
					packages = with pkgs; [
						go just zsh sqlite sqlc
					];
					shellHook = ''
						export APP_VERSION=${self'.packages.default.version}
						export SHELL=${pkgs.zsh}/bin/zsh
						exec $SHELL
						'';
				};
      };
    };
}
