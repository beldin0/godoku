package main

import "testing"

func TestVariables(t *testing.T) {
	if col[0] != [9]int{0, 9, 18, 27, 36, 45, 54, 63, 72} {
		t.Fail()
	}
	if row[0] != [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
		t.Fail()
	}
}

func TestSimpleIsIn(t *testing.T) {
	testSudoku := simpleSudokuSolver{[]int{
		0, 0, 3, 0, 2, 0, 6, 0, 0,
		9, 0, 0, 3, 0, 5, 0, 0, 1,
		0, 0, 1, 8, 0, 6, 4, 0, 0,

		0, 0, 8, 1, 0, 2, 9, 0, 0,
		7, 0, 0, 0, 0, 0, 0, 0, 8,
		0, 0, 6, 7, 0, 8, 2, 0, 0,

		0, 0, 2, 6, 0, 9, 5, 0, 0,
		8, 0, 0, 2, 0, 3, 0, 0, 9,
		0, 0, 5, 0, 1, 0, 3, 0, 0,
	}}
	tests := []struct {
		Zone     [9]int
		Value    int
		Expected bool
	}{
		{
			Zone:     row[1],
			Value:    3,
			Expected: true,
		},
		{
			Zone:     col[3],
			Value:    3,
			Expected: true,
		},
		{
			Zone:     col[4],
			Value:    3,
			Expected: false,
		},
		{
			Zone:     square[6],
			Value:    3,
			Expected: false,
		},
		{
			Zone:     square[8],
			Value:    7,
			Expected: false,
		},
		{
			Zone:     square[8],
			Value:    3,
			Expected: true,
		},
	}
	for _, test := range tests {
		actual := testSudoku.isIn(test.Value, test.Zone)
		if test.Expected != actual {
			t.Fatalf("Checking for %v in %v\nExpected: %v\tActual: %v\n", test.Value, test.Zone, test.Expected, actual)
		}
	}
}

func TestGetIndex(t *testing.T) {
	tests := []struct {
		Index    int
		ZoneList [9][9]int
		Expected int
	}{
		{
			Index:    0,
			ZoneList: row,
			Expected: 0,
		},
		{
			Index:    80,
			ZoneList: row,
			Expected: 8,
		},
		{
			Index:    32,
			ZoneList: row,
			Expected: 3,
		},
		{
			Index:    0,
			ZoneList: col,
			Expected: 0,
		},
		{
			Index:    80,
			ZoneList: col,
			Expected: 8,
		},
		{
			Index:    32,
			ZoneList: col,
			Expected: 5,
		},
		{
			Index:    0,
			ZoneList: square,
			Expected: 0,
		},
		{
			Index:    80,
			ZoneList: square,
			Expected: 8,
		},
		{
			Index:    32,
			ZoneList: square,
			Expected: 4,
		},
		{
			Index:    -1,
			ZoneList: row,
			Expected: -1,
		},
		{
			Index:    81,
			ZoneList: row,
			Expected: -1,
		},
	}
	for _, test := range tests {
		actual := getIndex(test.Index, test.ZoneList)
		if test.Expected != actual {
			t.Fatalf("Checking for %v in %v\nExpected: %v\tActual: %v\n", test.Index, test.ZoneList, test.Expected, actual)
		}
	}
}
