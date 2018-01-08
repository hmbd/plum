package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	//dup1()
	dup2()
}

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if strings.ToLower(line) == "exit" || strings.ToUpper(line) == "end"{
			break
		}
		counts[line] ++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("第 %d 行\t%s\t重复了\n", n, line)
		}
	}
}

func dup2(){
	counts := make(map[string]int)
	//files := os.Args[1:]
	fmt.Println("命令行参数： ", os.Args)
	files := []string{"D:\\GoProjects\\src\\plum\\test.go"}
	fmt.Println("文件名列表： ", files)
	if len(files) == 0{
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files{
			fmt.Println("文件名： ", arg)
			f, err := os.Open(arg)
			if err != nil{
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts{
		if n > 1{
			fmt.Printf("第 %d 行\t%s\t重复了\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan(){
		line := input.Text()
		if strings.ToLower(line) == "exit" || strings.ToUpper(line) == "end"{
			break
		}
		counts[line] ++
	}
}

func dup3(){
	counts := make(map[string]int)
	for _, filename := range os.Args[1:]{
		data, err := ioutil.ReadFile(filename)
		if err != nil{
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n"){
			counts[line] ++
		}
	}
	for line, n := range counts{
		if n > 1{
			fmt.Printf("第 %d 行\t%s\t重复了\n", n, line)
		}
	}
}