package ch01

import "fmt"

func Variables() {

	var x int64 = 10

	var y float64 = 10.2

	var z string = "Go Programming"

	//type casting & the implict variable declaration
	sum := x + int64(y)

	//String concatination
	fmt.Println(z + " with varibales ")

	//Adding string and integer
	fmt.Println(z + string(x))

	//Summation of int and type converted float
	fmt.Println(sum)
}
