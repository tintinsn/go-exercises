package main

import "fmt"

// * Beginner

// Exercise 1: Print numbers
func printNumber(num int) {
	for i:=1; i<= num; i++ {
		fmt.Println(i)
	}
}

func main() {
	printNumber(10)

}