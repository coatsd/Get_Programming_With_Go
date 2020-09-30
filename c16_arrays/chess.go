package main

import (
	"fmt"
)

type chessBoard [8][8]rune
type blackBoard [2][8]rune
type whiteBoard [2][8]rune

func initBoard(board *chessBoard, blackPieces *blackBoard, whitePieces *whiteBoard) {
	backRow := [8]rune{'r','n','b','k','q','b','n','r'}
	frontRow := [8]rune{'p','p','p','p','p','p','p','p'}

	for i := 0; i < len(blackPieces[0]); i++ {
		blackPieces[0][i] = backRow[i]
	}
	for i := 0; i < len(blackPieces[1]); i++ {
		blackPieces[1][i] = frontRow[i]
	}
	for i := 0; i < len(whitePieces[0]); i++ {
		whitePieces[0][i] = backRow[i] - 32
	}
	for i := 0; i < len(whitePieces[1]); i++ {
		whitePieces[1][i] = frontRow[i] - 32
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = 42
		}
	}
	// init black back row onto board
	board[0] = blackPieces[0]
	// init black front row
	board[1] = blackPieces[1]
	// init white front row
	board[6] = whitePieces[1]
	// init white back row
	board[7] = whitePieces[0]
}

func printBoard(board *chessBoard) {
	var acc string
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			acc += " " + string(board[i][j])
		}
		acc += "\n"
	}
	fmt.Printf(acc)
}

func main() {
	var board chessBoard

	var blackPieces blackBoard
	var whitePieces whiteBoard

	initBoard(&board, &blackPieces, &whitePieces)
	printBoard(&board)
}