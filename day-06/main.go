package main

import (
	"bufio"
	"fmt"
	"os"
)

func importData() []rune {
	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	return []rune(line)
}

func main() {
	data := importData()
	// fmt.Println(data)

	result := 0

	for i := 13; i < len(data); i++ {
		result = i
		// I'm guessing part two is going to be "actually it's 10 unique chars"
		sample := data[i-13:i+1]

		// I feel like there is a more efficient solution, but for now, doing found set in loop
		found := make(map[rune]bool)
		resultFound := false
		for i, char := range sample {
			if found[char] == true {
				break
			} else {
				found[char] = true
				if i == 13 {
					resultFound = true
				}
			}
		}

		if resultFound {
			break
		}
	}

	fmt.Println(result+1) // +1 because zero indexing
}
