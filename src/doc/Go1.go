/*
	 regexp 包为正则表达式实现了一个简单的库。

	该库接受的正则表达式语法为：

	正则表达式:
		串联 { '|' 串联 }
	串联:
		{ 闭包 }
	闭包:
		条目 [ '*' | '+' | '?' ]
	条目:
		'^'
		'$'
		'.'
		字符
		'[' [ '^' ] 字符遍历 ']'
		'(' 正则表达式 ')'
*/
package main

import (
	"fmt"
)

type T struct {
	name  string // 对象名
	value int    // 对象值
}

func main() {

	//owner := obj.Owner()
	//if owner != user {
	//	obj.SetOwner(user)
	//}

	var x, y int = 1, 2
	if x < y {
		fmt.Println(x)
	} else {
		fmt.Println(y)
	}

	if x = 3; x > y {
		fmt.Println(x)
	} else {
		fmt.Println(y)
	}

	//f, err := os.Open("name")
	//d, err := f.Stat()

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	oldMap := map[string]string{"Enone": "no error1", "Eio": "Eio1"}
	newMap := map[string]string{}
	for key, value := range oldMap {
		newMap[key] = value
	}
	fmt.Println(newMap)
	fmt.Println(oldMap)

	//for key := range oldMap{
	//	if key.expired() {
	//		delete(oldMap,key)
	//	}
	//}
	for _, value := range oldMap {
		fmt.Println(value)
	}

	for pos, char := range "学英语才能救中国\x80啊啊" {
		fmt.Printf("字符 %#U 始于字节位置 %d \n", char, pos)
	}

	//for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	//	a[i],a[j] = a[j],a[i]
	//}

}
