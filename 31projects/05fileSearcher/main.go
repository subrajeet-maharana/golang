package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
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
  resultsChan := make(chan SearchResult)
  done := make(chan struct{})
  var wg sync.WaitGroup

  go func() {
    for result := range resultsChan {
      results = append(results, result)
    }
    done <- struct{}{}
  }()

  scanner := bufio.NewScanner(file)
  lineNum := 0

  for scanner.Scan() {
    lineNum++
    line := scanner.Text()

    for _, searchTerm := range config.SearchTerms {
      wg.Add(1)
      go func(l string, ln int, t string) {
        defer wg.Done()
        if result := searchLine(l, ln, t, config.CaseSensitive); result != nil {
          log.Printf("Found match for '%s' on line %d", t, ln)
          resultsChan <- *result
        }
      }(line, lineNum, searchTerm)
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Printf("Error reading file: %v\n", err)
  }

  wg.Wait()
  close(resultsChan)
  <-done

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
  results, err := searchFile(config)
  if err != nil {
    log.Fatalf("Error: %v", err)
  }
  
  for _, result := range results {
    fmt.Printf("Line %d, pos %d (%s): %s\n",
      result.LineNum,
      result.Position,
      result.Term,
      result.Line,
    )
  }
}
