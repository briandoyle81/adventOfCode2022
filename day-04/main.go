package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func importData() [][]int{
	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([][]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		pairs := strings.Split(line, ",")

		first := strings.Split(pairs[0], "-")
		second := strings.Split(pairs[1], "-")

		firstLeftInt, _ := strconv.Atoi(first[0])
		firstRightInt, _ := strconv.Atoi(first[1])

		secondLeftInt, _ := strconv.Atoi(second[0])
		secondRightInt, _ := strconv.Atoi(second[1])


		entry := []int{firstLeftInt, firstRightInt, secondLeftInt, secondRightInt}

		data = append(data, entry)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return data
}

func main() {
	data := importData()

	overLappedPairs := 0

	for _, assignment := range data {
		if (assignment[0] <= assignment[2] && assignment[1] >= assignment[3]) ||
		   (assignment[2] <= assignment[0] && assignment[3] >= assignment[1]) {
			overLappedPairs++
		   }
	}

	fmt.Println(overLappedPairs)
}
