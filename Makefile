BINARY=telnetgo
BIN_DIR=bin

.PHONY: build lint clean

build:
	@mkdir -p ${BIN_DIR}
	go build -o ${BIN_DIR}/${BINARY} ./cmd/telnet

lint:
	go vet ./...
	golangci-lint run ./...

format:
	goimports -local github.com/aliskhannn/wget-go -w .

clean:
	rm -rf ${BIN_DIR}
	rm -f ${BINARY}