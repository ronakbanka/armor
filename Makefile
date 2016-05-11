COMMIT = $$(git describe --always)

deps:
	go get -v ./...

# build generate binary on './bin' directory.
build: deps
	go build -ldflags "-X main.GitCommit=$(COMMIT)" -o bin/armor

test: build
	go test -v ./...
