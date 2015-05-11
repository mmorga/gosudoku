package sudoku

func (b Board) ReduceCandidates() {
	for _, cell := range b.candidateCells() {
		seqCells := b.UnitCellsForCell(cell)
		valueCells := SelectCells(seqCells, isValueOrfixedCell)
		values := MapCellValues(valueCells)
		if candidateCell, ok := cell.(CandidateCell); ok {
			candidateCell.ReduceCandidates(values)
		}
	}
}
