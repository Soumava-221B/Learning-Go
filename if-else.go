package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter an integer: ")
    fmt.Scan(&num)

    if num%2 == 0 {
       return fmt.Sprintf("%d is even number", num)
    } 
        return fmt.Sprintf("%d is odd number", num)
    }
