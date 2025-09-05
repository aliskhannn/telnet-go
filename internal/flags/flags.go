package flags

import (
	"time"

	"github.com/spf13/pflag"
)

// Flags holds command-line options for the program.
type Flags struct {
	Timeout *time.Duration // connection timeout
}

// InitFlags initializes and parses command-line flags.
func InitFlags() Flags {
	return Flags{
		Timeout: pflag.DurationP("timeout", "t", 10*time.Second, "Timeout for TCP-server connection"),
	}
}
