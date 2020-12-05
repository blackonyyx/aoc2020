package main
import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)
type bits uint8

func set(b, flag bits) bits { return b | flag }
func quickRead(file string) []string {
	content, _ := ioutil.ReadFile(file)
	x := strings.Split(string(content), "\n")
	return x
}
func Part1() int {
const (
byr bits = 1 << iota
iyr
eyr
hgt
hcl
ecl
pid
cid
)
const pass = 127
var current bits
var fields []string
var count int
var prefix string
for _, line := range quickRead("./day4/input.txt") {
if line == "" {
if current == pass {
count++
}
current = 0
continue
}
fields = strings.Split(line, " ")
for field := range fields {
prefix = fields[field][0:3]
switch prefix {
case "byr":
current = set(current, byr)
case "iyr":
current = set(current, iyr)
case "eyr":
current = set(current, eyr)
case "hgt":
current = set(current, hgt)
case "hcl":
current = set(current, hcl)
case "ecl":
current = set(current, ecl)
case "pid":
current = set(current, pid)
default:
continue
}
}
}
// Just in case the last passport passes.
if current == pass {
count++
}
return count
}

func Part2() int {
const (
byr bits = 1 << iota
iyr
eyr
hgt
hcl
ecl
pid
cid
)
const pass = 127
var current bits
var fields []string
var count int
var prefix, val string
var num int
var err error
var b bool
isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
for _, line := range quickRead("./day4/input.txt") {
if line == "" {
if current == pass {
count++
}
current = 0
continue
}
fields = strings.Split(line, " ")
for field := range fields {
prefix = fields[field][0:3]
val = fields[field][4:]
switch prefix {
case "byr":
num, err = strconv.Atoi(val)
if err == nil && num >= 1920 && num <= 2002 {
current = set(current, byr)
}
case "iyr":
num, err = strconv.Atoi(val)
if err == nil && num >= 2010 && num <= 2020 {
current = set(current, iyr)
}
case "eyr":
num, err = strconv.Atoi(val)
if err == nil && num >= 2020 && num <= 2030 {
current = set(current, eyr)
}
case "hgt":
if strings.HasSuffix(val, "in") {
num, err = strconv.Atoi(val[0:2])
if err == nil && num >= 59 && num <= 76 {
current = set(current, hgt)
}
} else if strings.HasSuffix(val, "cm") {
num, err = strconv.Atoi(val[0:3])
if err == nil && num >= 150 && num <= 193 {
current = set(current, hgt)
}
}
case "hcl":
if strings.HasPrefix(val, "#") && len(val) == 7 {
b = true
for _, r := range val[1:] {
if !(r >= '0' && r <= '9') && !(r >= 'a' && r <= 'f') {
b = false
break
}
}
if b {
current = set(current, hcl)
}
}
case "ecl":
switch val {
case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
current = set(current, ecl)
default:
continue
}
case "pid":
if len(val) == 9 && strings.IndexFunc(val, isNotDigit) == -1 {
current = set(current, pid)
}
default:
continue
}
}
}
// Just in case the last passport passes.
if current == pass {
count++
}
return count
}

func main()  {
fmt.Println(Part1())
fmt.Println(Part2())
}
