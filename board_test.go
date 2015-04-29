package sudoku

import (
	"fmt"
	"strconv"
	"testing"
)

// This is TERRIBLE! Is there not a better way in Go to enter a multidimensional array???
//
// const EXPERT_BOARD_CANDIDATES := [][][]int{
//     [][]int{[4, 5, 9],       [6],          [1],                [2, 5],          [3, 4],       [2, 4, 5],    [2, 4, 5, 7, 8, 9], [2, 3, 4, 9],    [2, 3, 4, 5, 7, 8]]}
//     [][]int{[8],             [4, 7],       [4, 5, 7, 9],       [1, 2, 5, 6],    [1, 3, 4, 6], [2, 4, 5],    [2, 4, 5, 6, 7, 9], [2, 3, 4, 6, 9], [2, 3, 4, 5, 7]]}
//     [][]int{[2],             [3],          [4, 5],             [7],             [9],          [8],          [1],                [4, 6],          [4, 5]]}
//     [][]int{[1, 4, 5],       [9],          [2, 4, 5],          [3],             [8],          [7],          [2, 4],             [1, 2, 4],       [6]]}
//     [][]int{[1, 3, 4],       [1, 2, 4, 8], [2, 3, 4, 8],       [9],             [5],          [6],          [2, 4, 8],          [7],             [1, 2, 3, 4, 8]]}
//     [][]int{[3, 6],          [7, 8],       [3, 6, 7, 8],       [4],             [2],          [1],          [8, 9],             [5],             [3, 8]]}
//     [][]int{[7],             [5],          [2, 4, 6, 8, 9],    [1, 2, 6, 8],    [1, 4, 6],    [2, 4, 9],    [3],                [1, 2, 4, 6],    [1, 2, 4]]}
//     [][]int{[1, 3, 4, 6, 9], [1, 2, 4, 8], [2, 3, 4, 6, 8, 9], [1, 2, 5, 6, 8], [1, 4, 6, 7], [2, 4, 5, 9], [2, 4, 5, 6, 7],    [1, 2, 4, 6],    [1, 2, 4, 5, 7]]}
//     [][]int{[1, 4, 6],       [1, 2, 4],    [2, 4, 6],          [1, 2, 5, 6],    [1, 4, 6, 7], [3],          [2, 4, 5, 6, 7],    [8], [9]}

// func TestBoardFromArray(t *testing.T) {
// 	board := BoardFromArray(EXPERT_BOARD_CANDIDATES)

// 	actual := board.cell(1, 4)
// 	if actual != []int{1, 2, 4, 8} {
// 		t.Errorf("Cell value was not expected: %v", actual)
// 	}
// }

const (
	expertBoardSource = "_61______\n8________\n23_7981__\n_9_38___6\n___9_6_7_\n____21_5_\n75____3__\n_________\n_____3_89"
	sampleBoard       = "_47_8___6\n__2_3_489\n____5621_\n5____476_\n871__5_9_\n___79___1\n2_45_8___\n7_394____\n___3__148"
)

func CompareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}

func TestBoardFromString(t *testing.T) {
	board, err := BoardFromString(expertBoardSource)
	if err != nil {
		t.Errorf("Unable to create board from string (%v):\n%v", err, expertBoardSource)
	}
	actualCell := board.Cell(1, 4)
	if err != nil {
		t.Errorf("Unable to get cell from board (%v)", err)
	}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	candidateCell, ok := actualCell.(CandidateCell)
	if !ok {
		t.Errorf("Unable to cast Cell as a CandidateCell: %v", candidateCell)
	}

	candidates := candidateCell.Candidates()
	if !CompareSlices(candidates, expected) {
		t.Errorf("Cell Candidates were different than expected:\nExpected: %v\nActual: %v", expected, candidates)
	}
}

func TestfixedCells(t *testing.T) {
	board, err := BoardFromString(expertBoardSource)
	if err != nil {
		t.Errorf("Unable to create board from string (%v):\n%v", err, expertBoardSource)
	}

	fixedCells := board.fixedCells()
	fixedCellsLen := len(fixedCells)
	if fixedCellsLen != 25 {
		t.Errorf("Expected to find %v fixed cells, actually was %d", 25, fixedCellsLen)
	}
}

func TestReduce(t *testing.T) {
	board, _ := BoardFromString(sampleBoard)
	board.ReduceCandidates()
	// fmt.Printf("Board:\n%v\n\n", board.cells)
}

func TestNakedSingles(t *testing.T) {
	board, _ := BoardFromString(sampleBoard)
	board.NakedSingles()
	if !board.IsSolved() {
		t.Errorf("Expected Board to be solved\n%v", board)
	}
}

func TestHiddenSingles(t *testing.T) {
	board, _ := BoardFromString(expertBoardSource)
	// fmt.Printf("Initial Board:\n\n%v\n\n", board)
	board.NakedSingles()
	// fmt.Printf("Naked Singles Board:\n\n%v\n\n", board)
	board.HiddenSingles()
	// fmt.Printf("Final Board:\n\n%v\n\n", board)
	if board.IsSolved() {
		t.Errorf("Expected Board to NOT be solved\n%v", board)
	}
}

func TestExploration(t *testing.T) {
	dest := make([]int, 9, 9)
	for i := 0; i < 9; i++ {
		dest[i] = i
	}
	src := []int{10, 11, 12}
	dest = src
}

func TestExploration2(t *testing.T) {
	aRune := '4'
	str := fmt.Sprintf("%c", aRune)
	if str != "4" {
		t.Errorf("Rune convert didnt work like I thought, str is [%v]", str)
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		t.Errorf("Error converting str to int: %v", err)
	}
	if val != 4 {
		t.Errorf("Sscanf didn't return the expected value")
	}
}
