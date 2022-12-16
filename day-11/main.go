package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	id int
	items []big.Int
	operation string
	test big.Int
	ifTrue int
	ifFalse int
	count int
}

func importData() []Monkey{
	file, err := os.Open("test-data.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([]Monkey, 0)

	scanner := bufio.NewScanner(file)

	var newMonkey Monkey

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			switch line[0] {
			case 'M':
				split := strings.Fields(line)
				id, _ := strconv.Atoi(split[1][:1])
				newMonkey = Monkey{
					id: id,
					items: make([]big.Int, 0),
					operation: "",
					test: *big.NewInt(-1),
					ifTrue: -1,
					ifFalse: -1,
					count: 0,
				}
			case 'S':
				split := strings.Split(line, "items: ")
				items := strings.Split(split[1], ", ")
				for _, item := range items {
					num, _ := strconv.Atoi(item)
					bigNum := big.NewInt(int64(num))
					newMonkey.items = append(newMonkey.items, *bigNum)
				}
			case 'O':
				newMonkey.operation = line
			case 'T':
				split := strings.Fields(line)
				num, _ := strconv.Atoi(split[3])
				bigNum := big.NewInt(int64(num))
				newMonkey.test = *bigNum
			case 'I':
				split := strings.Fields(line)
				if split[1] == "true:" {
					split := strings.Fields(line)
					num, _ := strconv.Atoi(split[5])
					newMonkey.ifTrue = num
				} else {
					split := strings.Fields(line)
					num, _ := strconv.Atoi(split[5])
					newMonkey.ifFalse = num
				}

			}
		} else {
			data = append(data, newMonkey)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	// Hack because scanner seems to skip trailing newline and that's where append is
	data = append(data, newMonkey)
	return data
}

func main() {
	data := importData()

	// for _, monkey := range data {
	// 	fmt.Println(monkey)
	// }


	rounds := 20
	for rounds > 0 {
		for i := range data {
			// fmt.Println(monkey)
			monkey := &data[i]
			for _, item := range monkey.items {
				monkey.count++
				// fmt.Println(monkey)
				// Do the operation
				split := strings.Fields(monkey.operation)
				var first big.Int
				if split[3] == "old" {
					first = item
				} else {
					num, _ := strconv.Atoi(split[3])
					first = *big.NewInt(int64(num))
				}
				var second big.Int
				if split[5] == "old" {
					second = item
				} else {
					num, _ := strconv.Atoi(split[5])
					second = *big.NewInt(int64(num))
				}

				switch split[4] {
				case "+":
					item = *first.Add(&first, &second)
				case "*":
					item = *first.Mul(&first, &second)
				}

				// item = item / 3

				// With bigint, item.Mod(&item, &monkey.test) == big.NewInt(0) doesn't work
				if item.Mod(&item, &monkey.test).BitLen() == 0 {
					fmt.Println("TRUE HAPPENED")
					data[monkey.ifTrue].items = append(data[monkey.ifTrue].items, item)
				} else {
					data[monkey.ifFalse].items = append(data[monkey.ifFalse].items, item)
				}
			}

			// Items are all thrown
			monkey.items = make([]big.Int, 0)
		}

		rounds--
	}

	// var total int
	// counts := make([]big.Int, 0)

	// for _, monkey := range data {
	// 	counts = append(counts, monkey.count)
	// }

	// sort.Slice(counts, func(i, j int) bool {
	// 	returns big.
	// 	return counts[i] > counts[j]
	// })

	// total = counts[0] * counts[1]

	for _, monkey := range data {
		fmt.Println("Id", monkey.id)
		// fmt.Println("Items", monkey.items)
		// fmt.Println("Operation", monkey.operation)
		// fmt.Println("Test", &monkey.test)
		// fmt.Println("IfTrue", monkey.ifTrue)
		// fmt.Println("IfFalse", monkey.ifFalse)
		fmt.Println("Count", monkey.count)
		fmt.Println("*******")
	}

	// fmt.Println(total)
}
