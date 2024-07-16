package asciiart

import (
	"fmt"
	"strings"
)

/*
DisplayAsciiArt generates ASCII art from a given character map and input string. It:
1. Replaces CR with "\\n" and splits the input into lines.
2. Iterates through each line, using the map to convert characters to ASCII art.
3. Returns the complete ASCII art or an empty string if any character is not found.
*/
func DisplayAsciiArt(characterMap map[rune][]string, input string) string {
	inputRaw := strings.Replace(input, "\r\n", "\n", -1)
	inputSlice := strings.Split(inputRaw, "\n")
	var result strings.Builder
	for _, value := range inputSlice {
		for i := 0; i < 8; i++ {
			for _, char := range value {
				line, ok := characterMap[char]
				if ok {
					result.WriteString(fmt.Sprintf("%s ", line[i]))
				} else {
					result.WriteString("Characters not found")
					return ""
				}
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
