package main

import (
	"fmt"
)
// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan <- int) {
	
	fmt.Println("Add function initialized...")
	fmt.Println(a,"\n")

	sum := 0
	for i := range a{
		sum +=a[i]
	}
    res <- sum
}

func main() {
    p:= fmt.Println
    p("************** Task 4 - Two Part Sum **************")
    p("MAIN RUNNING...\n")
    a := []int{1, 2, 3, 4, 5, 6, 7}
    n := len(a)
    ch := make(chan int)
    p("FULL LIST: ",a)
    p("   - FIRST PART OF THE LIST: ",a[:n/2])
    go Add(a[:n/2], ch)
    p("   - SECOND PART OF THE LIST: ",a[n/2:],"\n")
    go Add(a[n/2:], ch)

    // TODO: Get the subtotals from the channel and print their sum.
    

    //---- alternative declaration ----
    //var sub int 
    //sub = <- ch
    //---------------------------------

    //---- shorter declaration --------
    sub1 := <- ch
    sub2 := <- ch
    //---------------------------------

   	//---- Add the two sub-sums together ---
   	totSum := sub1+sub2
   	p("-------------------")
   	p("SUBSUM 1:",sub1)
   	p("SUBSUM 2:", sub2)
    p("TOTAL SUM:",totSum)
}