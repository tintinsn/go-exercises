package main

import (
	"fmt"
)

//? Ex 1: Compound Interest Calculator

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

 //? Ex 2: Expense Filter and Sum
  
 /*
  * Problem Statement: 
  - You are given a list of transactions
  - print only the expenses 
  - calculate and print the total expense
  
  * Example
	Input: [100, -50, 200, -300, -150]
 	Expected Output:
 	  Expenses: -50 -300 -150
 	  Total Expense: -500.00
  
  * Condition:
  - Input
  	- How many inputs are required?
 		- a single input from user.
    - What the data type of input?
 		- slice
	- Are the inputs expected to be positive numbers onl?
 		- no, the input is list of slice that contain positive and negative number.
	- Are there any minimum or maximum limit?
 		- no.
	- Do the inputs need to be validated or converted before us?
 		- no.

  - Core Rule
  	- What exactly does the problem ask us to do?
  		- Print for each expense and total of expense 
	- Are there any specific rules or formulas given in the problem?
 		- must be a negative number.
 		- sum the expense by totalExpense += expense
	- Should the logic involve loops, conditions, or functions?
 		- use for loop for count sum total expense
	- Is there any domain-specific behavior (e.g., “compounded annually”, “even numbers only”)?
 		- no special conversion is needed.

  - Ouput
  	- What should be displayed as output?
 		- Expenses: -300 -140 -400
		- Total Expense: -840.00
	- Should the output include text labels or just values?
 		- include labels.
	- Does the output need to be formatted (e.g., 2 decimal places)?
 		- total expense must be 2 decimal places.
	- Is the output printed once or repeatedly (e.g., per loop iteration)?
 		- output appears only once.

  - Edge Cases
  	- What happens if inputs are zero or negative?
 		- the input is a slice that may contain both negative and positive numbers
	- What if inputs are missing or invalid (e.g., non-numeric characters)?
 		- should reject
	- Should the program handle extremely large or small values?
 		- no
	- Are there special cases that need custom handling?
  
  * Approach 1:
  	1. Read the input from user:
	2. if len of transaction > 0 run do task else warning to user
 	2. Initialize variables totalExpense use type float64 to keep sum of expense
 	3. loop through each transaction
	4. check if the value is negative
 	5. if negative, print the value
 	6. sum up the value to totalExpense
	7. after the loop, print the final totalExpense rounded 2 decimal places

	❌ Result: failed test case 2
 	   - Doesn't handle case where all values are positive
	   — prints empty "Expenses:" but still shows 0.00

 * Approach 2:
 	1. Same as above
	2. Add new slice variable to keep result for each expense
	3. if no expense found in new slice, print "No expense found."
	4. otherwise, print expense and total expense

    ✅ Result: Passes all test case
 	   - but use additional memory due to extra slice.
 	

 * Manual Test:
 	
 	Test Case 1:
 	Input: [100, -50, 200, -300, -150]
 	Expected Output:
 	  Expenses: -50 -300 -150
 	  Total Expense: -500.00
 	
 	Test Case 2:
 	Input: [100, 200, 300]
 	Expected Output:
 	  No expenses found.
 	
 	Test Case 3:
 	Input: []
 	Expected Output:
 	  No transactions provided.
 	
 	Test Case 4:
 	Input: [-1000]
 	Expected Output:
 	  Expenses: -1000
 	  Total Expense: -1000.00
 	
 	Test Case 5:
 	Input: [0, -0.01, 0.01]
 	Expected Output:
 	  Expenses: -0
 	  Total Expense: -0.01
 */

func expenseFilterSum(trans []float64) {

	/* <----------- Approach 1 ----------->*/
	// if len(trans) > 0 {
	// 	var totalExpense float64 
	// 	fmt.Print("Expenses: ")
		
	// 	for _ ,v := range trans {
	// 		if v < 0 {
	// 			fmt.Printf("%.0f ",v)
	// 			totalExpense += v
	// 		}
	// 	}
	// 	fmt.Println()
	// 	fmt.Printf("Total Expense: %.2f\n",totalExpense)
	// } else {
	// 	fmt.Println("No transactions provided.")
	// }

	/* <----------- Approach 2 ----------->*/
	if len(trans) == 0 {
		fmt.Println("No transactions provided.")
		return
	}

	var totalExpense float64
	var expenseList []float64
	
	for _, v := range trans {
		if v < 0 {
			expenseList = append(expenseList, v)
			totalExpense += v
		}
	}
	
	if len(expenseList) > 0 {
		fmt.Print("Expense: ")
		for _, v := range expenseList {
			fmt.Printf("%.0f ", v)
		}

		fmt.Println()
		fmt.Printf("Total Expense: %.2f\n",totalExpense)
	} else {
		fmt.Println("No expenses found")
	}
}




func main() {
	transactions := []float64{100, 200, -300, -490}
	expenseFilterSum(transactions)

}