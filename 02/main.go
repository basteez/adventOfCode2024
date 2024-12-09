package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	safeReports := 0

	for fileScanner.Scan() {
		fields := strings.Fields(fileScanner.Text())

		increasing, decreasing := true, true
		invalid := false
		for i := 0; i < len(fields)-1; i++ {
			current, err := strconv.Atoi(fields[i])
			checkError(err)
			next, err := strconv.Atoi(fields[i+1])
			checkError(err)

			if current > next {
				increasing = false
			} else if current < next {
				decreasing = false
			}
			if !increasing && !decreasing {
				invalid = true
				break
			}

			stepSize := math.Abs(float64(current - next))
			if stepSize == 0 || stepSize > 3 {
				invalid = true
				break
			}
		}
		if invalid {
			continue
		}

		safeReports++
	}

	fmt.Println("Safe reports:", safeReports)
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
