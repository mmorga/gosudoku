package sudoku

type Unit []Cell

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

func CandidateCellsWithValues(unit Unit, vals []int) (foundCells Unit) {
	for _, cell := range unit {
		if candidateCell, ok := cell.(CandidateCell); ok {
			if CandidateSetsEqual(candidateCell.Candidates(), vals) {
				foundCells = append(foundCells, cell)
			}
		}
	}
	return foundCells
}

func CandidateCellsWithAnyValues(unit Unit, vals []int) (foundCells Unit) {
	for _, cell := range unit {
		if candidateCell, ok := cell.(CandidateCell); ok {
			if CandidatesContainAny(candidateCell.Candidates(), vals) {
				foundCells = append(foundCells, cell)
			}
		}
	}
	return foundCells
}

func CellInUnit(cell Cell, unit Unit) bool {
	for _, c := range unit {
		if cell.Equal(c) {
			return true
		}
	}
	return false
}

func RemoveCandidatesFromUnitExceptForCells(pair []int, unit Unit, exceptCells Unit) bool {
	found := false
	for _, cell := range unit {
		if candidateCell, ok := cell.(CandidateCell); ok {
			if !CellInUnit(cell, exceptCells) {
				if CandidatesContainAny(candidateCell.Candidates(), pair) {
					candidateCell.ReduceCandidates(pair)
					found = true
				}
			}
		}
	}
	return found
}

func FlattenUnitSlice(nestedCells []Unit) (unit Unit) {
	for _, row := range nestedCells {
		for _, cell := range row {
			unit = append(unit, cell)
		}
	}
	return unit
}

func candidateVals(cells Unit) (vals []int) {
	canVals := make(map[int](bool), 9)
	for _, cell := range cells {
		if candidateCell, ok := cell.(CandidateCell); ok {
			for _, val := range candidateCell.Candidates() {
				canVals[val] = true
			}
		}
	}
	for val := range canVals {
		vals = append(vals, val)
	}
	return vals
}
