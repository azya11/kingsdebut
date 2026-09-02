// MIT License

// Copyright (c) 2026 Aziz Shamuratov

package chess

import (
	"fmt"
	"strconv"
)

// Converts current board into FEN string
func FEN(x Board) string {
	i := 7
	k := 0
	m := 0
	y := ""
	for i >= 0 {
		j := 7
		for j >= 0 {
			if x.Bpieces[m].Ptype != 0 {
				y = y + PieceTypeA(x.Bpieces[m])

			} else {
				k++
			}
			m++
			j--
		}
		if k > 0 {
			y = y + strconv.Itoa(k)
		}
		y = y + "/"
		i--
		k = 0
	}
	y = y + " "
	if x.Bmove == "white" {
		y = y + "w"
	} else {
		y = y + "b"
	}

	//ADD CASTLING RIGHTS, EN PAUSSANT, HALF-MOVE CLOCK, FULL MOVE
	fmt.Print(y)
	return y
}

// Converts FEN string into board
func ParseFEN(x string) (Board, error) {
	var b Board
	return b, nil
}
