// package main is special. It defines a standalone executable program, not a library.
// The main function in package main will be the entry point of our executable program.
// The fmt package from the standard library is imported to provide I/O functions.
package main

/*
import (
	"fmt"
	"math"
)

const inflationRate = 2.5 // constant declaration - float64 type is inferred

func main() {
	var investmentAmount float64 // value: 0.0 as initial value
	var years float64
	expectedReturnRate := 5.5	// type inference

	outputText("Investment Amount: ") // call function to print to console
	fmt.Scan(&investmentAmount) // read user input, & is used to get the address of the variable

	outputText("Expected Return Rate: ") // call function to print to console
	fmt.Scan(&expectedReturnRate) // read user input

	outputText("Years: ") // call function to print to console
	fmt.Scan(&years) // read user input

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	//futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	//futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	formattedFutureValue := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRealFutureValue := fmt.Sprintf("Future Value adjust for Inflation: %.1f\n", futureRealValue)
	// fmt.Print(formattedFutureValue, formattedRealFutureValue)

	// fmt.Println("Future Value: ", futureValue)
	//fmt.Printf("Future Value: %.1f \nFuture Value adjust for Inflation: %.1f", futureValue, futureRealValue)

	fmt.Print(formattedFutureValue, formattedRealFutureValue)

}
func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues (investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64){
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return fv, rfv
	// return
}
*/
