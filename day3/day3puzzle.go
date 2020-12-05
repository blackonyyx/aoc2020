package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./day3/input.txt")
	path, _ := os.Getwd()
	fmt.Println(path)
	if err != nil {
		log.Fatal(err)
	}
	x,y := ".", "#"
	something := strings.Split(string(content), "\n")
	xaxis := len(something[0])
	yaxis := len(something)
	fmt.Println(xaxis)
	fmt.Println(yaxis)
	grid := make([][]uint, yaxis)
	for i:=0; i < yaxis;i++ {
		grid[i] = make([]uint, xaxis)
		for j, _ := range something[i] {
			if something[i][j] == x[0] {
				grid[i][j] = 0
			}
			if something[i][j] == y[0] {
				grid[i][j] = 1
			}
		}
	}
	trees := make([]int, 0)
	checker := []int{1,3,5,7,1}
	ycheck := []int {1,2}
	for j:= 0; j < len(checker); j++ {
		if j > 3 {
			fmt.Println("hello here", checker[j], ycheck[1])
			trees = append(trees, treesPart1(grid, checker[j], ycheck[1], xaxis, yaxis))
		} else {
			fmt.Println("hello there", checker[j], ycheck[0])
			trees = append(trees, treesPart1(grid, checker[j], ycheck[0], xaxis, yaxis))
		}
	}
	prod := 1
	for _, v := range trees {
		prod *= v
		fmt.Println("hello",v)
	}
	fmt.Println(prod)
}

func treesPart1 (grid [][]uint, dx, dy, xaxis, yaxis int) int {
	xcoord,ycoord, count := 0,0,0
	for ycoord < yaxis {
		if grid[ycoord][xcoord] == 1 {
			count += 1
		}
		xcoord += dx
		xcoord = xcoord % xaxis
		ycoord += dy
	}
	return count
}