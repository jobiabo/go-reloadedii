package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BinW() {
	word := "(I have 10 (bin) years)"
	sword := strings.Fields(word)

	for i := 0; i < len(sword); i++ {
		if sword[i] == "(bin)" && i > 0 {

			num, err := strconv.ParseInt(sword[i-1], 2, 64)
			if err != nil {
				fmt.Println("an error occured")
				return
			}
			sword[i-1] = fmt.Sprintf("%d", num)
			sword = append(sword[:i], sword[i+1:]...)
			i--
		}
		result := strings.Join(sword, "")
		fmt.Println(result)
	}
}

func main() {
	BinW()
}
