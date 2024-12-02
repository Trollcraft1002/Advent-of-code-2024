package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println(t)
		arr := t.Split(" ")
		for _, v := range arr {
			fmt.Println(v)
		}
	}
}