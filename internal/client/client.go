package client

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

// Client represents a simple Telnet client with host, port and connection timeout.
type Client struct {
	Host    string
	Port    string
	Timeout time.Duration
}

// New creates a new Client with the specified host, port and timeout.
func New(host string, port string, timeout time.Duration) *Client {
	return &Client{
		Host:    host,
		Port:    port,
		Timeout: timeout,
	}
}

// Run establishes a TCP connection to the server and starts
// two goroutines: one for reading from the connection and
// writing to stdout, and one for reading from stdin and
// writing to the connection. It waits for both goroutines to finish.
func (c *Client) Run() error {
	wg := new(sync.WaitGroup)
	address := net.JoinHostPort(c.Host, c.Port)

	// Attempt to connect to the server with the specified timeout.
	conn, err := net.DialTimeout("tcp", address, c.Timeout)
	if err != nil {
		return fmt.Errorf("error connecting to %s: %w", address, err)
	}

	_, _ = fmt.Fprintf(os.Stdout, "Connected to %s\n", address)

	// Start two goroutines: one for reading from connection,
	// one for writing user input to the connection.
	wg.Add(2)
	go c.readFromConn(conn, wg)
	go c.writeToConn(conn, wg)

	// Wait until both goroutines finish.
	wg.Wait()

	return nil
}

// readFromConn reads data from the TCP connection and writes it to stdout.
// If the connection is closed by the server, it prints a message and returns.
func (c *Client) readFromConn(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	// Copy all data from connection to stdout until EOF.
	_, _ = io.Copy(os.Stdout, conn)

	// Print a message when the server closes the connection.
	_, _ = fmt.Fprintln(os.Stderr, "\nConnection closed")
}

// writeToConn reads user input from stdin and writes it to the TCP connection.
// When EOF (Ctrl+D) is reached, it closes the connection and exits.
func (c *Client) writeToConn(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	scanner := bufio.NewScanner(os.Stdin)

	// Read lines from stdin and send them to the server
	for scanner.Scan() {
		_, _ = fmt.Fprintln(conn, scanner.Text())
	}

	// Check for any error while reading stdin
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

	// Close the connection when user input ends (EOF)
	err := conn.Close()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error closing connection:", err)
	}
}
