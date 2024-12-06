package main

import (
	"bufio"
	"fmt"
	"os"
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
	var data [][]string

	for scanner.Scan() {

		t := scanner.Text()
		spData := strings.Split(t, "")
		data = append(data, spData)
	}
	sum := 0

	for v := range data {
		for i := range data[v] {
			//line
			if (data[v][i]) == "X" && len(data[v]) > i+3 && data[v][i+1] == "M" && data[v][i+2] == "A" && data[v][i+3] == "S" {
				//println("line")
				sum++
			}

			// left to right diagonal
			if (data[v][i]) == "X" && v+3 < len(data) && i+3 < len(data[v]) && data[v+1][i+1] == "M" && data[v+2][i+2] == "A" && data[v+3][i+3] == "S" {
				//println("lr diagonal")
				sum++
			}

			//column
			if (data[v][i]) == "X" && len(data) > v+3 && data[v+1][i] == "M" && data[v+2][i] == "A" && data[v+3][i] == "S" {
				//println("col")
				sum++
			}

			// right to left diagonal

			if (data[v][i]) == "X" && v+3 < len(data) && i-3 >= 0 && data[v+1][i-1] == "M" && data[v+2][i-2] == "A" && data[v+3][i-3] == "S" {
				//println("rl diagonal")
				sum++
			}

			//REVERSE ORDER

			//reverse line

			if (data[v][i]) == "X" && i-3 >= 0 && data[v][i-1] == "M" && data[v][i-2] == "A" && data[v][i-3] == "S" {
				//	println("r line")
				sum++
			}

			// reverse right to left diagonal

			if (data[v][i]) == "X" && v-3 >= 0 && i-3 >= 0 && data[v-1][i-1] == "M" && data[v-2][i-2] == "A" && data[v-3][i-3] == "S" {
				//println("r rl diagnoal")
				sum++
			}

			//reverse column

			if (data[v][i]) == "X" && v-3 >= 0 && data[v-1][i] == "M" && data[v-2][i] == "A" && data[v-3][i] == "S" {
				//println("r column")
				sum++
			}

			// reverse left to right diagonal

			if (data[v][i]) == "X" && v-3 >= 0 && i+3 < len(data[v]) && data[v-1][i+1] == "M" && data[v-2][i+2] == "A" && data[v-3][i+3] == "S" {
				//println("r lr diagonal")
				sum++
			}

		}
	}
	fmt.Println(sum)
	part2(data)
}

func part2(data [][]string) {
	for r := range data {
		for i := range data[r] {
			if data[r][i] == "A" {

			}
		}
	}
}
