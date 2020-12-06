package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main()  {
	content, err := ioutil.ReadFile("./day6/input.txt")
	path, _ := os.Getwd()
	fmt.Println(path)
	if err != nil {
		log.Fatal(err)
	}
	groups := strings.Split(string(content), "\n\n")
	part1(groups)
	part2(groups)
}

func part2(groups []string) {
	sum := 0
	for _, g := range groups {
		stuff := map[string]int{}
		persons := strings.Split(g, "\n")
		for _, person := range persons {
			for _, p := range person {
				stuff[string(p)] += 1
			}
		}
		for _,v := range stuff {
			if v == len(persons) {
				sum += 1
			}
		}
		fmt.Println(stuff, len(persons))
	}
	fmt.Println(sum)
}

func part1(groups []string)  {
	sum := 0
	for _, g := range groups {
		stuff := map[string]int{}
		persons := strings.Split(g, "\n")
		for _, person := range persons {
			for _, p := range person {
				if stuff[string(p)] == 0 {
					stuff[string(p)] = 1
				}
			}
		}
		for _,v := range stuff {
			sum += v
		}
		fmt.Println(stuff)
	}
	fmt.Println(sum)
}
