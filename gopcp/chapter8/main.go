package main

import (
	"log"
	"sync"
	"os"

	"github.com/seekplum/plum/gopcp/chapter8/thumbnail"
	"path/filepath"
	"strings"
)

// 顺序处理每张图片
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// 并行处理每张图片
// 问题在于没有等图片处理结束程序就退出了
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

// 通过channel发送事件进行计数
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

// 当遇到第一个非你nil的error时将返回到调用方，没有一个goroutine区排空errors channel
// 导致worker goroutine在想这个channel发送值时会永远阻塞而且不会退出，这种情况叫goroutine泄露
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}

	return nil
}

func makeThumbnials5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbFile string
		err       error
	}

	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbFile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, err
		}
		thumbfiles = append(thumbfiles, it.thumbFile)
	}
	return thumbfiles, nil
}

// 通过递增计数器确定goroutine退出，每个goroutine启动时加1，在goroutine退出时减1
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)

	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}

			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

func getCurrentDirectory() string {
	// TODO: 当编译成二进制程序时才生效，IDE debug时为生效？？
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {
	curr := getCurrentDirectory()
	imagePath := filepath.Join(curr, "test.jpg")
	filenames := []string{imagePath}
	makeThumbnails(filenames)
}
