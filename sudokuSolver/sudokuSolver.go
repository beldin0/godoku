package sudoku

import "log"

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
			30, 31, 32,
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

// NewSolver returns a sudoku Solver for the submitted puzzle
func NewSolver(puzzle []int) (solver Solver) {
	return &simpleSudokuSolver{puzzle}
}

// Solver is the interface for an implementation of a sudoku solving algorithm
type Solver interface {
	Solve() []int
}

type simpleSudokuSolver struct {
	puzzle []int
}

func (s *simpleSudokuSolver) Solve() []int {
	zeros := func() (zeros int) {
		for i := 0; i < 81; i++ {
			if s.puzzle[i] == 0 {
				zeros++
			}
		}
		return zeros
	}()
	for zeros > 0 {
		solvedAtLeastOne := false
		for i := 0; i < 81; i++ {
			if s.puzzle[i] != 0 {
				continue
			}
			possibles := s.getPossibles(i)
			if len(possibles) == 1 {
				s.puzzle[i] = possibles[0]
				solvedAtLeastOne = true
				zeros--
			}
		}
		if !solvedAtLeastOne {
			log.Fatalf("Unable to solve this puzzle! Still %v spaces left to fill.\n", zeros)
		}
	}
	return s.puzzle
}

func (s *simpleSudokuSolver) isIn(val int, zone [9]int) bool {
	for _, n := range zone {
		if s.puzzle[n] == val {
			return true
		}
	}
	return false
}

func (s *simpleSudokuSolver) getPossibles(index int) (possibles []int) {
	if s.puzzle[index] != 0 {
		return []int{s.puzzle[index]}
	}
	if index >= len(s.puzzle) {
		log.Fatalf("Unable to calculate zone indeces for index %v\n", index)
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

func getIndex(index int, zoneList [9][9]int) int {
	for i, zone := range zoneList {
		for _, j := range zone {
			if index == j {
				return i
			}
		}
	}
	log.Printf("Failed to find %v in %v\n", index, zoneList)
	return -1
}
