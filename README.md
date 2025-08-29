# telnet-go

A simple Telnet client implemented in Go.

This tool allows you to connect to any TCP server and interact with it in a bidirectional way — typing commands in your terminal and seeing the server responses in real-time.

---

## Features

- Connect to a specified host and port with a configurable timeout.
- Read from stdin and send to the TCP server.
- Print server responses to stdout.
- Exit when the user presses `Ctrl+D` (EOF).

---

## Installation

```bash
# Clone the repository
git clone https://github.com/aliskhannn/telnet-go.git
cd telnet-go

# Build the binary
make build

# Run the client
./bin/telnetgo <host> <port> [--timeout=5s]
````

---

## Usage Example with an Echo Server

You can test the client by connecting to a public echo server:

```bash
go run cmd/telnet/main.go tcpbin.com 4242
```

Example session:

```
Connected to tcpbin.com:4242
hello
hello
pint 324
pint 324

Connection closed
```

* Everything you type is echoed back by the server.
* Press `Ctrl+D` to exit the client.

---

## CLI Options

* `--timeout` (optional): connection timeout duration (default: 10s).

Example:

```bash
./bin/telnetgo example.com 80 --timeout=5s
```

---

## Development

* Lint the code:

```bash
make lint
```

* Clean build artifacts:

```bash
make clean
```

---

## Project Structure

```
telnet-go/
├── cmd/telnet/        # Main entry point
├── internal/config/   # Configuration (timeout, defaults)
├── internal/flags/    # CLI flags parsing
├── internal/client/   # Telnet client logic
├── Makefile           # Build, lint, clean commands
```
