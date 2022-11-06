package main

// Picks a move
func aiMove(board [5][5]string, aiColour string) ([2]int, [2]int) {
	possibleMoves := moveFinder(board, aiColour)
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
func findPawnMoves(board [5][5]string, aiColour string, x int, y int) [][2][2]int {
	moves := make([][2][2]int, 0)
	if board[x][y+1] == "ee" {
		moves = append(moves, [2][2]int{[2]int{x, y}, [2]int{x, y + 1}})
	}
	if x > 0 && string(board[x-1][y+1][0]) != aiColour && string(board[x-1][y+1][0]) != "ee" {
		moves = append(moves, [2][2]int{[2]int{x, y}, [2]int{x - 1, y + 1}})
	} else if x < 4 && string(board[x+1][y+1][0]) != aiColour && string(board[x+1][y+1][0]) != "ee" {
		moves = append(moves, [2][2]int{[2]int{x, y}, [2]int{x + 1, y + 1}})
	}
	return moves
}

func getKnightMoves(board [5][5]string, aiColour string, x int, y int) [][2][2]int {
	moves := make([][2][2]int, 0)

}

// Appends a list of possible moves for one square to the overall list of moves
func appendMoves(moves [][2][2]int, legalMoves [][2][2]int) [][2][2]int {
	for _, x := range legalMoves {
		moves = append(moves, x)
	}
	return moves
}
