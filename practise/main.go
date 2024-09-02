package main

import (
	"log"

	"github.com/sandeepgoutele/golang-coding/practise/modules"
)

func main() {
	log.Print("Hello world!!")
	// modules.ReverseWord()
	// modules.Sort012()
	// modules.BQueTester()
	// modules.BQueTester2()
	board := [][]int{
		{9, 5, 7, 0, 1, 3, 0, 8, 4},
		{4, 8, 3, 0, 5, 7, 1, 0, 6},
		{0, 1, 2, 0, 4, 9, 5, 3, 7},
		{1, 7, 0, 3, 0, 4, 9, 0, 2},
		{5, 0, 4, 9, 7, 0, 3, 6, 0},
		{3, 0, 9, 5, 0, 8, 7, 0, 1},
		{8, 4, 5, 7, 9, 0, 6, 1, 3},
		{0, 9, 1, 0, 3, 6, 0, 7, 5},
		{7, 0, 6, 1, 8, 5, 4, 0, 9},
	}
	log.Println(modules.SudokuSolver(board))
	modules.PrintSudoku(board)

	board2 := [][]int{
		{0, 0, 0, 2, 4, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 6, 0},
		{3, 6, 0, 0, 0, 0, 5, 7, 4},
		{0, 0, 3, 0, 8, 0, 0, 1, 0},
		{5, 0, 4, 0, 0, 0, 0, 0, 8},
		{0, 0, 0, 7, 0, 0, 0, 0, 0},
		{0, 0, 0, 6, 0, 9, 0, 0, 0},
		{0, 0, 8, 0, 0, 0, 6, 0, 0},
		{0, 7, 0, 0, 0, 4, 0, 9, 2},
	}
	log.Println(modules.SudokuSolver(board2))
	modules.PrintSudoku(board2)

	board3 := [][]int{
		{0, 0, 0, 0, 9, 0, 0, 0, 4},
		{4, 6, 0, 0, 2, 0, 0, 8, 1},
		{0, 0, 1, 0, 0, 0, 5, 0, 7},
		{0, 0, 0, 0, 4, 1, 7, 0, 0},
		{2, 0, 0, 0, 0, 0, 0, 3, 8},
		{8, 0, 0, 0, 7, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 4, 6, 0, 9},
		{0, 5, 4, 0, 0, 8, 0, 0, 0},
	}
	log.Println(modules.SudokuSolver(board3))
	modules.PrintSudoku(board3)
}
