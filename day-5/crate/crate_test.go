package crate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMove(t *testing.T) {
	st := Stacks{
		[]Crate{
			"A",
			"B",
		},
		[]Crate{
			"C",
		},
		[]Crate{},
	}
	st = st.Move(1, 0, 2) // B -> 2
	st = st.Move(1, 0, 2) // A -> 2
	st = st.Move(1, 1, 2) // C -> 2
	st = st.Move(3, 2, 1) // C, A, B -> 1

	assert.Empty(t, st[0])
	assert.Empty(t, st[2])
	assert.Equal(t, Crate("C"), st[1][0])
	assert.Equal(t, Crate("A"), st[1][1])
	assert.Equal(t, Crate("B"), st[1][2])
}

func TestMoveMultiple(t *testing.T) {
	st := Stacks{
		[]Crate{
			"A",
			"B",
		},
		[]Crate{},
		[]Crate{"C"},
	}
	st = st.MoveMultiple(2, 0, 1) // A, B -> 1
	st = st.MoveMultiple(1, 2, 1) // C -> 1
	st = st.MoveMultiple(3, 1, 0)

	assert.Empty(t, st[1])
	assert.Empty(t, st[2])
	assert.Equal(t, Crate("A"), st[0][0])
	assert.Equal(t, Crate("B"), st[0][1])
	assert.Equal(t, Crate("C"), st[0][2])
}
