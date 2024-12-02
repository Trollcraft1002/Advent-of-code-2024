package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func chcek(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.Open("./input.txt")
	chcek(err)

	defer data.Close()

	scanner := bufio.NewScanner(data)
	var l, r []int

	for scanner.Scan() {
		t := scanner.Text()
		str := strings.Split(t, "   ")
		_l, err := strconv.Atoi(str[0])
		if err != nil {
			panic(err)
		}

		_r, err := strconv.Atoi(str[1])
		if err != nil {
			panic(err)
		}

		l = append(l, _l)
		r = append(r, _r)

	}
	sort.Ints(l)
	sort.Ints(r)
	sum := 0
	for a := range l {

		calc := l[a] - r[a]
		sum += int(math.Abs(float64(calc)))

	}
	fmt.Println("part 1: ", sum)
	part2(l, r)
}

func part2(a, b []int) {
	sum := 0
	count := 0
	for l := range a {
		count = 0
		for r := range b {

			if a[l] == b[r] {
				count++
			}
		}
		sum += count * a[l]
	}
	fmt.Println("part 2: ", sum)
}
