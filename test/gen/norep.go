package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countUniqueTokens(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	tokenSet := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token := strings.TrimSpace(scanner.Text())
		if token != "" {
			tokenSet[token] = true
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return len(tokenSet), nil
}

func main() {
	filePath := "tokens.txt"

	uniqueTokensCount, err := countUniqueTokens(filePath)
	if err != nil {
		fmt.Println("Error counting unique tokens:", err)
		return
	}

	fmt.Printf("Generated %d unique tokens.\n", uniqueTokensCount)
}
