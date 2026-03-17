{self, ... }: {
	perSystem = { pkgs, ... }: let
		version = "1.0.0";
		ldflags = [
			"-s" "-w"
			"-X github.com/z3co/prot/cmd.Version=${version}"
			"-X github.com/z3co/prot/cmd.Commit=${self.rev or "dirty"}"
		];
	in {
		packages.default = pkgs.buildGoModule {
			pname = "prot";
			inherit version ldflags;
			src = ./.;
			vendorHash = "sha256-OxJd9m4+nM/yqpKyc/KQ8rZb14hC6YFWf+dA8PuIruE=";
		};
	};
}
