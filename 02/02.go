package main

import (
	"bufio"
	"os"
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

	for scanner.Scan() {
		t := scanner.Text()
		arr := strings.Split(t, " ")
		var int_arr []int

		for i := range arr {
			v, err := strconv.Atoi(arr[i])
			chcek(err)
			int_arr = append(int_arr, v)
		}

			for  i := range int_arr {
				for j := i; i < len(arr); i++ {
					if(j >)
				}
			}
	}
}
