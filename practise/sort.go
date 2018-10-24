package main

import (
	"fmt"
	"strings"
)

type Interface interface {
	Len() int
	Less(i, j int) bool // i小于j的比较结果
	Swap(i, j int)
}

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var names = StringSlice{"Python", "Go", "Java", "C"}

	fmt.Println(strings.Join(names, " "))
	//sort.Sort(names)
	for i := 0; i < names.Len()-1; i++ {
		if names.Less(i, i+1) {
			names.Swap(i, i+1)
		}
	}
	for _, item := range names[:names.Len()-1] {
		fmt.Println(item)
	}
	fmt.Println(strings.Join(names, " "))
}
