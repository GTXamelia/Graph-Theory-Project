package main


import (
	"fmt"
	"os"
	"bufio"
)

func intopost(infix string) string {
	specials := map[rune]int{'*':10, '.':9, '|':8}

	pofix, s:= []rune{}, []rune{}

	for _, r := range infix {
		switch {
		case r == '(':
			s = append(s, r)
		case r == ')':
			for s[len(s)-1] != '('{
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		default:
			pofix = append(pofix, r)
		}
	}

	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(pofix)
}

func TrimSuffix(s string) string {
    if len(s) > 0 {
		s = s[:len(s)-2]
	}
    return s
}


func main() {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter query: ")
	input, _ := reader.ReadString('\n')
	input = TrimSuffix(input)
	input = intopost(input) // Remove ending of string
	fmt.Println("Input", input)

}