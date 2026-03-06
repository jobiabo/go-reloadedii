package main

import (
	"fmt"
	"strings"
)

func CapToLowerMod() {
	text := "This is SO EXCITING (up)"
	stext := strings.Fields(text)
	fmt.Println(stext)

	for i := 0; i < len(stext); i++ {
		if stext[i] == "(up)" && i > 0 {

			stext[i-1] = strings.ToLower(stext[i-1])
			stext[i-2] = strings.ToLower(stext[i-2])
			stext = append(stext[:i], stext[i+1:]...)
			fmt.Println(stext)
		}
		result := strings.Join(stext, " ")
		fmt.Println(result)
	}

}

func main() {
	CapToLowerMod()
}
