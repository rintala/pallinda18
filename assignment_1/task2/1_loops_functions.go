//Exercise: Loops and Functions -- TASK 2

/************************* TASK DESCRIPTION ************************************/
	/*	- Implement a square root function
			- Given x, we want to find a z such that z^2 is most nearly x
			- Given z -= (z*z - x) / (2*z) run for-loop
			- Repeat the calculations 10 times and print each z along the way
			- Or repeat until diff is small enough
			- Try various values of x - see how quickly the guess improves
	*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	//x = 1
	//z := 1
	
	//STARTING GUESS = z, can change according to what we want
	z := float64(1)

 	fmt.Println("Sqrt calculation initiliazed")

	for math.Abs(z*z-x)/(2*z)>math.Pow(10,-6){
		z -= (z*z - x) /(2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println("\nExercise: Loops and Functions\n")
	fmt.Println("TRESHOLD SET AT: 10^-6\n")

	for x:=1.0;x<=10.0;x++{
		fmt.Println("\n--- CALC FOR x = ",x, "--------------------------")
		approxVal := Sqrt(x)
		fmt.Println("\nAPPROXIMATION:",approxVal)
		actualVal := math.Sqrt(x)
		fmt.Println("ACTUAL ANS:",actualVal)
		fmt.Println("DIFF FROM ACTUAL VALUE:",math.Abs(approxVal-actualVal))
	}

}

/**************************** QUESTIONS & ANSWERS ***************************************/
//How close are your function's results to the math.Sqrt in the standard library?
	/*	- See printouts for the different x's between 1-10
			- Calculating real or actual value with the math library's sqrt()-Functions
			- Greatest diff for x = 5, diff = 9.18143573613861e-07
			- So really close for all => good approx
	*/
