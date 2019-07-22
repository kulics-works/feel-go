package main

import . "fmt"

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

type work = func(string)

func test(x int, y string) (r1 string, r2 int) {
	a := x * 3
	var b int = 2
	return "hello", a + b
}
func testJudge() {
	if 1+1 == 2 {
	} else if 2*3 == 6 {
	} else {
	}
	a := 5
	switch a {
	case 1:
		{
			a += 1
		}

	case 2:
		{
			a += 2
		}
	case 3:
		{
			a += 2
		}
	case 4:
		{
			a += 2
		}

	default:
		{
			a += 0
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
	for i := 5; i > 0; i -= 1 {
		a += i
	}
	for a > 0 {
		a -= 1
	}
	for {
		break
	}
}

type human struct {
	name string
}

func (me *human) sayName() (n string) {
	return me.name
}

type man struct {
	human
	age int
}

func (me *man) doSomething(work string) {
	Println(work)
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
func testGo() (v int) {
	async := func() { Println("async") }
	go async()
	channel := make(chan int, 1)
	channel <- 2
	return <-channel
}
func testLambda(fn func(int) string) {
	fn(1)
}
func test_define() {
	a := 1
	if a == 1 {
		b := 2
		b := 3
		a := 3
	} else if a == 2 {
		b := 2
		b := 3
		if b == 2 {
			a := 1
		}
	} else {
		b := 2
		b := 3
		a := 3
	}
	switch a {
	case 1:
		{
			b := 2
			b := 3
			a := 3
			if a == 3 {
				b := 1
				a := 1
			} else {
				a := 2
			}
		}

		{
			c := 2
			a := 1
		}

	default:
		{
			b := 2
			b := 3
			a := 3
		}

	}
}
