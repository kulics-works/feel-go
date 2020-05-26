package main

import "fmt"

type work = func(string)

func main() {
	testGo()
	test(1, "2")
	testJudge()
	testLoop()
	m := man{}
	testProtocol(&m)
	testLambda(func(a int) (b string) {
		return ""
	})
}

const pi = 3.14
const eight int = 8

var num1 = 1
var num2 float64 = 12.345678

func test(x int, y string) (r1 string, r2 int) {
	a := x * 3
	b := 2
	return "hello", a + b
}
func testJudge() {
	a := 5
	if 1+1 == 2 {
		a = 2
	} else if 2*3 == 6 {
		a = 1
	} else {
		a = 5
	}
	switch a {
	case 1:
		{
			a = a + 1
		}

	case 2:
		{
			a = a + 2
		}
	case 3:
		{
			a = a + 2
		}
	case 4:
		{
			a = a + 2
		}

	default:
		{
			a = 0
		}
	}
}
func testLoop() {
	a := 0
	arr := []int{1, 2, 3, 4, 5}
	for _, i := range arr {
		a += i
	}
	for i, v := range arr {
		a += i + v
	}
	dic := map[string]int{"1": 1, "2": 2}
	for _, i := range dic {
		a += i
	}
	for i := 0; i <= 10; i += 1 {
		a += i
		if i == 7 {
			continue
		}
	}
	for i := 1; i < 3; i += 1 {
		println(i)
	}
	for i := 3; i > 1; i -= 1 {
		println(i)
	}
	for i := 5; i >= 0; i -= 1 {
		a += i
	}
	for a > 0 {
		a -= 1
		break
	}
}
func test_type(a []int, b map[string]int, c chan int) {
}

type human struct {
	name string
}

func (me *human) run() {
}
func (me *human) sayName() (n string) {
	return me.name
}

type man struct {
	human
	age int
}

func (me *man) doSomething(work string) {
	fmt.Println(work)
}

type person interface {
	sayName() (n string)
}
type worker interface {
	person
	doSomething(work string)
}

func testProtocol(w worker) (i interface{}) {
	w.doSomething("protocol")
	return w
}
func test_define() {
	a := 1
	if a == 1 {
		b := 2
		b = 3
		a = 3
		fmt.Println(b)
	} else if a == 2 {
		b := 2
		b = 3
		if b == 2 {
			a = 1
		}
		fmt.Println(b)
	} else {
		b := 2
		b = 3
		a = 3
		fmt.Println(b)
	}
	switch a {
	case 1:
		{
			b := 2
			b = 3
			a = 3
			if a == 3 {
				b = 1
				a = 1
			} else {
				a = 2
			}
			fmt.Println(b)
		}

	default:
		{
			b := 2
			b = 3
			a = 3
			fmt.Println(b)
		}
	}
}
func testGo() (v int) {
	channel := make(chan int)
	async := func() {
		fmt.Println("async")
		channel <- 2
	}
	go async()
	return <-channel
}
func testLambda(fn func(int) string) {
	fn(1)
}
