//Alarm Clock -- TASK 3
	//Alternative solution running the Remind function only one time with all 3 timers inside

package main

import (
	"fmt"
	"time"
	//"sync"
	//"math"
)

func Remind(text string, delay time.Duration){
		fmt.Println("Remind function called...\n")

		//NewTicker returns a new Ticker containing a channel that will send the
		//time with a period specified by the duration argument.
    threeHrTimer := time.NewTicker(3*time.Second)
    eightHrTimer := time.NewTicker(8*time.Second)
    twentyfourHrTimer := time.NewTicker(24*time.Second)

		//Run for-loop with select to identify what case/timer channel and return appropriate strings
    for {
    	select {
			case three := <-threeHrTimer.C:
    			fmt.Print("Klockan är: ", three.Hour(),".", three.Minute(),".", three.Second(),": ","Dags att äta","\n")
    		case eight := <-eightHrTimer.C:
    			fmt.Print("Klockan är: ", eight.Hour(),".", eight.Minute(),".", eight.Second(),": ","Dags att arbeta","\n")
    		case twentyfour := <-twentyfourHrTimer.C:
    			fmt.Print("Klockan är: ", twentyfour.Hour(),".", twentyfour.Minute(),".", twentyfour.Second(),": ","Dags att sova","\n")
    	}
    }
}

//use time.Duration(timeint) to convert int => time Duration
func main() {
	p := fmt.Println

	p("------- Alarm Clock - Main initialized -------")
	start := time.Now()
	p("START TIME: ", start.Format("15:04:05 (2006-01-02)"))
	p("---------------------------------")

	Remind("Custom alert", 3*time.Second)

}
