package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func importData() [][]string{
	file, err := os.Open("big-test-data.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([][]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		entry := strings.Fields(line)
		data = append(data, entry)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return data
}

// Part 1
// func main() {
// 	data := importData()
// 	// fmt.Println(data)
// 	registerX := 1
// 	cycle := 1
// 	lookAtCycle := 20
// 	signalSum := 0


// 	for _, instruction := range data {
// 		fmt.Println("Start of Cycle:", cycle, "RegisterX:", registerX)
// 		if cycle == lookAtCycle { // TODO: DRY
// 			lookAtCycle += 40
// 			signalSum += cycle * registerX
// 			fmt.Println("A signalSum:", cycle, "*", registerX, "=", cycle*registerX, "Total:", signalSum)
// 		}
// 		// fmt.Println(instruction)
// 		switch instruction[0] {
// 		case "noop":
// 			// Do nothing
// 		case "addx":
// 			cycle++

// 			if cycle == lookAtCycle { // TODO: DRY
// 				lookAtCycle += 40
// 				signalSum += cycle * registerX
// 				fmt.Println("B signalSum:", cycle, "*", registerX, "=", cycle*registerX, "Total:", signalSum)
// 			}

// 			value, _ := strconv.Atoi(instruction[1])
// 			registerX += value
// 		}

// 		cycle++
// 	}


// 	fmt.Println("registerX", registerX)
// 	fmt.Println("signalSum", signalSum)
// }

// Part 2
func drawScreen(pixels [][]bool) {
	for _, row := range pixels {
		for _, on := range row {
			if on {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}


func main() {
	data := importData()
	// fmt.Println(data)
	registerX := 1
	cycle := 1
	row := 0

	pixels := make([][]bool, 6)
	for i := range pixels {
		pixels[i] = make([]bool, 40)
	}

	drawScreen(pixels)

	for _, instruction := range data {
		// fmt.Println("Start of Cycle:", cycle, "RegisterX:", registerX)
		fmt.Println(registerX, cycle-(40*row), "diff", int(math.Abs(float64(registerX)-float64(cycle-(40*row)))))
		if int(math.Abs(float64(registerX)-float64(cycle-(40*row)))) <= 1 {
			pixels[row][cycle-(40*row)-1] = true
		}
		// fmt.Println(instruction)
		switch instruction[0] {
		case "noop":

		case "addx":


			cycle++
			fmt.Println(registerX, cycle-(40*row), "diff", int(math.Abs(float64(registerX)-float64(cycle-(40*row)))))
			if int(math.Abs(float64(registerX)-float64(cycle-(40*row)))) <= 1 {
				pixels[row][cycle-(40*row)-1] = true
			}

			value, _ := strconv.Atoi(instruction[1])
			registerX += value


		}

		cycle++
	}

	drawScreen(pixels)
	fmt.Println("registerX", registerX)
}
