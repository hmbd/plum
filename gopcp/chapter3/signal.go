package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
	"os/signal"
	"strconv"
	"strings"
	"errors"
	"bytes"
	"io"
)

func main() {
	go func() {
		time.Sleep(5 * time.Second)
		sendSignal()
	}()
	handleSignal()
}

func handleSignal() {
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("设置信号通知1: %s\n", sigs1)
	signal.Notify(sigRecv1, sigs1...)
	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("设置信号通知2: %s\n", sigs2)
	signal.Notify(sigRecv2, sigs2...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("接收到信号 1 的信息: %s\n", sig)
		}
		fmt.Printf("信号1 信息结束\n")
		wg.Done()
	}()

	go func() {

		for sig := range sigRecv2 {
			fmt.Printf("接收到信号 2 的信息: %s\n", sig)
		}
		fmt.Printf("信号2 信息结束\n")
		wg.Done()
	}()

	fmt.Println("等待 2 秒")
	time.Sleep(2 * time.Second)
	fmt.Printf("发送停止信号")
	signal.Stop(sigRecv1)
	close(sigRecv1)

	fmt.Printf("信号1 结束\n")
	wg.Wait()
}

func sendSignal() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("程序发生了致命错误: %s\n", err)
			debug.PrintStack()
		}
	}()

	cmds := []*exec.Cmd{
		exec.Command("ps", "aux"),
		exec.Command("grep", "signal"),
		exec.Command("grep", "-v", "grep"),
		exec.Command("grep", "-v", "grep run"),
		exec.Command("awk", "{print $2}"),
	}

	output, err := runCmds(cmds)
	if err != nil {
		fmt.Printf("执行系统命令出错： %s\n", err)
		return
	}

	pids, err := getPids(output)
	if err != nil {
		fmt.Printf("查询Pid 报错： %s\n", err)
		return
	}
	fmt.Printf("Target PID(s): \n%v\n", pids)
	for _, pid := range pids {
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Printf("查找进程失败: %s\n", err)
			return
		}
		sig := syscall.SIGQUIT
		fmt.Printf("发送信号 %s 给 PID(%d)..\n", sig, pid)
		err = proc.Signal(sig)
		if err != nil {
			fmt.Printf("发送信号错误:%s\n", err)
			return
		}
	}
}

func runCmds(cmds []*exec.Cmd) ([]string, error) {
	if cmds == nil || len(cmds) == 0 {
		return nil, errors.New("无效的系统命令列表")
	}
	first := true
	var output []byte
	var err error
	for _, cmd := range cmds {
		fmt.Printf("执行命令： %v\n", getCmdPlaintext(cmd))
		if !first {
			var stdinBuf bytes.Buffer
			stdinBuf.Write(output)
			cmd.Stdin = &stdinBuf
		}
		var stdoutBuf bytes.Buffer
		cmd.Stdout = &stdoutBuf

		if err = cmd.Start(); err != nil {
			return nil, getError(err, cmd)
		}

		if err = cmd.Wait(); err != nil {
			return nil, getError(err, cmd)
		}
		output = stdoutBuf.Bytes()
		if first {
			first = false
		}

	}

	var lines []string
	var outputBuf bytes.Buffer
	outputBuf.Write(output)
	for {
		line, err := outputBuf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, getError(err, nil)
			}
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

func getPids(strs [] string) ([]int, error) {
	var pids []int
	for _, str := range strs {
		pid, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		pids = append(pids, pid)
	}
	return pids, nil
}

func getCmdPlaintext(cmd *exec.Cmd) string {
	var buf bytes.Buffer
	buf.WriteString(cmd.Path)
	for _, arg := range cmd.Args[1:] {
		buf.WriteRune(' ')
		buf.WriteString(arg)
	}
	return buf.String()
}

func getError(err error, cmd *exec.Cmd, extraInfo ...string) error {
	var errMsg string
	if cmd != nil {
		errMsg = fmt.Sprintf("%s [%s %v]", err, (*cmd).Path, (*cmd).Args)
	} else {
		errMsg = fmt.Sprintf("%s", err)
	}
	if len(extraInfo) > 0 {
		errMsg = fmt.Sprintf("%s (%v)", errMsg, extraInfo)
	}
	return errors.New(errMsg)
}
