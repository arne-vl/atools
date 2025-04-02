package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var silent bool // Flag to silent "created" print

var presetVariables = []string{"year", "quarter", "month", "monthnumber", "weeknumber", "day", "daynumber", "hour", "minute"}

type Config struct {
	Construction struct {
		Directories []string `yaml:"directories"`
		Files       []File   `yaml:"files"`
	} `yaml:"construction"`
}

type File struct {
	Path    string `yaml:"path"`
	Content string `yaml:"content,omitempty"`
}

func getConfigDir() string {
	var homeDir string
	if runtime.GOOS == "windows" {
		homeDir = os.Getenv("USERPROFILE")
	} else {
		homeDir = os.Getenv("HOME")
	}
	return filepath.Join(homeDir, ".config", "atools", "blueprints")
}

func getConfigPath(filename string) (string, error) {
	configDir := getConfigDir()
	possibleFiles := []string{
		filepath.Join(configDir, filename),
		filepath.Join(configDir, filename+".yaml"),
		filepath.Join(configDir, filename+".yml"),
	}

	for _, file := range possibleFiles {
		if _, err := os.Stat(file); err == nil {
			return file, nil
		}
	}

	return "", fmt.Errorf("blueprint not found: %s", filename)
}

func parseYAMLFile(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data = bytes.ReplaceAll(data, []byte("\t"), []byte("    "))

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func stringInArray(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func promptForVariables(content string) map[string]string {
	variables := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	templateTags := map[string]struct{}{}
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		for {
			start := strings.Index(line, "{{")
			end := strings.Index(line, "}}")

			if start == -1 || end == -1 {
				break
			}

			variable := line[start+2 : end]
			templateTags[strings.TrimSpace(variable)] = struct{}{}
			line = line[:start] + line[end+2:]
		}
	}

	for variable := range templateTags {
		currentTime := time.Now()

		if !stringInArray(presetVariables, variable) {
			fmt.Printf("Enter value for %s: ", variable)
			val, _ := reader.ReadString('\n')
			variables[variable] = strings.TrimSpace(val)
			if !silent {
				fmt.Println()
			}
		} else {
			switch variable {
			case "year":
				variables["year"] = fmt.Sprintf("%d", currentTime.Year())
			case "quarter":
				month := currentTime.Month()
				quarter := (int(month)-1)/3 + 1
				variables["quarter"] = fmt.Sprintf("Q%d", quarter)
			case "month":
				variables["month"] = currentTime.Month().String()
			case "monthnumber":
				monthNumber := int(currentTime.Month())
				variables["monthnumber"] = fmt.Sprintf("%02d", monthNumber)
			case "weeknumber":
				dayOfYear := currentTime.YearDay()
				weekNumber := (dayOfYear-1)/7 + 1
				variables["weeknumber"] = fmt.Sprintf("%02d", weekNumber)
			case "day":
				variables["day"] = currentTime.Weekday().String()
			case "daynumber":
				variables["daynumber"] = fmt.Sprintf("%02d", currentTime.Day())
			case "hour":
				variables["hour"] = fmt.Sprintf("%02d", currentTime.Hour())
			case "minute":
				variables["minute"] = fmt.Sprintf("%02d", currentTime.Minute())
			}
		}
	}

	return variables
}

func applyBlueprint(content string, variables map[string]string) string {
	for variable, value := range variables {
		placeholder := fmt.Sprintf("{{ %s }}", variable)
		content = strings.ReplaceAll(content, placeholder, value)
	}

	return content
}

func constructFilesAndDirs(config *Config) error {
	var allContent []string
	allContent = append(allContent, config.Construction.Directories...)
	for _, file := range config.Construction.Files {
		allContent = append(allContent, file.Path)
		allContent = append(allContent, file.Content)
	}

	variables := promptForVariables(strings.Join(allContent, "\n"))

	for _, dir := range config.Construction.Directories {
		processedDir := applyBlueprint(dir, variables)
		if err := os.MkdirAll(processedDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", processedDir, err)
		}
		if !silent {
			fmt.Println("Created directory:", processedDir)
		}
	}

	for _, file := range config.Construction.Files {
		processedPath := applyBlueprint(file.Path, variables)
		fileDir := filepath.Dir(processedPath)
		if err := os.MkdirAll(fileDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory for file %s: %w", processedPath, err)
		}

		processedContent := applyBlueprint(file.Content, variables)

		if err := os.WriteFile(processedPath, []byte(processedContent), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", processedPath, err)
		}
		if !silent {
			fmt.Println("Created file:", processedPath)
		}
	}

	return nil
}

// constructCmd represents the construct command
var constructCmd = &cobra.Command{
	Use:   "construct [blueprint]",
	Short: "Construct your blueprints",

	Run: func(cmd *cobra.Command, args []string) {
		filename, err := getConfigPath(args[0])
		if err != nil {
			fmt.Println("Error locating config file:", err)
			os.Exit(1)
		}

		config, err := parseYAMLFile(filename)
		if err != nil {
			fmt.Println("Error parsing YAML:", err)
			os.Exit(1)
		}

		if err := constructFilesAndDirs(config); err != nil {
			fmt.Println("Error constructing directories and files:", err)
			os.Exit(1)
		}
	},
}

func init() {
	constructCmd.Flags().BoolVarP(&silent, "silent", "s", false, "Silent printing what is created")
	rootCmd.AddCommand(constructCmd)
}
