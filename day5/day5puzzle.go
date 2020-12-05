package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type seat struct {
	row, col string
}



func (s seat) PassID() (int, int, int) {
	row, col := 128, 8
	pr, pc := 0,0
	for _, v := range s.row {
		row = row / 2
		if string(v) == "B" {
			pr += row
		}
	}
	for _, v := range s.col {
		col = col / 2
		if string(v) == "R" {
			pc += col
		}
	}
	fmt.Println(pr,",",pc,",", (pr * 8) + pc )
	return pr, pc,(pr * 8) + pc

}

func main() {
	content, err := ioutil.ReadFile("./day5/input.txt")
	path, _ := os.Getwd()
	fmt.Println(path)
	if err != nil {
		log.Fatal(err)
	}
	idused := map[int]int{}
	max := 0
	something := strings.Split(string(content), "\n")
	passes := make([]seat, len(something))
	for i, v := range something {
		passes[i] = seat{v[:7], v[7:]}
	}
	minpr,maxpr := 129,0
	minimumid := 10000
	for _, v := range passes {
		pr,_,id := v.PassID()
		idused[id] = 1
		if minpr > pr {
			minpr = pr
		}
		if maxpr < pr {
			maxpr = pr
		}
		if id > max {
			max = id
		}
		if id < minimumid {
			minimumid = id
		}
	}
	fmt.Println(minpr,maxpr)
	fmt.Println(minimumid)
	//fmt.Println(max)
	for i:= minpr; i <= maxpr; i++ {
		for j:= 0; i < 8; j++ {
			temp := (i * 8) + j
			if i == minpr && j < 3 {
				continue
			}
			if idused[temp] == 0 {
				fmt.Println(i, j ,(i * 8) + j)
				goto out
			}
		}
	}
	out:
		fmt.Println(idused)
}