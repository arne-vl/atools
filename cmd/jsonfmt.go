package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"io"

	"github.com/spf13/cobra"
)

var compact bool // Flag to indicate if the output should be compact

// jsonfmtCmd represents the jsonfmt command
var jsonfmtCmd = &cobra.Command{
	Use:   "jsonfmt [document]",
	Short: "Pretty print JSON",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	var fileContent string
    	var err error

       	if len(args) == 1 {
		// Read from file if an argument is provided
		filePath := args[0]

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Println("Error: File does not exist")
			return
		}

		fileContent, err = readFileContent(filePath)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		} else {
			// Read from stdin if no argument is provided
			stat, _ := os.Stdin.Stat()
			if (stat.Mode() & os.ModeCharDevice) != 0 {
				fmt.Println("Error: No input provided. Provide a file or pipe JSON data.")
				return
			}

			stdinContent, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println("Error reading from stdin:", err)
				return
			}
			fileContent = string(stdinContent)
		}

		err = prettyPrintJSON(fileContent)
		if err != nil {
			fmt.Println("Error:", err)
		}
    },
}

func readFileContent(filePath string) (string, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
        return "", fmt.Errorf("could not read file: %v", err)
    }
    return string(fileContent), nil
}

func prettyPrintJSON(jsonStr string) (error) {
	var js map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &js)
	if err != nil {
		return fmt.Errorf("invalid JSON: %v", err)
	}

	if compact {
		compactJSON, err := json.Marshal(js)
		if err != nil {
			return fmt.Errorf("error compacting JSON: %v", err)
		}
		fmt.Println(string(compactJSON))
		return nil
	}
	prettyJSON, err := json.MarshalIndent(js, "", "  ")
	if err != nil {
		return fmt.Errorf("error pretty printing JSON: %v", err)
	}

	fmt.Println(string(prettyJSON))
	return nil
}

func init() {
	jsonfmtCmd.Flags().BoolVarP(&compact, "compact", "c", false, "Compact JSON output")
	rootCmd.AddCommand(jsonfmtCmd)
}
