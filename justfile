set shell := ["powershell.exe", "-c"]


build:
    go build cmd/learninggo/main.go

run: build
    .\main.exe

test:
    go test ./... -v


test-watch:
    watchexec -w . go test ./... -v
