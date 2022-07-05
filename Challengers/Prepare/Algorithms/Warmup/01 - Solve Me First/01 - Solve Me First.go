package main

import "fmt"

func solveMeFirst(a uint, b uint) uint {
	var res uint = 0
	res = a + b
	return res
}

func main() {
	var a, b, res uint
	fmt.Scanf("%v\n%v", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
}
