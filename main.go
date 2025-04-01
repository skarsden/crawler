package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]

	const maxConcurrency = 10
	cfg, err := configure(baseURL, maxConcurrency)
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
