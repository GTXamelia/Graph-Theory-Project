package automita

import (
	"github.com/fatih/color"
	"fmt"
	"bytes"
)

// Concatanates a regular expression before postfix
func ConcatAuto(infix string) string {

	// String variables
	var buffer bytes.Buffer
	Arr := []rune(infix)

	// Boolean used for '|' statement
	check := false
	
	// Loops through string character by character
	for infixChar := 0; infixChar < len(infix); infixChar++ {

		// If character == 0
		if infixChar == 0 {

			buffer.WriteString(string(Arr[infixChar])) // add character to buffer

		}else if string(Arr[infixChar]) == "." {	
			// Do nothing
		}else if specialsCheck(string(Arr[infixChar])) {

			// If character == "|" then check is true this is used later to concatanate or not
			if !(string(Arr[infixChar]) == "|") {
				check = false
			}else{
				check = true
			}

			buffer.WriteString(string(Arr[infixChar])) // add character to buffer
			
		}else if !(infixChar == 0) && string(Arr[infixChar]) == "(" {

			buffer.WriteString(".") // add to buffer
			buffer.WriteString(string(Arr[infixChar])) // Add character to buffer// Add character to buffer

		}else if !(infixChar == 0) && string(Arr[infixChar-1]) == "(" { // if 

			buffer.WriteString(string(Arr[infixChar])) // add character to buffer

		}else if string(Arr[infixChar]) == ")" {

			buffer.WriteString(string(Arr[infixChar])) // add character to buffer

		}else if check { // check is true if last

			buffer.WriteString(string(Arr[infixChar])) // Add character to buffer
			check = false
			
		}else {

			buffer.WriteString(".") // add to buffer
			buffer.WriteString(string(Arr[infixChar])) // Add character to buffer

		}
	}

	return buffer.String() // Return string
}

func specialsCheck(char string) bool {

	specials := []string{"*", "+", "?", ".", "|"}

	for specialChar := range specials {
		if char == specials[specialChar] {
			return true
		}
	}
	return false
}

// Remove new line and return so data is all on one line
func Remove(s []byte) []byte {
	
	s = bytes.Replace([]byte(s), []byte("\r\n"), []byte(" "), -1) // Replace '\r\n' which is return and line feed. These are replaced by ' ' 

    return s // Return byte array
}

// Returns a string which is taken in the the 'fmt.Scan'
func GetInput() string {
	var input string
	fmt.Scan(&input)
	return input
}

// ManinMenu function displays the menu options for the user
func MainMenu() {
	// Color setup used for 'Print' so input is on the same line as the prompt
	blueFmt := color.New(color.FgBlue)

	fmt.Println("")
	color.Blue("> 1. Compare regular expression to string")
	color.Blue("> 2. Compare regular expression to .txt file")
	color.Blue("> 3. Exit program")
	blueFmt.Print("Enter option: ")
}