package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

type Ship struct {
	Heading	int
	Horizontal int
	Vertical int
}

type Waypoint struct {
	Horizontal int
	Vertical int
}

func (w *Waypoint) Rotate(direction int) {
	if direction == 1 {
		// CW Rotation
		t := w.Horizontal
		w.Horizontal = w.Vertical
		w.Vertical = -t
	} else if direction == -1 {
		// CCW Rotation
		t := w.Vertical
		w.Vertical = w.Horizontal
		w.Horizontal = -t
	}
}

func main() {
	begin := time.Now().UnixNano()
	fileData, _ := ioutil.ReadFile("resources/input_12.txt")
	travelData := strings.Split(string(fileData), "\r\n")

	var ship Ship
	ship.Heading = 90

	for _, k := range travelData {
		instruction := strings.Split(k, "")
		command := instruction[0]
		value, _ := strconv.Atoi(strings.Join(instruction[1:], ""))

		if command == "N" {
			ship.Vertical += value
		} else if command == "S" {
			ship.Vertical -= value
		} else if command == "E" {
			ship.Horizontal += value
		} else if command == "W" {
			ship.Horizontal -= value
		} else if command == "R" {
			ship.Heading = (ship.Heading + value) % 360
		} else if command == "L" {
			ship.Heading = (360 + ship.Heading - value) % 360
		} else if command == "F" {
			if ship.Heading == 0 {
				ship.Vertical += value
			} else if ship.Heading == 90 {
				ship.Horizontal += value
			} else if ship.Heading == 180 {
				ship.Vertical -= value
			} else if ship.Heading == 270 {
				ship.Horizontal -= value
			}
		}
	}
	fmt.Println(math.Abs(float64(ship.Horizontal)) + math.Abs(float64(ship.Vertical)))
	part1Elapsed := float64(time.Now().UnixNano() - begin) / float64(time.Millisecond)
	fmt.Println("Part 1 elapsed time (ms):", part1Elapsed)
	begin = time.Now().UnixNano()

	var ship2 Ship
	var waypoint Waypoint
	waypoint.Horizontal = 10
	waypoint.Vertical = 1

	for _, k := range travelData {
		instruction := strings.Split(k, "")
		command := instruction[0]
		value, _ := strconv.Atoi(strings.Join(instruction[1:], ""))

		if command == "N" {
			waypoint.Vertical += value
		} else if command == "S" {
			waypoint.Vertical -= value
		} else if command == "E" {
			waypoint.Horizontal += value
		} else if command == "W" {
			waypoint.Horizontal -= value
		} else if command == "R" {
			for i := 0; i < (value / 90); i++ {
				waypoint.Rotate(1)
			}
		} else if command == "L" {
			for i := 0; i < (value / 90); i++ {
				waypoint.Rotate(-1)
			}
		} else if command == "F" {
			ship2.Vertical += value * waypoint.Vertical
			ship2.Horizontal += value * waypoint.Horizontal
		}
	}
	fmt.Println(math.Abs(float64(ship2.Horizontal)) + math.Abs(float64(ship2.Vertical)))
	part2Elapsed := float64(time.Now().UnixNano() - begin) / float64(time.Millisecond)
	fmt.Println("Part 2 elapsed time (ms):", part2Elapsed)
}
