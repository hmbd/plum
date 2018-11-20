package main

import (
	"fmt"
	"github.com/seekplum/plum/gopcp/chapter5/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	workList := make(chan []string)

	go func() { workList <- os.Args[1:] }()

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
