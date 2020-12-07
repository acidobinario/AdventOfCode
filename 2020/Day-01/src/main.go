package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*	readInput
	Read the filename input and returns a list of numbers.
*/
func readInput(filename string)([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []int{}, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return []int{}, err
		}

		numbers = append(numbers, i)
	}

	return numbers, nil
}

/*	checkEntries
	Finds the two entries that sum to 2020 and then multiplies the two numbers together and returns the product.
*/
func checkEntries(entries []int)(int, error) {
	for y, x := 0, 1; y < len(entries) - 2; {
		if entries[y] + entries[x] == 2020 {
			fmt.Println("The values are", entries[y], "and", entries[x])

			return entries[y] * entries[x], nil
		}

		if x == len(entries) - 1 {
			y++
			x = y + 1

			continue
		}

		x++
	}

	return 0, errors.New("couldn't find the required values")
}

func main() {
	numbers, err := readInput("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	answer, err := checkEntries(numbers)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(answer)
}
