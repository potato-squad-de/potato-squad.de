{
  outputs = {
    self,
    flake-utils,
    nixpkgs,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {inherit system;};

      hugoEnv = with pkgs; [
        hugo
        nodejs
        pnpm
      ];

      goEnv = with pkgs; [
        go
        gopls
        golangci-lint
        go-tools
        impl
        wgo
        delve

        goose
      ];
    in {
      devShells.default = pkgs.mkShell {
        packages = with pkgs; [
          git
          lefthook
          commitlint-rs
          just

          hugoEnv
          goEnv
          sqlite
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

          if [ -f go.mod ]; then
            go mod tidy
          fi
        '';
      };
    });

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };
}