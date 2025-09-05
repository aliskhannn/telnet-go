package main

import (
	"fmt"
	"os"

	"github.com/aliskhannn/telnet-go/internal/client"
	"github.com/aliskhannn/telnet-go/internal/config"
	"github.com/aliskhannn/telnet-go/internal/flags"
	"github.com/spf13/pflag"
)

func main() {
	// Initialize flags.
	options := flags.InitFlags()
	pflag.Parse()

	if pflag.NArg() == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "error: host and port are required")
		pflag.Usage()
		os.Exit(1)
	}

	cfg := config.New(*options.Timeout)

	args := pflag.Args()
	host := args[0]
	port := args[1]

	c := client.New(host, port, cfg.Timeout)
	if err := c.Run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "telnet error:", err)
		os.Exit(1)
	}
}
