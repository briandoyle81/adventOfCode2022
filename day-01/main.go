package main

import (
	"bufio"
	"container/heap"
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

// Part 1
// func main() {
//     data := importData()

// 	largestValue := 0
// 	var largestKey int

// 	totals := make(map[int]int)

// 	for key, val := range data {
// 		totals[key] = sumArray(val)
// 		if totals[key] > largestValue {
// 			largestValue = totals[key]
// 			largestKey = key
// 		}
// 	}

// 	fmt.Println(largestKey, largestValue)
// }

// Part 2

func main() {
    data := importData()

	// Store values in a min heap
	// TODO: Figure out how to make this a max heap instead
	h := &IntHeap{0, 0, 0}
	heap.Init(h)

	for _, val := range data {
		heap.Push(h, sumArray(val)*-1) // Hack to work with min heap
	}

	topThreeTotal := 0

	// There must be something more elegant than a for loop here
	for i := 0; i < 3; i++ {
		topThreeTotal += heap.Pop(h).(int) // TODO: What is this .(int) and why does it work?
	}

	fmt.Println(topThreeTotal * -1)
}
