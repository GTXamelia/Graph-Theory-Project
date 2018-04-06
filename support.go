package automita

import (
	"github.com/fatih/color"
	"fmt"
	"bytes"
	"strings"
)

// Auto concatanate a regular expression
// This part is still under development and not complete
func ConcatAuto(s string, n int) string {
	var buffer bytes.Buffer
	var n_1 = n - 1
	var l_1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
		}

		if (rune >= 65 && rune <= 122) && (rune+1 >= 65 && rune+1 <= 122){
			buffer.WriteRune('.')
		}
	}

	t := strings.Replace(buffer.String(), "..", ".", -1)
	t = strings.Replace(t, "..", ".", -1)

	s1 := t
    if last := len(s1) - 1; last >= 0 && s1[last] == '.' {
        s1 = s1[:last]
    }

	return s1
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
	blueFmt := color.New(color.FgBlue)

	fmt.Println("")
	color.Blue("> 1. Compare regular expression to string")
	color.Blue("> 2. Compare regular expression to .txt file")
	color.Blue("> 3. Exit program")
	blueFmt.Print("Enter option: ")
}