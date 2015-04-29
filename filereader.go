package sudoku

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func readPuzzleTitle(file *os.File) (string, error) {
	var grid, version string

	n, err := fmt.Fscanf(file, "%s %s\n", &grid, &version)
	if n != 2 || err != nil {
		// We're most likely at EOF
		return "", err
	}
	return fmt.Sprintf("%s %s", grid, version), err
}

func readPuzzle(file *os.File) (string, error) {
	var b bytes.Buffer
	var line string

	for i := 0; i < 9; i++ {
		n, err := fmt.Fscanln(file, &line)
		if n != 1 || err != nil {
			log.Fatalf("Unable to read Puzzle Line: %v", err)
			return b.String(), err
		}
		b.WriteString(line)
		b.WriteString("\n")
	}
	return b.String(), nil
}

func ReadFile(filename string) map[string](string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		panic("I regret nothing!")
	}

	puzzleSources := make(map[string](string), 0)
	done := false
	for !done {
		title, err := readPuzzleTitle(file)
		if err != nil {
			return puzzleSources
		}

		puzzle, err := readPuzzle(file)
		if err != nil {
			return puzzleSources
		}

		puzzleSources[title] = puzzle
	}

	return puzzleSources
}
