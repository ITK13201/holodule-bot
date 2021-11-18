.PHONY: build
build:
	go build -o ./bin/notify-daily notify-daily.go \
	&& go build -o ./bin/notify-coming-soon notify-coming-soon.go
goose-build:
	go get -u github.com/pressly/goose/v3/cmd/goose \
	&& go build -o ./bin/goose github.com/pressly/goose/v3/cmd/goose
