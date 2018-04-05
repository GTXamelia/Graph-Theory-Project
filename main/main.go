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

	blueFmt := color.New(color.FgBlue)

	count := 0
	i := 0

	fmt.Println("")
	blueFmt.Print("Enter regular expression: ")
	var text = automita.GetInput()
	regExp := automita.Intopost(text)

	blueFmt.Print("File name: ")
	var fileName = automita.GetInput()

	rawFile, err := ioutil.ReadFile(fileName + ".txt")

	refinedFile := strings.Split(string(automita.Remove(rawFile)), " ")

	for _, word := range refinedFile{
		i++
		if (automita.Pomatch(regExp, word) == true){
			count++
		}
	}

	if err != nil {
		color.Red("File not found. Please try again")
	} else if (count > 0) {
		color.Green("Expression \"" + regExp + "\" - Found: " + strconv.Itoa(count) + " matches")
	} else {
		color.Red("Expression \"" + regExp + "\" - Not found in file")
	}
}

func main() {
	blueFmt := color.New(color.FgBlue)

	for {
		//fmt.Println("Concater: ", automita.ConcatAuto(input1, 1))
		//fmt.Println("Postfix: ", automita.Intopost(input1))

		automita.MainMenu()
		var menuText = automita.GetInput()

		if menuText == "1"{

			fmt.Println("")
			blueFmt.Print("Enter regular expression: ")
			var input1 = automita.Intopost(automita.GetInput())

			// Get user input and store it in input2
			blueFmt.Print("Enter String: ")
			var input2 = automita.GetInput()

			// Send two inputs to the non-deterministic finite automaton
			nfa := automita.Pomatch(input1, input2)

			if nfa {
				color.Green("Match = " + "True")
			} else {
				color.Red("Match = " + "False")
			}

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