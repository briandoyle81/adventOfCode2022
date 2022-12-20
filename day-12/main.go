package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	row int
	col int
}

func importData() ([][]rune, Position, Position){
	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([][]rune, 0)
	var start Position
	var target Position

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, make([]rune, 0))
		for col, char := range line {
			data[row] = append(data[row], char)
			if char == 'S' {
				start = Position{row: row, col: col}
			}
			if char == 'E' {
				target = Position{row: row, col: col}
			}
		}
		row++
	}

	return data, start, target
}

func coordToString(coord Position) string {
	return strconv.Itoa(coord.row) + "," + strconv.Itoa(coord.col)
}

func stringToCoord(coords string) ([]int) {
	split := strings.Split(coords, ",")
	row, _ := strconv.Atoi(split[0])
	col, _ := strconv.Atoi(split[1])
	return []int{row, col}
}

func printGraph(graph [][]rune) {
	for _, rowData := range graph {
		for _, char := range rowData {
			fmt.Print(string(char))
		}
		fmt.Print("\n")
	}
}

func dfs(start Position, end Position, mapData [][]rune) int {
	qq := make([][]Position, 0)

	qq = append(qq, []Position{start})

	found := make(map[string]bool)
	found[coordToString(start)] = true

	for len(qq) > 0 {
		// Pop the end of the queue
		currentPath := qq[0]
		qq = qq[1:]
		current := currentPath[len(currentPath)-1]

		// For each neighbor not in visited
		for i := -1; i <= 1; i++ {
			for k := -1; k <= 1; k++ {
				// Skip diagonals
				if i != 0 && k != 0 {
					continue
				}

				neighborRow := current.row + i
				neighborCol := current.col + k
				neighborPos := Position{row: neighborRow, col: neighborCol}

				// skip out of bounds
				if neighborRow < 0 || neighborCol < 0 {
					continue
				}

				if neighborRow >= len(mapData) || neighborCol >= len(mapData[0]) {
					continue
				}

				// skip found
				if found[coordToString(neighborPos)] {
					continue
				}

				height := mapData[current.row][current.col]
				neighborHeight := mapData[neighborRow][neighborCol]
				adjustedHeight := neighborHeight

				// Deal with annoying E
				if neighborHeight == 'E' {
					adjustedHeight = 'z'
				}

				// Can only go up one, can go down any, or to end
				if adjustedHeight - height > 1 && height != 'S' {
					continue
				}

				newPath := make([]Position, 0)
				newPath = append(newPath, currentPath...)
				newPath = append(newPath, neighborPos)

				found[coordToString(neighborPos)] = true

				if neighborHeight == 'E' {
					return len(newPath)-1
				}

				qq = append(qq, newPath)
			}
		}
	}
	return math.MaxInt
}

func printPath(path []Position, mapData [][]rune) {
	for _, pos := range path {
		fmt.Println(string(mapData[pos.row][pos.col]), " ", pos, " -> ")
	}
}

// Part 1
// func main() {
// 	data, start, target := importData()
// 	// printGraph(data)

// 	length := dfs(start, target, data)
// 	fmt.Println(length)
// 	// printPath(path, data)

// 	// for _, coords := range path {
// 	// 	data[coords.row][coords.col] = '.'
// 	// }
// 	// printGraph(data)
// }

// Part 2
func main() {
	data, _, target := importData()
	// printGraph(data)

	shortest := math.MaxInt

	// Just do each of four edges one at a time

	for col := range data[0] {
		length := dfs(Position{row:0, col: col}, target, data)
		if length < shortest {
			shortest = length
		}
	}
	for col := range data[len(data)-1] {
		length := dfs(Position{row:len(data)-1, col: col}, target, data)
		if length < shortest {
			shortest = length
		}
	}
	for row := range data {
		length := dfs(Position{row:row, col: 0}, target, data)
		if length < shortest {
			shortest = length
		}
	}
	for row := range data {
		length := dfs(Position{row:row, col: len(data[0])-1}, target, data)
		if length < shortest {
			shortest = length
		}
	}

	fmt.Println(shortest)
}
