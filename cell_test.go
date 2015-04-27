package sudoku

import "testing"

func TestCandidateCellCandidates(t *testing.T) {
	cell := new(CandidateCell)
	actual := cell.Candidates()
	if len(actual) != 0 {
		t.Errorf("Candidates should have been size 1, was %v", len(actual))
	}
}
