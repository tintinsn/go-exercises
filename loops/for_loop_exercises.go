package main

import "fmt"

// * Easy

// Exercise 1: Print numbers
func printNumber(num int) {
	for i:=1; i<= num; i++ {
		fmt.Println(i)
	}
}

// Exercise 2: Print Even number
func evenNumber(num int) {
	for i:=0; i<= num; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	} 
}

// Exercise 3: Sum Of Number from 1 to n
func sumOfNumber(num int) {
	sum := 0
	for i:=1; i<= num; i++ {
		sum += i
	}
	fmt.Printf("Sum of 1 to %d = %d",num, sum)
}

func main() {
	sumOfNumber(100)
}