package htsimple

import (
	"fmt"
	"os"
	"strings"
)

func ApplyData(input_file string, data map[string]string) []byte {
	input, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Printf("Error while reading file: %v", err)
		os.Exit(1)
	}

	output := string(input)
	for key, value := range data {
		fmt.Printf("Data key: %v\n", key)
		output = strings.ReplaceAll(string(output), "{{"+key+"}}", value)
	}

	return []byte(output)
}
