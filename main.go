package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/ZhengXinwei-F/ProbeX/grpc"
	"github.com/ZhengXinwei-F/ProbeX/http"
	"github.com/ZhengXinwei-F/ProbeX/tcp"
)

func main() {
	var timeout time.Duration
	var silence bool

	rootCmd := &cobra.Command{
		Use:   "prober",
		Short: "A simple TCP, HTTP, and gRPC probe tool",
	}

	// define global flags in the root command
	rootCmd.PersistentFlags().DurationVarP(&timeout, "timeout", "t", 3*time.Second, "Timeout duration")
	rootCmd.PersistentFlags().BoolVarP(&silence, "silence", "s", false, "Silence mode, suppress all output")
	// Pass the flags of the root command to the child command
	tcpCmd := tcp.Cmd(&timeout, &silence)
	httpCmd := http.Cmd(&timeout, &silence)
	grpcCmd := grpc.Cmd(&timeout, &silence)

	// add subcommand
	rootCmd.AddCommand(tcpCmd, httpCmd, grpcCmd)

	// execute the root command
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
