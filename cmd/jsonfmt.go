package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var compact bool

// jsonfmtCmd represents the jsonfmt command
var jsonfmtCmd = &cobra.Command{
	Use:   "jsonfmt [document]",
	Short: "Pretty print JSON",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		fileContent, err := readFileContent(filePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
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
