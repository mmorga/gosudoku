package sudoku

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
