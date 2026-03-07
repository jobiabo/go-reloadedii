package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CapToUpperMod() {

	text := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	textSplit := strings.Fields(text)

	for i := 0; i < len(textSplit); i++ {

		// OLD CONDITION (only worked for "(up)")
		// if textSplit[i] == "(up)" && i > 0 {

		// NEW: detect modifiers with numbers like "(up, 2)"
		if strings.HasPrefix(textSplit[i], "(up") ||
			strings.HasPrefix(textSplit[i], "(low") ||
			strings.HasPrefix(textSplit[i], "(cap") {

			// REMOVE parentheses
			mod := strings.Trim(textSplit[i], "()")

			// split "up, 2"
			parts := strings.Split(mod, ",")

			action := strings.TrimSpace(parts[0])
			count := 1 // default if number not provided

			// NEW: read the number safely
			if len(parts) > 1 {
				n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil {
					count = n
				}
			}

			// NEW EDGE CASE FIX
			// Your code used:
			// textSplit[i-1] and textSplit[i-2]
			// which could panic if there are not enough words.

			for j := 1; j <= count; j++ {

				// stop if we reach start of slice
				if i-j < 0 {
					break
				}

				switch action {

				case "up":
					// OLD CODE
					// textSplit[i-1] = strings.ToUpper(textSplit[i-1])
					// textSplit[i-2] = strings.ToUpper(textSplit[i-2])

					// NEW: supports variable count
					textSplit[i-j] = strings.ToUpper(textSplit[i-j])

				case "low":
					textSplit[i-j] = strings.ToLower(textSplit[i-j])

				case "cap":
					word := strings.ToLower(textSplit[i-j])
					if len(word) > 0 {
						textSplit[i-j] = strings.ToUpper(string(word[0])) + word[1:]
					}
				}
			}

			// OLD removal (kept but improved)
			// textSplit = append(textSplit[:i], textSplit[i+1:]...)

			textSplit = append(textSplit[:i], textSplit[i+1:]...)

			// EDGE CASE FIX
			// When removing elements while looping,
			// we must step back so we don't skip elements.
			i--
		}
	}

	result := strings.Join(textSplit, " ")
	fmt.Println(result)
}

func main() {
	CapToUpperMod()
}
