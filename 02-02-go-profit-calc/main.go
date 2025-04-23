package main

import (
	"fmt"
	"strconv"
)

func reportExpenseBeforeTax(revenue float64, expenses float64, taxRate float64) {
	var ebt = revenue - expenses
	var profit = ebt * (1 - taxRate/100)
	var ratio = ebt / profit

	fmt.Printf("revenue  =%10.2f\n", revenue)
	fmt.Printf("expenses =%10.2f\n", expenses)
	fmt.Printf("taxRate  =%10.2f\n", taxRate)
	fmt.Printf("ebt      =%10.2f\n", ebt)
	fmt.Printf("profit   =%10.2f\n", profit)
	fmt.Printf("ratio    =%10.2f\n", ratio)

}

func readUserInputAsFloat64(fieldName string) (float64, error) {
	var s string
	var d float64
	var err error

	fmt.Printf("Enter %s:", fieldName)
	_, err = fmt.Scan(&s)
	d, err = strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("Error: Bad value (%s) specified ::: value must be a real number \n", s)
		return 0, err
	}

	return d, err
}

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	var err error

	revenue, err = readUserInputAsFloat64("Revenue")
	if err != nil {
		return
	}
	expenses, err = readUserInputAsFloat64("Expenses")
	if err != nil {
		return
	}
	taxRate, err = readUserInputAsFloat64("taxRate")
	if err != nil {
		return
	}

	reportExpenseBeforeTax(revenue, expenses, taxRate)
}

/**
cd 02-02-go-profit-calc
go run .
Enter Revenue:1100
Enter Expenses:100
Enter taxRate:10


revenue  =   1100.00
expenses =    100.00
taxRate  =     10.00
ebt      =   1000.00
profit   =    900.00
ratio    =      1.11

*/
