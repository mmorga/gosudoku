package sudoku

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
			candidateCell.ReduceCandidates(values)
			changedCells = append(changedCells, cell)
		}
	}
	return changedCells
}
