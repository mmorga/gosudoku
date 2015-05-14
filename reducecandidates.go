package gosudoku

func (b Board) ReduceCandidates() {
	for _, cell := range b.candidateCells() {
		seqCells := b.UnitCellsForCell(cell)
		valueCells := SelectCells(seqCells, isValueOrFixedCell)
		values := MapCellValues(valueCells)
		if candidateCell, ok := cell.(CandidateCell); ok {
			candidateCell.ReduceCandidates(values)
		}
	}
}

func ReduceCandidatesUnit(unit Unit) (changedCells Unit) {
	for _, cell := range unit {
		values := MapCellValues(SelectCells(unit, isValueOrFixedCell))
		if candidateCell, ok := cell.(CandidateCell); ok {
			candidates := candidateCell.Candidates()
			finalCandidates := subtractValsFromSlice(candidates, values)
			if len(candidates) > len(finalCandidates) {
				changedCells = append(changedCells, newCandidateCell(candidateCell.X(), candidateCell.Y(), finalCandidates))
			}
		}
	}
	return changedCells
}

func subtractValsFromSlice(initial []int, minus []int) (result []int) {
	for val := range initial {
		var includeVal = true
		for _, v := range minus {
			if val == v {
				includeVal = false
			}
		}
		if includeVal {
			result = append(result, val)
		}
	}
	return result
}
