package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) < 2 {
		fmt.Println("No concurrency maximum or page maximum provided")
		os.Exit(1)
	} else if len(args) < 3 {
		fmt.Println("No page maximum provided")
		os.Exit(1)
	} else if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Second argument is not a number")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Third argument is not a number")
		os.Exit(1)
	}

	cfg, err := configure(baseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Errorf("error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %v\n", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	for key, value := range cfg.pages {
		fmt.Printf("%s - %d\n", key, value)
	}

}
