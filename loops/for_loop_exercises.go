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

// * Intermediate

// Exercise 4: Print Multiplication Table Of N
func multiplicationTableOfN(num int) {
	
	for i:=1; i<=12; i++ {
		fmt.Printf("%d x %d = %d\n",num,i, i*num)
	
	}
}

// Exercise 5: Reverse Number
func reverseNum(num int) {
	for i:=num; i>=1; i-- {
		fmt.Println(i)
	}
}

// Exercise 6: Squared Number from 1 to n
func squaredNum(num int) {
	for i:=1; i<= num; i++ {
		square := i* i
		fmt.Printf("%d ^ 2 = %d\n",i,square)
	}
}


func main() {
	squaredNum(12)
}