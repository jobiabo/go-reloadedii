package main

import (
	"fmt"
	"strings"
)

func LettUp() {
	text := "Ready set go (up)"
	stext := strings.Fields(text)

	for i := 0; i < len(stext); i++ {
		if stext[i] == "(up)" && i > 0 {

			stext[i-1] = strings.ToUpper(stext[i-1])
			stext = append(stext[:i], stext[i+1:]...)
		}
		result := strings.Join(stext, " ")
		fmt.Println(result)
	}
}

func main() {
	LettUp()
}
