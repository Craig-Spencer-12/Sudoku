package main

import "fmt"

var GameBoard [9][9]int

func main() {
	initializeBoard()
	printBoard(GameBoard)
	fmt.Println(countZeros(&GameBoard))
	for i := 0; i < 100; i++ {
		for j := 1; j < 10; j++ {
			fillSlots(findPossible(j), j)
		}
	}

	printBoard(GameBoard)
	fmt.Println(countZeros(&GameBoard))
}

func initializeBoard() {
	for i := range GameBoard {
		for j := range GameBoard[i] {
			GameBoard[i][j] = 0
		}
	}

	GameBoard[0][1] = 1
	GameBoard[0][8] = 7
	GameBoard[1][0] = 8
	GameBoard[1][1] = 6
	GameBoard[1][3] = 4
	GameBoard[1][4] = 7
	GameBoard[1][7] = 1
	GameBoard[1][8] = 2
	GameBoard[2][0] = 7
	GameBoard[2][3] = 1
	GameBoard[2][4] = 3
	GameBoard[2][5] = 5
	GameBoard[2][6] = 8
	GameBoard[2][8] = 6
	GameBoard[3][2] = 8
	GameBoard[3][4] = 1
	GameBoard[3][8] = 5
	GameBoard[4][1] = 2
	GameBoard[4][4] = 4
	GameBoard[4][6] = 3
	GameBoard[4][7] = 8
	GameBoard[4][8] = 1
	GameBoard[5][3] = 9
	GameBoard[5][4] = 8
	GameBoard[5][6] = 2
	GameBoard[6][2] = 1
	GameBoard[7][0] = 4
	GameBoard[7][5] = 2
	GameBoard[7][6] = 1
	GameBoard[7][7] = 6
	GameBoard[8][0] = 6
	GameBoard[8][1] = 7
	GameBoard[8][2] = 9
	GameBoard[8][5] = 1
	GameBoard[8][6] = 4
	GameBoard[8][8] = 3

}

func printBoard(board [9][9]int) {
	for l := 0; l < 73; l++ {
		fmt.Print("-")
	}
	fmt.Print("\n")

	for i := range board {
		fmt.Print("|   ")
		for j := range board[i] {
			fmt.Printf("%v   |   ", board[i][j])
		}
		fmt.Print("\n")
		for l := 0; l < 73; l++ {
			fmt.Print("-")
		}
		fmt.Print("\n")
	}
}

func findPossible(n int) *[9][9]int {
	tempBoard := GameBoard

	for i := range tempBoard {
		for j := range tempBoard[i] {
			if tempBoard[i][j] == n {
				deleteRow(&tempBoard, i)
				deleteCol(&tempBoard, j)
				deleteBox(&tempBoard, (i/3)*3, (j/3)*3)
			}
		}
	}

	return &tempBoard
}

func deleteRow(board *[9][9]int, n int) {
	for i := 0; i < 9; i++ {
		board[n][i] = -1
	}
}

func deleteCol(board *[9][9]int, n int) {
	for i := 0; i < 9; i++ {
		board[i][n] = -1
	}
}

func deleteBox(board *[9][9]int, x int, y int) {
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			board[i][j] = -1
		}
	}
}

func fillSlots(board *[9][9]int, n int) {
	numZeros := 0
	locationX := 0
	locationY := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := i * 3; k < i*3+3; k++ {
				for l := j * 3; l < j*3+3; l++ {
					if board[k][l] == 0 {
						numZeros++
						locationX = k
						locationY = l
					}
				}
			}

			if numZeros == 1 {
				GameBoard[locationX][locationY] = n
			}

			numZeros = 0
		}
	}

	//printBoard(*board)
	//fmt.Println(n)
}

func countZeros(board *[9][9]int) int {
	count := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				count++
			}
		}
	}
	return count
}
