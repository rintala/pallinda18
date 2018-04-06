
package main

import("fmt"
		"time"
		"math/rand")

var (
	Web = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result struct{
	kind string
	query string
}
∂
type Search func(query string) Result

func fakeSearch(kind string) Search{
	return func(query string) Result{
		time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result){

	//Make goroutines out of these instead - run much faster
	/*results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))*/
	c := make(chan Result)
	go func(){ c <- Web(query)}()
	go func(){ c <- Image(query)}()
	go func(){ c <- Video(query)}()

	timeout := time.After(70*time.Millisecond)

	for i := range c{
		select {
		case result := <-c:
			results = append(results, result)
		
		case <-timeout:
			fmt.Println("timed out...")
			return
		}
	}

	return results
}

func main(){
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)

}