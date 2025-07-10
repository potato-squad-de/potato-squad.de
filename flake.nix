{
  outputs = {
    self,
    flake-utils,
    nixpkgs,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
      };
    in {
      devShells.default = pkgs.mkShell {
        packages = with pkgs; [
          git
          lefthook
          commitlint-rs
          just

          hugo
          nodejs
          pnpm
        ];

        shellHook = ''
          if [ ! -d .git ]; then
            git init && git add . && echo "chore: init" >.git/COMMIT_EDITMSG

            if [ -f lefthook.yaml ]; then
              lefthook install
            fi
          fi

          if [ ! -d node_modules ] && [ -f package.json ]; then
            pnpm install
          fi
        '';
      };
    });

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };
}
