// MIT License

// Copyright (c) 2026 Aziz Shamuratov

package main

import (
	"kingsdebut.com/xo/internal/chess"
)

func main() {

	//Just let it be here for now, soon I will come back and make it work.
	board, err := chess.NewBoard()
	if err != nil {
		println("Failed to initalize new board")
	}
	chess.DisplayBoard(board)

}
