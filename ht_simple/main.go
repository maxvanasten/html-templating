package ht_simple

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
		new_output := strings.ReplaceAll(string(output), "{{"+key+"}}", value)
		if new_output == output {
			fmt.Printf("[ht_simple]: Didn't find variable named \"%v\" in file \"%v\"\n", key, input_file)
		}
		output = new_output
	}

	return []byte(output), nil
}
