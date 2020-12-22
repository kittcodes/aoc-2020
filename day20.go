package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Tile struct {
	Id          int
	Values      [][]int
	TopEdge     []int
	RightEdge   []int
	BottomEdge  []int
	LeftEdge    []int
	TopMatch    int
	RightMatch  int
	BottomMatch int
	LeftMatch   int
	MatchCount  int
}

func (t Tile) trimEdges() [][]int {
	var newValues [][]int

	for rowIndex, row := range t.Values {
		if rowIndex == 0 || rowIndex == len(t.Values) - 1 {
			continue
		}
		var newRow []int
		for colIndex, value := range row {
			if colIndex == 0 || colIndex == len(row) - 1 {
				continue
			}
			newRow = append(newRow, value)
		}
		newValues = append(newValues, newRow)
	}
	return newValues
}

func reverse(slice []int) []int {
	var newSlice []int
	for i := len(slice) - 1; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}
	return newSlice
}

var (
	tileSet    map[int]*Tile
	digitRegex *regexp.Regexp
	tileMap 	[][]Tile
	fileName	string
)

func (t *Tile) IncrementMatchCount() {
	t.MatchCount++
}

func FlipVIntSlice(t [][]int) [][]int {
	var response [][]int
	for i := len(t) - 1; i >= 0; i-- {
		response = append(response, t[i])
	}
	return response
}

func (t *Tile) FlipV(findEdges bool) {
	newTile := Tile{}
	for i := len(t.Values) - 1; i >= 0; i-- {
		newTile.Values = append(newTile.Values, t.Values[i])
	}
	t.Values = newTile.Values
	if findEdges {
		temp := t.TopMatch
		t.TopMatch = t.BottomMatch
		t.BottomMatch = temp
		t.FindEdges(true)
	}
}

func FlipHIntSlice(t [][]int) [][]int {
	var response [][]int
	for _, row := range t {
		var newRow []int
		for i := len(row) - 1; i >= 0; i-- {
			newRow = append(newRow, row[i])
		}
		response = append(response, newRow)
	}
	return response
}

func (t *Tile) FlipH(findEdges bool) {
	newTile := Tile{}
	for _, line := range t.Values {
		var newLine []int
		for i := len(line)-1; i >= 0; i-- {
			newLine = append(newLine, line[i])
		}
		newTile.Values = append(newTile.Values, newLine)
	}
	t.Values = newTile.Values
	if findEdges {
		t.FindEdges(true)
	}
}


func (t *Tile) Save() {
	tileSet[t.Id] = t
}

func IntSlicesMatch(input, target []int) bool {
	if len(input) != len(target) {
		return false
	}

	for i := 0; i < len(target); i++ {
		if input[i] != target[i]  {
			return false
		}
	}
	return true
}

func (t *Tile) RotateMatchingEdges() {
	temp := t.LeftMatch
	t.LeftMatch = t.BottomMatch
	t.BottomMatch = t.RightMatch
	t.RightMatch = t.TopMatch
	t.TopMatch = temp
	t.Save()
}

func (t *Tile) FindMatchingEdges() {
	t.MatchCount = 0
	t.LeftMatch, t.TopMatch, t.RightMatch, t.BottomMatch = 0, 0, 0, 0
	t.FindMatchingEdge("top")
	t.FindMatchingEdge("right")
	t.FindMatchingEdge("bottom")
	t.FindMatchingEdge("left")
	t.Save()
}

func (t *Tile) FindMatchingEdge(edgeName string) {
	var edge []int
	switch edgeName {
	case "top":
		edge = t.TopEdge
	case "right":
		edge = t.RightEdge
	case "bottom":
		edge = t.BottomEdge
	case "left":
		edge = t.LeftEdge
	default:
		fmt.Println("Something fucked up")
		return
	}

	revEdge := reverse(edge)

	for _, tile := range tileSet {
		if t.Id == tile.Id {
			continue
		}
		for _, compEdge := range [][]int{tile.TopEdge, tile.RightEdge, tile.BottomEdge, tile.LeftEdge} {
			isMatch := true
			isReverseMatch := true

			if !IntSlicesMatch(edge, compEdge) {
				isMatch = false
			}
			if !IntSlicesMatch(revEdge, compEdge) {
				isReverseMatch = false
			}
			if isMatch || isReverseMatch {
				t.IncrementMatchCount()
				switch edgeName {
				case "top":
					t.TopMatch = tile.Id
				case "right":
					t.RightMatch = tile.Id
				case "bottom":
					t.BottomMatch = tile.Id
				case "left":
					t.LeftMatch = tile.Id
				}
				t.Save()
				break
			}
		}
	}
}

