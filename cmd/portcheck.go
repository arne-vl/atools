package cmd

import (
	"fmt"
	"net"
	"strconv"

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

		addr := fmt.Sprintf(":%s", port)
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			fmt.Printf("Port %s is occupied.\n", port)
		} else {
			fmt.Printf("Port %s is free.\n", port)
			ln.Close()
		}
	},
}

func init() {
	rootCmd.AddCommand(portCheckCmd)
}
