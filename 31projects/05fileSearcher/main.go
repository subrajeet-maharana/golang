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

func searchLine(line string, lineNum int, term string, caseSensitive bool) *SearchResult {
  searchLine := line
  searchTerm := term

  if !caseSensitive {
    searchLine = strings.ToLower(line)
    searchTerm = strings.ToLower(term)
  }

  pos := strings.Index(searchLine, searchTerm)
  if pos >= 0 {
    return &SearchResult{
      LineNum: lineNum,
      Line: line,
      Term: term, 
      Position: pos, 
    }
  }
  
  return nil
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
      if result := searchLine(line, lineNum, searchTerm, config.CaseSensitive); result != nil {
        results = append(results, *result)
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
