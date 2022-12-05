package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Part 1
// func importData() [][]string{
// 	file, err := os.Open("data.txt")

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer file.Close()

// 	data := make([][]string, 0)

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		entry := []string{line[:len(line)/2], line[len(line)/2:]}

// 		data = append(data, entry)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println(err)
// 	}
// 	return data
// }

// Part 1
// func main() {
// 	data := importData()

// 	pairs := make(map[rune]int)

// 	// Why isn't there a set?
// 	for _, pack := range data {
// 		foundInBoth := make(map[rune]bool)
// 		frontContent := make(map[rune]bool)
// 		for _, letter := range pack[0] {
// 			frontContent[letter] = true
// 		}

// 		for _, letter := range pack[1] {
// 			if frontContent[letter] == true {
// 				foundInBoth[letter] = true
// 			}
// 		}

// 		for key := range foundInBoth {
// 			pairs[key]++
// 		}
// 	}

// 	total := 0
// 	for key, count := range pairs {
// 		if (unicode.IsLower((key))) {
// 			total += count * (int(key) - 96)
// 		} else {
// 			total += count * (int(key) - 64 + 26) // Unicode to score
// 		}
// 	}

// 	fmt.Println((total))
// }

func importDataFlat() []string{
	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return data
}

func comparePacks(foundInOne map[rune]bool, foundInTwo map[rune]bool, foundInThree map[rune]bool) rune{
	// Find keys in both 1 and 2 at O(n) where n == len(foundInOne)
	foundInOneAndTwo := make(map[rune]bool)
	for key1 := range foundInOne {
		if foundInTwo[key1] == true {
			foundInOneAndTwo[key1] = true
		}
	}

	// Find keys also in the third at O(n) where n = len(foundInOneAndTwo)
	for key1n2 := range foundInOneAndTwo {
		if foundInThree[key1n2] == true {
			// fmt.Println(key1n2, string(key1n2))
			return key1n2
		}
	}

	// Shouldn't not find an answer
	panic("Didn't find a badge")
}

// Part 2
func main() {
	data := importDataFlat()

	badges := make(map[rune]int)

	for i := 0; i < len(data); i += 3 {
		// fmt.Println(data[i][0], string(data[i][0]))
		foundInOne := make(map[rune]bool)
		foundInTwo := make(map[rune]bool)
		foundInThree := make(map[rune]bool)

		for _, char := range data[i] {
			foundInOne[char] = true
		}
		for _, char := range data[i+1] {
			foundInTwo[char] = true
		}
		for _, char := range data[i+2] {
			foundInThree[char] = true
		}

		badges[comparePacks(foundInOne, foundInTwo, foundInThree)] += 1
	}
	fmt.Println(badges)
	total := 0
	for key, count := range badges {
		if (unicode.IsLower((key))) {
			total += count * (int(key) - 96)
		} else {
			total += count * (int(key) - 64 + 26) // Unicode to score
		}
	}
	fmt.Println(total)
}
