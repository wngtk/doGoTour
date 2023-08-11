package main

import "fmt"

// Go's return values may be named. If so, they are treated as variables **defined at the top of the function**.
// return values 是命名为 x, y 相当于在函数开头定义了 var x, y int
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
