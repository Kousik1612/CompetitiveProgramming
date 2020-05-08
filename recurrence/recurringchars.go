package recurrence

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// RecurringChars : Program to find out the number of recurring characters in a given string.
func RecurringChars() {	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	occurence := make(map[string]int, 0)
	result := make(map[string]int, 0)
	for _, r := range text {
		val := string(r)
		if _, ok := occurence[val]; ok {
			occurence[val]++
			result[val] = occurence[val]
		} else {
			occurence[val] = 1
		}
	}
	// Prints only the recurrent character(s).
	fmt.Println(result)
	// Prints all the character input with the number of recurrence of each character.
	fmt.Println(occurence)
}
