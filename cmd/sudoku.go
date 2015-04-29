package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mmorga/sudoku"
)

func main() {
	app := cli.NewApp()
	app.Name = "sudoku"
	app.Author = "Mark Morga"
	app.Email = "https://github.com/mmorga/gosudoku"
	app.Version = sudoku.Version()
	app.Usage = "Command Line Utility for solving Sudoku puzzles"
	app.Action = func(c *cli.Context) {
		sudoku.Solve(c.Args()[0])
	}
	app.Run(os.Args)
}
