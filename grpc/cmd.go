package grpc

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Probe checks whether the gRPC server is accessible and verifies the service's health status
func Probe(address string, service string, timeout time.Duration) error {
	// Connect to gRPC server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(timeout))
	if err != nil {
		return fmt.Errorf("failed to connect to gRPC server: %w", err)
	}
	defer conn.Close()

	// Create a health check client
	client := healthpb.NewHealthClient(conn)

	// Send a health check request
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := client.Check(ctx, &healthpb.HealthCheckRequest{Service: service})
	if err != nil {
		return fmt.Errorf("gRPC health check failed: %w", err)
	}

	// Output health status
	fmt.Println("gRPC Health Status:", resp.Status)
	return nil
}

func Cmd(timeout *time.Duration, silence *bool) *cobra.Command {
	var service string

	cmd := &cobra.Command{
		Use:   "grpc address port",
		Short: "Check if a gRPC server is reachable and optionally verify service health",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			address := args[0] + ":" + args[1]

			if err := Probe(address, service, *timeout); err != nil {
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

	// 添加 --service 选项，短写 -S
	cmd.Flags().StringVarP(&service, "service", "S", "", "Specify the gRPC service to check")

	return cmd
}
