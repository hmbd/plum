package main

import (
	"math/rand"
	"time"
	"fmt"
)

func mirroredQuery() string {
	responses := make(chan string, 3)

	go func() { responses <- request("asia.gopli.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()

	return <-responses
}

func request(hostname string) (response string) {
	rand.Seed(time.Now().Unix()) // 不给时间戳种子会导致每次生成的都是一样的值
	sleepTime := rand.Intn(3) + 1
	fmt.Printf("hostname: %s, sleep time: %d\n", hostname, sleepTime)
	time.Sleep(time.Second * time.Duration(sleepTime))
	return hostname
}

func main() {
	result := mirroredQuery()
	fmt.Println(result)
}
