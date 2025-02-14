{
  description = "urban-tui";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  outputs = { self,nixpkgs }: {
    packages."x86_64-linux".default = with import nixpkgs { system = "x86_64-linux";};
      stdenv.mkDerivation {
        name = "urban-tui";
        src = self;
        buildInputs = [
          pkgs.go
        ];
        buildPhase = "go mod tidy && go build -o urban-tui";
        installPhase = "mkdir -p $out/bin; install -t $out/bin urban-tui";
      };
  };

}
