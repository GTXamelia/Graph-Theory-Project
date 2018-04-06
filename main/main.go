package main

import (
	".."
	"github.com/fatih/color"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"os"
)

// Function used to read from a file and check it for strings that match a regular expression
func FileReader(){

	// Color setup used for 'Print' so input is on the same line
	blueFmt := color.New(color.FgBlue)

	// Takes in user input and passed it to get postfix
	fmt.Println("")
	blueFmt.Print("Enter regular expression: ")
	var text = automita.GetInput()
	regExp := automita.Intopost(text)

	// Takes in user input to find a '.txt. file
	blueFmt.Print("File name: ")
	var fileName = automita.GetInput()

	// Find file and store value in 'rawFile' if no errors are encountered and 'err' if file encounters any errors
	rawFile, err := ioutil.ReadFile(fileName + ".txt")

	// Send 'rawFile' var to the 'Remove' function and converts return data to string
	refinedFile := strings.Split(string(automita.Remove(rawFile)), " ")

	// Loops though 'refinedFile' adds the index to 'word'
	// 'word' is then send to the 'Pomatch' function
	// if the condition is true then 'count' is incremented
	count := 0
	for _, word := range refinedFile{
		if (automita.Pomatch(regExp, word) == true){
			count++
		}
	}

	// If 'err' doesn't equal 'nil' then it has been written to, and if it has, then an error was encountered
	// else if 'count' is greater than 0 then at least 1 string matched the regular expression
	// else if no statement is true it means that no string matched the regular expression
	if err != nil {
		color.Red("File not found. Please try again") // Red text tells the user that a file was not found
	} else if (count > 0) {
		color.Green("Expression \"" + regExp + "\" - Found: " + strconv.Itoa(count) + " matches") // Green text tells the user a match was found
	} else {
		color.Red("Expression \"" + regExp + "\" - Not found in file") // Red text tells the user 
	}
}

func main() {

	// Color setup used for 'Print' so input is on the same line
	blueFmt := color.New(color.FgBlue)

	// Infinite loop that doesn't end until the user enters 3 and the program ends
	for {
		// Calls 'MainMenu' function which displays the menu options
		automita.MainMenu()
		var menuText = automita.GetInput() // Get user input
		
		if menuText == "1"{ // If 'menuText' equals 1

			// Inform the user to enter a regular expression
			// Takes in user input and seds it to the 'Intopost' function
			fmt.Println("")
			blueFmt.Print("Enter regular expression: ")
			var input1 = automita.Intopost(automita.GetInput())

			// Inform the user to enter a string
			// Get user input and store it in input2
			blueFmt.Print("Enter String: ")
			var input2 = automita.GetInput()

			// Send two inputs to the non-deterministic finite automaton
			nfa := automita.Pomatch(input1, input2)

			// If nfa returns true then the user is told the string and 
			if nfa {
				color.Green("Match = " + "True") // 
			} else {
				color.Red("Match = " + "False")
			}

		} else if menuText == "2" { // If 'menuText' equals 2
			FileReader()
		} else if menuText == "3" { // If 'menuText' equals 3
			fmt.Println("")
			color.Blue("Exiting...")
			os.Exit(2)
		} else { // else if no option is ture then the user has selected and invalid option
			fmt.Println("")
			color.Red("Please select a valid option")
		}
	}
}