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

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, make([]int, 0))

		for _, char := range line {
			data[row] = append(data[row], int(char)-48) // Need to find the correct way to do this
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return data
}

func printGraph(graph [][]int) {
	for _, line := range graph {
		fmt.Println(line)
	}
}

func coordToString(row int, col int) string {
	return strconv.Itoa(row) + "," + strconv.Itoa(col)
}

func stringToCoord(coords string) (int, int) {
	split := strings.Split(coords, ",")
	row, _ := strconv.Atoi(split[0])
	col, _ := strconv.Atoi(split[1])
	return row, col
}

// Part 1
// func main() {
// 	data := importData()

// 	// printGraph(data)
// 	// Edges are visible
// 	// visible := len(data) * 2 + len(data[0]) * 2 - 4

// 	visible := make(map[string]bool)

// 	for row, rowData := range data {
// 		// Might be slightly naive but o(n*4) == o(n)

// 		// Visible from left edge
// 		tallest := rowData[0]
// 		visible[coordToString(row, 0)] = true
// 		// Only the tallest tree will be visible
// 		for col := 1; col < len(rowData); col++ {
// 			if data[row][col] > tallest {
// 				visible[coordToString(row, col)] = true
// 				tallest = data[row][col]
// 			}
// 		}

// 		// Visible from right edge
// 		tallest = rowData[len(rowData)-1]
// 		visible[coordToString(row, len(rowData)-1)] = true
// 		// Only the tallest tree will be visible
// 		for col := len(rowData)-1; col >0; col-- {
// 			if data[row][col] > tallest {
// 				visible[coordToString(row, col)] = true
// 				tallest = data[row][col]
// 			}
// 		}
// 	}

// 	for col := 0; col < len(data); col++ {
// 		// Visible from top edge
// 		tallest := data[0][col]
// 		visible[coordToString(0, col)] = true
// 		for row := 1; row < len(data); row++ {
// 			if data[row][col] > tallest {
// 				visible[coordToString(row, col)] = true // Should DRY this and practice pointers
// 				tallest = data[row][col]
// 			}
// 		}

// 		// Visible from bottom edge
// 		tallest = data[len(data)-1][col]
// 		visible[coordToString(len(data)-1, col)] = true
// 		for row := len(data)-1; row >0; row-- {
// 			if data[row][col] > tallest {
// 				visible[coordToString(row, col)] = true // Should DRY this and practice pointers
// 				tallest = data[row][col]
// 			}
// 		}
// 	}

// 	visibleCount := 0

// 	for range visible {
// 		visibleCount++
// 	}

// 	fmt.Println(visibleCount)
// }

// Part 2
func main() {
	data := importData()

	// printGraph(data)
	// Edges are visible
	// visible := len(data) * 2 + len(data[0]) * 2 - 4

	scenicValue := make(map[string][]int)

	for row := 1; row < len(data)-1; row++ {
		// I can't think of a non brute force way to do part 2
		rowData := data[row]
		// Visible from left edge
		for col := 0; col < len(rowData); col++ {
			scenicValue[coordToString(row, col)] = make([]int, 0)
			if col == 0 {
				scenicValue[coordToString(row, col)] = append(scenicValue[coordToString(row, col)], 0)
			} else {
				height := rowData[col]
				// Go back left
				for cursor := col-1; cursor >= 0; cursor-- {
					if data[row][cursor] >= height || cursor == 0 {
						scenicValue[coordToString(row, col)] = append(scenicValue[coordToString(row, col)], col - cursor)
						break
					}
				}
			}
		}

		// Visible from right edge
		for col := len(rowData)-1; col >= 0; col-- {
			if col == rowData[len(rowData)-1] {
				scenicValue[coordToString(row, col)] = append(scenicValue[coordToString(row, col)], 0)
			} else {
				height := rowData[col]
				// go back right
				for cursor := col+1; cursor < len(rowData); cursor++ {
					if data[row][cursor] >= height || cursor == len(rowData)-1 {
						scenicValue[coordToString(row, col)] = append(scenicValue[coordToString(row, col)], cursor - col)
						break
					}
				}
			}
		}
	}

	for col := 1; col < len(data[0])-1; col++ {
		// Visible from top edge
		for row := 0; row < len(data); row++ {
			if row == 0 {
				scenicValue[coordToString(row, col)] = append(scenicValue[coordToString(row, col)], 0)
			} else {
				height := data[row][col]
				// go up
				for cursor := row-1; cursor >= 0; cursor-- {
					if data[cursor][col] >= height || cursor == 0 {
						scenicValue[coordToString(row, col)] = append(scenicValue[coordToString(row, col)], row - cursor)
						break
					}
				}
			}
		}

		// Visible from bottom edge
		for row := len(data)-1; row >= 0; row-- {
			if row == len(data)-1 {
				scenicValue[coordToString(row, col)] = append(scenicValue[coordToString(row, col)], 0)
			} else {
				height := data[row][col]
				// go down
				for cursor := row+1; cursor < len(data); cursor++ {
					if data[cursor][col] >= height || cursor == len(data)-1 {
						scenicValue[coordToString(row, col)] = append(scenicValue[coordToString(row, col)], cursor - row)
						break
					}
				}
			}
		}
	}

	mostScenic := 0
	mostScenicCoords := ""
	mostScenicCoords = mostScenicCoords + ""
	for key, values := range scenicValue {
		// fmt.Println(key, values)
		acc := 1
		for _, value := range values {
			acc *= value
		}

		if acc > mostScenic {
			mostScenic = acc
			mostScenicCoords = key
		}
	}

	fmt.Println(mostScenic)
	fmt.Println(mostScenicCoords)
}
