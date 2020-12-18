package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseFormula(formula string) int {
	// Find all indexes (ranges) of open/close parents
	openIndexes := openParenRE.FindAllStringIndex(formula, -1)
	closeIndexes := closeParenRE.FindAllStringIndex(formula, -1)

	// Grab just the first index from each into an int slice
	var openIndexList, closeIndexList []int
	for _, k := range openIndexes {
		openIndexList = append(openIndexList, k[0])
	}
	for _, k := range closeIndexes {
		closeIndexList = append(closeIndexList, k[0])
	}

	var subFormulas []string

	// For every close paren...
	for _, k := range closeIndexList {
		matchingOpenParenIndex, removeIndex := -1, -1
		// Find the nearest open paren before it
		for i, l := range openIndexList {
			if l > k { break }
			if l > matchingOpenParenIndex {
				matchingOpenParenIndex, removeIndex = l, i
			}
		}
		// Remove that open paren from the list of indexes
		openIndexList = append(openIndexList[:removeIndex], openIndexList[removeIndex+1:]...)

		// Add the contents to a slice called subFormulas
		subFormulas = append(subFormulas, formula[matchingOpenParenIndex:k+1])
	}

	var searchList []string
	var replaceList []int

	for true {
		var newSubFormulas []string

		// For every subFormula
		for _, k := range subFormulas {
			// Trim out the opening parens
			eval := k[1:len(k)-1]

			// If this contains parens, it must contain a subformula
			if strings.Contains(eval, "(") {
				// Hold onto it for later
				newSubFormulas = append(newSubFormulas, k)
			} else {
				//If this is just a vanilla formula, then add it to the list to search/replace
				searchList = append(searchList, k)

				// Evaluate the value for the replacement
				replaceList = append(replaceList, EvalReverseOrder(eval))
			}
		}

		// If we have no containing subformulas, kick out of the infinite loop
		if len(newSubFormulas) == 0 {
			break
		}

		// For each of our remaining subformulas...
		for subFormulaIndex, subFormula := range newSubFormulas {
			for searchIndex, searchItem := range searchList {
				// Replace the subformulas we evaluated above
				subFormula = strings.ReplaceAll(subFormula, searchItem, fmt.Sprint(replaceList[searchIndex]))
			}
			newSubFormulas[subFormulaIndex] = subFormula
		}
		subFormulas = newSubFormulas
	}

	// Now that we have a comprehensive list (in order of evaluation and replacement)
	for searchIndex, searchItem := range searchList {
		// Replace every subformula in the original formula with the search/replace items
		formula = strings.ReplaceAll(formula, searchItem, fmt.Sprint(replaceList[searchIndex]))
	}
	// Evaluate the remaining simple formula
	return EvalReverseOrder(formula)
}

func EvalReverseOrder(formula string) int {
	// While there's a + in the string
	for strings.Contains(formula, "+") {

		// Split into fields
		fields := strings.Fields(formula)

		// Loop through fields
		for i, k := range fields {
			if k == "+" {
				// Once we find a +, we add those two numbers, and replace the value in the formula
				search := fmt.Sprintf("%s + %s", fields[i-1], fields[i+1])

				prev, _ := strconv.Atoi(fields[i-1])
				next, _ := strconv.Atoi(fields[i+1])
				replace := fmt.Sprintf("%d", prev + next)
				formula = strings.Replace(formula, search, replace, 1)
				// Break the first loop to repeat the fields split
				break
			}
		}
	}
	// Now we should have # * # * # ..., simply evaluate using the Part 1 formula
	return EvaluateFormula(formula)
}

func EvaluateFormula(formula string) int {
	result := 0
	operator := ""

	// Split simple formula into fields
	fields := strings.Fields(formula)
	for i, field := range fields {
		if i == 0 {
			// starting result is the first number
			numVal, _ := strconv.Atoi(field)
			result = numVal
		} else {
			if field == "+" || field == "*" {
				operator = field
			} else {
				numVal, _ := strconv.Atoi(field)
				if operator == "+" {
					// Add next number
					result += numVal
				} else if operator == "*" {
					// Multiply next number
					result *= numVal
				}
			}
		}
	}
	return result
}

var (
	openParenRE = regexp.MustCompile(`[(]`)
	closeParenRE = regexp.MustCompile(`[)]`)
)

func main() {
	begin := time.Now().UnixNano()
	fileData, _ := ioutil.ReadFile("resources/input_18.txt")
	formulas := strings.Split(string(fileData), "\r\n")

	sum := 0
	for _, formula := range formulas {
		val := ParseFormula(formula)
		sum += val
		//fmt.Println("Evaluating:", formula, " =", ParseFormula(formula))
	}
	fmt.Println(sum)
	elapsed := float64(time.Now().UnixNano() - begin) / float64(time.Millisecond)
	fmt.Println("Elapsed time (ms):", elapsed)
}
