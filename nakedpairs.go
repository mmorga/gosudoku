package gosudoku

// This one is a unit oriented strategy

// * unit: row, col, box

// For each unique pair of candidates in a unit
// 	If unique pair exists exclusively twice in a unit
// 		Found!
// 		Remove pair values from other cells in unit

func (b Board) NakedPairs() bool {
	foundAny := false
	found := true
	var allUpdatedCells Unit

	for found {
		b.ReduceCandidates()
		found = false
		for _, unit := range b.Units() {
			for _, pair := range UniquePairsInUnit(unit) {
				nakedCells := CandidateCellsWithValues(unit, pair)
				if len(nakedCells) == 2 {
					updatedCells := RemoveCandidatesFromUnitExceptForCells(pair, unit, nakedCells)
					// for cell := range updatedCells {
					allUpdatedCells = append(allUpdatedCells, updatedCells...)
					// }
					if len(updatedCells) > 0 {
						found = true
						foundAny = true
					}
				}
			}
		}
	}
	return foundAny
}

func NakedPairsUnit(unit Unit) (updatedCells Unit) {
	for _, pair := range UniquePairsInUnit(unit) {
		nakedCells := CandidateCellsWithValues(unit, pair)
		if len(nakedCells) == 2 {
			updatedCells = append(updatedCells,
				RemoveCandidatesFromUnitExceptForCells(pair, unit, nakedCells)...)
		}
	}
	return updatedCells
}
