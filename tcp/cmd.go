package tcp

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// Probe check if the tcp port is available
func Probe(address string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

// Cmd returns a subcommand of the tcp command
func Cmd(timeout *time.Duration, silence *bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tcp address port",
		Short: "Check if a TCP port is open",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			address := args[0] + ":" + args[1]
			if err := Probe(address, *timeout); err != nil {
				if !*silence {
					fmt.Printf("TCP probe failed: %v\n", err)
				}
				os.Exit(1)
			} else {
				if !*silence {
					fmt.Println("TCP probe succeeded")
				}
			}
		},
	}

	return cmd
}
