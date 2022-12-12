package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		split := strings.Fields(line)
		data = append(data, split)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return data
}

func coordToString(coord []int) string {
	return strconv.Itoa(coord[0]) + "," + strconv.Itoa(coord[1])
}

func stringToCoord(coords string) ([]int) {
	split := strings.Split(coords, ",")
	row, _ := strconv.Atoi(split[0])
	col, _ := strconv.Atoi(split[1])
	return []int{row, col}
}

// func printGraph(visited map[string]bool) {

// }

func moveTail(head []int, tail []int, visited map[string]bool, pos int) []int {
	// This is going to be like collision, need lots of ifs (could probably consolidate)

	if tail[0] - head[0] == 2 {
		tail[0]--
		// diagonal
		if tail[1] < head[1] {
			tail[1]++
		}
		if tail[1] > head[1] {
			tail[1]--
		}
	}
	if tail[0] - head[0] == -2 {
		tail[0]++
		// diagonal
		if tail[1] < head[1] {
			tail[1]++
		}
		if tail[1] > head[1] {
			tail[1]--
		}
	}
	if tail[1] - head[1] == 2 {
		tail[1]--
		// diagonal
		if tail[0] < head[0] {
			tail[0]++
		}
		if tail[0] > head[0] {
			tail[0]--
		}
	}
	if tail[1] - head[1] == -2 {
		tail[1]++
		// diagonal
		if tail[0] < head[0] {
			tail[0]++
		}
		if tail[0] > head[0] {
			tail[0]--
		}
	}

	// overlap
	// only update the tail as visited
	if pos == 8 {
		visited[coordToString(tail)] = true
	}

	return tail
}

// Part 1
// func main() {
// 	data := importData()

// 	head := []int{0, 0}
// 	tail := []int{0, 0}

// 	visited := make(map[string]bool)

// 	visited[coordToString(tail)] = true

// 	for _, instruction := range data {
// 		dir := instruction[0]
// 		amount, _ := strconv.Atoi(instruction[1])

// 		switch dir {
// 		case "U":
// 			for amount > 0 {
// 				head[0]--
// 				tail = moveTail(head, tail, visited)
// 				amount--
// 			}
// 		case "D":
// 			for amount > 0 {
// 				head[0]++
// 				tail = moveTail(head, tail, visited)
// 				amount--
// 			}
// 		case "L":
// 			for amount > 0 {
// 				head[1]--
// 				tail = moveTail(head, tail, visited)
// 				amount--
// 			}
// 		case "R":
// 			for amount > 0 {
// 				head[1]++
// 				tail = moveTail(head, tail, visited)
// 				amount--
// 			}
// 		}
// 	}

// 	fmt.Println(len(visited))
// }

// Part 2
func main() {
	data := importData()

	head := []int{0, 0}
	tails := make([][]int, 0)

	amount := 9
	for amount > 0 {
		amount--
		tails = append(tails, []int{0, 0})
	}

	visited := make(map[string]bool)

	visited[coordToString([]int{0, 0})] = true

	for _, instruction := range data {
		dir := instruction[0]
		amount, _ := strconv.Atoi(instruction[1])

		switch dir {
		case "U":
			for amount > 0 {
				head[0]--
				tails[0] = moveTail(head, tails[0], visited, 0)
				for tailIndex := 1; tailIndex < len(tails); tailIndex++ {
					tails[tailIndex] = moveTail(tails[tailIndex-1], tails[tailIndex], visited, tailIndex)
				}
				amount--
			}
		case "D":
			for amount > 0 {
				head[0]++
				tails[0] = moveTail(head, tails[0], visited, 0)
				for tailIndex := 1; tailIndex < len(tails); tailIndex++ {
					tails[tailIndex] = moveTail(tails[tailIndex-1], tails[tailIndex], visited, tailIndex)
				}
				amount--
			}
		case "L":
			for amount > 0 {
				head[1]--
				tails[0] = moveTail(head, tails[0], visited, 0)
				for tailIndex := 1; tailIndex < len(tails); tailIndex++ {
					tails[tailIndex] = moveTail(tails[tailIndex-1], tails[tailIndex], visited, tailIndex)
				}
				amount--
			}
		case "R":
			for amount > 0 {
				head[1]++
				tails[0] = moveTail(head, tails[0], visited, 0)
				for tailIndex := 1; tailIndex < len(tails); tailIndex++ {
					tails[tailIndex] = moveTail(tails[tailIndex-1], tails[tailIndex], visited, tailIndex)
				}
				amount--
			}
		}
	}
	fmt.Println(visited)
	fmt.Println(len(visited))
}
