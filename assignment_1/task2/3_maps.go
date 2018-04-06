//Exercise: Maps -- TASK 2

/************************* TASK DESCRIPTION ************************************/
	/*	- Implement the function WordCount
			- It should return a map of the counts of each "rod" in the string s
			- The wc.Test function runs a test suite against the provided func
					- And prints success or failure
			- You might find strings.Fields helpful
	*/

package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {

	//declare word vector out of sentence, split on space
	wordVec := strings.Fields(s)

	//declare map for outputting words and their counts
	totalCount := make(map[string]int)

	//looping over each word in the sentence
	//adding 1 for each occurance of the word thus counting freq
	for word := range wordVec{
		//COULD ALSO LOOP OVER index, word
		totalCount[wordVec[word]] +=1
	}

	return totalCount
}

func main() {
	fmt.Println("\n--------------- Pre existing test suite on the word counter ---------------")
	wc.Test(WordCount)
	fmt.Println("-----------------------------------------------------------------------------")
	ownTestString := "hey, this is my own test string, hey hey hey this go language is so fun lets go"

	outputMap := WordCount(ownTestString)
	fmt.Println("\n--------------- My own test of the word counter ---------------")
	for outWord, outCount := range outputMap{
		fmt.Println("- ",outWord,outCount)
	}
	fmt.Println("---------------------------------------------------------------")
}

/**************************** QUESTIONS & ANSWERS ***************************************/
//No questions
