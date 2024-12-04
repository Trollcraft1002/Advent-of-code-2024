package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// --------------------------PART 1----------------------------
func main() {
	data, err := os.Open("./input.txt")
	check(err)

	defer data.Close()

	scanner := bufio.NewScanner(data)

	safe := 0
	for scanner.Scan() {
		t := scanner.Text()
		arr := strings.Split(t, " ")
		var int_arr []int

		for i := range arr {
			v, err := strconv.Atoi(arr[i])
			check(err)
			int_arr = append(int_arr, v)
		}
		isBroken := false
		isAscending := false
		isDescending := false
		for i := 0; i < len(int_arr)-1; i++ {

			if int_arr[i] == int_arr[i+1] {
				isBroken = true
				continue
			}

			if int_arr[i] < int_arr[i+1] {
				isAscending = true
			} else if int_arr[i] > int_arr[i+1] {
				isDescending = true
			}

			if isDescending && int_arr[i] > int_arr[i+1] && int_arr[i+1] <= int_arr[i]-4 {
				isBroken = true
			}
			if isAscending && int_arr[i] < int_arr[i+1] && int_arr[i+1] >= int_arr[i]+4 {
				isBroken = true
			}

			if isAscending == isDescending {
				isBroken = true
			}

			if isBroken {
				//println("Unsafe")
				break
			}
		}
		if !isBroken {
			//fmt.Println("Safe")
			safe++
		}
		//fmt.Println("NEW INPUT")
		//fmt.Println(int_arr)
	}
	fmt.Println("PART 1:", safe)
	//fmt.Println("PART 2:", part2())
	part2()
}

// --------------------------PART 2----------------------------

func part2() {
	data, err := os.Open("./input.txt")
	check(err)
	defer data.Close()

	scanner := bufio.NewScanner(data)
	count := 0
	for scanner.Scan() {
		t := scanner.Text()
		arr := strings.Split(t, " ")
		var intArr []int

		for _, s := range arr {
			v, err := strconv.Atoi(s)
			check(err)
			intArr = append(intArr, v)
		}

		if isSafe(intArr) {
			count++
			continue
		}

		for i := 0; i < len(intArr); i++ {
			tempArr := append([]int{}, intArr[:i]...)
			tempArr = append(tempArr, intArr[i+1:]...)
			if isSafe(tempArr) {
				count++
				break
			}
		}
	}
	fmt.Println("PART 2:", count)
}

func isSafe(intArr []int) bool {
	if len(intArr) < 2 {
		return true
	}

	direction := 0 // 1 for increasing, -1 for decreasing
	for i := 0; i < len(intArr)-1; i++ {
		diff := intArr[i+1] - intArr[i]

		if diff == 0 || abs(diff) > 3 {
			return false
		}

		if direction == 0 {
			if diff > 0 {
				direction = 1
			} else {
				direction = -1
			}
		} else {
			if (direction == 1 && diff < 0) || (direction == -1 && diff > 0) {
				return false
			}
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
