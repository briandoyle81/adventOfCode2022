package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func importData() map[int][]int{
	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make(map[int][]int)

	scanner := bufio.NewScanner(file)

	elfNumber := 1

	for scanner.Scan() {
		if scanner.Text() == "" {
			elfNumber++
			continue
		} else {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			if data[elfNumber] == nil {
				data[elfNumber] = []int{calories}
			}	else {
				data[elfNumber] = append(data[elfNumber], calories)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return data
}

// So sum isn't built in?
func sumArray(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
    data := importData()

	largestValue := 0
	var largestKey int

	totals := make(map[int]int)

	for key, val := range data {
		totals[key] = sumArray(val)
		if totals[key] > largestValue {
			largestValue = totals[key]
			largestKey = key
		}
	}

	fmt.Println(largestKey, largestValue)
}
