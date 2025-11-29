set shell := ["powershell.exe", "-c"]


build:
    go build ./...

run: build
    .\learninggo.exe

