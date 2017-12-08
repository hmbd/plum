package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile0(filePath string) string {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	return string(f)
}

func readFile1(filePath string) string {
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

func readFile2(filePath string) string {
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 1024, 1024)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

func readFile3(filePath string) string {
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func readMain() {
	file := "test.txt"

	start := time.Now()

	readFile0(file)
	t0 := time.Now()
	fmt.Printf("readFile0 time: %v\n", t0.Sub(start))

	readFile1(file)
	t1 := time.Now()
	fmt.Printf("readFile1 time: %v\n", t1.Sub(t0))

	readFile2(file)
	t2 := time.Now()
	fmt.Printf("readFile2 time: %v\n", t2.Sub(t1))

	readFile3(file)
	t3 := time.Now()
	fmt.Printf("readFile3 time: %v\n", t3.Sub(t2))

}

func writeFile0(filePath string) {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(filePath, d1, 0644)
	check(err)
}
func writeFile1(filePath string) {
	f, err := os.Create(filePath)
	check(err)

	defer f.Close()

	d2 := []byte{97, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("writeFile1 写入 %d 个字节\n", n2)

	n3, err := f.WriteString("writeFile1\n")
	fmt.Printf("writeFile1 写入 %d 个字节\n", n3)

	f.Sync()
}

func writeFile2(filePath string) {
	f, err := os.Create(filePath)
	check(err)

	defer f.Close()
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("writeFile2\n")
	fmt.Printf("writeFile2 写入 %d 个字节\n", n4)

	w.Flush()

}
func writeFile3(filePath string) {
	f, err1 := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664) //打开文件
	check(err1)
	defer f.Close()

	buffer := []byte("writeFile3\n")
	n, err1 := f.Write(buffer) // 字节写入
	check(err1)
	fmt.Printf("writeFile3 写入 %d 个字节\n", n)

	n1, err2 := f.WriteString("writeFile3\n") //写入文件(字符串)
	check(err2)
	fmt.Printf("writeFile3 写入 %d 个字节\n", n1)
	f.Sync()
}
func writeMain() {
	file := "test.txt"
	writeFile0(file)
	writeFile1(file)
	writeFile2(file)
	writeFile3(file)
}

func main() {
	writeMain()
	readMain()
}
