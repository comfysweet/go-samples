package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	verbose = flag.Bool("v", false, "вывод промежуточных результатов")
	// подсчитывающий семафор для ограничения параллельности
	sema = make(chan struct{}, 20)
	done = make(chan struct{})
)

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//walkDir рекурсивно обходит дерево файлов с корнем в dir
// и отправляет размер каждого найденного файла в fileSizes
func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			time.Sleep(2 * time.Second)
			walkDir(subdir, fileSizes, wg)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents возвращает записи каталога dir
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // захват маркера
	case <-done:
		return nil // отмена
	}

	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d файлов %.1f KB\n", nfiles, float64(nbytes)/1024)
}

// команда du вычисляет суммарные размер всех файлов в каталоге
func main() {
	//определят исходные каталоги
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"../"}
	}

	go func() {
		os.Stdin.Read(make([]byte, 1)) // чтение 1 байта
		close(done)
	}()

	// обход дерева файлов
	filesizes := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, filesizes, &wg)
	}
	go func() {
		wg.Wait()
		close(filesizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// вывод результатов
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// опустошение канала folesizes, чтобы позволить
			// завершиться существующим go-подпрограммам
			for range filesizes {
				//ничего не делаем
			}
			return
		case size, ok := <-filesizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}
