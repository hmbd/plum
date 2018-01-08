package main

import "fmt"

// 下列情况可以考虑使用指针：
// 		1. 需要改变参数的值
// 		2. 避免复制操作
// 		3. 节省内存

type abc struct {
	v int
}

type S map[string][]string

func Summary1(params string) (s *S) { // 等价,函数中不使用 s func Summary(params string) *S{
	s = &S{
		"name":           []string{params},
		"focus(project)": []string{"UE", "AgileMethodology", "SoftwareEngineering"},
		"hobby(life)":    []string{"Basketball", "Movies", "Travel"},
	}
	return s
}

func Summary2(params string) (s S) {
	s = S{
		"name":           []string{params},
		"profrssion":     []string{"Javaprogrammer", "ProjectManager"},
		"interest(lang)": []string{"Clojure", "Python", "Golang"},
	}
	return s
}
func (a abc) test1() { // 传入的值，不能修改到原值
	a.v = 1
	fmt.Printf("test1: %d\n", a.v)
}

func (a *abc) test2() { // 传入的是引用，可以修改原值
	fmt.Printf("test2 before: %d\n", a.v)
	a.v = 2
	fmt.Printf("test2 after: %d\n", a.v)
}

func (a *abc) test3() { // 传入的是引用，可以修改原值
	fmt.Printf("test3: %d\n", a.v)
}

func pointer_test() {
	var p *int // 空指针，输出为 nil
	fmt.Printf("空指针 p: %v\n", p)

	var i int
	p = &i // 指向局部变量，变量初始值为 变量的值
	fmt.Printf("指向变量后值 p: %v, *p: %v\n", p, *p)

	*p = 4 // 通过指针修改原变量数值
	fmt.Printf("修改后值 p: %v, *p: %v i: %d \n", p, *p, i)

	m := [3]int{3, 4, 5} // 数组的初始化
	fmt.Printf("数组m: %v---%d, %d, %d\n", m, m[0], m[1], m[2])

	x := [3]*int{&m[0], &m[1], &m[2]} // 指针数组的初始化
	fmt.Printf("指针数组1 x: %v---%d, %d, %d\n", x, x[0], x[1], x[2])
	fmt.Printf("指针数组2 *x: %v---%d, %d, %d\n", x, *x[0], *x[1], *x[2])

	var n [3] *int
	n = x
	fmt.Printf("指针数组3 n: %v---%d, %d, %d\n", n, n[0], n[1], n[2])

	y := []*[3]*int{&x} // 指向数组的指针(即二级指针)
	fmt.Printf("指向数组的指针1 y: %v, %d\n", y, y[0])
	fmt.Printf("指向数组的指针2 *y[0]: %v, %d\n", y, *y[0])
	fmt.Printf("指向数组的指针3 *y[][]: %v, %v, %d\n", *y[0][0], *y[0][1], *y[0][2])
}

type Student struct {
	name  string
	id    int
	score uint
}

func memery_test() {
	p := new(Student) // new 分配出来的数据是指针形式
	p.name = "China"
	p.id = 12334
	p.score = 99
	fmt.Println("结构体指针 p: ", *p)

	var st Student // var 定义的变量是数值形式
	st.name = "Chinese"
	st.id = 5678
	st.score = 100
	fmt.Println("结构体对象 st: ", st)

	var ptr *[]Student // make 分配 slice、map、channel的空间，并且返回的不是指针
	fmt.Println("声明的指针 ptr: ", ptr)
	ptr = new([]Student)
	fmt.Println("new 声明的指针 ptr: ", ptr)
	*ptr = make([]Student, 3, 100)
	fmt.Println("make 声明的指针 ptr: ", ptr)
	ptr1 := make([]Student, 3, 100)
	fmt.Println("make 声明的指针 ptr1: ", ptr1)

	stu := []Student{{"China", 1234, 66}}
	fmt.Println("直接初始化 stu: ", stu)
}
func main() {
	var i int  // i 的类型是 int 类型
	i = 4      // i 的值为 4
	var p *int // p 的类型是 `int类型的指针`
	p = &i     // p 的值为 `i 的地址`

	fmt.Printf("i=%d, p=%d, *p=%d\n", i, p, *p)

	*p = 2 // *p 的值为 ``i的地址`的指针`(其实就是 i),等价于 i = 2
	fmt.Printf("i=%d, p=%d, *p=%d\n", i, p, *p)

	i = 3
	fmt.Printf("i=%d, p=%d, *p=%d\n", i, p, *p)

	a := abc{} // a := new(abc)
	a.test1()
	a.test2()
	a.test3()

	s1 := Summary1("Harry1")
	fmt.Printf("Summery1 地址: %v\n", s1)
	fmt.Printf("Summery1 内容: %v\n", *s1)

	s2 := Summary2("Harry2")
	fmt.Printf("Summery2 地址: %v\n", &s2)
	fmt.Printf("Summery2 内容: %v\n", s2)

	pointer_test()
	memery_test()
}
