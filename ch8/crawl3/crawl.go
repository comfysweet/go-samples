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
	workList := make(chan []string)  // список url, могут быть дубли
	unseenLinks := make(chan string) // удаление дублей

	// запуск с аргументами командной строки
	go func() {
		workList <- os.Args[1:]
	}()

	// создание 20 go-подпрограмм сканирования для выборки
	// всех непросмотренных ссылок
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { workList <- foundLinks }()
			}
		}()
	}

	// главная go-подпрограмма удаляет дубликаты из списска
	// и отправляет непросмотренные ссылки сканерам
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
