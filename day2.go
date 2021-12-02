package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	coordinatesFirstPart  = make(map[string]int)
	coordinatesSecondPart = make(map[string]int)

	previous_value int
	prev_sum       int
	direction      string
	aim            = 0
)

func main() {

	file, err := os.Open("day2_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		distance, _ := strconv.Atoi(line[1])
		direction := line[0]
		// fmt.Println(distance)

		if sumOfDirections, found := coordinatesFirstPart[direction]; found {
			sumOfDirections += distance
			coordinatesFirstPart[direction] = sumOfDirections

		} else {
			coordinatesFirstPart[direction] = distance

		}

		if direction == "down" || direction == "up" {
			if direction == "down" {
				aim += distance
			} else {
				aim -= distance
			}
		} else {
			coordinatesSecondPart["horizontal"] += distance
			coordinatesSecondPart["depth"] += distance * aim

		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	multipliedCoordinatesPart1 := (coordinatesFirstPart["down"] - coordinatesFirstPart["up"]) * coordinatesFirstPart["forward"]
	multipliedCoordinatesPart2 := coordinatesSecondPart["horizontal"] * coordinatesSecondPart["depth"]
	fmt.Println("First part: ", multipliedCoordinatesPart1)
	fmt.Println("First part: ", multipliedCoordinatesPart2)

}
