package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	f, err := os.Open("./input.txt")
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		t := scanner.Text()
		regex := regexp.MustCompile(`mul\((\d+,\d+)\)`)
		res := regex.FindAllStringSubmatch(t, -1)
		for v := range res {
			temp := strings.Split(res[v][1], ",")
			a, err := strconv.Atoi(temp[0])
			check(err)
			b, err := strconv.Atoi(temp[1])
			check(err)
			sum += a * b
		}
	}
	fmt.Println(sum)
	part2()
}

// -----------------------------PART 2 ------------------
func part2() {

	f, err := os.Open("./input.txt")
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	regex := regexp.MustCompile(`(do\(\))|(don't\(\))|mul\((\d+,\d+)\)`)
	mul := true
	for scanner.Scan() {
		t := scanner.Text()

		res := regex.FindAllStringSubmatch(t, -1)

		for v := range res {
			if res[v][0] == "do()" {
				fmt.Println(res[v][0])
				mul = true
				continue
			} else if res[v][0] == `don't()` {
				fmt.Println(res[v][0])
				mul = false
				continue
			}
			if !mul {
				continue
			}
			temp := strings.Split(res[v][3], ",")
			a, err := strconv.Atoi(temp[0])
			if err != nil {
				check(err)
			}

			b, err := strconv.Atoi(temp[1])
			check(err)
			sum += a * b
		}
	}
	fmt.Println(sum)

}
