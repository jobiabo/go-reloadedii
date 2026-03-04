package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hexBin() {
	text := "I have 1E (Hex) files"
	tsplit := strings.Fields(text)
	fmt.Println(text)
	fmt.Println(tsplit)

	for i := 0; i < len(tsplit); i++ {
		if tsplit[i] == "(Hex)" && i > 0 {
			num, err := strconv.ParseInt(tsplit[i-1], 16, 64)
			if err != nil {
				fmt.Println("invalid hex")
				return
			}
			tsplit[i-1] = fmt.Sprintf("%d", num)
			tsplit = append(tsplit[:i], tsplit[i+1:]...)
			i--
		}
		result := strings.Join(tsplit, "")
		fmt.Println(result)

	}
}

func main() {
	hexBin()
}