func RotateInts90CW(t [][]int) {
	var N = len(t)
	for i := 0; i < N/2; i++ {
		for j := i; j < N - i - 1; j++ {
			temp := t[i][j]
			t[i][j] = t[N - 1 - j][i]
			t[N - 1 - j][i] = t[N - 1 - i][N - 1 - j]
			t[N - 1 - i][N - 1 - j] = t[j][N - 1 - i]
			t[j][N - 1 - i] = temp
		}
	}
}

func IsIntSliceNessie(t [][]int) bool {
	searchRows := []int{1, 2, 0}
	for _, i := range searchRows {
		for j := 0; j < len(t[i]); j++ {
			if t[i][j] & Nessie[i][j] != Nessie[i][j] {
				return false
			}
		}
	}
	return true
}

func (t *Tile) Rotate90CW(findEdges bool) {
	var N = len(t.Values)
	for i := 0; i < N/2; i++ {
		for j := i; j < N - i - 1; j++ {
			temp := t.Values[i][j]
			t.Values[i][j] = t.Values[N - 1 - j][i]
			t.Values[N - 1 - j][i] = t.Values[N - 1 - i][N - 1 - j]
			t.Values[N - 1 - i][N - 1 - j] = t.Values[j][N - 1 - i]
			t.Values[j][N - 1 - i] = temp
		}
	}
	if findEdges {
		t.FindEdges(true)
	}
}


func (t *Tile) FindEdges(save bool) {
	var rightEdge, leftEdge []int
	for _, line := range t.Values {
		leftEdge = append(leftEdge, line[0])
		rightEdge = append(rightEdge, line[len(line)-1])
	}
	t.TopEdge = t.Values[0]
	t.LeftEdge = leftEdge
	t.RightEdge = rightEdge
	t.BottomEdge = t.Values[len(t.Values)-1]
	if save {
		t.Save()
	}
}

var NessieTop = []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0}
var NessieMid = []int{1,0,0,0,0,1,1,0,0,0,0,1,1,0,0,0,0,1,1,1}
var NessieBot = []int{0,1,0,0,1,0,0,1,0,0,1,0,0,1,0,0,1,0,0,0}

var Nessie = [][]int{NessieTop, NessieMid, NessieBot}

func CountNessie(t [][]int) int {
	foundNessies := 0

	for rowIndex, row := range t {
		if rowIndex >= len(t) - len(Nessie) {
			break
		}
		for colIndex := range row {
			if colIndex >= len(row) - len(NessieMid) {
				break
			}
			var suspectSection [][]int
			suspectSection = append(suspectSection, t[rowIndex][colIndex:colIndex+20])
			suspectSection = append(suspectSection, t[rowIndex+1][colIndex:colIndex+20])
			suspectSection = append(suspectSection, t[rowIndex+2][colIndex:colIndex+20])
			if IsIntSliceNessie(suspectSection) {
				foundNessies++
			}
		}
	}
	return foundNessies
}

func init() {
	flag.StringVar(&fileName, "input", "input20.txt", "Input file for this problem")
	flag.Parse()
}

