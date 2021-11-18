.PHONY: build
build:
	go build -o ./bin/notify-daily notify-daily.go utils.go \
	&& go build -o ./bin/notify-coming-soon notify-coming-soon.go utils.go
goose-build:
	go get -u github.com/pressly/goose/v3/cmd/goose \
	&& go build -o ./bin/goose github.com/pressly/goose/v3/cmd/goose
debug-daily:
	go run notify-daily.go utils.go
debug-coming-soon:
	go run notify-coming-soon.go utils.go
