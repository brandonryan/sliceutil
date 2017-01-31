package sliceutil

import "testing"

func TestInsert(t *testing.T) {
	tslice := []int{0, 1, 2, 3}
	tslice = InsertInt(tslice, 0, -3, -2, -1)
	if len(tslice) != 7 {
		t.Fail()
	}
	tslice = InsertInt(tslice, 7, 4, 5, 6)
	if len(tslice) != 10 {
		t.Fail()
	}
	tslice = InsertInt(tslice, 3, -999)
	if len(tslice) != 11 {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	tslice := []int{0, 1, 2, 3}
	tslice = DeleteInt(tslice, 0, 1)
	if len(tslice) != 3 {
		t.Fail()
	}
}

func TestCut(t *testing.T) {
	tslice := []int{0, 1, 2, 3}
	tslice, rslice := CutInt(tslice, 1, 2)
	if len(tslice) != 2 {
		t.Fail()
	}
	if len(rslice) != 2 {
		t.Fail()
	}
}
