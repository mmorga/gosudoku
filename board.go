package sudoku

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Board struct {
	cells [][]Cell
}

func (b Board) Cell(x, y int) (c Cell) {
	return b.cells[y][x]
}

func BoardFromArray(a [][]int) (b *Board, err error) {
	b = new(Board)
	for y, row := range a {
		for x, val := range row {
			b.cells[y][x] = CellFactory(b, x, y, val)
		}
	}
	return b, err
}

func BoardFromString(s string) (b *Board, err error) {
	b = new(Board)
	b.cells = make([][]Cell, 9, 9)
	splice := strings.Split(s, "\n")
	for row, rowString := range splice {
		if len(rowString) > 0 {
			b.cells[row] = make([]Cell, 9, 9)
			for col, cellString := range rowString {
				var val interface{}
				sval := fmt.Sprintf("%c", cellString)
				val, err := strconv.Atoi(sval)
				if err != nil {
					val = sval
				}
				b.cells[row][col] = CellFactory(b, col, row, val)
			}
		}
	}
	return b, err
}

func isCompleteSet(vals []int) bool {
	if len(vals) != 9 {
		return false
	}
	sort.Ints(vals)

	for i := 1; i < 10; i++ {
		if i != vals[i-1] {
			return false
		}
	}
	return true
}

func (b Board) IsSolved() bool {
	for _, seq := range b.Sequences() {
		vals := MapCellValues(SelectCells(seq, isValueOrfixedCell))
		if !isCompleteSet(vals) {
			return false
		}
	}
	return true
}

type cellFilter func(Cell) bool

func isfixedCell(cell Cell) bool {
	_, ok := cell.(fixedCell)
	return ok
}

func (b Board) fixedCells() (cells []Cell) {
	cells = b.SelectBoardCells(isfixedCell)
	return cells
}

func isValueCell(cell Cell) bool {
	_, ok := cell.(valueCell)
	return ok
}

func isValueOrfixedCell(cell Cell) bool {
	return isValueCell(cell) || isfixedCell(cell)
}

func (b Board) ValueAndfixedCells() (cells []Cell) {
	cells = b.SelectBoardCells(isValueOrfixedCell)
	return cells
}

func (b Board) valueCells() (cells []Cell) {
	cells = b.SelectBoardCells(isValueCell)
	return cells
}

func isCandidateCell(cell Cell) bool {
	_, ok := cell.(candidateCell)
	return ok
}

func (b Board) candidateCells() (cells []Cell) {
	cells = b.SelectBoardCells(isCandidateCell)
	return cells
}

func (b Board) AllCells() (cells []Cell) {
	return flatten(b.cells)
}

func (b Board) SelectBoardCells(cellMatches cellFilter) (cells []Cell) {
	return SelectCells(b.AllCells(), cellMatches)
}

func SelectCells(fromCells []Cell, cellMatches cellFilter) (cells []Cell) {
	for _, cell := range fromCells {
		if cellMatches(cell) {
			cells = append(cells, cell)
		}
	}
	return cells
}

func MapCellValues(fromCells []Cell) (cells []int) {
	for _, cell := range fromCells {
		cells = append(cells, cell.Value())
	}
	return cells
}

func (b Board) RowForCell(c Cell) []Cell {
	return b.cells[c.Y()][:]
}

func (b Board) column(colIdx int) (seq []Cell) {
	for _, row := range b.cells {
		seq = append(seq, row[colIdx])
	}
	return seq
}

func (b Board) ColForCell(c Cell) (seq []Cell) {
	return b.column(c.X())
}

func groupBounds(idx int) (xMin, xMax, yMin, yMax int) {
	y := idx / 3
	x := idx % 3

	return x * 3, x*3 + 2, y * 3, y*3 + 2
}

func groupIdxFor(x, y int) int {
	return (y / 3 * 3) + x/3
}

func (b Board) group(idx int) (seq []Cell) {
	// TODO: Better way to do this with slices?
	xMin, xMax, yMin, yMax := groupBounds(idx)
	for row := yMin; row <= yMax; row++ {
		for col := xMin; col <= xMax; col++ {
			seq = append(seq, b.Cell(col, row))
		}
	}
	return seq
}

func (b Board) GroupForCell(c Cell) (seq []Cell) {
	return b.group(groupIdxFor(c.X(), c.Y()))
}

func (b Board) Sequences() (seqs [][]Cell) {
	// Rows
	for _, r := range b.cells {
		seqs = append(seqs, r)
	}
	for colIdx := 0; colIdx < 9; colIdx++ {
		seqs = append(seqs, b.column(colIdx))
	}
	for groupIdx := 0; groupIdx < 9; groupIdx++ {
		seqs = append(seqs, b.group(groupIdx))
	}

	return seqs
}

func (b Board) SequencesForCell(c Cell) (seqs [][]Cell) {
	seqs = append(seqs, b.ColForCell(c))
	seqs = append(seqs, b.RowForCell(c))
	seqs = append(seqs, b.GroupForCell(c))
	return seqs
}

func (b Board) SequenceCellsForCell(c Cell) (seqs []Cell) {
	return flatten(b.SequencesForCell(c))
}

func flatten(nestedCells [][]Cell) (cells []Cell) {
	for _, row := range nestedCells {
		for _, cell := range row {
			cells = append(cells, cell)
		}
	}
	return cells
}

func (b Board) String() string {
	buf := bytes.NewBufferString("")
	buf.WriteString("\n")
	for rowIdx, row := range b.cells {
		if rowIdx == 3 || rowIdx == 6 {
			buf.WriteString(" ------+-------+------\n")
		}
		for colIdx, cell := range row {
			if colIdx == 3 || colIdx == 6 {
				buf.WriteString(" |")
			}
			cs := strconv.Itoa(cell.Value())
			if cs == "0" {
				cs = "."
			}
			buf.WriteString(" ")
			buf.WriteString(cs)
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func (b Board) ReduceCandidates() {
	for _, cell := range b.candidateCells() {
		seqCells := b.SequenceCellsForCell(cell)
		valueCells := SelectCells(seqCells, isValueOrfixedCell)
		values := MapCellValues(valueCells)
		if candidateCell, ok := cell.(CandidateCell); ok {
			candidateCell.ReduceCandidates(values)
		}
	}
}

func (b Board) ReplaceCellWithNewValueCell(oldCell Cell, val int) {
	x := oldCell.X()
	y := oldCell.Y()
	b.cells[y][x] = newValueCell(&b, x, y, val)
}

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

func hasVal(vals []int, val int) bool {
	for _, v := range vals {
		if v == val {
			return true
		}
	}
	return false
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
