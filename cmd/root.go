package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "atools",
	Short: "Arne's CLI Tools",
	Long: `   _____   __                .__
  /  _  \_/  |_  ____   ____ |  |   ______
 /  /_\  \   __\/  _ \ /  _ \|  |  /  ___/
/    |    \  | (  <_> |  <_> )  |__\___ \
\____|__  /__|  \____/ \____/|____/____  >
        \/                             \/

Arne's CLI Tools contains some fun / useful commands.
For suggestions: open an issue at https://github.com/arne-vl/atools-cli.

Created by Arne Van Looveren.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// TODO: add version command
}
