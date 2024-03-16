{
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.11";

  outputs = { self, nixpkgs }:
    let
      goVersion = "21";
      overlays = [ (final: prev: { go = prev."go_1_${toString goVersion}"; }) ];
      supportedSystems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forEachSupportedSystem = f: nixpkgs.lib.genAttrs supportedSystems (system: f {
        pkgs = import nixpkgs { inherit overlays system; };
      });
    in
    {
      devShells = forEachSupportedSystem ({ pkgs }: {
        default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gotools
            golangci-lint
          ];
          shellHook = ''
			      # Code ENVs
            export MQTT_BROKER="10.0.0.108"
            export MQTT_PORT="1883"
			      export MQTT_USER="root"
			      export MQTT_PASSWORD="coreflux"
         '';
        };
      });
    };
}

