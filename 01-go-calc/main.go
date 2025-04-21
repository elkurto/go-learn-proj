package main

import (
	"fmt"
	"math"
)

func computeFutureValue() {
	const inflationRate = 6.5
	var investmentAmount float64 = 1000
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Printf("starting principle (1000 default): ")
	_, err := fmt.Scan(&investmentAmount)
	if err != nil {
		fmt.Printf("Input error for starting principle %s", err)
		return
	}

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("futureValue =%f\n", futureValue)
	fmt.Printf("futureRealValue =%f\n", futureRealValue)
}

func main() {
	computeFutureValue()
}
