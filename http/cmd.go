package http

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// Probe checks whether the HTTP endpoint is reachable
func Probe(url string, timeout time.Duration, skipTLS bool, headers map[string]string) error {
	client := &http.Client{
		Timeout: timeout,
	}
	if skipTLS {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// 设置 Headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
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
	var headers []string

	cmd := &cobra.Command{
		Use:   "http url",
		Short: "Check if an HTTP endpoint is reachable",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			// 解析 Headers
			headerMap := make(map[string]string)
			for _, h := range headers {
				parts := strings.SplitN(h, ":", 2)
				if len(parts) == 2 {
					headerMap[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
				}
			}

			if err := Probe(url, *timeout, skipTLS, headerMap); err != nil {
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
	cmd.Flags().StringArrayVarP(&headers, "header", "H", []string{}, "Add custom HTTP headers (e.g., -H 'Authorization: Bearer token')")

	return cmd
}
