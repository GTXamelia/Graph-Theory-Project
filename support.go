package automita

import (
	"github.com/fatih/color"
	"fmt"
	"bytes"
)

func ConcatAuto(infix string) string {

	var buffer bytes.Buffer
	Arr := []rune(infix)
	
	for infixChar := 0; infixChar < len(infix); infixChar++ {

		if infixChar == 0 {
			
			buffer.WriteString(string(Arr[infixChar]))
			continue
		}
		if specialsCheck(string(Arr[infixChar])) {

			buffer.WriteString(string(Arr[infixChar]))
			continue
		}
		if !(infixChar == 0) && string(Arr[infixChar]) == "(" {

			buffer.WriteString(".")
			buffer.WriteString(string(Arr[infixChar]))
			continue
		}
		if !(infixChar == 0) && string(Arr[infixChar-1]) == "(" {

			buffer.WriteString(string(Arr[infixChar]))
			continue
		}
		if string(Arr[infixChar]) == ")" {

			buffer.WriteString(string(Arr[infixChar]))
			continue
		}
		if (Arr[infixChar] >= 65 && Arr[infixChar] <= 122) {

			buffer.WriteString(".")
			buffer.WriteString(string(Arr[infixChar]))
			continue
		}
	}

	fmt.Println(buffer.String())
	return buffer.String()
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