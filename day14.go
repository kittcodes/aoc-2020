package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func ReturnMemoryAddresses(mask []string, addr int) []int {
	var addresses, XIndexes []int
	XCount := strings.Count(strings.Join(mask, ""), "X")
	for i, k := range mask {
		if k == "X" {
			// Grab the index of every "X" in the mask
			XIndexes = append(XIndexes, i)
		}
	}

	// Perform the mask+addr work
	addrStr := strings.Split(fmt.Sprintf("%036s", strconv.FormatInt(int64(addr), 2)), "")
	for i, k := range mask {
		if k == "0" {
			mask[i] = addrStr[i]
		} else if k == "1" {
			mask[i] = "1"
		} else {
			continue
		}
	}

	for i := 0; i < int(math.Pow(float64(2), float64(XCount))); i++ {
		// Convert the counter to a 0-padded binary representation in a string slice.
		// The string length is XCount so the length is the same as XIndexes.
		filterString := fmt.Sprint("%0", XCount, "s")
		indexesToFlip := strings.Split(fmt.Sprintf(filterString, strconv.FormatInt(int64(i), 2)), "")

		// For each X in the mask, set to corresponding value from indexesToFlip
		for j, k := range XIndexes {
			mask[k] = indexesToFlip[j]
		}
		newInt, _ := strconv.ParseInt(strings.Join(mask, ""), 2, 64)
		addresses = append(addresses, int(newInt))
	}
	return addresses
}


func main() {
	begin := time.Now().UnixNano()
	fileData, _ := ioutil.ReadFile("resources/input_14.txt")
	inputData := strings.Split(string(fileData), "\r\n")

	part1MemoryValues := make(map[int]int)
	part2MemoryValues := make(map[int]int)
	var currentMask []string
	for _, line := range inputData {
		record := strings.Split(line, " = ")
		if record[0] == "mask" {
			// Set the current mask and move to the next line
			currentMask = strings.Split(record[1], "")
			continue
		}

		// Grab the address value (trim out "mem[" and "]", and conver to int
		stringAddressValue := strings.Split(record[0], "")[4:len(record[0])-1]
		intAddressValue, _ := strconv.Atoi(strings.Join(stringAddressValue, ""))

		// Convert the target value and the memory address to 36-char, 0-padded binary value
		valueToSet, _ := strconv.Atoi(record[1])
		binaryStringVal := strings.Split(fmt.Sprintf("%036s", strconv.FormatInt(int64(valueToSet), 2)), "")
		binaryMemoryAddressStringVal := strings.Split(fmt.Sprintf("%036s", strconv.FormatInt(int64(intAddressValue), 2)), "")

		var part1Target, part2Target []string

		for i, k := range currentMask {
			part1Result := ""
			part2Result := ""
			if k == "1" {
				part1Result = "1"
				part2Result = "1"
			} else if k == "0" {
				part1Result = "0"
				part2Result = binaryMemoryAddressStringVal[i]
			} else if k == "X" {
				part1Result = binaryStringVal[i]
				part2Result = "X"
			}
			part1Target = append(part1Target, part1Result)
			part2Target = append(part2Target, part2Result)
		}

		// Part 1
		filteredIntVal, _ := strconv.ParseInt(strings.Join(part1Target, ""), 2, 64)
		part1MemoryValues[intAddressValue] = int(filteredIntVal)

		// Part 2
		addrsToSet := ReturnMemoryAddresses(part2Target, intAddressValue)
		for _, k := range addrsToSet {
			part2MemoryValues[k] = valueToSet
		}

	}

	sum := 0
	for _, k := range part1MemoryValues {
		sum += k
	}
	fmt.Println("Part 1:", sum)		// Part 1

	sum = 0
	for _, k := range part2MemoryValues {
		sum += k
	}
	fmt.Println("Part 2:", sum)		// Part 2
	elapsed := float64(time.Now().UnixNano() - begin) / float64(time.Millisecond)
	fmt.Println("Elapsed time (ms):", elapsed)
}
