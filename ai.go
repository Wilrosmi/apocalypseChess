package main

import (
	"fmt"
	"math/rand"
)

// Picks a move
func aiMove(board [5][5]string, aiColour string) ([2]int, [2]int) {
	possibleMoves := moveFinder(board, aiColour)
	moveNumber := rand.Intn(len(possibleMoves))
	fmt.Println(possibleMoves)
	return possibleMoves[moveNumber][0], possibleMoves[moveNumber][1]
}

// Finds all the legal moves for the ai
func moveFinder(board [5][5]string, aiColour string) [][2][2]int {
	var moves [][2][2]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			legalMoves := getMovesForSquare(board, aiColour, i, j)
			moves = appendMoves(moves, legalMoves)
		}
	}
	return moves
}

// Gets the legal moves for a player for a given square
func getMovesForSquare(board [5][5]string, aiColour string, x int, y int) [][2][2]int {
	moves := make([][2][2]int, 0)
	if string(board[x][y][0]) != aiColour {
		return moves
	} else if string(board[x][y][1]) == "0" {
		legalMoves := getPawnMoves(board, aiColour, x, y)
		moves = appendMoves(moves, legalMoves)
	} else {
		legalMoves := getKnightMoves(board, aiColour, x, y)
		moves = appendMoves(moves, legalMoves)
	}
	return moves
}

// Finds the moves for a pawn square
func getPawnMoves(board [5][5]string, aiColour string, x int, y int) [][2][2]int {
	moves := make([][2][2]int, 0)
	forwardOrBack := 1
	if aiColour == "b" {
		forwardOrBack *= -1
	}
	if board[x][y+forwardOrBack] == "ee" {
		moves = append(moves, [2][2]int{{x, y}, {x, y + forwardOrBack}})
	}
	if x > 0 && string(board[x-1][y+forwardOrBack][0]) != aiColour && board[x-1][y+forwardOrBack] != "ee" {
		moves = append(moves, [2][2]int{{x, y}, {x - 1, y + forwardOrBack}})
	} else if x < 4 && string(board[x+1][y+forwardOrBack][0]) != aiColour && board[x+1][y+forwardOrBack] != "ee" {
		moves = append(moves, [2][2]int{{x, y}, {x + 1, y + forwardOrBack}})
	}
	return moves
}

// Finds the moves for a knight square
func getKnightMoves(board [5][5]string, aiColour string, x int, y int) [][2][2]int {
	moves := make([][2][2]int, 0)
	if x+2 <= 4 && y+1 <= 4 && string(board[x+2][y+1][0]) != aiColour {
		moves = append(moves, [2][2]int{{x, y}, {x + 2, y + 1}})
	}
	if x+2 <= 4 && y-1 >= 0 && board[x+2][y-1] != aiColour {
		moves = append(moves, [2][2]int{{x, y}, {x + 2, y - 1}})
	}
	if x-2 >= 0 && y+1 <= 4 && string(board[x-2][y+1][0]) != aiColour {
		moves = append(moves, [2][2]int{{x, y}, {x - 2, y + 1}})
	}
	if x-2 >= 0 && y-1 >= 0 && string(board[x-2][y-1][0]) != aiColour {
		moves = append(moves, [2][2]int{{x, y}, {x - 2, y - 1}})
	}
	if x+1 <= 4 && y+2 <= 4 && string(board[x+1][y+2][0]) != aiColour {
		moves = append(moves, [2][2]int{{x, y}, {x + 1, y + 2}})
	}
	if x+1 <= 4 && y-2 >= 0 && string(board[x+1][y-2][0]) != aiColour {
		moves = append(moves, [2][2]int{{x, y}, {x + 1, y - 2}})
	}
	if x-1 >= 0 && y-2 >= 0 && string(board[x-1][y-2][0]) != aiColour {
		moves = append(moves, [2][2]int{{x, y}, {x - 1, y - 2}})
	}
	if x-1 >= 0 && y+2 <= 4 && string(board[x-1][y+2][0]) != aiColour {
		moves = append(moves, [2][2]int{{x, y}, {x - 1, y + 2}})
	}
	return moves
}

// Appends a list of possible moves for one square to the overall list of moves
func appendMoves(moves [][2][2]int, legalMoves [][2][2]int) [][2][2]int {
	for _, x := range legalMoves {
		moves = append(moves, x)
	}
	return moves
}
