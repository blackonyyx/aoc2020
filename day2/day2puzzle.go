package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Policy struct {
	lower, higher int
	char, password string
}
func (p Policy) isValidPart2() bool {
	val := 0
	if (p.password[p.lower-1] == p.char[0]) {
		val = val ^ 1
	}
	if (p.password[p.higher-1] == p.char[0]) {
		val = val ^ 1
	}

	return val != 0
}
func (p Policy) isValidPart1() bool {
	count := 0
	for i , _ := range p.password {
		if p.char[0] == p.password[i] {
			count ++
			if count > p.higher {
				return false
			}
		}
	}
	if count >= p.lower {
		return true
	}
	return false
}
func main() {
	content, err := ioutil.ReadFile("./day2/input1.txt")
	path, _ := os.Getwd()
	fmt.Println(path)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	something := strings.Split(string(content), "\n")
	policies := make([]Policy, len(something))
	pattern := regexp.MustCompile(`-| |: `)
	for i,v := range something {
		parsed := pattern.Split(v, -1)
		// Parse to a policy
		lo, _ := strconv.Atoi(parsed[0])
		hi, _ := strconv.Atoi(parsed[1])
		policies[i] = Policy{lo,hi, parsed[2], parsed[3] }
		if policies[i].isValidPart2() {
			count++
		}
	}
	fmt.Println("Number of valid policies: " , count)
}