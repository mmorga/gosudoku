package sudoku

import (
	"fmt"
	"sort"
)

func Solve(filename string) {
	puzzles := ReadFile(filename)
	var titles []string
	for title := range puzzles {
		titles = append(titles, title)
	}
	sort.Strings(titles)
	total := len(titles)
	solved := 0
	for _, title := range titles {
		board, _ := BoardFromString(puzzles[title])
		for (board.NakedPairs() || board.NakedSingles() || board.HiddenSingles()) && !board.IsSolved() {
		}
		if board.IsSolved() {
			solved++
		}
	}
	fmt.Printf("\nResults: Solved %d of %d puzzles.\n", solved, total)
}
