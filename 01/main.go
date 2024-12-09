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

func main() {
	col1, col2, numberOfOccurrences := readInput("input.txt")

	sort.Ints(col1)
	sort.Ints(col2)

	sum, score := calculateSumAndScore(col1, col2, numberOfOccurrences)

	fmt.Println("Sum:", sum)
	fmt.Println("Score:", score)
}

func readInput(fileName string) ([]int, []int, map[int]int) {
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var col1, col2 []int
	numberOfOccurrences := make(map[int]int)

	for fileScanner.Scan() {
		fields := strings.Fields(fileScanner.Text())

		n1, err := strconv.Atoi(strings.TrimSpace(fields[0]))
		checkError(err)

		n2, err := strconv.Atoi(strings.TrimSpace(fields[1]))
		checkError(err)

		col1 = append(col1, n1)
		col2 = append(col2, n2)

		numberOfOccurrences[n2]++
	}

	return col1, col2, numberOfOccurrences
}

func calculateSumAndScore(col1, col2 []int, occurrences map[int]int) (int, int) {
	sum, score := 0, 0

	for i := 0; i < len(col1); i++ {
		n1 := col1[i]
		n2 := col2[i]

		sum += int(math.Abs(float64(n1 - n2)))

		score += n1 * occurrences[n1]
	}

	return sum, score
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