func main() {
	begin := time.Now().UnixNano()
	digitRegex = regexp.MustCompile(`([0-9]+)`)
	fileData, _ := ioutil.ReadFile(fileName)
	tileData := strings.Split(string(fileData), "\r\n\r\n")

	tileNum := -1
	tileSet = make(map[int]*Tile)

	for _, tile := range tileData	{
		var newTile Tile
		lines := strings.Split(tile, "\r\n")
		for i, line := range lines {
			if i == 0 {
				subMatch := digitRegex.FindStringSubmatch(line)
				if len(subMatch) > 0 {
					tileNum, _ = strconv.Atoi(subMatch[0])
				}
				newTile.Id = tileNum
			} else {
				tokens := strings.Split(line, "")
				var intLine []int
				for _, token := range tokens {
					tokVal := 0
					if token == "#" {
						tokVal = 1
					}
					intLine = append(intLine, tokVal)
				}
				newTile.Values = append(newTile.Values, intLine)
			}
		}
		newTile.FindEdges(false)
		newTile.Save()
	}
	elapsed := float64(time.Now().UnixNano() - begin) / float64(time.Millisecond)
	step := time.Now().UnixNano()
	fmt.Println("Tiles parsed and loaded, edges defined. Time (ms):", elapsed)

	for _, tile := range tileSet {
		tile.FindMatchingEdges()
	}

	elapsed = float64(time.Now().UnixNano() - step) / float64(time.Millisecond)
	step = time.Now().UnixNano()
	fmt.Println("Matching edges found. Time (ms):", elapsed)

	originTileId := -1
	for _, tile := range tileSet {
		if tile.MatchCount == 2 {
			for tile.TopMatch == 0 || tile.RightMatch == 0 {
				tile.Rotate90CW(true)
				tile.RotateMatchingEdges()
			}
			originTileId = tile.Id
			break
		}
	}
	elapsed = float64(time.Now().UnixNano() - step) / float64(time.Millisecond)
	step = time.Now().UnixNano()
	fmt.Println("Origin tile found. Time (ms):", elapsed)

	workingTile := tileSet[originTileId]

	for true {
		var row []Tile
		for true {

			row = append(row, *workingTile)

			nextId := workingTile.RightMatch
			if nextId == 0 {
				break
			}
			nextTile := tileSet[nextId]

			// Line up the match
			for nextTile.LeftMatch != workingTile.Id {
				nextTile.Rotate90CW(true)
				nextTile.RotateMatchingEdges()
			}

			// FlipV
			if !IntSlicesMatch(nextTile.LeftEdge, workingTile.RightEdge) {
				nextTile.FlipV(true)
			}

			workingTile = tileSet[nextTile.Id]
		}
		if len(row) > 0 {
			tileMap = append(tileMap, row)
			nextId := row[0].TopMatch

			if nextId == 0 {
				break
			}
			top := tileSet[nextId]

			// Line up the match
			for top.BottomMatch != row[0].Id {
				top.Rotate90CW(true)
				top.FindMatchingEdges()
			}

			// FlipV
			if !IntSlicesMatch(top.BottomEdge, row[0].TopEdge) {
				top.FlipH(true)
				top.FindMatchingEdges()
			}

			workingTile = tileSet[top.Id]
		} else {
			break
		}
		row = nil
	}
	elapsed = float64(time.Now().UnixNano() - step) / float64(time.Millisecond)
	step = time.Now().UnixNano()
	fmt.Println("Tiles mapped. Time (ms):", elapsed)

	for rowIndex, row := range tileMap {
		for colIndex, tile := range row {
			var newTile Tile
			newTile.Id = tile.Id
			newTile.Values = tile.trimEdges()
			tileMap[rowIndex][colIndex] = newTile
		}
	}

	var fullMap [][]int
	totalOn := 0

	for _, row := range tileMap {
		var newRow1, newRow2, newRow3, newRow4, newRow5, newRow6, newRow7, newRow8 []int
		for _, tile := range row {
			for valRowIndex, rowVal := range tile.Values {
				for _, cell := range rowVal {
					totalOn += cell
					if valRowIndex == 0 {
						newRow1 = append(newRow1, cell)
					} else if valRowIndex == 1 {
						newRow2 = append(newRow2, cell)
					} else if valRowIndex == 2 {
						newRow3 = append(newRow3, cell)
					} else if valRowIndex == 3 {
						newRow4 = append(newRow4, cell)
					} else if valRowIndex == 4 {
						newRow5 = append(newRow5, cell)
					} else if valRowIndex == 5 {
						newRow6 = append(newRow6, cell)
					} else if valRowIndex == 6 {
						newRow7 = append(newRow7, cell)
					} else if valRowIndex == 7 {
						newRow8 = append(newRow8, cell)
					}
				}
			}
		}
		fullMap = append(fullMap, newRow8, newRow7, newRow6, newRow5, newRow4, newRow3, newRow2, newRow1)
	}

	elapsed = float64(time.Now().UnixNano() - step) / float64(time.Millisecond)
	step = time.Now().UnixNano()
	fmt.Println("Final image defined. Time (ms):", elapsed)
	maxNessie := 0
	for j := 0; j < 2; j++ {
		for i := 0; i < 4; i++ {
			nessie := CountNessie(fullMap)
			if nessie > maxNessie {
				maxNessie = nessie
				break
			}
			RotateInts90CW(fullMap)
		}
		if maxNessie > 0 {
			break
		}
		fullMap = FlipHIntSlice(fullMap)
	}

	fmt.Println(totalOn - maxNessie*15)
	elapsed = float64(time.Now().UnixNano() - step) / float64(time.Millisecond)
	fmt.Println("Finished. Time (ms):", elapsed)
	elapsed = float64(time.Now().UnixNano() - begin) / float64(time.Millisecond)
	fmt.Println("Elapsed time (ms):", elapsed)
}
