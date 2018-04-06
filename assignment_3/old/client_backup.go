package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	server := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}

	// Add a time limit for all requests made by this client.
	client := &http.Client{Timeout: 10 * time.Second}

	//for {
		before := time.Now()
		//res := Get(server[0], client)
		res := MultiGet(server, client)
		after := time.Now()
		fmt.Println("Response:", res)
		fmt.Println("Time:", after.Sub(before))
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	//}
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
	//fmt.Println(urls)

	outputCh := make(chan *Response)
	before := time.Now()
	timeout := time.After(500 * time.Millisecond)
	timeOutBool := false

	for _, url := range urls{
		go func(url string, client *http.Client){
			//before := time.Now()
			//outputCh := <- Get(url,client)
			//after := time.Now()

			//result := Get(url,client)
			//fmt.Println("STATCODE:",result.StatusCode)
			//Remember that statusCode = 200 is generated if the service is unavailable
			
			/*if(result.StatusCode==200){
				fmt.Println("- Connection OK")
			}*/

			select{

				case outputCh <- Get(url,client):
					fmt.Println("OUTPUT RECEIVED")		

				//case time.Now().Sub(before)>(time.Second*5):
				/*
				case <- time.After(time.Second*5):
					fmt.Println("TOOK T0O LONG")
				*/

				case thetime := <- timeout:
					fmt.Println("Timeout, took too long",thetime.Sub(before))
					timeOutBool = true

					//FASTNAR IBLAND HÄR... - se tilla tt den verkligen quitar allt om denna aktiveras
					//return
			}
			
		}(url, client)

		//outputCh <- go Get(url, client)
	}

	//res := <-outputCh
	ansCounter :=0

	for res := range outputCh {
		
		if (timeOutBool == true){
			fmt.Println("timeout changing statuscode...")
			res.StatusCode = 503
			fmt.Println("NOW THE PROGRAM SHOULD TERMINATE")
			return res

		} else{
			fmt.Println("no timeout recognized..")

		}

		if res.StatusCode == 200 {
			fmt.Println("YES = 200")
			fmt.Println("THE PROGRAM SHOULD NOW RETURN AND TERMINATE")
			return res
		} else if res.StatusCode == 503 {
			ansCounter++
			fmt.Println("NO = 503, ANSCOUNTER:", ansCounter)
		}

		if ansCounter > 2 {
			fmt.Println("503 - Timeout No Answers")
			//return 
		}

	}
	

	//println(<-outputCh)
	//return res// TODO

	return nil
}
