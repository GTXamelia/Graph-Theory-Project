package main

import (
	"fmt"
	".."
)

func main() {

	// Get user input and send the R.E to postfix
	fmt.Print("Enter regular expression: ")
	var input1 = automita.GetInput()
	fmt.Println("Postfix: ", automita.Intopost(input1))
	input1 = automita.Intopost(input1)

	// Get user input and store it in input2
	fmt.Print("Enter String: ")
	var input2 = automita.GetInput()

	// Send two inputs to the non-deterministic finite automaton
	nfa := automita.Pomatch(input1, input2)
	fmt.Println("Match = ", nfa)
	fmt.Println("")
}