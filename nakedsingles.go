package gosudoku

func (b Board) NakedSingles() bool {
	foundAny := false
	found := true
	for found {
		b.ReduceCandidates()
		found = false
		for _, cell := range b.candidateCells() {
			if candidateCell, ok := cell.(CandidateCell); ok {
				if len(candidateCell.Candidates()) == 1 {
					b.ReplaceCellWithNewValueCell(cell, (candidateCell.Candidates())[0])
					found = true
					foundAny = true
				}
			}
		}
	}
	return foundAny
}

func NakedSinglesUnit(unit Unit) (updatedCells Unit) {
	for _, cell := range unit {
		if candidateCell, ok := cell.(CandidateCell); ok {
			if len(candidateCell.Candidates()) == 1 {
				updatedCells = append(updatedCells, ValueCellAt(candidateCell))
			}
		}
	}
	return updatedCells
}
