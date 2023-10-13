.PHONY: build
build:
	sh ./script/build.sh

.PHONY: run
run: build
	./build/mercury-darwin-arm64 --mode demo