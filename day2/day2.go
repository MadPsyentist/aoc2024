package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	day2Part1()
	day2Part2()
}

func day2Part1() {
	reportList := parseInput("day2Example")

	validReports := 0
	for reportNumber, report := range reportList {
		valid := true
		if report[0] > report[1] {
			for i := 1; i < len(report); i++ {
				score := report[i-1] - report[i]
				if score <= 0 || score > 3 {
					valid = false
					break
				}
			}
			if valid {
				fmt.Printf("report %d is valid\n", reportNumber)
				validReports++
			}
		} else {
			for i := 1; i < len(report); i++ {
				score := report[i] - report[i-1]
				if score <= 0 || score > 3 {
					valid = false
					break
				}
			}
			if valid {
				fmt.Printf("report %d is valid\n", reportNumber)
				validReports++
			}
		}
	}

	fmt.Printf("day 2 part 1: %d\n", validReports)
}

func day2Part2() {
	reportList := parseInput("day2Example")

	validReports := 0
	for _, report := range reportList {
		if validateReport(report) {
			validReports++
		} else {
			for index := range report {
				newReport := removeElementFromSlice(index, report)
				if validateReport(newReport) {
					validReports++
					break
				}
			}
		}
	}

	fmt.Printf("day 2 part 2: %d\n", validReports)
}

func validateReport(report []int) bool {
	if len(report) < 2 {
		return false
	}

	valid := true
	if report[0] > report[1] {
		for i := 1; i < len(report); i++ {
			score := report[i-1] - report[i]
			if score <= 0 || score > 3 {
				valid = false
				break
			}
		}
	} else {
		for i := 1; i < len(report); i++ {
			score := report[i] - report[i-1]
			if score <= 0 || score > 3 {
				valid = false
				break
			}
		}
	}
	return valid
}

func removeElementFromSlice(indexToRemove int, slice []int) []int {
	newSlice := make([]int, len(slice)-1)
	i := 0
	for j, ele := range slice {
		if j != indexToRemove {
			newSlice[i] = ele
			i++
		}
	}
	return newSlice
}

func parseInput(fileName string) (reportList [][]int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	reportList = make([][]int, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		report := make([]int, 0)
		for _, val := range line {
			value, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}

			report = append(report, value)
		}

		reportList = append(reportList, report)
	}

	return reportList
}
