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

			// add frag1 to frag2
			frag1.accept.edge1 = frag2.initial

			// append frag1 and frag2 to the top of stack
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			// take 2 elements off stack (frag2, frag1), same as previous
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '*':
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '+':
			frag := nfastack[len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = &initial

			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})
		case '?':
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag.initial, edge2: frag.accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}

	return nfastack[0]
}

func Pomatch(po string, s string) bool {
	ismatch := false
	ponfa := Poretonfa(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		current, next = next, []*state{}
	}

	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	} 

	return ismatch
}