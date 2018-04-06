package automita

import (
)

// This function changes an infix to a postfix which can be read by the nfa
func Intopost(infix string) string {

	// Map of special characters with presedence
	specials := map[rune]int{'*': 10, '+': 9, '?': 8, '.': 7, '|': 6,}

	// Run array setup for 'pofix' and 's'
	pofix, s := []rune{}, []rune{}

	// Loop over infix character by character
	for _, r := range infix {
		switch {
		case r == '(': // if r == '(' 
			s = append(s, r)
		case r == ')': // if r == ')' 
			// Loop through stack until open bracket is found
			for s[len(s)-1] != '('{
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = s[:len(s)-1]
		case specials[r] > 0: // if one of the special characters is found
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			// Appends character to stack
			s = append(s, r)
		default: // Anything else. Not one of the specials or a bracket
			pofix = append(pofix, r)
		}
	}

	// Appends leftovers to pofix
	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(pofix) // Return pofix in string format
}