package main

import (
	"fmt"
	".."
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/fatih/color"
	"os"
)

func FileReader(){
	
	count := 0
	i := 0
	
	var text string
	fmt.Print("Enter regular expression: ")
	fmt.Scan(&text)

	regExp := automita.Intopost(text)

	b, err := ioutil.ReadFile("../file.txt")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	s := strings.Split(str, " ")

	for _, word := range s{
		i++
		if (automita.Pomatch(regExp, word) == true){
			count++
		}
	}

	if (count > 0) {
		fmt.Println("Expression \"" + regExp + "\" - Found: " + strconv.Itoa(count))
	} else {
		fmt.Println("Expression \"" + regExp + "\" - Not found in file")
	}

	count = 0
	i = 0

}

func main() {
	for {
		// Get user input and send the R.E to postfix
		fmt.Println("")
		
		
		//fmt.Println("Concater: ", automita.ConcatAuto(input1, 1))
		//fmt.Println("Postfix: ", automita.Intopost(input1))

		

		var menuText = automita.GetInput()

		if menuText == "1"{
			fmt.Println("")
			fmt.Print("Enter regular expression: ")
			var input1 = automita.Intopost(automita.GetInput())

			// Get user input and store it in input2
			fmt.Print("Enter String: ")
			var input2 = automita.GetInput()

			// Send two inputs to the non-deterministic finite automaton
			nfa := automita.Pomatch(input1, input2)
			color.Green("Match = %b", nfa)
			fmt.Println("")

		} else if menuText == "2" {
			FileReader()
		} else if menuText == "3" {
			os.Exit(2)
		} else {
			fmt.Println("Please select a valid option")
		}
	}
}