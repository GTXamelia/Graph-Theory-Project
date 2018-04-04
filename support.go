package automita

import (
	"fmt"
	"bytes"
	"strings"
	"io/ioutil"
)

func FileReader(fileName string, regExp string) string{
	
	count := 0

	fmt.Print("Enter name of text file: ")
	var input1 = GetInput()

	b, err := ioutil.ReadFile(fileName) // just pass the file name

	// Print error
	if err != nil {
		return "File not found"
	}
	
	str := string(b)

	//split the string into token words
	s := strings.Split(str, " ")

	for i, word := range s{

		// fmt.Println(pomatch(text, word))
		if(pomatch(regExp, s[i]) == true){
			count++
		// st[count] = strconv.Itoa(i)
		fmt.Println("The word" +  " " + word + " appears " + strconv.Itoa(i) + " words in ")
		}
	}

	if (count > 0) {
	
		fmt.Println("Yes! The expression " + text + " exists in the text file " + strconv.Itoa(count) + " times.")

	

	} else {
		fmt.Println("No! The expression " + text + " does not exist in the text document.")
	}
	// fmt.Println(pomatch("I+", "I"))
		

	//fmt.Println(str)
	
	//fmt.Println(pomatch("ab.*c*|", "abab"))

	count = 0
	i = 0

}

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

func GetInput() string {
	var input string
	fmt.Scan(&input)
	return input
}