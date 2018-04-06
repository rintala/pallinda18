// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {

	fmt.Println("\nORACLE INITIALIZEDDD")

	questions := make(chan string)
	answers := make(chan string)
	
	// TODO: Answer questions.
	// TODO: Make prophecies.
	// TODO: Print answers.

	//until channel 'questions' has been closed

	//go generatePredictions("I predict u will live long boi!", answers)

	//First go routine - to handle prophecies
	go func(){
	
			for{

				time.Sleep(15 * time.Second)
				//fmt.Println("Im about to give you a prophecy..")
				var qInput string
				select{
					case qInput = <-questions:
						fmt.Println("USE question as input to prophecy..")
					default:
						qInput = ""
						//fmt.Println("THE PROPECHY IS",<-answers)
				}
				go prophecy(qInput,answers)
		}

	}()

	//Second go routine - to handle answers on incoming questions
	go func(){

		for question := range questions{
			fmt.Println("IM about to ans your question")
			go generateAnswers(question, answers)
			
			//ansTemp := <-answers
			//answers<-ansTemp
			//fmt.Println("THE ANSWER IS:",ansTemp)
		}
		
	}()

	//Third go routine - to handle printing of the answers
	go func(){
		//fmt.Print(prompt)
		for answer := range answers{
			fmt.Print(answer)
		}

	}()
	
	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.

// TODO: make it better.

func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.

	//fmt.Println("\nProphecy is initialized...")
	//fmt.Print("\n",prompt)
	//time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)
	time.Sleep(time.Duration(10+rand.Intn(10)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"We shall live.",
		"This is why we do this.",
		"Living on Pluto will soon be a fact.",
	}

	prophecyOutput := "\n"+longestWord + "... " + nonsense[rand.Intn(len(nonsense))]+ "\n"
	//answer <- prophecyOutput
	for l := range prophecyOutput{
		time.Sleep(80*time.Millisecond)
		answer<-string(prophecyOutput[l])
	}
	//fmt.Println("\nProphecy finished")
	fmt.Print("\n"+prompt)
}

func generateAnswers(pred string, answer chan<- string){
	//time.Sleep(time.Duration(10+rand.Intn(10)) * time.Second)
	//fmt.Println(pred, "AND MY ANSWER IS u SUCK!")
	fmt.Println("\ngenerating answer...\n")
	time.Sleep(5 * time.Second)
	//pred = "----"+pred + "AND MY ANSWER IS u SUCK!"+"-----"
	//myAns := "this is my answer now..\n"

	myAnswers := []string{
		"Sry, dat wont work u fattie.",
		"Yes, bich pls.",
		"Lol, what u trying to say to me dawg?",
		"OMG dont u try me!",
		"No, yall hoes can cassh me outsidee!",
		"Remember who's tha baws here, I decide ur future.",
		"ITS ERRRYDAYY BROOO!!",
	}

	myAns := myAnswers[rand.Intn(len(myAnswers))]

	for l := range myAns{
		time.Sleep(80*time.Millisecond)
		//fmt.Print(string(myAns[l]))
		answer <-string(myAns[l])
	}
	answer<-"\n> "

	//answer <- pred
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	fmt.Println("Init initialized...")
	rand.Seed(time.Now().Unix())
}
