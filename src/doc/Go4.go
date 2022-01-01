package main

import "fmt"

/**
 跟踪例程可针对反跟踪例程设置实参
entering: b
in b
entering: a
in a
leaving: a
leaving: b
*/
func trace(s string) string {
	fmt.Println("entering:", s) // 1
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b") // 2
	a()
}

func main() {
	b()
}