package main

import (
	"fmt"
	"io"
	"os"
	"time"
	"reflect"
)

func main() {
	fileBasedPipe()
	inMemorySyncPipe()
}

func fileBasedPipe() {
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Printf("1. 创建管道失败了，失败原因: %s\n", err)
	}
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("1. 从管道中读取数据失败了，失败原因：%s\n", err)
		}
		fmt.Printf("1. 从管道中读取的字节数: %d byte(s)\n", n)
	}()

	input := make([] byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	fmt.Printf("1. 写入的数据类型:%s , 值: %s\n", reflect.TypeOf(input), input)
	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("1. 从管道中读取数据失败了，失败原因：%s\n", err)
	}
	fmt.Printf("1. 往管道中写入字节数 %d byte(s)\n", n)
	time.Sleep(200 * time.Millisecond)
}

func inMemorySyncPipe() {
	reader, writer := io.Pipe()
	go func() {
		output := make([] byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("2. 从管道中读取数据失败了，失败原因：%s\n", err)
		}
		fmt.Printf("2. 从管道中读取的字节数: %d byte(s)\n", n)
	}()

	input := make([] byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	fmt.Printf("2. 写入的数据类型:%s , 值: %s\n", reflect.TypeOf(input), input)
	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("2. 从管道中读取数据失败了，失败原因：%s\n", err)
	}
	fmt.Printf("2. 往管道中写入字节数 %d byte(s)\n", n)
	time.Sleep(200 * time.Millisecond)
}
