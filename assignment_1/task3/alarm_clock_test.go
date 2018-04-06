//Alarm Clock -- TASK 3
	//Solution with three different concurrent goroutines running and sending to a channel

package main

import (
	"fmt"
	"time"
	"testing"
	"strings"
)
/*
func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}
*/

func Remind(text string, delay time.Duration, outCh chan string){
		fmt.Println("Remind function called...\n")
		//NewTicker returns a new Ticker containing a channel that will send the
		//time with a period specified by the duration argument.
    	newTimer := time.NewTicker(delay)

		//Run for-loop with select to identify what case/timer channel and return appropriate strings
		//for{
			select{
				case <-newTimer.C:
					current := time.Now()
					currentString := fmt.Sprintf("Klockan är: "+"%d.%d.%d"+": "+"%s"+"\n", current.Hour(), current.Minute(), current.Second(), text)
					outCh<-currentString
			}
		//}
}

func TestRemind(t *testing.T) {
	//p := fmt.Println
	outCh := make(chan string)

	//p("------- Alarm Clock - Main initialized -------")
	
	start := time.Now()
	
	//p("START TIME: ", start.Format("15:04:05 (2006-01-02)"))
	//p("---------------------------------")
	message1 := "Dags att äta"
	message2 := "Dags att arbeta"
	message3 := "Dags att sova"

	go Remind(message1, 3*time.Second, outCh)
	go Remind(message2, 8*time.Second, outCh)
	go Remind(message3, 24*time.Second, outCh)

	//for outp := range outCh{
	for i:=0;i<3;i++{
		theOutput:= <-outCh
		fmt.Println(theOutput)
		theTextSlice := strings.Fields(theOutput)[3:]
		//var theText string
		theText := strings.Join(theTextSlice," ")
		
		fmt.Println(theText)
		fmt.Println("TIMDIFF",time.Since(start))
		//timeDiff := time.Since(start)
		
		//|| timeDiff!=(3*time.Second)

		if (theText != message1 && i==0){
			t.Error("EXPECTED:",message1,"...", "GOT:", theOutput)
		}

		if (theText != message2 && i==1){
			t.Error("EXPECTED:",message2,"...", "GOT:", theOutput)
		} 

		if (theText != message3 && i==2) {
			t.Error("EXPECTED:",message3,"...", "GOT:", theOutput)
		}
	}

}
