BINARY=telnetgo
BIN_DIR=bin

.PHONY: build lint clean

build:
	@mkdir -p ${BIN_DIR}
	go build -o ${BIN_DIR}/${BINARY} ./cmd/telnet

lint:
	go vet ./...
	golangci-lint run ./...

clean:
	rm -rf ${BIN_DIR}
	rm -f ${BINARY}