package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows, cols, boxes := make([][]int, 9), make([][]int, 9), make([][]int, 9)
	for k := 0; k < 9; k++ {
		rows[k] = make([]int, 9)
		cols[k] = make([]int, 9)
		boxes[k] = make([]int, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j] - '1'
			if val >= 0 && val < 9 {
				r, c := i/3, j/3
				b := r*3 + c
				if rows[i][val] > 0 || cols[j][val] > 0 || boxes[b][val] > 0 {
					return false
				} else {
					rows[i][val] = 1
					cols[j][val] = 1
					boxes[b][val] = 1
				}
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '3', '.', '.', '5', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '8', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '1', '1', '6', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '1', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '7'},
		{'.', '.', '.', '.', '.', '.', '.', '4', '.'},
	}

	fmt.Println(isValidSudoku(board))
}
