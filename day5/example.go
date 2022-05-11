package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Pair struct {
	Word  string
	Count int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}
func (p PairList) Less(i, j int) bool {
	return p[i].Count < p[j].Count
}
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	words := readFile()
	counter := getSortedPairList(words)
	makeResultFile(counter)
	fmt.Println(counter)
}

func makeResultFile(counter PairList) {
	f, err := os.Create("stat.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, pair := range counter {
		f.WriteString(fmt.Sprintf("%v\t%v\n", pair.Word, pair.Count))
	}
}

func getSortedPairList(words []string) PairList {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	counter := make(PairList, len(m))
	i := 0
	for word, count := range m {
		//counter = append(counter, Pair{word, count})
		counter[i] = Pair{word, count}
		i++
	}
	//sort.Sort(sort.Reverse(counter))
	//либо сделать так
	sort.Slice(counter, func(i, j int) bool {
		return counter[i].Count > counter[j].Count
	})
	return counter
}

func readFile() []string {
	data, err := ioutil.ReadFile("./day5/pg1661.txt")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	text := strings.ToLower(string(data))
	reg, err := regexp.Compile("[a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	words := reg.FindAllString(text, -1)
	return words
}
