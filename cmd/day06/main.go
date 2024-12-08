package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lab, guard, err := readData("inputs/day06.txt")
	if err != nil {
		fmt.Println("error reading data:", err)
		return
	}

	visited := simulateGuardMovement(lab, guard)

	fmt.Println(visited)

}

type pos struct {
	x, y int
}

type guard struct {
	pos       pos
	direction string
}

func simulateGuardMovement(lab map[pos]string, guard guard) int {
	visited := 1

	for {
		// Store the current position
		prevPos := guard.pos

		moveGuard(&guard)

		char, exists := lab[guard.pos]
		if !exists {
			// Guard has left the lab
			break
		}

		if char == "#" {
			// Turn the guard and revert to the previous position
			turnGuard(&guard)
			guard.pos = prevPos
			continue
		}

		if char == "." {
			visited++
			continue
		}
	}

	return visited
}

func moveGuard(guard *guard) {
	switch guard.direction {
	case "^":
		guard.pos.y -= 1 // Move up
	case ">":
		guard.pos.x += 1 // Move right
	case "v":
		guard.pos.y += 1 // Move down
	case "<":
		guard.pos.x -= 1 // Move left
	}
}

func turnGuard(guardData *guard) {
	switch guardData.direction {
	case "^":
		guardData.direction = ">"
	case ">":
		guardData.direction = "v"
	case "v":
		guardData.direction = "<"
	case "<":
		guardData.direction = "^"
	default:
		fmt.Println("Invalid direction:", guardData.direction)
	}
}

func readData(filename string) (map[pos]string, guard, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error reading file:", err)
		return nil, guard{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var guardData guard
	lab := make(map[pos]string)
	yPos := 0

	for scanner.Scan() {
		line := scanner.Text()

		for i, char := range line {

			position := pos{x: i, y: yPos}

			if strings.Contains("<^>v", string(char)) {
				guardData = guard{pos: position, direction: string(char)}
				lab[position] = "."
				continue
			}

			lab[position] = string(char)
		}
		yPos++
	}

	if err := scanner.Err(); err != nil {
		return nil, guard{}, err
	}

	return lab, guardData, nil
}
