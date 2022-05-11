package main

import (
	"fmt"
	links "github.com/comfysweet/go-samples/ch5"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // захват маркера
	list, err := links.Extract(url)
	<-tokens // освобождение маркера
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)
	var n int // количество ожидающих отправки в список

	// запуск с аргументами командной строки
	n++
	go func() {
		workList <- os.Args[1:]
	}()

	// параллельное сканирование
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-workList
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
