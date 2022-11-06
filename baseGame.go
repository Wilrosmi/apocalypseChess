package main

import (
	"strings", 
	"fmt"
)

func main() {
	moveCounter := 0
	gameOverCheck := "x"
	board := createNewBoard()
	for moveCounter < 100 && gameOverCheck == "x" {
		printBoard(board)
		var whiteRawOld string
		var whiteRawNew string
		var blackRawOld string
		var blackRawNew string
		var whiteOldMove [2]int
		var whiteNewMove [2]int
		var blackOldMove [2]int
		var blackNewMove [2]int
		fmt.Println("Enter the square for white to move from: ")
		fmt.Scanln(&whiteRawOld)
		fmt.Println("Enter the square for white to move to: ")
		fmt.Scanln(&whiteRawNew)
		fmt.Println("Enter the square for black to move from: ")
		fmt.Scanln(&blackRawOld)
		fmt.Println("Enter the squaere for black to move to: ")
		fmt.Scanln(&blackRawNew)
		if validateUserInput(whiteRawNew) && validateUserInput(whiteRawOld) {
			whiteNewMove = cleanUserInput(whiteRawNew)
			whiteOldMove = cleanUserInput(whiteRawOld)
		} else {
			gameOverCheck := "b"
			break
		}
		if validateUserInput(blackRawNew) && validateUserInput(blackRawOld) {
			blackNewMove = cleanUserInput(blackRawNew)
			blackOldMove = cleanUserInput(blackRawOld)
		} else {
			gameOverCheck := "w"
			break
		} 
		if !checkValidMove(board, "w", whiteOldMove, whiteNewMove) {
			gameOverCheck := "b"
			break
		} else if !checkValidMove(board, "b", blackOldMove, blackNewMove) {
			gameOverCheck := "w"
			break
		}
		board = resolveMoves(board, whiteOldMove, whiteNewMove, blackOldMove, blackNewMove)
		gameOverCheck = gameOverCheck(board)
		moveCounter++
	} 
}

// Creates a new starting board state
func createNewBoard() [5][5]string {
	firstAndLastColumns := [5]string{"w1", "w0", "e", "b0", "b1"}
	middleColumns := [5]string{"w0", "e", "e", "e", "b0"}
	board := [5][5]string{firstAndLastColumns, middleColumns, middleColumns, middleColumns, firstAndLastColumns}
	return board
}

// Prints the board to the terminal
func printBoard(board [5][5]string) {
	for i : 0; i < 5; i++ {
		fmt.Println(board[0][i] + " " + boaboard[1][i] + " " + board[2][i] + " " + board[3][i] + " " + board[4][i])
	}
}

// Checks if the users input is in the right format
func validateUserInput(input string) bool {
	validNums := "01234"
	if len(input) != 2 {
		return false
	} else if !strings.Contains(validNums, string(input[0])) || !strings.Contains(validNums, string(input[1])) {
		return false
	}
	return true
}

// Changes the users input from its raw form to the int array form we need it in
func cleanUserInput(input string) [2]int {
	return [2]int{int(input[0]), int(input[1])}
}

// Checks if a proposed move is valid or not
func checkValidMove(board [5][5]string, whosTurn string, currentSquare [2]int, newSquare [2]int) bool {
	if !canPlayerMoveThatPiece(board, whosTurn, currentSquare) {
		return false
	} else if currentSquare == newSquare {
		// Player needs to select a new square to move to
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

// Creates the new board given a valid move for each player
func resolveMoves(board [5][5]string, wOldSquare [2]int, wNewSquare [2]int, bOldSquare [2]int, bNewSquare [2]int) [5][5]string {
	var newBoard [5][5]string
	if wNewSquare[0] == bNewSquare[0] && wNewSquare[1] == wNewSquare[1] {
		newBoard = moveToSameSquare(board, wOldSquare, wNewSquare, bOldSquare, bNewSquare)
	} else {
		newBoard = moveToDifferentSquares(board, wOldSquare, wNewSquare, bOldSquare, bNewSquare)
	}
	return newBoard
}

// Creates the new board state when both players have selected to move to the same square
func moveToSameSquare(board [5][5]string, wOldSquare [2]int, wNewSquare [2]int, bOldSquare [2]int, bNewSquare [2]int) [5][5]string {
	newBoard := board
	whitePiece := string(board[wOldSquare[0]][wOldSquare[1]][1])
	blackPiece := string(board[bOldSquare[0]][bOldSquare[1]][1])
	if whitePiece == blackPiece {
		// Given we know both players can move to this new square we know it must be empty and dont have to change it
	} else if whitePiece == "0" {
		// Implies black piece is 1 as they cant be the same at this point so black wins the fight
		newBoard[wNewSquare[0]][wNewSquare[1]] = "b1"
	} else {
		// Only option left is inverse of above
		newBoard[wNewSquare[0]][wNewSquare[1]] = "w1"
	}
	// Pieces are gone from their original squares
	newBoard[wOldSquare[0]][wOldSquare[1]] = "e"
	newBoard[bOldSquare[0]][bOldSquare[1]] = "e"
	return newBoard
}

// Creates the new board state when the players have selected to move to different squares
func moveToDifferentSquares(board [5][5]string, wOldSquare [2]int, wNewSquare [2]int, bOldSquare [2]int, bNewSquare [2]int) [5][5]string {
	newBoard := board
	whitePieceToMove := getPieceToMove(board, wOldSquare, wNewSquare)
	blackPieceToMove := getPieceToMove(board, bOldSquare, bNewSquare)
	newBoard[wOldSquare[0]][wOldSquare[1]] = "e"
	newBoard[bOldSquare[0]][bOldSquare[1]] = "e"
	newBoard[wNewSquare[0]][wNewSquare[1]] = whitePieceToMove
	newBoard[bNewSquare[0]][bNewSquare[1]] = blackPieceToMove
	return newBoard
}

// Gets the piece that should be "moved" - i.e. should be put into the new square
func getPieceToMove(board [5][5]string, oldSquare [2]int, newSquare [2]int) string {
	if string(board[oldSquare[0]][oldSquare[1]][1]) == "0" && newSquare[1] == 4 {
		// If pawn is moving onto the last rank, make it a knight
		return string(board[oldSquare[0]][oldSquare[1]][0]) + "1"
	} else {
		return board[oldSquare[0]][oldSquare[1]]
	}
}

// Returns "b" if black wins, "w" if white wins, "d" if its a draw, and "x" if the game isnt over
func checkGameOver(board [5][5]string) string {
	whitePawnsLeft := countPawns(board, "w")
	blackPawnsLeft := countPawns(board, "b")
	if whitePawnsLeft == 0 && blackPawnsLeft == 0 {
		return "d"
	} else if whitePawnsLeft == 0 {
		return "b"
	} else if blackPawnsLeft == 0 {
		return "w"
	} else {
		return "x"
	}
}

// Counts the number of pawns a player has
func countPawns(board [5][5]string, player string) int {
	count := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] == player+"0" {
				count++
			}
		}
	}
	return count
}
