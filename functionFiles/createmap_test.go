package asciiart

import (
	"os"
	"testing"
)

/*
TestReadMap evaluates the CreateMap function by creating a temporary file with ASCII art content,
and verifying if CreateMap accurately parses this content into a character map. This test checks the correctness
of the output map in terms of size, key presence, and line accuracy against expected values.
*/
func TestReadMap(t *testing.T) {
	testFilePath := "test2"
	content := "\n1\n2\n3\n4\n5\n6\n7\n8\n\n9\n10\n11\n12\n13\n14\n15\n16"
	err := os.WriteFile(testFilePath, []byte(content), 0o644)
	if err != nil {
		t.Errorf("Error writing to file %s :", err)
	}
	expectedMap := map[rune][]string{
		' ': {"1", "2", "3", "4", "5", "6", "7", "8"},
		'!': {"9", "10", "11", "12", "13", "14", "15", "16"},
	}
	resultMap, err := CreateMap(testFilePath)
	if err != nil {
		t.Errorf("Error reading file %v", err)
		return
	}
	if len(resultMap) != len(expectedMap) {
		t.Errorf("Expected map size %d , got %d", len(expectedMap), len(resultMap))
		return
	}
	for key, expectedValue := range expectedMap {
		actualValue, ok := resultMap[key]
		if !ok {
			t.Errorf("Missing key %v in resultMap", key)
			continue
		}
		if len(actualValue) != len(expectedValue) {
			t.Errorf("Expected %d lines for key %v got %d", len(actualValue), key, len(resultMap))
			continue
		}
		for index, line := range expectedValue {
			if actualValue[index] != line {
				t.Errorf("Expected line %d for key %v : %s ,got : %s", index+1, key, line, actualValue[index])
			}
		}

	}
}
