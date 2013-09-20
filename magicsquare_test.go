// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

package magicsquare

import (
	"github.com/bmizerany/assert"
	"testing"
)

func sumRow(row []int) int {
	var sum int
	for _, v := range row {
		sum += v
	}
	return sum
}

func sumColumn(s Square, column int) int {
	var sum int
	for _, row := range s {
		sum += row[column]
	}
	return sum
}

// TestSums ensures that all rows and columns sum up to the same value.
func TestSums(t *testing.T) {
	size := 5
	sq, err := MagicSquare(size)
	if err != nil {
		t.Fatal(err)
	}
	// First we figure out what row zero sums up to
	sum0 := sumRow(sq[0])
	// Check the remaining rows
	for r := 1; r < size; r++ {
		sum1 := sumRow(sq[r])
		assert.Equal(t, sum0, sum1)
	}
	// Check the columns
	for c := 0; c < size; c++ {
		sum1 := sumColumn(sq, c)
		assert.Equal(t, sum0, sum1)

	}
	// TODO: Test diagonal sums
}

// TestEvenSize ensures that an error is returned if we request a magic square
// of even size.
func TestEvenSize(t *testing.T) {
	_, err := MagicSquare(4)
	assert.NotEqual(t, nil, err)
}
