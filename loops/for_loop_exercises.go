package main

import "fmt"

// * Beginner

// Exercise 1: Print numbers
// func printNumber(num int) {
// 	for i:=1; i<= num; i++ {
// 		fmt.Println(i)
// 	}
// }

// Exercise 2: Print Even number
func evenNumber(num int) {
	for i:=0; i<= num; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	} 
}



func main() {
	evenNumber(15)
}