package main

import (
	"bytes"
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
	"io"
	"math/rand"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMITER       = '\t'
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go serverGo()
	time.Sleep(500 * time.Millisecond)
	go clientGo(1)
	wg.Wait()
}

func printLog(role string, sn int, format string, args ...interface{}) {
	if !strings.HasSuffix("format", "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]: %s", role, sn, fmt.Sprintf(format, args...))
}

func printServerLog(format string, args ...interface{}) {
	printLog("Server", 0, format, args...)
}

func printClientLog(sn int, format string, args ...interface{}) {
	printLog("Client", sn, format, args...)
}

func strToInt32(str string) (int32, error) {
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("强制类型转换成 int32 报错： %s", err)
	}

	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0, fmt.Errorf("数据： %s 不在 int32 范围内", num)
	}
	return int32(num), nil
}

func cbrt(param int32) float64 {
	return math.Cbrt(float64(param))
}

func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}

		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}

func serverGo() {
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerLog("注册监听失败: %s", err)
		return
	}

	defer listener.Close()

	printServerLog("注册监听成功，本地地址: %s", listener.Addr())
	for {
		conn, err := listener.Accept() // 阻塞到有连接进来
		if err != nil {
			printServerLog("阻塞失败： %s", err)
		}
		printServerLog("客户端(远程地址: %s)连接成功", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
		wg.Done()
	}()

	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printServerLog("连接已经关闭了")
			} else {
				printServerLog("读取数据错误： %s", err)
			}
			break
		}
		printServerLog("接收到客户端的数据: %s", strReq)

		intReq, err := strToInt32(strReq)
		if err != nil {
			n, err := write(conn, err.Error())
			printServerLog("发送信息错误： %s ( %d bytes)", err, n)
			continue
		}
		floatResp := cbrt(intReq)
		respMsg := fmt.Sprintf("%d 的立方根是 %f", intReq, floatResp)
		n, err := write(conn, respMsg)
		if err != nil {
			printServerLog("写入数据错误: %s", err)
		}
		printServerLog("发送结果给客户端: %s (%d bytes)", respMsg, n)
	}
}

func clientGo(id int) {
	defer wg.Done()

	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)
	if err != nil {
		printServerLog("连接失败： %s", err)
		return
	}
	defer conn.Close()

	printClientLog(id, "连接服务器(远程地址： %s, 本地地址: %s)", conn.RemoteAddr(), conn.LocalAddr())
	time.Sleep(200 * time.Millisecond)
	requestNumber := 5
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))
	for i := 0; i < requestNumber; i++ {
		req := rand.Int31()
		n, err := write(conn, fmt.Sprintf("%d", req))
		if err != nil {
			printClientLog(id, "写入数据错误： %s", err)
			continue
		}
		printClientLog(id, "发送数据给服务端： %d (%d bytes)", req, n)
	}

	for j := 0; j < requestNumber; j++ {
		strResp, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printClientLog(id, "连接关闭了")
			} else {
				printClientLog(id, "读取数据错误： %s", err)
			}
			break
		}
		printClientLog(id, "接收到返回的数据: %s", strResp)
	}
}
