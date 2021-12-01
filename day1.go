package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	depth_measured = []int{}
	previous_value int
	prev_sum       int
)

func sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {

	file, err := os.Open("day1_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		depth_measured = append(depth_measured, i)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	counter_part1 := 0
	counter_part2 := 0

	for idx, depth := range depth_measured {

		if idx == 0 {
			previous_value = depth
			prev_sum = sum(depth_measured[idx : idx+3])
			continue
		}

		if previous_value < depth {
			counter_part1 += 1
		}

		actual_sum := sum(depth_measured[idx : idx+3])

		if prev_sum < actual_sum {
			counter_part2 += 1
		}

		previous_value = depth
		prev_sum = actual_sum

	}

	fmt.Println("First part result:", counter_part1)
	fmt.Println("Second part result:", counter_part2)
}
