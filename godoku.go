package main

import "fmt"

// Sudoku Solver Exercise
// ======================
//
// The aim is to complete the solve() function so that it returns a completed
// sudoku puzzle. We've setup the execise to use an array to represent the
// state of a puzzle but if you think there is a more suitable data structure
// then feel free to change it.
//
// Please do not spend more than two hours on this, this can be a difficult
// exercise so it's not expected that you complete it within the given time. We
// want to see how you approach coding problems, so please show your working as
// best you can - bonus points if you use git to record your working history!

// row, col, square are arrays of array of integers representing the indexes in the puzzle slice that correspond to that 'zone'
var (
	row, col = func() (row [9][9]int, col [9][9]int) {
		v := 0
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				row[i][j] = v
				col[j][i] = v
				v++
			}
		}
		return
	}()
	square = [9][9]int{
		{
			0, 1, 2,
			9, 10, 11,
			18, 19, 20,
		},
		{
			3, 4, 5,
			12, 13, 14,
			21, 22, 23,
		},
		{
			6, 7, 8,
			15, 16, 17,
			24, 25, 26,
		},
		{
			27, 28, 29,
			36, 37, 38,
			45, 46, 47,
		},
		{
			31, 31, 32,
			39, 40, 41,
			48, 49, 50,
		},
		{
			33, 34, 35,
			42, 43, 44,
			51, 52, 53,
		},
		{
			54, 55, 56,
			63, 64, 65,
			72, 73, 74,
		},
		{
			57, 58, 59,
			66, 67, 68,
			75, 76, 77,
		},
		{
			60, 61, 62,
			69, 70, 71,
			78, 79, 80,
		},
	}
)

// SudokuSolver is the interface for an implementation of a sudoku solving algorithm
type SudokuSolver interface {
	Solve() []int
}

type simpleSudokuSolver struct {
	puzzle []int
}

func (s simpleSudokuSolver) Solve() []int {
	// TODO
	return nil
}

func solve(puzzle []int) []int {
	s := simpleSudokuSolver{puzzle}
	return s.Solve()
}

func compare(puzzle []int, completePuzzle []int) bool {
	for i, n := range puzzle {
		if completePuzzle[i] != n {
			return false
		}
	}
	return true
}

func (s simpleSudokuSolver) isIn(val int, zone [9]int) bool {
	for _, n := range zone {
		if s.puzzle[n] == val {
			return true
		}
	}
	return false
}

func getIndex(index int, zoneList [9][9]int) int {
	for i, zone := range zoneList {
		for _, j := range zone {
			if index == j {
				return i
			}
		}
	}
	return -1
}

func (s simpleSudokuSolver) getPossibles(index int) (possibles []int) {
	if s.puzzle[index] != 0 {
		return []int{s.puzzle[index]}
	}
	rowIndex := getIndex(index, row)
	colIndex := getIndex(index, col)
	sqrIndex := getIndex(index, square)
	for i := 1; i < 10; i++ {
		if !s.isIn(i, row[rowIndex]) && !s.isIn(i, col[colIndex]) && !s.isIn(i, square[sqrIndex]) {
			possibles = append(possibles, i)
		}
	}
	return
}

var easySudoku = []int{
	0, 0, 3, 0, 2, 0, 6, 0, 0,
	9, 0, 0, 3, 0, 5, 0, 0, 1,
	0, 0, 1, 8, 0, 6, 4, 0, 0,

	0, 0, 8, 1, 0, 2, 9, 0, 0,
	7, 0, 0, 0, 0, 0, 0, 0, 8,
	0, 0, 6, 7, 0, 8, 2, 0, 0,

	0, 0, 2, 6, 0, 9, 5, 0, 0,
	8, 0, 0, 2, 0, 3, 0, 0, 9,
	0, 0, 5, 0, 1, 0, 3, 0, 0,
}

var hardSudoku = []int{
	6, 0, 0, 0, 0, 0, 1, 5, 0,
	9, 5, 4, 7, 1, 0, 0, 8, 0,
	0, 0, 0, 5, 0, 2, 6, 0, 0,

	8, 0, 0, 0, 9, 4, 0, 0, 6,
	0, 0, 3, 8, 0, 5, 4, 0, 0,
	4, 0, 0, 3, 7, 0, 0, 0, 8,

	0, 0, 6, 9, 0, 3, 0, 0, 0,
	0, 2, 0, 0, 4, 7, 8, 9, 3,
	0, 4, 9, 0, 0, 0, 0, 0, 5,
}

func main() {
	var easyComplete = compare(solve(easySudoku), []int{
		4, 8, 3, 9, 2, 1, 6, 5, 7,
		9, 6, 7, 3, 4, 5, 8, 2, 1,
		2, 5, 1, 8, 7, 6, 4, 9, 3,
		5, 4, 8, 1, 3, 2, 9, 7, 6,
		7, 2, 9, 5, 6, 4, 1, 3, 8,
		1, 3, 6, 7, 9, 8, 2, 4, 5,
		3, 7, 2, 6, 8, 9, 5, 1, 4,
		8, 1, 4, 2, 5, 3, 7, 6, 9,
		6, 9, 5, 4, 1, 7, 3, 8, 2,
	})

	var hardComplete = compare(solve(hardSudoku), []int{
		6, 3, 2, 4, 8, 9, 1, 5, 7,
		9, 5, 4, 7, 1, 6, 3, 8, 2,
		1, 7, 8, 5, 3, 2, 6, 4, 9,

		8, 1, 7, 2, 9, 4, 5, 3, 6,
		2, 9, 3, 8, 6, 5, 4, 7, 1,
		4, 6, 5, 3, 7, 1, 9, 2, 8,

		7, 8, 6, 9, 5, 3, 2, 1, 4,
		5, 2, 1, 6, 4, 7, 8, 9, 3,
		3, 4, 9, 1, 2, 8, 7, 6, 5,
	})

	fmt.Printf("Easy Sudoku Solved?: %v\n", easyComplete)
	fmt.Printf("Hard Sudoku Solved?: %v\n", hardComplete)
}
