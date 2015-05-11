package sudoku

import "sort"

// This one is a unit oriented strategy

// * unit: row, col, box

// For each unique pair of candidates in a unit
// 	If unique pair exists exclusively twice in a unit
// 		Found!
// 		Remove pair values from other cells in unit

func (b Board) NakedPairs() bool {
	foundAny := false
	found := true
	for found {
		b.ReduceCandidates()
		found = false
		for _, unit := range b.Units() {
			for _, pair := range UniquePairsInUnit(unit) {
				nakedCells := PairAppearsNakedInUnit(pair, unit)
				if len(nakedCells) == 2 {
					if RemoveCandidatesFromUnitExceptForCells(pair, unit, nakedCells) {
						found = true
						foundAny = true
					}
				}
			}
		}
	}
	return foundAny
}

func UniquePairsInUnit(unit Unit) (pairs [][]int) {
	for _, cell := range unit {
		if candidateCell, ok := cell.(CandidateCell); ok {
			candidates := candidateCell.Candidates()
			if len(candidates) == 2 {
				pairs = append(pairs, candidates)
			}
		}

	}
	return pairs
}

func candidateSetsEqual(set1 []int, set2 []int) bool {
	if len(set1) != len(set2) {
		return false
	}
	sort.Ints(set1)
	sort.Ints(set2)
	for i, val := range set1 {
		if val != set2[i] {
			return false
		}
	}
	return true
}

func PairAppearsNakedInUnit(pair []int, unit Unit) (nakedCells Unit) {
	for _, cell := range unit {
		if candidateCell, ok := cell.(CandidateCell); ok {
			candidates := candidateCell.Candidates()
			if candidateSetsEqual(candidates, pair) {
				nakedCells = append(nakedCells, cell)
			}
		}
	}
	return nakedCells
}

func CellInUnit(cell Cell, unit Unit) bool {
	for _, c := range unit {
		if cell.Equal(c) {
			return true
		}
	}
	return false
}

func CandidatesContainAny(candidates []int, vals []int) bool {
	for _, val := range vals {
		for _, candidate := range candidates {
			if val == candidate {
				return true
			}
		}
	}
	return false
}

func RemoveCandidatesFromUnitExceptForCells(pair []int, unit Unit, nakedCells Unit) bool {
	found := false
	for _, cell := range unit {
		if candidateCell, ok := cell.(CandidateCell); ok {
			if !CellInUnit(cell, nakedCells) {
				if CandidatesContainAny(candidateCell.Candidates(), pair) {
					candidateCell.ReduceCandidates(pair)
					found = true
				}
			}
		}
	}
	return found
}
