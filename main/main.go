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

	fmt.Println("")
	color.Blue("Enter regular expression: ")
	var text = automita.GetInput()
	regExp := automita.Intopost(text)

	color.Blue("File name: ")
	var fileName = automita.GetInput()

	rawFile, err := ioutil.ReadFile(fileName + ".txt")
	if err != nil {
		fmt.Print(err)
	}

	test := automita.Remove(rawFile, 10)
	test = automita.Remove(test, 13)

	refinedFile := strings.Split(string(test), " ")

	for _, word := range refinedFile{
		i++
		if (automita.Pomatch(regExp, word) == true){
			count++
			fmt.Println("Found: ", count)
		}
	}

	if (count > 0) {
		color.Green("Expression \"" + regExp + "\" - Found: " + strconv.Itoa(count) + " matches")
	} else {
		color.Red("Expression \"" + regExp + "\" - Not found in file")
	}
}

func main() {
	for {
		//fmt.Println("Concater: ", automita.ConcatAuto(input1, 1))
		//fmt.Println("Postfix: ", automita.Intopost(input1))

		automita.MainMenu()
		var menuText = automita.GetInput()

		if menuText == "1"{

			fmt.Println("")
			color.Blue("Enter regular expression: ")
			var input1 = automita.Intopost(automita.GetInput())

			// Get user input and store it in input2
			color.Blue("Enter String: ")
			var input2 = automita.GetInput()

			// Send two inputs to the non-deterministic finite automaton
			nfa := automita.Pomatch(input1, input2)
			fmt.Println("")

			if nfa {
				color.Green("Match = " + "True")
			} else {
				color.Red("Match = " + "False")
			}

			fmt.Println("")

		} else if menuText == "2" {
			FileReader()
		} else if menuText == "3" {
			fmt.Println("")
			color.Blue("Exiting...")
			os.Exit(2)
		} else {
			fmt.Println("")
			color.Red("Please select a valid option")
		}
	}
}