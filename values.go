package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	float64 n := 7.0/3.0
	fmt.Sprintf("7.0/3.0 = &.1f", n)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
