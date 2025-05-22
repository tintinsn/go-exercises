package main

import (
	"fmt"
)

// Ex 1: Compound Interest Calculator

/*
Write a program that calculates compound interest over a number of years.
Given a initiala amount, annual interest rate (as a percent), and number of years,
use a for loop to print the balance at the end of each year.

  * Example
	Input:
		amount = 10000
		rate = 3
		years = 3

	Output:
		Year 1: 10300.00
		Year 2: 10609.00
		Year 3: 10927.27

  * Condition:
  - amount, rate, years must be non-negative
  - rate is given as a percentage
  - interest compounded annually
  - output should be rounded to 2 decimal places

  * Approach:
  - check if amount, rate, year are less than 0 then stop program
  - Use a loop that runs for the number of years provided by the user.
	- for each year, compute the interest based on the current amount and update the amount to
	  include the compounded interest.
  - print the balance at the end of each year.
*/

 func compoundInterestCal(amount float64,rate float64 , year int) {
	if amount <= 0 || rate <= 0 || year <= 0 {
		fmt.Println("You should type amount, rate, year greater than 0 !")
		return
	}
	

	for i:=1; i<=year; i++ {
		amount +=  (amount * rate)/100
		fmt.Printf("Year %d: %.2f\n", i,amount) 
	}
 }


func main() {
	compoundInterestCal(10000,3,3)

}