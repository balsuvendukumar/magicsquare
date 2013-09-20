// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

// Package magicsquare generates magic squares of odd size.
package magicsquare

import "errors"

type square [][]int

type position struct {
	row int
	col int
}

// MagicSquare generates a magic square of a given size.
func MagicSquare(size int) (square, error) {
	if size%2 == 0 {
		return nil, errors.New("Size must be an odd number")
	}
	// Initialize the square
	s := make(square, size)
	for i := 0; i < size; i++ {
		s[i] = make([]int, size)
	}
	// Start rightmost column, middle row
	p := position{size / 2, size - 1}
	for i := 1; true; i++ {
		s[p.row][p.col] = i
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
		if s[n.row][n.col] == 0 {
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
		if s[n.row][n.col] == 0 {
			p = n
			continue
		}
		// We're done
		break

	}
	return s, nil
}
