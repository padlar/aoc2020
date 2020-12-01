package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return lines, err
		}
		lines = append(lines, val)
	}
	return lines, scanner.Err()
}

func main() {
	input, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("something went wrong: %s\n", err)
	}

	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })

	find2numbers(input)

	find3numbers(input)
}

func find3numbers(input []int) {
	i := 0
	j := len(input) - 1
	k := 1
	for {
		if i > j {
			log.Println("didnt find 2 numbers adding upto 2020")
			k++
			i = 0
			j = len(input) - 1
		}
		sum := input[i] + input[j] + input[k]
		if sum == 2020 {
			log.Printf("found 3 nrs - %d + %d + %d = 2020, %d\n", input[i], input[j], input[k], input[i]*input[j]*input[k])
			break
		} else if sum < 2020 {
			i++
		} else {
			j--
		}
	}
}

func find2numbers(input []int) {
	i := 0
	j := len(input) - 1
	for {
		if i > j {
			log.Println("didnt find 2 numbers adding upto 2020")
			break
		}
		sum := input[i] + input[j]
		if sum == 2020 {
			log.Printf("found 2 nrs - %d + %d = 2020, %d\n", input[i], input[j], input[i]*input[j])
			break
		} else if sum < 2020 {
			i++
		} else {
			j--
		}
	}
}
