.PHONY:
.SILENT:

build:
	GOARCH=amd64 GOOS=windows go build -o ./.bin/bot ./cmd/main.go

run: build
		./.bin/bot