package sudoku

import (
	"fmt"
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
		b.cells[row] = make([]Cell, 9, 9)
		for col, cellString := range rowString {
			var val interface{}
			sval := fmt.Sprintf("%c", cellString)
			val, err := strconv.Atoi(sval)
			if err != nil {
				val = sval
			}
			b.cells[row][col] = CellFactory(b, col, row, val)
			// 	b.cells[row][col] = NewAllCandidatesCell(b, col, row)
			// } else {
			// 	b.cells[row][col] = NewFixedCell(b, col, row, val)
			// }
		}
	}
	return b, err
}

type cellFilter func(Cell) bool

func isFixedCell(cell Cell) bool {
	_, ok := cell.(FixedCell)
	return ok
}

func (b Board) FixedCells() (cells []Cell) {
	cells = b.SelectBoardCells(isFixedCell)
	return cells
}

func isCandidateCell(cell Cell) bool {
	_, ok := cell.(CandidateCell)
	return ok
}

func (b Board) CandidateCells() (cells []Cell) {
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

func (b Board) ColForCell(c Cell) (seq []Cell) {
	for i, _ := range b.cells {
		seq = append(seq[:], b.cells[i][c.X()])
	}
	return seq
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

func (b Board) ReduceCandidates() {
	for _, candidateCell := range b.CandidateCells() {
		seqCells := b.SequenceCellsForCell(candidateCell)
		fixedCells := SelectCells(seqCells, isFixedCell)
		fixedValues := MapCellValues(fixedCells)
		candidateCell.ReduceCandidates(fixedValues)
		/*
			For each candidate cell
				Get set of fixed values for each cell row, column, and group
				delete each fixed value from cell's candidates
				if len(cell.candidates) == 1 { convert cell to ValueCell, found = true }
			end
		*/
	}
}
