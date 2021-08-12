package main

import (
	"fmt"
	"strings"
)

func main() {
	inputStr := `"To be or not to be," quoth the Bard, "that
is the question". 
The programming contestant replied: "I must disagree. 
To ` + "`" + `C' or not to ` + "`" + `C', that is The Question!"`

	fmt.Println(inputStr)

	//space between input and output
	fmt.Println("")

	arr := strings.Split(inputStr, "\"")
	var count int
	var result string
	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] == "" {
			continue
		}
		if count%2 == 0 {
			result += arr[i] + "``"
		} else {
			result += arr[i] + "''"
		}
		count++
	}

	fmt.Println(result)
}
