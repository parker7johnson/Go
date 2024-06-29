package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
	// "math/rand"
	// "rsc.io/quote"
)

type GameBoard struct {
	x, y  int
	board [][]int
}

const boardSize = 40

func main() {

	gameBoard := GameBoard{boardSize, boardSize, nil}
	gameBoard = generateRandState(gameBoard)
	for {
		render(gameBoard)
		// displayBoardState(gameBoard)
		fmt.Println("--------------------------")
		updateState(gameBoard)
		time.Sleep(1 * time.Second)
		clearOutput()
		// displayBoardState(gameBoard)
	}

}

func updateState(gb GameBoard) {
	prevState := gb.board
	for i := 0; i < gb.x; i++ {
		for j := 0; j < gb.y; j++ {
			gb.board[i][j] = calculateNextState(prevState, i, j)
		}
	}
}

func calculateNextState(board [][]int, x int, y int) int {
	// possible positions = board[x+1][y], board[x-1][y], board[x+1][y+1], board[x+1][y-1], board[x-1][y-1], board[x-1][y+1], board[x][y+1], board[x][y-1]
	nextState := 0

	for i := x - 1; i <= x+1; i++ {
		if i < 0 || i >= boardSize {
			continue
		}
		for j := y - 1; j <= y+1; j++ {
			if j < 0 || j >= boardSize || (j == y && i == x) {
				continue
			} else {
				nextState += board[i][j]
			}
		}
	}

	live := 0
	if nextState == 3 || nextState == 2 {
		live = 1
	}
	return live
}

func generateRandState(gb GameBoard) GameBoard {
	gb.board = deadCells(gb)
	for i := 0; i < gb.x; i++ {
		for j := 0; j < gb.y; j++ {
			if rand.Float32() < .5 {
				gb.board[i][j] = 1
			}
		}
	}
	return gb
}

func deadCells(gb GameBoard) [][]int {
	gb.board = make([][]int, gb.x)
	for i := 0; i < gb.x; i++ {
		gb.board[i] = make([]int, gb.y)
	}
	return gb.board
}

func render(gb GameBoard) {
	// fmt.Println(gb.x, gb.y)
	// fmt.Println(gb.board)
	for i := 0; i < gb.x; i++ {
		for j := 0; j < gb.y; j++ {
			if gb.board[i][j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func displayBoardState(gb GameBoard) {
	for i := 0; i < gb.x; i++ {
		for j := 0; j < gb.y; j++ {
			fmt.Print(gb.board[i][j])
		}
		fmt.Println()
	}
}

func clearOutput() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
