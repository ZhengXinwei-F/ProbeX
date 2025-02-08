package grpc

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Probe Check whether the gRPC server is accessible
func Probe(address string, timeout time.Duration) error {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(timeout))
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

// Cmd returns the subcommand of the grpc command
func Cmd(timeout *time.Duration, silence *bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grpc address port",
		Short: "Check if a gRPC server is reachable",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			address := args[0] + ":" + args[1]
			if err := Probe(address, *timeout); err != nil {
				if !*silence {
					fmt.Printf("gRPC probe failed: %v\n", err)
				}
				os.Exit(1)
			} else {
				if !*silence {
					fmt.Println("gRPC probe succeeded")
				}
			}
		},
	}

	return cmd
}
