package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("InvestmentAmount: ")
	outputText("InvestmentAmount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("futureValue: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("futureRealValue: %.0f\n", futureRealValue)
	// fmt.Println("futureValue:", futureValue)
	// fmt.Printf(`futureValue: %.1f\n
	// 			futureRealValue: %.0f\n`,
	// 	futureValue,
	// 	futureRealValue)
	// fmt.Printf("futureRealValue: %f", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}
