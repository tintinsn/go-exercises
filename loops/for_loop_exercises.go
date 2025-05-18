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

// Exercise 7: Prime Number from 2 to n
/* 
 * condition of prime number:
 - divisivle by 1 and itself
 - 1 it not prime number

 * Approach:
 - loop through number from 2 to n.
 - for each number, assume it's prime number until proven otherwise.
 - check if number is divisible by any number from 2 to n - 1.
 	- If divisivle, it not a prime number -> break
 	- If not divisible by any number -> it's a prime number ->  print the number.
*/
func primeNumbers(num int) {
	for i:=2; i<= num; i++ { 
		isPrime := true 
		for j:=2; j<i; j++ {
			if i % j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}


func main() {
	primeNumbers(1)
}