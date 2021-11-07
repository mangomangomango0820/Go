// @Time    : 2021.11.6
// @Author  : Xueshan Zhang
// @File    : F:\Software\go\go\Project\src\GoBasics3.go

package main

import "fmt"

// 1. clarify global variable
//    it's acceptable if not all are used later
var (
	g0 = 2.5
	g1 = 8
)

func main() {
	// 1. clarify local variable
	//    while not be initialized with new value, then use default value
	fmt.Println("*** 1")
	fmt.Println("g0 = ", g0) // acceptable: here do not use global variable g1
	var i int
	fmt.Println("i0 = ", i)
	i = 5
	fmt.Println("i1 = ", i)
	var variable string
	fmt.Print("variable =", variable, "\n")
	fmt.Println()

	// 2. judge variable type based on initialized value
	fmt.Println("*** 2")
	var num = 7.1
	fmt.Println("num =", num)
	fmt.Println()

	// 3. clarify a new variable without 'var'
	fmt.Println("*** 3")
	name := "tom"
	fmt.Println("name =", name)
	fmt.Println()

	// 4. clarify multiple variable
	fmt.Println("*** 4")
	var n1, n2, n3 int
	fmt.Println("n1 =", n1, "; n2 =", n2, "; n3 =", n3)
	var n4, sex = 3, "female"
	fmt.Println("n4 =", n4, "; sex =", sex)
	//    clarify multiple new variables          // both n5 and date are new variables
	n5, date := 4, "11/7"
	fmt.Println("n5 =", n5, "; date =", date)
	fmt.Println()

	// Results:
	/*
		PS F:\Software\go\go\Project> go run .\src\GoBasics3.go
		*** 1
		g0 =  2.5
		i0 =  0
		i1 =  5
		variable =

		*** 2
		num = 7.1

		*** 3
		name = tom

		*** 4
		n1 = 0 ; n2 = 0 ; n3 = 0
		n4 = 3 ; sex = female
		n5 = 4 ; date = 11/7
	*/

	// 5. variable type cannot be changed
	//var j int = 3
	//j = 3.2                      // j was int, so j cannot be float now
	//fmt.Print("j = ", j)

	// Results:
	/*
		# command-line-arguments
		src\GoBasics3.go:12:4: constant 3.2 truncated to integer
	*/

	// 6. variable cannot be named again
	//j := 99
	//fmt.Print("j = ", j)

	// Results:
	/*
		# command-line-arguments
		src\GoBasics3.go:13:4: no new variables on left side of :
	*/
}
