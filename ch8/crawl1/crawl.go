package main

import (
	"fmt"
	links "github.com/comfysweet/go-samples/ch5"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)
	go func() {
		workList <- os.Args[1:]
	}()

	// параллельное сканирование
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}
