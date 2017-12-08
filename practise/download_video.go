package main

import (
	"regexp"
	"sync"
	"net/http"
	"io/ioutil"
	"os"
	"bufio"
	"io"
	"fmt"
	"strings"
	"time"
	"path"
	"path/filepath"
	"log"
)

var indexUrl = "https://www.ted.com/" // 视频网址,自动识别数据类型
/*
 href="/talks/martina_flor_the_secret_language_of_letter_design?language=es" title="Martina Flor: The secret language of letter design">

href="/talks/martina_flor_the_secret_language_of_letter_design"
    title="The secret language of letter design"
*/
var patterIndex = regexp.MustCompile(`href="(.+)"[\t\s]*title="(.+)"`) // 视频 播放 地址
/*
{"uri":"https://download.ted.com/talks/RayDalio_2017-950k.mp4?apikey=489b859150fc58263f17110eeb44ed5fba4a3b22","filesize_bytes":118773264,"mime_type":"video/mp4"
*/
//var patterVideo = regexp.MustCompile(`{"uri":"([\w/\?_\-:]+)","filesize_bytes":(\d+),"mime_type":"video/mp4"}`) // 视频 下载 地址
var patterVideo = regexp.MustCompile(`{"uri":"(.+?\.mp4\?apikey=\w+)","filesize_bytes":(\d+),"mime_type":"video\/mp4"}`) // 视频 下载 地址
var wg sync.WaitGroup                                                                                                    // 等待组，提供给线程使用

type DownLoadList struct {
	Data map[string][]int64 // 下载进度
	Lock sync.Mutex         // 锁
}

/*
查询当前目录
TODO: 查询出来是 /tmp, 确认下为什么？
*/
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

/*
param:
	url 网页链接

return:
	content 网页内容
	statusCode 访问url网页连接的状态码
通过url获取网页的源码，并返回code标识请求是否成功
*/
func getUrlContent(url string) (content string, statusCode int) {
	response, err1 := http.Get(url)
	if err1 != nil {
		statusCode = 202
		return
	}
	defer response.Body.Close()
	data, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		statusCode = 202
		return
	}
	statusCode = response.StatusCode
	content = string(data)
	return
}

/*
检查文件是否存在
param:
	filename: 文件路径

return:
	exists: 文件是否存在
	isDir: 文件是否是目录
*/
func checkFileIsExist(filename string) (exists bool, isDir bool) {
	exists = true
	f, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			exists = false
		} else {
			// 非文件不存在错误
			panic(err)
		}
	}

	isDir = false
	if exists {
		if f.IsDir() {
			isDir = true
		}
	}
	return exists, isDir
}

/*
读取文件的内容
param:
	filename: 文件路径
return:
	content: 文件内容
*/
func readFileContent(filename string) (content string) {
	fmt.Printf("读取文件 %s 内容\n", filename)
	fi, err := os.Open(filename)
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
	content = string(chunks)
	return

}

/*
param:
	filename: 文件名
	url: 要写入的字符串, 视频链接

return:
	true: 写入成功  false: 写入失败
*/
func writeContentFile(filename string, url string) bool {
	existsFile, isDir := checkFileIsExist(filename)
	if isDir {
		fmt.Printf("%s必须是文件类型，不能是目录\n", filename)
		os.Exit(1)
	}
	fmt.Println("文件中追加内容: ", filename)
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	exists := false
	result := true
	if existsFile {
		data := readFileContent(filename)
		exists = strings.Contains(data, url)
	}

	if !exists {
		fmt.Println("url不存在，写入文件中")
		url += "\n"
		_, err := f.WriteString(url)
		if err != nil {
			result = false
		}
	}
	f.Sync()
	return result
}

/*
打印下载进度
*/
func (downloadList *DownLoadList) printProgress() {
	for {
		downloadList.Lock.Lock()
		for key, arr := range downloadList.Data {
			fmt.Printf("%s progress: [%-50s] %d%% Done\n", key, strings.Repeat("#", int(arr[0]*50/arr[1])), arr[0]*100/arr[1])
		}
		downloadList.Lock.Unlock()
		time.Sleep(time.Second * 3)
		fmt.Printf("\033[2J")
	}
}

