package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type entry struct {
	min   int
	max   int
	char  string
	pass  string
	count int
}

/*	readInput
	Read the filename and returns a list.
*/
func readInput(filename string) ([]string, error) {
	filePath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var passwords []string

	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	return passwords, nil
}

/*	checkPasswords
	Checks each password for policy compliance in the given list and returns the compliant passwords.
*/
func checkPasswords(passwords []string) ([]entry, error) {
	var compliant []entry
	for _, line := range passwords {
		s := strings.Split(line, " ")

		min, err := strconv.Atoi(strings.Split(s[0], "-")[0])
		if err != nil {
			return nil, err
		}

		max, err := strconv.Atoi(strings.Split(s[0], "-")[1])
		if err != nil {
			return nil, err
		}

		e := entry{
			min:   min,
			max:   max,
			char:  strings.Trim(s[1], ":"),
			pass:  s[2],
			count: strings.Count(s[2], strings.Trim(s[1], ":")),
		}

		if e.count <= e.max && e.count >= e.min {
			compliant = append(compliant, e)
		}
	}

	return compliant, nil
}

/*	checkPasswordsPartTwo
	Checks each password for policy compliance in the given list and returns the compliant passwords.
*/
func checkPasswordsPartTwo(passwords []string) ([]entry, error) {
	var compliant []entry
	for _, line := range passwords {
		s := strings.Split(line, " ")

		min, err := strconv.Atoi(strings.Split(s[0], "-")[0])
		if err != nil {
			return nil, err
		}

		max, err := strconv.Atoi(strings.Split(s[0], "-")[1])
		if err != nil {
			return nil, err
		}

		e := entry{
			min:   min,
			max:   max,
			char:  strings.Trim(s[1], ":"),
			pass:  s[2],
			count: strings.Count(s[2], strings.Trim(s[1], ":")),
		}

		if string(e.pass[e.min-1]) == e.char && string(e.pass[e.max-1]) != e.char {
			compliant = append(compliant, e)
		}

		if string(e.pass[e.min-1]) != e.char && string(e.pass[e.max-1]) == e.char {
			compliant = append(compliant, e)
		}
	}

	return compliant, nil
}

func main() {
	passwords, err := readInput("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	answer, err := checkPasswords(passwords)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Part one:", len(answer), "\n-------")

	answerTwo, err := checkPasswordsPartTwo(passwords)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Part two:", len(answerTwo))
}
