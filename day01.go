package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	FinalTarget    int
	ComponentCount int
	InputFile      string
	StartTime      int64
)

func init() {
	StartTime = time.Now().UnixNano()
	flag.IntVar(&FinalTarget, "target", 2020, "Target total value for subset sum")
	flag.IntVar(&ComponentCount, "count", 2, "Number of components that will equal the target")
	flag.StringVar(&InputFile, "input", "inputs/input_01.txt", "Puzzle input file")
	flag.Parse()

	if ComponentCount < 2 {
		log.Fatal("Cannot have a count of less than 2")
	}
}

func subsetSumSolver(sumTarget, count int, candidates []int) ([]int, bool) {
	var solution []int
	if len(candidates) < count {
		//fmt.Println("sanity check", len(candidates), count)
		// Can't possibly make an n-component sum from set of size m where n > m
		return nil, false
	}
	if count > 2 {
		// 3-or-more-component subset sum problem
		for true {
			if len(candidates) == 0 {
				// Sanity check
				return nil, false
			}

			// Grab the first element from the slice as a candidate
			firstValue := candidates[0]
			if firstValue == 0 {
				// Edge case for parsing errors where a 0 value comes in
				candidates = candidates[1:]
				continue
			}

			// Try to solve for n-1 components with the remainder of the data set (and a new target)
			innerSolution, isSolved := subsetSumSolver(sumTarget-firstValue, count-1, candidates[1:])

			if isSolved {
				// If we have a valid solution, add the original value
				innerSolution = append(innerSolution, firstValue)

				// Sort for my sanity
				sort.Ints(innerSolution)
				return innerSolution, true
			} else {
				// If we don't have a valid solution, pop out the first value and repeat
				candidates = candidates[1:]
			}
		}
	} else {
		// Now we're looking at a 2-component subset sum problem
		for true {
			if len(candidates) < count {
				// count exceeds the number of candidate values, meaning no subset makes the sum
				return nil, false
			} else if len(candidates) == 0 {
				return nil, false
			}
			// Find the pair target number by subtracting the first element from the sumTarget
			pairTarget := sumTarget - candidates[0]

			// The index of where we'll trim the slice
			sliceMax := len(candidates)-1

			// Iterate through the remainder of the set
			for index, candidateValue := range candidates[1:] {
				if candidateValue > pairTarget && index <= sliceMax {
					// If the candidateValue exceeds the pairTarget, we've gone too far.
					// This value (and all larger values) cannot possibly be a component of the sum.
					sliceMax = index
					break
				}
				if candidateValue == pairTarget {
					// The subset is the first element and this candidateValue
					solution = []int{candidates[0], candidateValue}
					return solution, true
				}
			}
			// If we didn't find a match, repeat by popping the first element out, and trimming off any value too large to be a component
			candidates = candidates[1: sliceMax+1]
		}
	}
	// This line is never reached, but the golang compiler whines if I don't put it here
	return nil, false
}

func main() {
	data, _ := ioutil.ReadFile(InputFile)
	numbersFromFile := strings.Split(string(data), "\r\n")
	var dataSetNumbers []int

	// Build out an int slice of the numbers from the data set
	for _, k := range numbersFromFile {
		v, err := strconv.Atoi(k)
		if err != nil { continue }
		dataSetNumbers = append(dataSetNumbers, v)
	}
	// Sort initial values so we can eliminate some values from comparisons
	sort.Ints(dataSetNumbers)

	solutionSet, isSolved := subsetSumSolver(FinalTarget, ComponentCount, dataSetNumbers)
	if isSolved {
		fmt.Println("Solution found:", solutionSet)
	} else {
		fmt.Println("No solution found in the dataset")
	}
	elapsedTime := float64(time.Now().UnixNano() - StartTime) / float64(time.Millisecond)
	fmt.Println("Execution Time Elapsed (ms):", elapsedTime)
}