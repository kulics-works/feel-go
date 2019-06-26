package main

import . "fmt"

func main() {
	test_go()
	test(1, "2")
	test_judge()
	test_loop()
	m := man{}
	test_protocol(&m)
	test_lambda(func(a int) (b string) {
		return ""
	})
}

const pi = 3.14
const eight int = 8

var num1 = 1
var num2 float64 = 12.345678

type work = func(string)
type do_work func(string)

func test(x int, y string) (r1 string, r2 int) {
	a := x * 3
	var b int = 2
	return "hello", a + b
}
func test_judge() {
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
	default:
		{
			a += 0
		}
	}
}
func test_loop() {
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

func (me *human) say_name() (n string) {
	return me.name
}

type man struct {
	human
	age int
}

func (me *man) do_something(work string) {
	Println(work)
}

type person interface {
	say_name() (n string)
}
type worker interface {
	person
	do_something(work string)
}

func test_protocol(w worker) (i interface{}) {
	w.do_something("protocol")
	return w
}
func test_go() (v int) {
	async := func() { Println("async") }
	go async()
	channel := make(chan int, 1)
	channel <- 2
	return <-channel
}
func test_lambda(fn func(int) string) {
	fn(1)
}
