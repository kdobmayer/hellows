BINARY = hello
OUTPUT_DIR = out
DOCKER_USER = kdobmayer

all: clean test lint $(BINARY)

$(BINARY):
	mkdir -p $(OUTPUT_DIR)
	go build -o ./$(OUTPUT_DIR)/$(BINARY) ./cmd/hello

install:
	go install -o $(BINARY) ./...

test:
	go test -v ./...

lint:
	golint -set_exit_status ./...

clean:
	@rm -rf $(OUTPUT_DIR)

docker-build:
	docker build -f ./build/Dockerfile -t $(DOCKER_USER)/$(BINARY) .

distclean: clean
	docker image rm $(DOCKER_USER)/$(BINARY)
	docker image prune --force

.PHONY: install test lint clean docker-build distclean
