package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func importData() [][]string{
	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([][]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		entry := []string{line[:len(line)/2], line[len(line)/2:]}

		data = append(data, entry)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return data
}

func main() {
	data := importData()

	pairs := make(map[rune]int)

	// Why isn't there a set?
	for _, pack := range data {
		foundInBoth := make(map[rune]bool)
		frontContent := make(map[rune]bool)
		for _, letter := range pack[0] {
			frontContent[letter] = true
		}

		for _, letter := range pack[1] {
			if frontContent[letter] == true {
				foundInBoth[letter] = true
			}
		}

		for key := range foundInBoth {
			pairs[key]++
		}
	}

	total := 0
	for key, count := range pairs {
		if (unicode.IsLower((key))) {
			total += count * (int(key) - 96)
		} else {
			total += count * (int(key) - 64 + 26) // Unicode to score
		}
	}

	fmt.Println((total))
}
