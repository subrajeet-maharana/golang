package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type SearchConfig struct {
  CaseSensitive bool
  FilePath string
  SearchTerms []string
}

type SearchResult struct {
  LineNum int
  Line string
  Term string
  Position int
}

func searchFile(config SearchConfig) ([]SearchResult, error ){
  file, err := os.Open(config.FilePath)
  if err != nil {
    return nil, fmt.Errorf("Failed to open the file: %v", err)
  }
  defer file.Close()

  var results []SearchResult

  scanner := bufio.NewScanner(file)
  lineNum := 0
  matchCount := 0

  for scanner.Scan() {
    lineNum++
    line := scanner.Text()

    for _, searchTerm := range config.SearchTerms {
      searchLine := line

      if !config.CaseSensitive {
        searchLine = strings.ToLower(line)
        searchTerm = strings.ToLower(searchTerm)
      }

      if strings.Contains(searchLine, searchTerm) {
        fmt.Printf("Line is: %d: %s\n", lineNum, line)
        results = append(results, SearchResult{lineNum, line, searchTerm, 0})
        matchCount++
      }
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Printf("Error reading file: %v\n", err)
  }

  fmt.Printf("\nFound %d matches\n", matchCount)
  return results, nil
}

func main() {
  config := SearchConfig{}

  flag.StringVar(&config.FilePath, "file", "", "Path to the file to search")
  flag.BoolVar(&config.CaseSensitive, "case-sensitive", false, "Enable case sensitive search")
  flag.Parse()

  config.SearchTerms = flag.Args()
  if len(config.SearchTerms) ==0 || config.FilePath == "" {
    fmt.Println("Usage: program -file=<filepath> [-case-sensitive] <search_term1> [search_term2 ...]")
    flag.PrintDefaults()
    os.Exit(1)
  }

  log.Printf("Starting search in file: %s\n", config.FilePath)
  if err := searchFile(config); err != nil {
    fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    os.Exit(1)
  }
}
