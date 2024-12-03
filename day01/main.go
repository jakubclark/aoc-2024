package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	leftList, rightList := getInput()

	// Part 1
	result := int64(0)
	for i := range leftList {
		distance := abs(leftList[i] - rightList[i])
		result += distance
	}
	fmt.Printf("Part 1: %d\n", result)

	// Part 2
	result = int64(0)
	occurrences := count(rightList)
	for _, left := range leftList {
		result += occurrences[left] * left
	}
	fmt.Printf("Part 2: %d\n", result)
}

func getInput() ([]int64, []int64) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Can't open input.txt", err)
	}

	leftList := []int64{}
	rightList := []int64{}

	for _, line := range strings.Split(string(input), "\n") {
		nums := strings.Fields(line)

		left, err := strconv.ParseInt(nums[0], 10, 64)
		if err != nil {
			log.Fatal("Can't convert number", err)
		}
		right, err := strconv.ParseInt(nums[1], 10, 64)
		if err != nil {
			log.Fatal("Can't convert number", err)
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	return leftList, rightList
}

func count(nums []int64) map[int64]int64 {
	result := make(map[int64]int64)
	for _, v := range nums {
		result[v]++
	}
	return result
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
