package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"yaml-parser/internal"
)

const (
	version = "1.0.0"
	appName = "yaml-parser"
)

func main() {
	// Define flags
	showVersion := flag.Bool("version", false, "show program version")
	showHelp := flag.Bool("help", false, "show help")
	flag.Parse()

	// Show version
	if *showVersion {
		fmt.Printf("%s version %s\n", appName, version)
		os.Exit(0)
	}

	// Show help
	if *showHelp {
		printHelp()
		os.Exit(0)
	}

	// Get arguments (files to compare)
	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Error: at least 2 files required for comparison\n")
		printHelp()
		os.Exit(1)
	}

	configs := make([]internal.Config, 0, len(args))
	fileNames := make([]string, 0, len(args))

	for _, arg := range args {
		fmt.Printf("Loading %s...\n", arg)
		config, err := internal.LoadConfig(arg)
		if err != nil {
			log.Fatalf("Error loading %s: %v", arg, err)
		}
		configs = append(configs, config)
		fileNames = append(fileNames, arg)
	}

	fmt.Println("\nComparing configurations...")
	internal.CompareAllConfigs(configs, fileNames)
}

func printHelp() {
	fmt.Printf(`%s - YAML configuration files comparison tool

Usage:
  %s [options] file1.yml file2.yml [file3.yml ...]

Options:
  -version     Show program version
  -help        Show this help message

Arguments:
  file1.yaml    Base file for comparison
  file2.yaml    File to compare with base
  file3.yaml    Additional files to compare

Examples:
  %s config1.yaml config2.yaml
  %s config1.yaml config2.yaml config3.yaml
  %s -version
  %s -help

Description:
  This tool loads YAML files and compares their structure.
  The first file is used as the reference for comparison.
  Output shows missing keys in the compared files.

Exit codes:
  0 - All configurations are identical
  1 - Differences found or error occurred

`, appName, appName, appName, appName, appName, appName)
}
