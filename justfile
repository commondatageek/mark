[private]
default:
    just --list

build:
    go build .
    go build -o ./bin/ ./cmd/...