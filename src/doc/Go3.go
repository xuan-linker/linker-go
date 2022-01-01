package main

import "fmt"

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// 像这样使用它们：
func a() {
	trace("a")
	defer untrace("a")
	// 做一些事情....
}
func b() {
	defer untrace("b")
	trace("b")
	// 做一些事情....
}
func main() {
	a()
	b()

}
