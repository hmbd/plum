package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast            // supports broadcast access capability
	FlagLoopback             // is a loopback interface
	FlagPointToPoint         // belongs to a point-to-point link
	FlagMulticast            // supports multicast access capability
)

const (
	_ = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

func main(){
	fmt.Println(Friday)
	fmt.Println(FlagBroadcast)
	fmt.Println(MiB)
	fmt.Println(1 << 3) // 左移(移一位等价于 乘以 2) 右边补零  00000001(二进制) 补 3 个零 00001000(二进制)
	fmt.Println(16 >> 3) // 右移(移一位等价于 除以 2) 左边补零 00010000(二进制) 补 3 个零 00000010(二进制)
	// 只有常量可以是无类型的
}