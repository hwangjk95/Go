package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetString(fileName string) ([]string, error) {

	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err

	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

/*
func main() {

	lines, err := GetString("vote.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)

		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
  기존방식*/

func main() {
	lines, err := GetString("vote.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("%s: %d\n", name, count)
	}
}
