package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
)

func main() {
	day1Part1()
	day1Part2()
}

func day1Part1() {
	list1, list2 := parseInput("day1Example")

	slices.SortFunc(list1, func(i, j int) int {
		return i - j
	})

	slices.SortFunc(list2, func(i, j int) int {
		return i - j
	})

	answer := 0
	for i := 0; i < len(list1); i++ {
		answer = answer + int(math.Abs(float64(list1[i])-float64(list2[i])))
	}

	fmt.Printf("day 1 part 1: %d\n", answer)
}

func day1Part2() {
	list1, list2 := parseInput("day1Example")
	map1 := make(map[int]int)
	map2 := make(map[int]int)

	for index, ele := range list1 {
		map1[ele]++
		map2[list2[index]]++
	}

	answer := 0
	for k := range map1 {
		answer = answer + ((map1[k] * map2[k]) * k)
	}

	fmt.Printf("day 1 part 2: %d\n", answer)
}

func parseInput(fileName string) (numberList1, numberList2 []int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return make([]int, 0), make([]int, 0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	numberList1 = make([]int, 1)
	numberList2 = make([]int, 1)

	i := 1
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if (i % 2) == 1 {
			numberList1 = append(numberList1, value)
		} else {
			numberList2 = append(numberList2, value)
		}
		i++
	}

	return numberList1, numberList2
}
