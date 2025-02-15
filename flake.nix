{
  description = "urban-tui";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  outputs = { self,nixpkgs }: {
    packages."x86_64-linux".default = with import nixpkgs { system = "x86_64-linux";};
      buildGoModule rec {
        pname = "urban-tui";
        version = "0.1.0";
        vendorHash = "sha256-93xVsTemE8HIvWKjBcArsyram1Z1WQhbF5yyig2yPLI=";
        CGO_ENABLED = 0;
        src = ./.;
      nativeBuildInputs = [
          pkgs.go
          ];
      postBuild = ''
         go build -o urban-tui
          '';
      installPhase = ''
         mkdir -p $out/bin
         mv urban-tui $out/bin
        '';
      };
      
  };

}
