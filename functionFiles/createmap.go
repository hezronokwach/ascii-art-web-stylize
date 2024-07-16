package asciiart

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"os"
	"path/filepath"
)

/*
CreateMap reads from a specified .txt file to construct a map of ASCII art representations for characters.
It validates the file's existence, checks the checksum to ensure integrity, and parses the content into a map.
If errors occur (file not found, wrong format, modified content), it returns an error.
Each character's ASCII art should be structured in up to 8 lines of text, separated by empty lines for different characters.
*/
func CreateMap(fileName string) (map[rune][]string, error) {
	fileString := fmt.Sprintf("%s%s", fileName, ".txt")
	file, err := os.Open(fileString)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist:", fileName)
		} else {
			fmt.Println("Error checking file status:", err)
		}
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var data []byte
	for _, value := range lines {
		for _, value2 := range value {
			data = append(data, byte(value2))
		}
	}
	crc32Table := crc32.MakeTable(crc32.IEEE)
	checksum := crc32.Checksum(data, crc32Table)
	if !(checksum == 0x9ffd59bc || checksum == 0x2f465361 || checksum == 0x6ee86a07 || checksum == 1501001935) {
		return nil, fmt.Errorf("file modified")
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	if filepath.Ext(fileString) != ".txt" {
		fmt.Println("Wrong extension, use .txt")
		return nil, err
	}
	characterMap := make(map[rune][]string)
	var currentChar rune = ' '
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if len(characterMap[currentChar]) == 8 {
			currentChar++
			characterMap[currentChar] = []string{}
		}
		characterMap[currentChar] = append(characterMap[currentChar], line)
	}
	return characterMap, nil
}
