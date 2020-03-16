package function

import (
	"errors"
	"fmt"
)

// Hello : コンソールにhelloを出力
func Hello() {
	fmt.Println("hello")
}

// Add : 整数2つを受け取り、足した結果を返す
func Add(x, y int) int {
	return x + y
}

// Div : 整数x, yを受け取り、xをyで割った結果を返す yが0の時はエラー
func Div(x, y int) (float64, error) {
	if y == 0 {
		return 0, errors.New("divide by zero")
	}

	return float64(x) / float64(y), nil
}
