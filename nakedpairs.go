package sudoku

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
				nakedCells := CandidateCellsWithValues(unit, pair)
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
