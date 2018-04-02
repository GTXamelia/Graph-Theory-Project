package main


import (
	"fmt"
)

func intopost(infix string) string {
	pofix := []rune{}
	s:= []rune{}




	return string(pofix)
}


func main() {
	
	fmt.Println("Infix:		","a.b.c")
	fmt.Println("Postfix 	",intopost("a.b.c"))

	fmt.Println("Infix:		","(a.(b|d))*")
	fmt.Println("Postfix 	",intopost("(a.(b|d))*"))

	fmt.Println("Infix:		","a.(b|d).c*")
	fmt.Println("Postfix 	",intopost("a.(b|d).c*"))

	fmt.Println("Infix:		","a.(b.b)+.c")
	fmt.Println("Postfix 	",intopost("a.(b.b)+.c"))

}