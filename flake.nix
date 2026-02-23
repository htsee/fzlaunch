{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
  };

  outputs =
    {
      self,
      nixpkgs,
      ...
    }@inputs:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs {
        inherit system;
      };
    in
    {
      devShells."x86_64-linux".default = pkgs.mkShell {
        packages = with pkgs; [
          go
          gopls
          golangci-lint
          golangci-lint-langserver
          fzf
          television
        ];
      };
    };
}
