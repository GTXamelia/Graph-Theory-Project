package main

import (
	"fmt"
	".."
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/fatih/color"
)

func FileReader(regExp string) string{
	
	count := 0

	fmt.Print("Enter name of text file: ")
	var fileName = automita.GetInput()

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File not found")
	}
	
	str := string(b)

	s := strings.Split(str, " ")

	for _, s := range s{
		if (automita.Pomatch(regExp, s) == true){
			count++
		}
	}

	if (count > 0) {
		return ("Expression \"" + regExp + "\" - Found: " + strconv.Itoa(count))
	}
		
	return ("Expression \"" + regExp + "\" - Not found in file")
}

func main() {

	// Get user input and send the R.E to postfix
	fmt.Println("")
	fmt.Print("Enter regular expression: ")
	var input1 = automita.GetInput()
	input1 = automita.ConcatAuto(input1, 1)
	fmt.Println("Concater: ", automita.ConcatAuto(input1, 1))
	fmt.Println("Postfix: ", automita.Intopost(input1))
	fmt.Println(FileReader(input1))
	input1 = automita.Intopost(input1)

	// Get user input and store it in input2
	fmt.Print("Enter String: ")
	var input2 = automita.GetInput()

	// Send two inputs to the non-deterministic finite automaton
	nfa := automita.Pomatch(input1, input2)
	color.Green("Match = %b", nfa)
	fmt.Println("")
}