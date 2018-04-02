package main

import (
	"fmt"
	"os"
	"bufio"
	".."
)

// This function removed last two values from a string
// Getting user input adds two ascii codes that mess up post fix
// The 13 and 10 are for next line 13 = Carriage return, 10 = Line feed
// These two number would still be evaluated by postfix even if they are invisible to the user
// Example:
// input = a.b.c*
// ascii before = [97 46 98 46 99 42 '13 10']
// ascii after  = [97 98 46 99 '13 10' 42 46]
// output = ab.c
// *.
func TrimSuffix(s string) string {
    if len(s) > 0 {
		s = s[:len(s)-2]
	}
    return s
}

type state struct {
	symbol rune
	edge1 *state
	edge2 *state
}

func main() {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter query: ")
	input, _ := reader.ReadString('\n')
	input = TrimSuffix(input)
	input = postfix.Intopost(input) // Remove ending of string
	fmt.Println("Result", input)

}