/*
下载视频函数

param:
	url: 视频下载地址
	filename: 文件名
	downList:

return:
	true: 下载成功 false: 下载失败了
*/
func DownLoadVideo(url string, filename string, downList *DownLoadList) bool {
	buffer := make([]byte, 1024)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("创建文件%s失败了\n", filename)
		return false
	}
	defer f.Close()
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("获取视频资源失败, 链接： %s\n", url)
	}
	defer response.Body.Close()
	bufferRead := bufio.NewReader(response.Body)
	for {
		index, err := bufferRead.Read(buffer)
		if err == io.EOF {
			break
		}
		f.Write(buffer[:index])
		fileInfo, err := os.Stat(filename)
		fileSize := fileInfo.Size()
		downList.Lock.Lock()
		downList.Data[filename] = []int64{fileSize, response.ContentLength}
		downList.Lock.Unlock()
	}
	wg.Done()
	return true
}

func initDownPath(filePath string) {
	exists, isDir := checkFileIsExist(filePath)
	if exists {
		if !isDir {
			fmt.Printf("%s 文件已经存在，且格式非目录类型", filePath)
			os.Exit(1)
		}
	} else {
		os.Mkdir(filePath, 0755)
	}
}
func downLoadMain() {
	var downLoadList DownLoadList
	downLoadList.Data = make(map[string][]int64)
	downList := &downLoadList

	//currPath := getCurrentDirectory()
	currPath, err := os.Getwd()
	if err != nil {
		fmt.Println("查询当前所在目录失败")
		os.Exit(1)
	}
	downPath := path.Join(currPath, "go-video")
	//downPath := "/tmp/go-video"
	initDownPath(downPath)
	// 获取网页内容
	content, statusCode := getUrlContent(indexUrl)
	fmt.Printf("statusCode: %d\n", statusCode)

	if statusCode != 200 {
		fmt.Println("获取网页内容失败了")
		return
	}

	htmlResult := patterIndex.FindAllStringSubmatch(content, -1)
	length := len(htmlResult)
	if length == 0 {
		fmt.Printf("url： %s 中没有找到视频的链接\n", indexUrl)
		return
	}
	go downList.printProgress()

	// 只下载一个视频
	if length > 1 {
		length = 1
	}

	// 记录下载url的文件路径
	filename := path.Join(downPath, "url.txt")
	exists, isDir := checkFileIsExist(filename)
	if exists && !isDir {
		os.Remove(filename)
	} else if exists && isDir{
		fmt.Printf("文件 %s 必须是文本的格式\n", filename)
		os.Exit(1)
	}

	// 对视频播放列表的每一个视频进行解析
	for i := 0; i < length; i ++ {
		videoInfo := htmlResult[i]
		newUrl := videoInfo[1]
		exists := strings.Contains(newUrl, indexUrl)
		if !exists {
			newUrl = indexUrl + newUrl
		}
		title := videoInfo[2]
		title = strings.Replace(title, " ", "_", -1)
		fmt.Printf("title: %s, url: %s\n", title, newUrl)
		context, videoStatus := getUrlContent(newUrl)
		if videoStatus != 200 {
			fmt.Printf("视频url: %s 访问失败了\n", newUrl)
			continue
		}

		videoResult := patterVideo.FindAllStringSubmatch(context, -1)
		length := len(videoResult)
		ok := writeContentFile(filename, newUrl)
		fmt.Printf("视频数量： %d, 记录是否成功： %t\n", length, ok)
		if length > 0 && ok {
			videoUrl := videoResult[0][1] // 下载url
			fileSize := videoResult[0][2] // 文件大小

			// 找到大小最大(即清晰度最高)的视频
			for i := 0; i < length; i ++ {
				if videoResult[i][2] > fileSize {
					videoUrl = videoResult[i][1]
				}
			}
			wg.Add(1)
			filePath := path.Join(downPath, title)
			fmt.Printf("videoUrl: %s, fileSize: %s, filePath: %s\n", videoUrl, fileSize, filePath)
			go DownLoadVideo(videoUrl, filePath, downList)
		} else {
			fmt.Println("没有视频需要下载")
		}
	}
	wg.Wait()
	return
}
func main() {
	downLoadMain()
}
