#!/bin/bash

root="$PWD";

build() {

	go_os="windows"
	go_arch="$1"

	cd "${root}/source";

	env CGO_ENABLED=0 GOOS="${go_os}" GOARCH="${go_arch}" go build -o "${root}/build/fixcrowdstrike_${go_arch}.exe";

}


build "386";
build "amd64";

