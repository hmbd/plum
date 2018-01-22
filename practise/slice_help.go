package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d, len=%d\t%v\n", i, cap(y), len(y), y)
		x = y
	}
	s := []int{1,2,3,4,5}
	fmt.Println(remove(s, 2))
	fmt.Println(remove1(s, 2))

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)
}

func appendInt(x []int, y ... int) []int {
	var z []int // 声明一个切片
	zlength := len(x) + 1
	if zlength <= cap(x) {
		z = x[:zlength]
	} else {
		zcap := zlength
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlength, zcap)
		copy(z, x)
	}
	//z[len(x)] = y
	copy(z[len(x):], y)
	return z
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i ++
		}
	}
	// 在原有slice内存空间之上返回不包含空字符串的列表：
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove1(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
