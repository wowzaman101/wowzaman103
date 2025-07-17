GO := GO111MODULE=on go
DOCKER := DOCKER_DEFAULT_PLATFORM=linux/amd64

.PHONY: ci
ci:
	$(GO) mod tidy && \
	$(GO) mod download && \
	$(GO) mod verify && \
	$(GO) mod vendor && \
	$(GO) fmt ./... \

.PHONY: build
build:
	$(GO) build -mod=vendor -a -installsuffix cgo -tags musl -o main ./cmd/.

start:
	go run --tags dynamic $(shell pwd)/cmd/.

dev: 
	nodemon --exec go run --tags dynamic $(shell pwd)/cmd/wire_gen.go --signal SIGTERM

test:
	$(GO) test -cover ./... -coverprofile=cover.out -covermode count && go tool cover -html cover.out -o cover.html

wire:
	wire cmd/main.go

.PHONY: clean
clean:
	@rm -rf main ./vendor