package automita

import (
)

type state struct {
	symbol rune
	edge1 *state
	edge2 *state
}

type nfa struct {
	initial *state
	accept  *state
}

func addState(l []*state, s *state, a *state) []*state {
	l = append(l,s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)

		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}

	return l
}

func Poretonfa(pofix string) *nfa {
	nfastack := []*nfa{}

	for _, r := range pofix {
		switch r {
		case '.':
			// take 2 elements off stack (frag2, frag1)
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// Join frag1 to frag2
			frag1.accept.edge1 = frag2.initial

			// append frag1 and frag2 to the top of stack
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			// take 2 elements off stack (frag2, frag1), same as previous
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// Create a state which point to fragments
			accept := state{}
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			
			// Get both states of the fragment
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			// Add to stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '*':
			// Take single element 'frag' off stack
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// Create a state which point to fragments
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			// 'edge1' Points to initial state
			// 'edge2' loops around state
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			// Add to stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '+':
			// Take single element 'frag' off stack
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// Create a state which point to fragments
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			// 'edge1' Points to initial state
			frag.accept.edge1 = &initial

			// Add to stack
			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})
		case '?':
			// Take single element 'frag' off stack
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// Create new state that points to current element and accept state
			initial := state{edge1: frag.initial, edge2: frag.accept}

			// Add to stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})
		default:
			// Create a state which point to fragments
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			// Add to stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}
	return nfastack[0] // Return 
}

// Takes user entered regular expression after it has been turned to postfix and string to compare
func Pomatch(po string, s string) bool {

	// Initial state
	ismatch := false

	// Send postfix regular expression to 'Poretonfa' function
	ponfa := Poretonfa(po)

	// Initial pointers for structs (linked list)
	current := []*state{}
	next := []*state{}

	// Send current, initial and accept state
	current = addState(current[:], ponfa.initial, ponfa.accept)

	// Loop through the users string character by character
	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		// Swaps between 'current' and 'next' to create a new array to search
		current, next = next, []*state{}
	}

	// Loops through current array to find a match
	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	} 

	return ismatch // Return boolean value for if a match is found
}