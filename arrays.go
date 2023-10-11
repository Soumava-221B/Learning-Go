package main

import "fmt"

func main() {
	var c [5]int
	fmt.Println("emp:", c)

	c[4] = 100
	c[2] = 102
	fmt.Println("set: ", c)
	fmt.Println("get: ", c[4])

	fmt.Println("len: ", len(c))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
	     for j := 0; j < 3; j++ {
	     twoD[i][j] = i + j
     }
	}
	fmt.Println("2D: ", twoD)
	}
