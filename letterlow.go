package main

import (
	"fmt"
	"strings"
)

func LettDown() {
	text := "I should stop SHOUTING (low)"
	stext := strings.Fields(text)
	fmt.Println(text)
	fmt.Println(stext)

	for i := 0; i < len(stext); i++ {
		if stext[i] == "(low)" && i > 0 {

			stext[i-1] = strings.ToLower(stext[i-1])
			stext = append(stext[:i], stext[i+1:]...)
			i--
		}
		result := strings.Join(stext, " ")
		fmt.Println(result)
	}
}

func main() {
	LettDown()
}
