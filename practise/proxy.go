package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"

	docker "github.com/fsouza/go-dockerclient"
)
// https://github.com/UESTC-BBS/socket-tcp-proxy/blob/master/proxy.go

var VERSION = "0.0.0-src"

var count uint64

type Conf struct {
	Logfile string `json:"logfile"`
	Proxy []Proxy `json:"proxy"`
}

type Proxy struct{
	Docker string `json:"docker"`
	Socket string `json:"socket"`
	Port string `json:"port"`
	DockerIp string `json:"docker_ip"`
	DockerAddr string `json:"docker_addr"`
}

func (p *Proxy) StartForward(){
	log.Println("[INFO] 转发 " + p.Socket + " 给 " + p.DockerAddr)
	listen, err := net.Listen("unix", p.Socket)
	exec.Command("chmod", "777", p.Socket).Run()
	if err != nil{
		log.Fatal(err)
	}
	for {
		uconn, err := listen.Accept()
		if err != nil{
			log.Println("[ERROR] 发送 " + p.Socket + " 给 " + p.DockerAddr + " 发生了错误：  " + err.Error())
			continue
		}
		go forward(p, uconn)
	}
}

func forward(p *Proxy, uconn net.Conn){
	id := atomic.AddUint64(&count, 1)

	tconn, err := net.Dial("tcp", p.DockerAddr)

	if err != nil{
		log.Printf("[ERROR] 本地连接TCP出错： %s" + p.Socket + " 远程地址: " + p.DockerAddr + "\n", err)
		return
	}
	log.Printf("[%d] 连接来自： " + p.Socket + "和" + p.DockerAddr, id)

	var wg sync.WaitGroup

	go func(uconn net.Conn, tconn net.Conn){
		wg.Add(1)
		defer wg.Done()
		io.Copy(uconn, tconn)
		uconn.Close()
	}(uconn, tconn)

	go func(uconn net.Conn, tconn net.Conn){
		wg.Add(1)
		defer wg.Done()
		io.Copy(tconn, uconn)
		tconn.Close()
	}(uconn, tconn)

	wg.Wait()
}

func ReadToString(filePath string)(string, error){
	b, err := ioutil.ReadFile(filePath)
	if err != nil{
		return "", err
	}

	return string(b), nil
}

var conf Conf
var configFilePath string
var debug bool


var dockerclient *docker.Client

func init(){
	debug = (runtime.GOOS == "darwin")
	configFilePath = "/etc/socket-proxy/proxy.json"
	jsonString, err := ReadToString(configFilePath)
	if err != nil{
		if !debug{
			log.Fatal(err)
		}
	}
	json.Unmarshal([]byte(jsonString), &conf)

	dockerclient, err = docker.NewClientFroEnv()
	if err != nil{
		if !debug{
			log.Fatal(err)
		}
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main(){
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	logFile, err := os.OpenFile(conf.Logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil && !debug{
		log.Println("打开文件错误： ", err)
		os.Exit(1)
	}else if !debug{
		log.SetOutput(logFile)
	}
	defer logFile.Close()

	for k, v := range conf.Proxy{
		name, err := dockerclient.InspectContainer(v.Docker)
		if err != nil{
			log.Fatal(err)
		}
		log.Println(v.Docker, name.NetworkSettings.IPAdderss)
		v.DockerIp = name.NetworkSettings.IPAddress
		v.DockerAddr = name.NetworkSettings.IPAddess + ":" + v.Port
		conf.Proxy[k] = v
	}
	log.Println(conf.Proxy)
	for _, v := range conf.Proxy{
		go v.StartForward()
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP, syscall.SIGTERM)
	<- ch
	for _, v := range conf.Proxy{
		os.Remove(v.Socket)
	}
	log.Println("关闭监听信号")
	os.Exit(0)
}
