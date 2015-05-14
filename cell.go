package gosudoku

import "sort"

type Cell interface {
	X() int
	Y() int
	Value() int
	Equal(Cell) bool
}

type CandidateCell interface {
	Cell
	Candidates() []int
	ReduceCandidates([]int) bool
}

type baseCell struct {
	x, y int
}

type candidateCell struct {
	baseCell
	candidates map[int]bool
}

type valueCell struct {
	baseCell
	value int
}

type fixedCell struct {
	valueCell
}

func (b baseCell) X() int {
	return b.x
}

func (b baseCell) Y() int {
	return b.y
}

func (b baseCell) Value() int {
	return 0
}

func (b baseCell) Equal(c1 Cell) bool {
	return b.X() == c1.X() &&
		b.Y() == c1.Y()
}

func (c candidateCell) ReduceCandidates(vals []int) bool {
	var result = false
	for _, v := range vals {
		if c.candidates[v] {
			result = true
		}
		delete(c.candidates, v)
	}
	return result
}

func (c candidateCell) Candidates() (candidates []int) {
	for candidate := range c.candidates {
		candidates = append(candidates, candidate)
	}
	sort.Ints(candidates)
	return candidates
}

func (v valueCell) Value() int {
	return v.value
}

func CellFactory(x int, y int, val interface{}) (cell Cell) {
	switch val := val.(type) {
	case []int:
		cell = newCandidateCell(x, y, val)
	case int:
		if val >= 1 && val <= 9 {
			cell = newFixedCell(x, y, val)
		} else {
			cell = newAllCandidatesCell(x, y)
		}
	case string:
		cell = newAllCandidatesCell(x, y)
	}
	return cell
}

func newFixedCell(x int, y int, val int) (cell Cell) {
	return fixedCell{
		valueCell: valueCell{
			baseCell: baseCell{
				x: x,
				y: y,
			},
			value: val,
		},
	}
}

func newValueCell(x int, y int, val int) (cell Cell) {
	return valueCell{
		baseCell: baseCell{
			x: x,
			y: y,
		},
		value: val,
	}
}

func newCandidateCell(x int, y int, val []int) (cell Cell) {
	candidates := make(map[int]bool, len(val))
	for _, key := range val {
		candidates[key] = true
	}
	return candidateCell{
		baseCell: baseCell{
			x: x,
			y: y,
		},
		candidates: candidates,
	}
}

func newAllCandidatesCell(x int, y int) (cell Cell) {
	candidates := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
		7: true,
		8: true,
		9: true,
	}
	return candidateCell{
		baseCell: baseCell{
			x: x,
			y: y,
		},
		candidates: candidates,
	}
}

func ValueCellAt(cell CandidateCell) Cell {
	return valueCell{
		baseCell: baseCell{
			x: cell.X(),
			y: cell.Y(),
		},
		value: cell.Candidates()[0],
	}
}
