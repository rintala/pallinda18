package main

import (
	"fmt"
	"strconv"
	"time"
)

func Remind(text string, delay time.Duration) string {
	custom := time.Tick(delay)
	eat := time.Tick(3 * time.Second)
	work := time.Tick(8 * time.Second)
	sleep := time.Tick(24 * time.Second)

	hour := strconv.Itoa(time.Now().Hour())
	minute := strconv.Itoa(time.Now().Minute())
	sec := strconv.Itoa(time.Now().Second())
	fmt.Println("Klockan är START: " + hour + "." + minute + "."+sec+": " + text)

	for {
		select {
		case <-custom:
			hour := strconv.Itoa(time.Now().Hour())
			minute := strconv.Itoa(time.Now().Minute())
			sec := strconv.Itoa(time.Now().Second())
			fmt.Println("Klockan är " + hour + "." + minute + "."+sec+": " + text)
		case <-eat:
			hour := strconv.Itoa(time.Now().Hour())
			minute := strconv.Itoa(time.Now().Minute())
			sec := strconv.Itoa(time.Now().Second())
			fmt.Println("Klockan är " + hour + "." + minute + "."+sec+": Dags att äta")
		case <-work:
			hour := strconv.Itoa(time.Now().Hour())
			minute := strconv.Itoa(time.Now().Minute())
			sec := strconv.Itoa(time.Now().Second())
			fmt.Println("Klockan är " + hour + "." + minute + "."+sec+": Dags att arbeta")
		case <-sleep:
			hour := strconv.Itoa(time.Now().Hour())
			minute := strconv.Itoa(time.Now().Minute())
			sec := strconv.Itoa(time.Now().Second())
			fmt.Println("Klockan är " + hour + "." + minute + "."+sec+": Dags att sova")
			return "done"
		default:
			time.Sleep(10 * time.Millisecond)

		}
	}

}

func main() {
	fmt.Println(Remind("hej", 5*time.Second))
}
