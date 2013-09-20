package main

import (
	"github.com/kr/pretty"
)

type square [][]int

type pition struct {
	row int
	col int
}

const size = 3

func main() {
	// Initialize the square
	square := make([][]int, size)
	for i := 0; i < size; i++ {
		square[i] = make([]int, size)
	}
	// Start rightmost column, middle row
	col := size - 1
	row := (size / 2)
	p := pition{row, col}
	for i := 1; true; i++ {
		square[p.row][p.col] = i
		//
		// Find next position
		//
		// Up and to the right
		n := p
		if n.col < size-1 {
			n.col++
		} else {
			n.col = 0
		}
		if n.row > 0 {
			n.row--
		} else {
			n.row = size - 1
		}
		if square[n.row][n.col] == 0 {
			p = n
			continue
		}
		// Left in the same row
		n = p
		if n.col > 0 {
			n.col--
		} else {
			n.col = size - 1
		}
		if square[n.row][n.col] == 0 {
			p = n
			continue
		}
		// We're done
		break

	}
	pretty.Println(square)
}
