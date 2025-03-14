// concurrency and parallelism are different things in golang
// concurrency is context switching in rapid pace so that it seems like parallel execution while parallelism is completely cpu bound multi-core implementation
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer
var signals = []string{"test"}

func main() {
	// go greeter("Subrajeet")
	// greeter("Rakesh")
	websiteList := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
	}

  go func() {
    fmt.Println("Anonymous go routine running...")
    time.Sleep(5*time.Second)
    fmt.Println("Anonymous go routine ending...")
  }()

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}
func greeter(s string) {
	// for i := 0; i < 6; i++ {
	// 	time.Sleep(5 * time.Millisecond)
	// 	fmt.Println(s)
	// }
}

func getStatusCode(webOp string) {
	defer wg.Done()
	res, err := http.Get(webOp)
	if err != nil {
		fmt.Println("Error in connection...")
	} else {
		mut.Lock()
		signals = append(signals, webOp)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, webOp)
	}
}
