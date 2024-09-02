package modules

import "log"

func SudokuSolver(board [][]int) bool {
	for idx := 0; idx < 9; idx++ {
		for jdx := 0; jdx < 9; jdx++ {
			if board[idx][jdx] == 0 {
				for num := 1; num <= 9; num++ {
					if isValidSudoku(board, idx, jdx, num) {
						board[idx][jdx] = num
						if SudokuSolver(board) {
							// printSudoku(board)
							return true
						}
					}
					board[idx][jdx] = 0
				}
				return false
			}
		}
	}
	return true
}

func isValidSudoku(board [][]int, row, col, elem int) bool {
	for idx := 0; idx < 9; idx++ {
		if board[idx][col] == elem ||
			board[row][idx] == elem ||
			board[3*(row/3)+idx/3][3*(col/3)+idx%3] == elem {
			return false
		}
	}
	return true
}

func PrintSudoku(board [][]int) {
	for _, value := range board {
		log.Println(value)
	}
	log.Println()
}
