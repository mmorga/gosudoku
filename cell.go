package sudoku

import "sort"

type Cell interface {
	X() int
	Y() int
	Value() int
	Candidates() []int
	ReduceCandidates([]int)
}

type BaseCell struct {
	board *Board
	x, y  int
}

type CandidateCell struct {
	BaseCell
	candidates map[int]bool
}

type ValueCell struct {
	BaseCell
	value int
}

type FixedCell struct {
	ValueCell
}

func (b BaseCell) X() int {
	return b.x
}

func (b BaseCell) Y() int {
	return b.y
}

func (b BaseCell) Value() int {
	// TODO: What's the right thing to do here?
	return 0
}

func (BaseCell) Candidates() []int {
	return make([]int, 0, 0)
}

func (BaseCell) ReduceCandidates([]int) {
	panic("Invalid to Reduce Candidates on me")
}

func (c CandidateCell) ReduceCandidates(vals []int) {
	for _, v := range vals {
		delete(c.candidates, v)
	}
}

func (c CandidateCell) Candidates() (candidates []int) {
	for k, _ := range c.candidates {
		candidates = append(candidates, k)
	}
	sort.Ints(candidates)
	return candidates
}

func (v ValueCell) Value() int {
	return v.value
}

// func (b *BaseCell) RowForCell() []*Cell {
// 	return b.board.RowForCell(b)
// }

func CellFactory(b *Board, x int, y int, val interface{}) (cell Cell) {
	switch val := val.(type) {
	case []int:
		cell = NewCandidateCell(b, x, y, val)
	case int:
		cell = NewFixedCell(b, x, y, val)
	case string:
		cell = NewAllCandidatesCell(b, x, y)
	}
	return cell
}

func NewFixedCell(b *Board, x int, y int, val int) (cell Cell) {
	return FixedCell{
		ValueCell: ValueCell{
			BaseCell: BaseCell{
				board: b,
				x:     x,
				y:     y,
			},
			value: val,
		},
	}
}

func NewValueCell(b *Board, x int, y int, val int) (cell Cell) {
	return ValueCell{
		BaseCell: BaseCell{
			board: b,
			x:     x,
			y:     y,
		},
		value: val,
	}
	// c := new(ValueCell)
	// c.board = b
	// c.x = x
	// c.y = y
	// c.value = val
	// if cell, ok := (*c).(*Cell); ok {
	// 	return cell
	// }
	// return nil
}

func NewCandidateCell(b *Board, x int, y int, val []int) (cell Cell) {
	candidates := make(map[int]bool, len(val))
	for _, key := range val {
		candidates[key] = true
	}
	return CandidateCell{
		BaseCell: BaseCell{
			board: b,
			x:     x,
			y:     y,
		},
		candidates: candidates,
	}
	// c := new(CandidateCell)
	// c.board = b
	// c.x = x
	// c.y = y
	// c.candidates = make(map[int]bool, len(val))
	// for _, key := range val {
	// 	c.candidates[key] = true
	// }
	// return c
}

func NewAllCandidatesCell(b *Board, x int, y int) (cell Cell) {
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
	return CandidateCell{
		BaseCell: BaseCell{
			board: b,
			x:     x,
			y:     y,
		},
		candidates: candidates,
	}
	// c := new(CandidateCell)
	// c.board = b
	// c.x = x
	// c.y = y
	// c.candidates = map[int]bool{
	// 	1: true,
	// 	2: true,
	// 	3: true,
	// 	4: true,
	// 	5: true,
	// 	6: true,
	// 	7: true,
	// 	8: true,
	// 	9: true,
	// }

	// return c
}
