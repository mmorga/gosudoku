package sudoku

func (b Board) HiddenSingles() bool {
	foundAny := false
	found := true
	for found {
		b.ReduceCandidates()
		found = false
		for _, row := range b.Sequences() {
			candidateCells := SelectCells(row, isCandidateCell)
			candidateVals := candidateVals(candidateCells)
			for _, val := range candidateVals {
				hasValCells := hasCandidateVal(candidateCells, val)
				if len(hasValCells) == 1 {
					b.ReplaceCellWithNewValueCell(hasValCells[0], val)
					// fmt.Printf("Hidden Single Found in row %d, col: %d\n", rowIdx, hasValCells[0].X())
					found = true
					foundAny = true
				}
			}
		}
	}
	return foundAny
}

func hasCandidateVal(cells []Cell, val int) (foundCells []Cell) {
	for _, cell := range cells {
		if candidateCell, ok := cell.(CandidateCell); ok {
			if hasVal(candidateCell.Candidates(), val) {
				foundCells = append(foundCells, cell)
			}
		}
	}
	return foundCells
}

func hasVal(vals []int, val int) bool {
	for _, v := range vals {
		if v == val {
			return true
		}
	}
	return false
}

func candidateVals(cells []Cell) (vals []int) {
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

// # Hidden n strategy
// 1. set status to false
// 2. For each in board sequences
//   a. is there a set of size n for any array entries that those n values
// only exist together in exists in only n place(s)?
//     1. set that cell value to n
//     2. remove all n values from the other array entries
//     3. set status to true
// 3. return status
//
// For a. above what does this entail?
// 1. collect the unique set of values that occur in all array entries
// 2. for each combination(n)
//   a. collect the index of every array entry in which it occurs
// 3. return the collection of value-index pairs with a count of n
//
//
// Hidden n means in a sequence, n values occur in n cells and not in any
// other cells of the sequence.
// For n > 2, the n cells must each contain 2..n values and no other values.
// Then remove all other candidate values from the n cells.
