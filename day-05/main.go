package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Modified from golangcookbook.com
func reverse(numbers []rune) []rune {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func importData() ([]string, map[int][]rune, int) {
	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	crates := make(map[int][]rune)
	instructions := make([]string, 0)

	scanner := bufio.NewScanner(file)

	readBoardMode := true
	lastIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		if readBoardMode {
			if len(line) == 0 {
				readBoardMode = false
				continue
			}
			// Find letters and convert the spacial position of the data to a map of stacks (literally!)
			for i, char := range line {
				if !unicode.IsLetter(char) {
					continue
				} else {
					index := ((i-1) / 4) + 1 // I must be over thinking this

					//[Z] [M] [P] [X] [Y]
 					// 1   2   3   4   5
					// 1   5   9   13  17
					crates[index] = append(crates[index], char)
					lastIndex = index
				}
			}
		} else {
			instructions = append(instructions, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// Reverse the order, because we're appending to the stack top to bottom,
	// resulting in a backwards order
	stackIndex := 1
	for stackIndex <= lastIndex {
		crates[stackIndex] = reverse(crates[stackIndex])
		stackIndex++
	}

	return instructions, crates, lastIndex
}

func printCrates(crates map[int][]rune, lastIndex int) {
	stackIndex := 1
	for stackIndex <= lastIndex {
		output := strconv.Itoa(stackIndex)

		for _, char := range crates[stackIndex] {
			output = output + string(char)
		}
		stackIndex++
		fmt.Println(output)
	}
}

func main() {
	instructions, crates, lastIndex := importData()

	printCrates(crates, lastIndex)

	for _, instruction := range instructions {
		split := strings.Split(instruction, " ")

		number, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])

		if to > lastIndex {
			lastIndex = to
		}

		for i := 0; i < number; i++ {
			// Is there not a better way?
			moved := crates[from][len(crates[from])-1]
			crates[from] = crates[from][:len(crates[from])-1]
			crates[to] = append(crates[to], moved)
		}
	}

	final := ""

	// Need some work to iterate through numerical keys in order
	stackIndex := 1

	for stackIndex <= lastIndex {
		final = final + string(crates[stackIndex][len(crates[stackIndex])-1])

		stackIndex++
	}

	fmt.Println(" ")
	printCrates(crates, lastIndex)
	fmt.Println(" ")
	fmt.Println(final)
}
