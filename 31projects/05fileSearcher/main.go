package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type SearchConfig struct {
  CaseSensitive bool
  FilePath string
  SearchTerm string
}

func searchFile(config SearchConfig) error {
  file, err := os.Open(config.FilePath)
  if err != nil {
    return fmt.Errorf("Failed to open the file: %v", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  lineNum := 0
  matchCount := 0

  for scanner.Scan() {
    lineNum++
    line := scanner.Text()

    searchLine := line
    searchTerm := config.SearchTerm
    
    if !config.CaseSensitive {
      searchLine = strings.ToLower(line)
      searchTerm = strings.ToLower(config.SearchTerm)
    }

    if strings.Contains(searchLine, searchTerm) {
      fmt.Printf("Line is: %d: %s\n", lineNum, line)
      matchCount++
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Printf("Error reading file: %v\n", err)
  }

  fmt.Printf("\nFound %d matches\n", matchCount)
  return nil
}

func main() {
  config := SearchConfig{}

  flag.StringVar(&config.FilePath, "file", "", "Path to the file to search")
  flag.BoolVar(&config.CaseSensitive, "case-sensitive", false, "Enable case sensitive search")
  flag.Parse()

  args := flag.Args()
  if len(args) !=1 || config.FilePath == "" {
    fmt.Println("Usage: program -file=<filepath> [-case-sensitive] <search_term>")
    flag.PrintDefaults()
    os.Exit(1)
  }
  config.SearchTerm = args[0]

  if err := searchFile(config); err != nil {
    fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    os.Exit(1)
  }
}
