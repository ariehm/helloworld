
main:
	go build -o dist/hw ./cmd/hw

release-build:
	goreleaser build --snapshot --clean

release-check:
	goreleaser check
