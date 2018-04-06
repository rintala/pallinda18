//Pythia, the Oracle of Delphi -- TASK 3

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

	//First go routine - to handle prophecies
	go func(){
		for{
				//inifinite for loop running with delay of 15s
				time.Sleep(15 * time.Second)
				//fmt.Println("Im about to give you a prophecy..")

				go prophecy("",answers)
		}

	}()

	//Second go routine - to handle answers on incoming questions
	//runs until channel 'questions' has been closed
	go func(){
		for question := range questions{
			fmt.Println(star+" is about to ans your question..")
			go generateAnswers(question, answers)
		}

	}()

	//Third go routine - to handle printing of the answers
	go func(){
		for answer := range answers{
			characters := strings.Split(answer, "")
			for _, c := range characters {
				fmt.Printf(c)

				//Randomizing the delay within 60 - 300 ms
				time.Sleep(time.Millisecond*time.Duration((60*rand.Intn(5))))
			}
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

	prophecyOutput := "\n"+longestWord + "... " + nonsense[rand.Intn(len(nonsense))]+ "\n"+prompt

	//Send the random prophecy on the answer channel
	answer <- prophecyOutput
}

func generateAnswers(pred string, answer chan<- string){
	//time.Sleep(time.Duration(10+rand.Intn(10)) * time.Second)
	//fmt.Println(pred, "AND MY ANSWER IS u SUCK!")
	fmt.Println("\ngenerating answer...\n")
	time.Sleep(5 * time.Second)

	//Declaring and initializing an array with the different answers
	myAnswers := []string{
		"Sry, that wont work.",
		"Yes, what did you think?",
		"Lol, what u trying to say to me?",
		"OMG dont u try me!",
		"No, yall can cashh me outsidee!",
		"Just remember I decide ur future.",
		"ITS ERRRYDAYY BROOO!!",
	}

	myAns := myAnswers[rand.Intn(len(myAnswers))]
	myAns = myAns+"\n"+prompt

	//Send the random answer on the answer channel
	answer <- myAns

}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	fmt.Println("Init initialized...")
	rand.Seed(time.Now().Unix())
}
