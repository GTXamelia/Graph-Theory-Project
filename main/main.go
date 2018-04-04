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
	
	color.Blue("Enter regular expression: ")
	var text = automita.GetInput()

	regExp := automita.Intopost(text)

	b, err := ioutil.ReadFile("file.txt")
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
		color.Green("Expression \"" + regExp + "\" - Found: " + strconv.Itoa(count))
	} else {
		color.Red("Expression \"" + regExp + "\" - Not found in file")
	}

	count = 0
	i = 0
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

			if nfa {
				color.Green("Match = " + "True")
			} else {
				color.Red("Match = " + "False")
			}

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