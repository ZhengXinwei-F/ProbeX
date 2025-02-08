package http

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// Probe checks whether the HTTP endpoint is reachable
func Probe(url string, timeout time.Duration, skipTLS bool) error {
	client := &http.Client{
		Timeout: timeout,
	}
	if skipTLS {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	return fmt.Errorf("HTTP probe failed, status code: %d", resp.StatusCode)
}

// Cmd returns a subcommand of the HTTP command
func Cmd(timeout *time.Duration, silence *bool) *cobra.Command {
	var skipTLS bool

	cmd := &cobra.Command{
		Use:   "http url",
		Short: "Check if an HTTP endpoint is reachable",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			if err := Probe(url, *timeout, skipTLS); err != nil {
				if !*silence {
					fmt.Printf("HTTP probe failed: %v\n", err)
				}
				os.Exit(1)
			} else {
				if !*silence {
					fmt.Println("HTTP probe succeeded")
				}
			}
		},
	}

	// Add flags to the command
	cmd.Flags().BoolVarP(&skipTLS, "skip-tls", "k", false, "Skip TLS verification")

	return cmd
}
