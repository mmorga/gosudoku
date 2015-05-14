package gosudoku

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	cells []Unit
}

func (b Board) Cell(x, y int) (c Cell) {
	return b.cells[y][x]
}

func BoardFromArray(a [][]int) (b *Board, err error) {
	b = new(Board)
	for y, row := range a {
		for x, val := range row {
			b.cells[y][x] = CellFactory(x, y, val)
		}
	}
	return b, err
}

func BoardFromString(s string) (b *Board, err error) {
	b = new(Board)
	b.cells = make([]Unit, 9, 9)
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
				b.cells[row][col] = CellFactory(col, row, val)
			}
		}
	}
	return b, err
}

func (b Board) IsSolved() bool {
	for _, seq := range b.Units() {
		vals := MapCellValues(SelectCells(seq, isValueOrFixedCell))
		if !isCompleteSet(vals) {
			return false
		}
	}
	return true
}

// BEGIN REFACTOR AREA

type cellFilter func(Cell) bool

func isfixedCell(cell Cell) bool {
	_, ok := cell.(fixedCell)
	return ok
}

func isValueCell(cell Cell) bool {
	_, ok := cell.(valueCell)
	return ok
}

func isValueOrFixedCell(cell Cell) bool {
	return isValueCell(cell) || isfixedCell(cell)
}

func isCandidateCell(cell Cell) bool {
	_, ok := cell.(candidateCell)
	return ok
}

// END REFACTOR AREA

// TODO - get rid of this
func (b Board) candidateCells() (cells Unit) {
	cells = b.SelectBoardCells(isCandidateCell)
	return cells
}

// TODO - get rid of this
func (b Board) AllCells() (cells Unit) {
	return FlattenUnitSlice(b.cells)
}

// TODO - get rid of this
func (b Board) SelectBoardCells(cellMatches cellFilter) (cells Unit) {
	return SelectCells(b.AllCells(), cellMatches)
}

func SelectCells(fromCells Unit, cellMatches cellFilter) (cells Unit) {
	for _, cell := range fromCells {
		if cellMatches(cell) {
			cells = append(cells, cell)
		}
	}
	return cells
}

func MapCellValues(fromCells Unit) (cells []int) {
	for _, cell := range fromCells {
		cells = append(cells, cell.Value())
	}
	return cells
}

func (b Board) RowForCell(c Cell) Unit {
	return b.cells[c.Y()][:]
}

func (b Board) column(colIdx int) (seq Unit) {
	for _, row := range b.cells {
		seq = append(seq, row[colIdx])
	}
	return seq
}

func (b Board) ColForCell(c Cell) (seq Unit) {
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

func (b Board) group(idx int) (seq Unit) {
	// TODO: Better way to do this with slices?
	xMin, xMax, yMin, yMax := groupBounds(idx)
	for row := yMin; row <= yMax; row++ {
		for col := xMin; col <= xMax; col++ {
			seq = append(seq, b.Cell(col, row))
		}
	}
	return seq
}

func (b Board) GroupForCell(c Cell) (seq Unit) {
	return b.group(groupIdxFor(c.X(), c.Y()))
}

func (b Board) Units() (seqs []Unit) {
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

func (b Board) UnitsForCell(c Cell) (seqs []Unit) {
	seqs = append(seqs, b.ColForCell(c))
	seqs = append(seqs, b.RowForCell(c))
	seqs = append(seqs, b.GroupForCell(c))
	return seqs
}

func (b Board) UnitCellsForCell(c Cell) (seqs Unit) {
	return FlattenUnitSlice(b.UnitsForCell(c))
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

func (b Board) ReplaceCellWithNewValueCell(oldCell Cell, val int) {
	x := oldCell.X()
	y := oldCell.Y()
	b.cells[y][x] = newValueCell(x, y, val)
}
