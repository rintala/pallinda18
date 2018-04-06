//Exercise: Slices -- TASK 2

/************************* TASK DESCRIPTION ************************************/
	/*	- Implement the Pic Function
			- It should return a slice of length dy
			- Each element is a slice of dx 8-bit unsigned ints
			- When running the program it will display the picture
			- The choice of image is up to you (function)
			- You need a loop to allocate each []uint8 inside the [][]uint8
	*/

package main

import (
	"golang.org/x/tour/pic"
	"fmt"
	"math"
)

func Pic(dx, dy int) [][]uint8 {

	//declare and assign s to 1d slice with type []uint8 of its elements
		//make allocates a zeroed array, returns a slice that refers to that array
		//make([] []uint8)
	
	//var s [][]uint8

	//make a slice of length dy, with place for elemnts of type []uint8 (slice)
	s := make([][]uint8, dy)

	//for i:=0; i<dy;i++{
	for i := range s{
		//printSlice(s)

		s[i] = make([]uint8,dx)
		fmt.Println("INDEX i:",i)
		//s[i] = append(innerSlice,3)

		//Loop over slice elements and fill them w values from a funct of our choice
		for j:=0;j<len(s[i]);j++{
			//Using uint(..) for converting int => unsigned int

      //--------------FUNCTIONS--------------

	      	//Function 1
	      	//s[i][j] = uint8((i+j)/2)

	      	//Function 2
	      	//s[i][j] = uint8(i*j)

			//Function 3
			s[i][j] = uint8(math.Pow(float64(i),float64(j)))

    	//-------------------------------------
      }
	}

	return s

}

func main() {
	pic.Show(Pic)
}

/**************************** QUESTIONS & ANSWERS ***************************************/
//No questions, but see result folder for the different image outputs
