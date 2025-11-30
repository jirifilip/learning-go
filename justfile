build:
    go build cmd/learninggo/main.go


run: build
    ./main


test:
    go test ./... -v


test-watch:
    watchexec -w . go test ./... -v
