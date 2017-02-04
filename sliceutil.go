package sliceutil

import (
	"github.com/cheekybits/genny/generic"
)

//TODO: think about how structures with pointers in them can be handled

type SType generic.Type

//InsertSName inserts data into slice
func InsertSName(a []SType, i int, newdata ...SType) []SType {
	l := len(newdata)
	a = append(a, make([]SType, l)...)
	copy(a[i+l:], a[i:])
	copy(a[i:], newdata)
	return a
}

//DeleteSName deetes data from a slice and returns the new slice.
func DeleteSName(a []SType, i int, count int) []SType {
	copy(a[i:], a[i+count:])
	for j := i + count; j < len(a); j++ {
		//What do i do here?
	}
	return a[:len(a)-count]
}

//CutSName cuts data out of a slice then returns the new cut array and the data cut from it
func CutSName(a []SType, i int, count int) ([]SType, []SType) {
	b := make([]SType, count)
	copy(b, a[i:i+count])
	a = DeleteSName(a, i, count)
	return a, b
}
