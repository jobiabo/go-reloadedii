package main

import (
	"fmt"
	"strings"
)

func CapFirstLetter() {
	text := "Welcome to the Brooklyn bridge (cap)"
	textsplit := strings.Fields(text)

	for i := 0; i < len(textsplit); i++ {
		if textsplit[i] == "(cap)" && i > 0 {

			textsplit[i-1] = strings.Title(textsplit[i-1])
			fmt.Println(textsplit[i-1])
			textsplit = append(textsplit[:i], textsplit[i+1:]...)
		}
		fmt.Println(textsplit)
	}
	result := strings.Join(textsplit, " ")
	fmt.Println(result)
}

func main() {
	CapFirstLetter()

}
