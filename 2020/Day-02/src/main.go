package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	min		int
	max		int
	char	string
	pass	string
	count	int
}

/*	readInput
	Read the filename and returns a list.
*/
func readInput(filename string)([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
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
func checkPasswords(passwords []string)([]entry, error){
	var compliant []entry
	for _, line := range passwords {
		s := strings.Split(line, " ")

		min, err := strconv.Atoi(strings.Split(s[0], "-")[0])
		if err != nil {
			return []entry{}, err
		}

		max, err := strconv.Atoi(strings.Split(s[0], "-")[1])
		if err != nil {
			return []entry{}, err
		}

		e := entry{
		min,
		max,
		strings.Trim(s[1], ":"),
		s[2],
		strings.Count(s[2], strings.Trim(s[1], ":")),
		}

		if e.count <= e.max && e.count >= e.min {
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

	compliant, err := checkPasswords(passwords)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(compliant))
}