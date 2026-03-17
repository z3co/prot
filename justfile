version := var_env_or_default("APP_VERSION", "dev")
flags := "-s -w -X 'github.com/z3co/prot/cmd.Version={{version}}' -X 'github.com/z3co/prot/cmd.Commit=$(git rev-parse --short HEAD)' -X 'github.com/z3co/prot/cmd.BuildTime=$(date -u +%Y-%d-%mT%H:%M:%SZ)'"
release:
	#!/usr/bin/env bash
	mkdir -p dist
	for os in linux darwin; do
		for arch in amd64 arch64; do
			echo "Building $os-$arch..."
			GOOS=$os GOARCH=$arch go build -ldflags="{{flags}}" -o prot
			tar -czf "dist/prot-{{version}}-$os-$arch.tar.gz" prot
			rm prot
		done
	done
	echo "Building windows-amd64..."
	GOOS=windows GOARCH=amd64 go build -ldflags="{{flags}}" -o prot.exe
	zip "dist/prot-{{version}}-windows-amd64.tar.gz" prot.exe
	rm prot.exe

	echo "Generating checksums..."
	cd dist sha256sum *.tar.gz *.zip > checksums.txt
	echo "Done!"


