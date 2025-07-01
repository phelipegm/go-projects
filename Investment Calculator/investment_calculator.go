package main

import ( 
	"math"
	"fmt"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow((1 + expectedReturnRate / 100), years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1 + expectedReturnRate / 100), years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	
	return fv, frv
}