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

func main() {
	// Setup reader
	reader := bufio.NewReader(os.Stdin)

	// Take in regular expression
	fmt.Print("\nEnter regular expression: ")
	input1, _ := reader.ReadString('\n')
	input1 = automita.TrimSuffix(input1)
	input1 = automita.Intopost(input1)

	// Take in string to compare to regular expression
	fmt.Print("Enter string: ")
	input2, _ := reader.ReadString('\n')
	input2 = automita.TrimSuffix(input2)

	// Send two inputs to the non-deterministic finite automaton
	nfa := automita.Pomatch(input1, input2)
	fmt.Println("Match = ", nfa)
	fmt.Println("")
}