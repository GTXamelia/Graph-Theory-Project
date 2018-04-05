package automita

import (
	"fmt"
	"bytes"
	"strings"
	"github.com/fatih/color"
)

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

func Remove(s []byte, r byte) []byte {
	for _, k := range s {
		for i, v := range s {
			if v == r {
				k--
				s = append(s[:i], 32)
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
    return s
}

func GetInput() string {
	var input string
	fmt.Scan(&input)
	return input
}

func MainMenu() {
	fmt.Println("")
	color.Blue("> 1. Compare regular expression to string")
	color.Blue("> 2. Compare regular expression to .txt file")
	color.Blue("> 3. Exit program")
	color.Blue("Enter option: ")
}