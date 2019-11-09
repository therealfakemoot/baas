{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
	buildInputs = with pkgs; [ go figlet ];
	shellHook = ''
		go run ./cmd/
	'';
}
