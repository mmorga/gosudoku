package sudoku

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func ReadFile(filename string) (puzzleSources map[string](string)) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("I regret nothing!")
	}

	buf := bytes.NewBuffer(fileBytes)
	var line string
	fmt.Sscanln(buf.String(), line)
	fmt.Printf("Scanned Line:\n%s\n", line)
	fmt.Printf("Whole Buf:\n%s\n", buf.String())

	return puzzleSources
}
