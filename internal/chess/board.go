// MIT License

// Copyright (c) 2026 Aziz Shamuratov

package chess

import (
	"errors"
	"fmt"
	"strconv"
)

// Basic type for Pieces in chess logic. (Just basic OOP, nothing too crazy here)
type PieceType struct {
	Pcolor string // "White", "Black", ""
	Ptype  int8   // 1-Pawn 2-Rook 3-Knight 4-Bishop 5-Queen 6-King
}

func newPieceType() PieceType {
	return PieceType{
		Pcolor: "", Ptype: 0,
	}
}

// ??? idk why but I will need it in future
type Square int

type Board struct {
	Bpieces [64]PieceType
	Bmove   string // White ,Blacks
}

// Convert chess notation into number: e4 -> 28 (Implemented manipulating on ASCII conversion)
// I am not sure if I want to return plain int or obj Square, note taken.
func Str2Sqr(str string) (Square, error) {
	if len(str) != 2 {
		return Square(-1), errors.New("string < 2")
	}
	runes := []rune(str)
	// I wrote this line without any help I promise
	x := (((int(runes[1]) - 49) * 8) + ((int(runes[0])) - 97))
	sqr := Square(x)
	return sqr, nil
}

// 8.22 3:20 Changes to creation of new board.
func NewBoard() (Board, error) {
	i := 0
	var b Board
	b.Bmove = "white"
	var k_white int8 = 2
	var j_white int8 = 4
	var k_black int8 = 2
	var j_black int8 = 4
	for i < 64 {
		piece := newPieceType()
		// Create for White: Rook(0) -> Knight(1) -> Bishop(2) -> Queen(3) -> King(4)
		if i < 5 {
			piece.Ptype = k_white
			var color string = "white"
			piece.Pcolor = color
			k_white++
			b.Bpieces[i] = piece
		}

		// Create for White: Bishop(5) -> Knight(6) -> Rook(7)
		if (i > 4) && (i < 8) {
			piece.Ptype = j_white
			piece.Pcolor = "white"
			j_white--
			b.Bpieces[i] = piece
		}

		// Create for White: Set of Pawns (8-15)
		if (i > 7) && (i < 16) {
			piece.Ptype = 1
			piece.Pcolor = "white"
			b.Bpieces[i] = piece
		}

		// Create empty pieces (16-47)
		if (i > 15) && (i < 48) {
			piece.Ptype = 0
			piece.Pcolor = ""
			b.Bpieces[i] = piece
		}

		// Create for Black: Set of Pawns (48-55)
		if (i > 47) && (i < 56) {
			piece.Ptype = 1
			piece.Pcolor = "black"
			b.Bpieces[i] = piece
		}

		// Create for Black: Rook(56) -> Knight(57) -> Bishop(58) -> Queen(59) -> King(60)
		if (i > 55) && (i < 61) {
			piece.Pcolor = "black"
			piece.Ptype = k_black
			k_black++ //I am either idiot or genious for this line of code
			b.Bpieces[i] = piece
		}

		// Create for Black: Bishop(61) -> Knight(62) -> Rook (63)
		if (i > 60) && (i < 64) {
			piece.Pcolor = "black"
			piece.Ptype = j_black
			j_black--
			b.Bpieces[i] = piece
		}

		i++
	}
	return b, nil

}

func White() string {
	return "White"
}

func Black() string {
	return "Black"
} // this is dumb bro, why did wrote it?!

func DisplayBoard(x Board) {
	i := 7
	k := 0
	m := 1
	println("a b c d e f g h ")
	println()
	for i >= 0 {
		j := 7
		for j >= 0 {
			PrintE(x.Bpieces[k])
			k++
			j--
		}
		fmt.Printf("  [%v]", m)
		println()
		m++
		i--
	}
}

func ParseFEN(x Board) string {
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

func PrintE(p PieceType) error {
	str := '0'

	if p.Pcolor == "white" {
		if p.Ptype == 1 {
			str = 'P'
		}
		if p.Ptype == 2 {
			str = 'R'
		}
		if p.Ptype == 3 {
			str = 'N'
		}
		if p.Ptype == 4 {
			str = 'B'
		}
		if p.Ptype == 5 {
			str = 'Q'
		}
		if p.Ptype == 6 {
			str = 'K'
		}
	}
	if p.Pcolor == "black" {
		if p.Ptype == 1 {
			str = 'p'
		}
		if p.Ptype == 2 {
			str = 'r'
		}
		if p.Ptype == 3 {
			str = 'n'
		}
		if p.Ptype == 4 {
			str = 'b'
		}
		if p.Ptype == 5 {
			str = 'q'
		}
		if p.Ptype == 6 {
			str = 'k'
		}
	}
	fmt.Printf("%c ", str)
	return nil
}

func PieceTypeA(p PieceType) string {
	str := "0"

	if p.Pcolor == "white" {
		if p.Ptype == 1 {
			str = "P"
		}
		if p.Ptype == 2 {
			str = "R"
		}
		if p.Ptype == 3 {
			str = "N"
		}
		if p.Ptype == 4 {
			str = "B"
		}
		if p.Ptype == 5 {
			str = "Q"
		}
		if p.Ptype == 6 {
			str = "K"
		}
	}
	if p.Pcolor == "black" {
		if p.Ptype == 1 {
			str = "p"
		}
		if p.Ptype == 2 {
			str = "r"
		}
		if p.Ptype == 3 {
			str = "n"
		}
		if p.Ptype == 4 {
			str = "b"
		}
		if p.Ptype == 5 {
			str = "q"
		}
		if p.Ptype == 6 {
			str = "k"
		}
	}
	return str
}
