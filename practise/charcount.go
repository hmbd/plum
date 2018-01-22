package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid ++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 charactes\n", invalid)
	}
}

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}
func Add(list []string) {
	m[k(list)] ++
}

func Count(list []string) int {
	return m[k(list)]
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string){
	edges := graph[from]
	if edges == nil{
		// 惰性初始化map是一个惯用方式
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool{
	return graph[from][to]
}