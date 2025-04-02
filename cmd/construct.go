package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var silent bool // Flag to silent "created" print

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

func constructFilesAndDirs(config *Config) error {
	for _, dir := range config.Construction.Directories {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
		if silent {
			fmt.Println("Created directory:", dir)
		}
	}

	for _, file := range config.Construction.Files {
		fileDir := filepath.Dir(file.Path)
		if err := os.MkdirAll(fileDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory for file %s: %w", file.Path, err)
		}

		if err := os.WriteFile(file.Path, []byte(file.Content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", file.Path, err)
		}
		if silent {
			fmt.Println("Created file:", file.Path)
		}
	}

	return nil
}

// constructCmd represents the construct command
var constructCmd = &cobra.Command{
	Use:   "construct [blueprint]",
	Short: "Construct your blueprints",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a blueprint to construct")
			os.Exit(1)
		}

		filename, err := getConfigPath(args[0])
		if err != nil {
			fmt.Println("Error locating blueprint:", err)
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
	constructCmd.Flags().BoolVarP(&silent, "silent", "s", true, "Silent printing what is created")
	rootCmd.AddCommand(constructCmd)
}
