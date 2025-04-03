package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// jsonfmtCmd represents the jsonfmt command
var jsonfmtCmd = &cobra.Command{
	Use:   "jsonfmt",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jsonfmt called")
	},
}

func init() {
	rootCmd.AddCommand(jsonfmtCmd)
}
