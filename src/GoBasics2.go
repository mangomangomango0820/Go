// @Time    : 2021.11.6
// @Author  : Xueshan Zhang
// @File    : F:\Software\go\go\Project\src\GoBasics2.go

package main

import "fmt"

func getVal(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num2 - num1
	return sum, sub
}

func main() {
	sum, sub := getVal(20, 10)
	fmt.Println("sum = ", sum, "; sub = ", sub)
	sum, _ = getVal(5, 10) // only keep 1 var
	fmt.Println("sum = ", sum)
}

// Results:
/*
PS F:\Software\go\go\Project> go run F:\Software\go\go\Project\src\GoBasics2.go
sum =  30 ; sub =  -10
sum =  15
*/
