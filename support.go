package automita

import (
	"github.com/fatih/color"
	"fmt"
	"bytes"
)

func ConcatAuto(infix string) string {

	var buffer bytes.Buffer
	strArr := []rune(infix)
	
	for infixChar := 0; infixChar < len(infix); infixChar++ {

		if infixChar ==0 {
			buffer.WriteString(string(strArr[infixChar]))
			continue
		}
		
		if specialsCheck(string(strArr[infixChar])) {

			buffer.WriteString(string(strArr[infixChar]))
			continue
		}
		if !(infixChar == 0) && string(strArr[infixChar]) == "(" {
			buffer.WriteString(".")
			buffer.WriteString(string(strArr[infixChar]))
			continue
		}
		if !(infixChar == 0) && string(strArr[infixChar-1]) == "(" {

			buffer.WriteString(string(strArr[infixChar]))
			continue
		}
		if string(strArr[infixChar]) == ")" {

			buffer.WriteString(string(strArr[infixChar]))
			continue
		}
		buffer.WriteString(".")
		buffer.WriteString(string(strArr[infixChar]))

	}

	fmt.Println(buffer.String())
	return buffer.String()
}

func specialsCheck(char string) bool {

	specials := []string{"*", "+", "?", ".", "|"}

	for spec := range specials {
		if char == specials[spec] {
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