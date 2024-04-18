package hte

import (
	"fmt"
	"os"
	"strings"
)

func ApplyData(input_file string, data map[string]string) ([]byte, error) {
	input, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Printf("Error while reading file: %v", err)
		return nil, err
	}

	output := string(input)
	for key, value := range data {
		output = strings.ReplaceAll(string(output), "{{"+key+"}}", value)
	}

	return []byte(output), nil
}
