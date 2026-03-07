package main

import (
	"fmt"
	"strings"
	"strconv"
)

func reCode() {
	word := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	splword := strings.Fields(word)

	for i := 0; i < len(splword); i++ {
		if strings.HasPrefix(splword[i], "(up") ||
			strings.HasPrefix(splword[i], "(low") ||
			strings.HasPrefix(splword[i], "(cap)") {

				cleanedi := strings.Trim(splword[i], "()")
				sepi := strings.Split(cleanedi, ",")
				pointi := strings.TrimSpace(sepi[0])

				count := 0

				if len(sepi) > 1 {
					num, err := strconv.Atoi(strings.TrimSpace(sepi[1]))
					if err == nil {
						count = num


					}
					
					for j := 0; j <= count; j++ {
						if i-j < 0 {
							break
						}
						switch pointi {

						case "up" :
							splword[i-j] = strings.ToUpper(splword[i-j])

						case "low" :
							splword[i-j] = strings.ToLower(splword[i-j])

						case "cap" :
							posi := splword[i-j]
							if len(posi) > 0 {
								splword[i-j] = strings.ToUpper(string(posi[0])) + posi[1:]
							}

						}
					}
					splword = append(splword[:i], splword[i+1:]...)
					i--
				}
			}
	}

	result := strings.Join(splword, " ")
	fmt.Println(result)
}

func main() {
	reCode()
}