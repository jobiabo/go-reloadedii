package main

import (
	"fmt"
	"strings"
)

func CapToUpperMod() {
	text := "Ready, set, go (up) !" 
	textSplit := strings.Fields(text)

	for i := 0; i < len(textSplit); i++ {
		if textSplit[i] == "(up)" && i > 0 {

			textSplit[i-1] = strings.ToUpper(textSplit[i-1])
			fmt.Println(textSplit[i-1])
			textSplit[i-2] = strings.ToUpper(textSplit[i-2])
			fmt.Println(textSplit[i-2])
			textSplit = append(textSplit[:i], textSplit[i+1:]...)
			fmt.Println(textSplit)
		}
	}
	result := strings.Join(textSplit, " ")
	fmt.Println(result)
}

func main() {
	CapToUpperMod()
}