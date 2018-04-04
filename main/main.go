package main

import (
	"fmt"
	".."
	"io/ioutil"
	"strings"
)

func FileReader(regExp string) string{
	
	count := 0

	fmt.Print("Enter name of text file: ")
	var fileName = automita.GetInput()

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "File not found"
	}
	
	str := string(b)

	s := strings.Split(str, " ")

	for i, word := range s{
		if (automita.Pomatch(regExp, s[i]) == true){
			count++
		}
	}

	if (count > 0) {
		
	} else {
		
	}

	return "nil"
}

func main() {

	// Get user input and send the R.E to postfix
	fmt.Println("")
	fmt.Print("Enter regular expression: ")
	var input1 = automita.GetInput()
	input1 = automita.ConcatAuto(input1, 1)
	fmt.Println("Concater: ", automita.ConcatAuto(input1, 1))
	fmt.Println("Postfix: ", automita.Intopost(input1))
	input1 = automita.Intopost(input1)



	// Get user input and store it in input2
	fmt.Print("Enter String: ")
	var input2 = automita.GetInput()

	// Send two inputs to the non-deterministic finite automaton
	nfa := automita.Pomatch(input1, input2)
	fmt.Println("Match = ", color.Green(nfa))
	fmt.Println("")
}