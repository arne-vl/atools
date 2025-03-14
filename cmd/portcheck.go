package cmd

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// portCheckCmd represents the portcheck command
var portCheckCmd = &cobra.Command{
	Use:   "portcheck [port]",
	Short: "Check if a port is occupied or free",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		port := args[0]
		if _, err := strconv.Atoi(port); err != nil {
			fmt.Println("Invalid port number")
			return
		}

		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", port), timeout)
		if err != nil {
			fmt.Printf("Port %s is free.\n", port)
		}
		if conn != nil {
			defer conn.Close()
			fmt.Printf("Port %s is occupied.\n", port)
		}
	},
}

func init() {
	rootCmd.AddCommand(portCheckCmd)
}
