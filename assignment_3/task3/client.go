//Weather station -- TASK 3

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	//Declare and initialize the slice with the 3 different severs to connect to
	server := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}

	// Add a time limit for all requests made by this client.
	client := &http.Client{Timeout: 10 * time.Second}

	for {
		before := time.Now()
		//res := Get(server[0], client)
		res := MultiGet(server, client)
		after := time.Now()
		fmt.Println("Response:", res)
		fmt.Println("Time:", after.Sub(before))
		fmt.Println()

		//Sleep 2s in between each Multiget run
		time.Sleep(2000 * time.Millisecond)
	}
}

type Response struct {
	Body       string
	StatusCode int
}

func (r *Response) String() string {
	return fmt.Sprintf("%q (%d)", r.Body, r.StatusCode)
}

// Get makes an HTTP Get request and returns an abbreviated response.
// The response is empty if the request fails.
func Get(url string, client *http.Client) *Response {
	res, err := client.Get(url)
	if err != nil {
		return &Response{}
	}
	// res.Body != nil when err == nil
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}
	return &Response{string(body), res.StatusCode}
}

// MultiGet makes an HTTP Get request to each url and returns
// the response from the first server to answer with status code 200.
// If none of the servers answer before timeout, the response is 503
// – Service unavailable.
func MultiGet(urls []string, client *http.Client) *Response {

	outputCh := make(chan *Response)
	before := time.Now()
	timeout := time.After(500 * time.Millisecond)
	finalOut := make(chan *Response)

	//Running through the the URLs - goroutine for each URL - run GET
	//go func(){

	for _, url := range urls{
		//fmt.Println("url loop..")
		go func(url string, client *http.Client){

			//Use existing Get function for each URL => MultiGet
			outputCh <- Get(url,client)
			fmt.Println("OUTPUT RECEIVED")
			
		}(url, client)
	}
	//}()
	//fmt.Println("url loop done or started")
	
	//Start new goroutine to run continously - identifying timeout OR if something is on the outputCh-channel
	go func(){

		theLoop:
		for{
			select{

				//Constantly run check if total time is out - if so break loop and return error code
				case thetime := <-timeout:
					fmt.Println("TIMEOUT TOOK TOO LONG!!", thetime.Sub(before))
					
					//Edit response obj
					failedRes := &Response{"Timeout error", 503}
					finalOut <-failedRes

					break theLoop
					
				//If value on output channel (from server) - check what statuscode and take action
				case res := <-outputCh:
					if(res.StatusCode == 200){
						
						//Break the for loop if correct statuscode '200' on channel & thus terminate the search 
						fmt.Println("YES OK - LETS RETURN")
						println("WE HAVE SOMETHING ON THE OUTPUT CHANNEL", res.StatusCode)
						finalOut <- res
						break theLoop

					//If 503 - Service is unavailable - try another server
					} else if(res.StatusCode == 503){
						fmt.Println("Service unavailable... Let's try another server!")
					}	
			}
		}
	}()
	
	//Block until a value sent on the finalOut channel (by the running goroutine)
	return <-finalOut
}

/****************** CONCLUSIONS & OBSERVATIONS *********************************/
/*
	- Might want to run the entire URL for loop in a goroutine, wouldn't really change the behaviour though?
	- My timeout runs from the beginning of the MultiGet function (unclear when to start it)
*/