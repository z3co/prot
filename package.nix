{
	perSystem = { pkgs, ... }: {
		packages.default = pkgs.buildGoModule {
			pname = "prot";
			version = "0.1.0";
			src = ./.;
			vendorHash = "sha256-OxJd9m4+nM/yqpKyc/KQ8rZb14hC6YFWf+dA8PuIruE=";
		};
	};
}
