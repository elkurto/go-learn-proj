package main

import (
	"fmt"
	"os"
	"strconv"
)

func readAndReportCmdLineArgs() {
	for i := 0; i < len(os.Args); i++ {

		fmt.Printf("%d : %s\n", i, os.Args[i])
	}
}

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

func printUsage() {
	fmt.Println(
		`
	Usage:
		go run . <revenue> <expense> <taxRate>
	
	example:
		go run . 1100.00 100.00 10.00`,
	)
}

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	var err error

	revenue, err = strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("bad value for revenue")
		printUsage()
		return
	}

	expenses, err = strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("bad value for expenses")
		printUsage()
		return
	}

	taxRate, err = strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("bad value for taxRate")
		printUsage()
		return
	}

	reportExpenseBeforeTax(revenue, expenses, taxRate)
}
