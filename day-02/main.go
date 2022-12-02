package main

import (
	"bufio"
	"fmt"
	"os"
)



func importData() [][]rune{
	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([][]rune, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		entry := []rune{[]rune(line)[0], []rune(line)[2]}
		data = append(data, entry)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return data
}

// Part 1
// func main() {
// 	WIN_POINTS := 6
// 	TIE_POINTS := 3

// 	points := map[rune]int {
// 		'A': 1, // Rock
// 		'B': 2, // Paper
// 		'C': 3, // Scissors
// 		'X': 1, // Rock
// 		'Y': 2, // Paper
// 		'Z': 3, // Scissors
// 	}

// 	tie := map[rune]rune {
// 		'X': 'A',
// 		'Y': 'B',
// 		'Z': 'C',
// 	}

// 	win := map[rune]rune {
// 		'X': 'C',
// 		'Y': 'A',
// 		'Z': 'B',
// 	}

// 	data := importData()
// 	score := 0

// 	for _, pair := range data {
// 		score += points[pair[1]]
// 		if pair[0] == tie[pair[1]] {
// 			score += TIE_POINTS
// 		} else if pair[0] == win[pair[1]] {
// 			score += WIN_POINTS
// 		}
// 	}

// 	fmt.Println(score)
// }

// Part 2
func main() {
	WIN_POINTS := 6
	TIE_POINTS := 3

	points := map[rune]int {
		'A': 1, // Rock
		'B': 2, // Paper
		'C': 3, // Scissors
	}

	win := map[rune]rune {
		'A': 'B',
		'B': 'C',
		'C': 'A',
	}

	lose := map[rune]rune {
		'A': 'C',
		'B': 'A',
		'C': 'B',
	}

	data := importData()
	score := 0

	for _, pair := range data {
		if pair[1] == 'X' { // Lose
			selection := lose[pair[0]]
			score += points[selection]
		} else if pair[1] == 'Y' { // Draw
			selection := pair[0]
			score += points[selection] + TIE_POINTS
		} else if pair[1] == 'Z' { // Win
			selection := win[pair[0]]
			score += points[selection] +  WIN_POINTS
		} else {
			fmt.Println("Error: Bad data")
		}
	}

	fmt.Println(score)
}
