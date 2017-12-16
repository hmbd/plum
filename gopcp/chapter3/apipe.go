package main

import (
	"fmt"
	"os/exec"
	"bytes"
	"io"
	"bufio"
)

func main() {
	runCmd()
	fmt.Println("\n===========================")
	runCmdWithPipe()
}


func runCmd() {
	useBufferIO := false
	message := "My first command comes from goland"
	fmt.Printf("执行命令 `echo -n \"%s\"`:\n", message)
	cmd0 := exec.Command("echo", "-n", message)
	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("获取第一条命令的管道失败了，失败原因： %s\n", err)
		return
	}
	if err := cmd0.Start(); err != nil {
		fmt.Printf("启动第一条命令失败了，失败原因： %s\n", err)
		return
	}
	if useBufferIO {
		outputBuf0 := bufio.NewReader(stdout0)
		output0, _, err := outputBuf0.ReadLine()
		if err != nil {
			fmt.Printf("从管道读取数据失败了，失败原因： %s\n", err)
			return
		}
		fmt.Printf("输出结果： %s\b", string(output0))

	} else {
		var outputBuf0 bytes.Buffer
		for {
			tempOutput := make([] byte, 5)
			n, err := stdout0.Read(tempOutput)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Printf("从管道读取数据失败了，失败原因： %s\n", err)
					return
				}
			}
			if n > 0 {
				outputBuf0.Write(tempOutput[:n])
			}
		}
		fmt.Printf("输出结果： %s\b", outputBuf0.String())
	}

}



func runCmdWithPipe() {
	fmt.Println("执行命令 `ps aux | grep apipe`:")
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")
	var outputBuf1 bytes.Buffer
	cmd1.Stdout = &outputBuf1
	if err := cmd1.Start(); err != nil {
		fmt.Printf("启动命令失败了，失败原因: %s\n", err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Printf("等待第一条命令执行失败，失败原因： %s\n", err)
		return
	}
	//fmt.Printf("输出结果1: %s\n", outputBuf1.Bytes())
	cmd2.Stdin = &outputBuf1
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	if err := cmd2.Start(); err != nil {
		fmt.Printf("第二条命令启动失败了，失败原因： %s\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil{
		fmt.Printf("等待第二条命令执行失败，失败原因： %s\n", err)
		return
	}
	fmt.Printf("输出结果2: %s\n", outputBuf2.Bytes())
}
