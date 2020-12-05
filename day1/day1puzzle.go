package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main()  {
	content, err := ioutil.ReadFile("./day1/input1.txt")
	path, _ := os.Getwd()
	fmt.Println(path)
	if err != nil {
		log.Fatal(err)
	}
	something := strings.Split(string(content), "\n")
	lst := make([]int,len(something))
	for i, v := range something {
		lst[i], _ = strconv.Atoi(v)
	}
	sort.Ints(lst)
	target := 2020
	for _,v := range lst {
		errata, thing := sum2(lst,target-v)
		if errata {
			fmt.Println(v ,thing[0], thing[1])
			fmt.Println(v * thing[0] * thing[1])
			break
		}
	}

}

func sum2(lst []int, target int) (bool,[]int) {
	x := 0
	l,r := 0, len(lst) - 1;
	ans := make([]int, 2)
	required := target
	success := false
	for x != required {
		x = lst[l] + lst[r]
		if l == r {
			break
		}
		if x == required {
			fmt.Println("Required answer = ", lst[l]*lst[r])
			ans[0], ans[1] = lst[l], lst[r]
			success = true
			break
		}
		if x < required {
			l ++
		}
		if x > required {
			r --
		}
	}
	return success, ans
}