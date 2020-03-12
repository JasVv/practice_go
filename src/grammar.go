package main

import (
	"fmt"
	"strconv"
)

var str string = "hello world!"
var x int = 10

const y int = 20

func main() {

	fmt.Println(str)

	// 変数宣言を関数の内部で行う場合は，varと型宣言を省略し，:= と書ける。
	z := x * y

	fmt.Println(z)

	// if文
	xMessage, yMessage := "x!", "y!"
	if x > y {
		fmt.Println(xMessage)
	} else {
		fmt.Println(yMessage)
	}

	// for文
	for i := 1; i <= 3; i++ {
		fmt.Println(strconv.Itoa(i) + "回目")
	}
	n := 4
	for n <= 6 {
		fmt.Println(strconv.Itoa(n) + "回目")
		n++
	}
	for {
		n++
		if n > 10 {
			break
		}

		if n%2 == 0 {
			continue
		}

		fmt.Println(strconv.Itoa(n) + "回目")
	}

}
