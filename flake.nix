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
      packages.${system}.default = pkgs.buildGoModule {
        name = "fzlaunch";
        version = "0.1.0";
        src = self;
        goSum = ./go.sum;
        vendorHash = "sha256-7K17JaXFsjf163g5PXCb5ng2gYdotnZ2IDKk8KFjNj0=";
      };

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
