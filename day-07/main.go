package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func importData() []string{
	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return data
}

func upOneDir(currentDir string) string {
	if currentDir == "/" {
		return ""
	}
	currentDir = currentDir[1:len(currentDir)-1] // Otherwise split adds empty strings to the array
	dirSplit := strings.Split(currentDir, "/")
	if len(dirSplit) == 1 {
		currentDir = "/"
	} else {
		dirSplit = dirSplit[:len(dirSplit)-1]
		currentDir = "/" + strings.Join(dirSplit, "/") + "/"
	}

	return currentDir
}

func main() {
	data := importData()

	sizes := make(map[string]int)
	currentDir := "/"

	for i, line := range data {
		// hack to skip cd /
		if i == 0 {
			continue
		}
		switch {
		case line[0] == 'd':
			// No action needed, will capture dirs with cd
		case line[0] == '$':
			split := strings.Fields(line)
			if split[1] == "ls" {
				continue
			} else if split[2] == ".." {
				// Go up a dir
				fmt.Println("Going up from", currentDir)
				currentDir = upOneDir(currentDir)
				fmt.Println("to", currentDir)
			} else {
				// Go in a dir
				fmt.Println("Going in from", currentDir)
				currentDir = currentDir + split[2] + "/"
				fmt.Println("to", currentDir)
			}
		case unicode.IsNumber(rune(line[0])):
			split := strings.Fields(line)
			number, _ := strconv.Atoi(split[0])

			// Add size to the terminal directory
			sizes[currentDir] += number

			// Add size to all directories above
			remainingPath := upOneDir(currentDir)

			for len(remainingPath) > 0 {
				sizes[remainingPath] += number

				remainingPath = upOneDir(remainingPath)
			}
		}
	}

	// Part 1
	// total := 0
	// for _, size := range sizes {
	// 	if size <= 100000 {
	// 		total += size
	// 	}
	// }

	// Part 2
	unusedSpace := 70000000 - sizes["/"]
	spaceNeeded := 30000000 - unusedSpace
	fmt.Println("Space Needed", spaceNeeded)
	toDelete := 70000000
	for _, size := range sizes {
		fmt.Println(spaceNeeded, "vs", size)
		if size >= spaceNeeded && size < toDelete {
			toDelete = size
		}
	}

	// fmt.Println(sizes)
	fmt.Println("Size:", toDelete)
}
