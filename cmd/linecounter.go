package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var listFiles bool   // Flag to list files by line count
var spreadFiles bool // Flag to show the amount of files
var directory string // Flag to specify directory
var recursive bool   // Flag to enable recursive search

// lineCounterCmd represents the linecounter command
var lineCounterCmd = &cobra.Command{
	Use:   "linecounter [extension]",
	Short: "Count the number of lines in files with a specific extension",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ext := args[0]
		if directory == "" {
			directory = "."
		}

		ext = strings.TrimPrefix(ext, ".")

		fileLines, totalLines, err := countLinesInFiles(directory, ext, recursive)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Printf("Total lines in *.%s files: %d\n", ext, totalLines)

		if spreadFiles {
			fmt.Println("Spread over", len(fileLines), "file(s)")
		}

		if listFiles {
			sort.Slice(fileLines, func(i, j int) bool {
				return fileLines[i].lines > fileLines[j].lines
			})
			fmt.Println("\nFiles sorted by line count:")
			for _, file := range fileLines {
				fmt.Printf("%d\t%s\n", file.lines, file.path)
			}
		}
	},
}

type fileLineCount struct {
	path  string
	lines int
}

func countLinesInFiles(dir, ext string, recursive bool) ([]fileLineCount, int, error) {
	totalLines := 0
	var fileLines []fileLineCount

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && !recursive && path != dir {
			return filepath.SkipDir
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), "."+ext) {
			lines, err := countLines(path)
			if err != nil {
				return err
			}
			totalLines += lines
			fileLines = append(fileLines, fileLineCount{path, lines})
		}
		return nil
	})

	if err != nil {
		return nil, 0, err
	}
	return fileLines, totalLines, nil
}

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines, scanner.Err()
}

func init() {
	lineCounterCmd.Flags().BoolVarP(&listFiles, "list", "l", false, "List files sorted by line count")
	lineCounterCmd.Flags().BoolVarP(&spreadFiles, "spread", "s", false, "Show amount of files scanned")
	lineCounterCmd.Flags().StringVarP(&directory, "directory", "d", "", "Directory to scan (default: current directory)")
	lineCounterCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Enable recursive search")
	rootCmd.AddCommand(lineCounterCmd)
}
