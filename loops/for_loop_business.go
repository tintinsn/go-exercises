package main

import (
	"fmt"
	"math"
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
 		- no, the input is a list of slice that contain positive and negative number.
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

 //? Ex 3: Bank Statement with Balance

/*
 * Problem Statement:
 - Given a list of transactions (each with a description and amount), 
 - print a full bank statement showing the transaction and running balance.

 * Example
 Input: 


	transactions := []Transaction {
		{"Deposit", 1000},
		{"Withdraw", -300},
		{"Transfer", -200},
	}
 Expected Output: 
		Deposit: 1000.00 | Balance: 1000.00  
		Withdraw: -300.00 | Balance: 700.00  
		Transfer: -200.00 | Balance: 500.00  

 * Condition:
 - Input:
   - How many inputs are required?
		- A single input: a list of transactions.
   - What is the data type of each input?
		- transactions is a slice of Transaction structs, so type is []Transacrion
		- []Transaction contains 2 fields
			1. Description field type string
			2. Amount field type float64
			3. type Transaction struct {
				Description string
				Amount      float64
				}
   - Are the inputs expected to be positive numbers only?
		- no, the inputs expected to be a slice of []Transaction
   - Are there any minimum or maximum limits?
		- no
   - Do the inputs need to be validated or converted before use?
		- no

 - Core Rule:
   - What exactly does the problem ask us to do?
		- print statement and the running balance.
   - Are there any specific rules or formulas?
		- no,
   - Should the logic involve loops, conditions, or functions?
		- use a loop.
   - Is there any domain-specific behavior?
		- no

 - Output:
   - What should be displayed as output?
		- Deposit: 1000.00 | Balance: 1000.00  
		Withdraw: -300.00 | Balance: 700.00  
		Transfer: -200.00 | Balance: 500.00  
   - Should the output include labels or just values?
		- include labels
   - Does the output need to be formatted (e.g., decimal places)?
		- print the number of amount and balance rounded 2 decimal places
   - Is the output printed once or repeatedly?
		- Each line is printed once for each transaction.

 - Edge Cases:
   - What happens if inputs are zero or negative?
		- the input is a slice that may contain both negative and positive numbers
   - What if inputs are missing or invalid?
		- warnning to user and stop running the program
   - Should the program handle extremely large or small values?
		- no
   - Are there special cases that need custom handling?
		- no

 * Approach 1:
  1. check if the transactions list is empty
		- if empty print a warning and return.
  2. create and initialize variable call `balance` type float64 to keep track of the running balance.
  3. loop through each transactions
		- add the amount to balance
		- Print description, amount, and current balance in the required format

✅ Result: Passes all test cases


 * Manual Test:
  Test Case 1:
	Input:
		transactions := []Transaction{
		{"Deposit", 1000},
		{"Withdraw", -300},
		{"Transfer", -200},
		}

	Expected Output:
		Deposit: 1000.00 | Balance: 1000.00  
		Withdraw: -300.00 | Balance: 700.00  
		Transfer: -200.00 | Balance: 500.00


  Test Case 2:
	Input:
		transactions := []Transaction{
		{"Deposit", 5000},
		{"Withdraw", -1500},
		{"Deposit", 2000},
		{"Withdraw", -1000},
		}

	Expected Output:
		Deposit: 5000.00 | Balance: 5000.00  
		Withdraw: -1500.00 | Balance: 3500.00  
		Deposit: 2000.00 | Balance: 5500.00  
		Withdraw: -1000.00 | Balance: 4500.00


  Test Case 3:
	Input:
		transactions := []Transaction{
		{"Deposit", 300},
		{"Withdraw", -100},
		{"Withdraw", -50},
		{"Deposit", 250},
		}

	Expected Output:
		Deposit: 300.00 | Balance: 300.00  
		Withdraw: -100.00 | Balance: 200.00  
		Withdraw: -50.00 | Balance: 150.00  
		Deposit: 250.00 | Balance: 400.00

*/

type Transaction struct{
	Description string
	Amount      float64
}

func bankStatementBalance(trans []Transaction) {
	if len(trans) == 0 || trans == nil {
		fmt.Print("No transactions provided.")
		return
	}

	var balance float64

	for _,v := range trans {
		balance += v.Amount
		fmt.Printf("%s: %.2f | Balance: %.2f\n",v.Description,v.Amount, balance)
	}
}

 //? Ex 4: Cart Discount Calculator
/*
 * Problem Statement:
- Write a program that takes a list of products (with price and discount%) 
- and uses a for loop to calculate the price after discount for each item and the total price of the cart.

 * Example
 Input:
	products := []Product {
		{"Shoes", 1200, 10},
		{"T-shirt", 500, 0}
	}
 Expected Output: 
	Shoes: 1080.00
	T-Shirt: 500.00
	Total: 1580.00

 * Condition:
 - Input:
   - How many inputs are required?
		- A single input: a list of products
   - What is the data type of each input?
		- products is a slice of Product struct,
		- []Product contain 3 fields
			1. Name type string
			2. Price type float64
			3. Discount type float64
   - Are the inputs expected to be positive numbers only?
		- No, input expected to be a non-empty list of products.
		- Each product should have a price greater than 0 and a discount between 0 and 100.
   - Are there any minimum or maximum limits?
		- price must be greater than 0
		- discount must be between 0 and 100 (inclusive)
   - Do the inputs need to be validated or converted before use?
		- Yes. negative price must be coverted to positive values
		- discount should be validate to ensure they are between 0 and 100.

 - Core Rule:
   - What exactly does the problem ask us to do?
		- print price of each product after using discount.
		- print total price of all item in cart
   - Are there any specific rules or formulas?
		- Use the formula: `price * (1 - discount / 100)`
		- The discount percentage must be between 0 and 100 (inclusive)
   - Should the logic involve loops, conditions, or functions?
		- Use a loop to iterate over each product
		- Apply the discount calculation per item
		- Sum the final prices into a total
		
   - Is there any domain-specific behavior?
		- none

 - Output:
   - What should be displayed as output?
		Shoes: 1080.00
		T-Shirt: 500.00
		Total: 1580.00
   - Should the output include labels or just values?
		- include labels
   - Does the output need to be formatted (e.g., decimal places)?
		- All prices should be rounded to 2 decimal places
   - Is the output printed once or repeatedly?
		- One line per product
  		- One line for the total

 - Edge Cases:
   - What happens if inputs are zero or negative?
		- the inputs cannot be nil or empty list.
   - What if inputs are missing or invalid?
		- if list of products are empty or nil warning to user and return.
		- If any product has a negative price:
  			- Convert it to a positive number before calculation
		- If any product has a discount less than 0% or more than 100%:
  			- Show a warning and skip or terminate the process
   - Should the program handle extremely large or small values?
		- Yes, the program should correctly process very high prices or small decimal discounts without rounding errors or overflow.


 * Approach 1:
  1. check if the products is empty or nil, then warning and return
  2. create and initialize variable call `total` type float64 for keep track total price of products in cart
  3. loop through each product
		- check discount if less than 0 or greater than 100 then return.
		- check price is negative number if negative then convert to positive number
		- compute price with discount by price * discount / 100 then update new price
		- print product name and price after use discount
		- accumulate the `total` by adding each item price with discount
  4. print total price
 ✅ Result: Passes all test cases


 * Manual Test:
  Test Case 1:
  Input: 
	products := []Product{
	{"Shoes", 1200, 10},
	{"T-Shirt", 500, 0},
	{"Bag", 1500, 5},
	}
  Expected Output:
	Shoes: 1080.00
	T-Shirt: 500.00
	Bag: 1425.00
	Total: 3005.00

  Test Case 2:
  Input: 
	products := []Product{
	{"Notebook", 100, 0},
	{"Pen", 50, 0},
	}
  Expected Output:
	Notebook: 100.00
	Pen: 50.00
	Total: 150.00


  Test Case 3:
  Input: 
  	products := []Product{
	{"Promo Code", 1000, 100},
	{"Free Gift", 500, 100},
	}
  Expected Output:
	Promo Code: 0.00
	Free Gift: 0.00
	Total: 0.00

  Test Case 4:
  Input: 
  	products := []Product{}
  Expected Output:
	No products provided.

  Test Case 5:
  Input: 
  	products := []Product{
	{"Smartphone", 999.99, 12.5},
	}
  Expected Output:
	Smartphone: 874.99
	Total: 874.99

  Test Case 6:
  Input: 
  	products := []Product{
	{"Bug Item", 200, 150}, // Discount > 100%
	}
  Expected Output:
	Smartphone: 874.99
	Total: 874.99
*/

type Product struct {
	Name string
	Price float64
	Discount float64
}

func cartDiscountCal(products []Product) {
	if products == nil || len(products) == 0 {
		fmt.Println("No products provided.")
		return
	}

	var total float64
	for _, product := range products {
		
		if product.Discount < 0 || product.Discount > 100  {
			fmt.Println("discount must be between 0 and 100 percent")
			return
		}

		if product.Price < 0 {
			product.Price = math.Abs(product.Price)
		}

		discountedPrice := product.Price - (product.Price * product.Discount / 100)
		fmt.Printf("%s: %.2f\n", product.Name, discountedPrice)
		total += discountedPrice
	}
	fmt.Printf("Total: %.2f\n", total)
	}

func main() {
	products := []Product{
	{"Shoes", 1200, 10},
	{"T-Shirt", 500, 0},
	{"Bag", -1500, 5},
	}
	cartDiscountCal(products)
}