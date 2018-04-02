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

// This function removed last two values from a string
// Getting user input adds two ascii codes that mess up post fix
// The 13 and 10 are for next line 13 = Carriage return, 10 = Line feed
// These two number would still be evaluated by postfix even if they are invisible to the user
// Example:
// s input = a.b.c*
// s ascii before = [97 46 98 46 99 42 '13 10']
// s ascii after  = [97 98 46 99 '13 10' 42 46]
// s output = ab.c
// *.
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
	fmt.Println("Result", input)

}