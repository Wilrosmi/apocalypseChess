package main

func main() {
}

// Creates a new starting board state
func createNewBoard() [5][5]string {
	firstAndLastColumns := [5]string{"w1", "w0", "e", "b0", "b1"}
	middleColumns := [5]string{"w0", "e", "e", "e", "b0"}
	board := [5][5]string{firstAndLastColumns, middleColumns, middleColumns, middleColumns, firstAndLastColumns}
	return board
}

// Checks if a proposed move is valid or not
func checkValidMove(board [5][5]string, whosTurn string, currentSquare [2]int, newSquare [2]int) bool {
	if !canPlayerMoveThatPiece(board, whosTurn, currentSquare) {
		return false
	} else if currentSquare == newSquare {
		// Player needs to select a new square to move to
		return false
	} else if newSquare[0] < 0 || newSquare[0] > 4 || newSquare[1] < 0 || newSquare[1] > 4 {
		// Selected square off the board
		return false
	}
	var validPieceMove bool
	switch string(board[currentSquare[0]][currentSquare[1]][1]) {
	case "0":
		validPieceMove = checkPawnMove(board, whosTurn, currentSquare, newSquare)
	case "1":
		validPieceMove = checkKnightMove(board, whosTurn, currentSquare, newSquare)
	default:
		validPieceMove = false
	}
	if !validPieceMove {
		return false
	}
	return true
}

// Checks if there is a piece of the players colour on the square they have selected
func canPlayerMoveThatPiece(board [5][5]string, whosTurn string, currentSquare [2]int) bool {
	if string(board[currentSquare[0]][currentSquare[1]][0]) != whosTurn {
		return false
	} else {
		return true
	}
}

// Checks if the proposed move is possible for a pawn
func checkPawnMove(board [5][5]string, whosTurn string, currentSquare [2]int, newSquare [2]int) bool {
	pieceOnSquare := string(board[newSquare[0]][newSquare[1]][0])
	if newSquare[0] == currentSquare[0] && newSquare[1] == currentSquare[1]+1 && pieceOnSquare == "e" {
		return true
	} else if (newSquare[0] == currentSquare[0]+1 || newSquare[0] == currentSquare[0]-1) && newSquare[1] == currentSquare[1]+1 && (pieceOnSquare != "e" && pieceOnSquare != whosTurn) {
		return true
	} else {
		return false
	}
}

// Checks if the proposed move is possible for a knight
func checkKnightMove(board [5][5]string, whosTurn string, currentSquare [2]int, newSquare [2]int) bool {
	if string(board[newSquare[0]][newSquare[1]][0]) == whosTurn {
		return false
	}
	if (newSquare[0] == currentSquare[0]+2 || newSquare[0] == currentSquare[0]-2) && (newSquare[1] == currentSquare[1]+1 || newSquare[1] == currentSquare[1]-1) {
		return true
	} else if (newSquare[0] == currentSquare[0]+1 || newSquare[0] == currentSquare[0]-1) && (newSquare[1] == currentSquare[1]+2 || newSquare[1] == currentSquare[1]-2) {
		return true
	} else {
		return false
	}
}
