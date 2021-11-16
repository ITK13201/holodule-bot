.PHONY: build
build:
	go build -o ./bin/holodule-bot main.go
goose-build:
	go get -u github.com/pressly/goose/v3/cmd/goose \
	&& go build -o ./bin/goose github.com/pressly/goose/v3/cmd/goose
