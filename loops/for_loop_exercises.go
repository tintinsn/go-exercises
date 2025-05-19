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

// Exercise 8: Star Pyramid

/* 
 Ex. n = 5
 row number(i)		spaces(j)		stars(k)
 i = 1 				4				1
 i = 2 				3  				3
 i = 3 				2  				5
 i = 4 				1  				7
 i = 5   			0				9

 - Spaces = n - i
 - Stars =  (2 * i) - 1
 
 * condition:
 	- number of pyramid row equals the input values
 	- each row start with spaces with n - i for align star to center 
 	- each row contains (2*i - 1) star

 * Approach:
	- Use 3 nested loops:
		1. outer loop: control the number of row from 1 to n
		2. first inner loop: print (n-1) spaces without a newline.
		3. second inner loop: print (2*i - 1) star without a newline.
	- after print each row use Println() to newline.

*/


func starPyramid(n int) {
	// number of row
	for i:=1; i<= n; i++ { 
		// print spaces
	 	for j:=1; j<= n-i; j++ {  
		fmt.Print(" ")
		}

		// print star
		for k:=1; k<= (2*i) -1; k++ {
		fmt.Print("*")
		}

		// new line
	 	fmt.Println()
	}
}

// Exercise 8: FibonacciSequence

/* 
 What is Fibonacci?
	- each number is the sum of two previous number.
	- Example. 0, 1, 1, 2, 3, 5, 8, 13, ...

 * condition:
	- first and second number are 0, 1
	- Every number is sum of previous two number.
	- How many times should this loop run? --> n - 2 because the first two number are already printed.

 * Approach:
	- Initialize two variables a = 0, b = 1.
	- Print the first two numbers.
	- Use a loop to run  (n-2) times:
		- Calculate the next number by --> next = a+b
		- Update values -> a = b, b = next
		- Print  the next number.
*/

func fibonacciSequence(n int) {
	a := 0
	b := 1

	fmt.Print(a,b, " ")
	for i:=3; i<= n; i++ {
		next := a+b
		a = b
		b = next
		fmt.Print(next, " " )
	}
}


func main() {
	fibonacciSequence(8)
}