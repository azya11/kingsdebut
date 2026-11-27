// MIT License

// Copyright (c) 2026 Aziz Shamuratov

package chess

import (
	"fmt"
	"strings"
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

// NoSquare represents the absence of an en passant target.
const NoSquare Square = -1

type Board struct {
	Bpieces   [64]PieceType
	Bmove     string // White ,Blacks
	Bcastling string
	Benpass   Square // en passant target, or NoSquare when unavailable
	Bhalf     int
	Bfull     int
}

// Convert chess notation into number: e4 -> 28 (Implemented manipulating on ASCII conversion)
// I am not sure if I want to return plain int or obj Square, note taken.
func Str2Sqr(str string) (Square, error) {
	if len(str) != 2 || str[0] < 'a' || str[0] > 'h' || str[1] < '1' || str[1] > '8' {
		return Square(-1), fmt.Errorf("invalid square %q: expected a1 through h8", str)
	}
	return Square((str[1]-'1')*8 + (str[0] - 'a')), nil
}

// Sqr2Str converts an index from 0 to 63 into algebraic notation.
func Sqr2Str(sq Square) (string, error) {
	if sq < 0 || sq >= 64 {
		return "", fmt.Errorf("invalid square index %d", sq)
	}
	return string([]byte{byte(sq%8) + 'a', byte(sq/8) + '1'}), nil
}

// 8.22 3:20 Changes to creation of new board.
func NewBoard() (Board, error) {
	i := 0
	var b Board
	b.Bmove = "white"
	b.Bcastling = "KQkq"
	b.Benpass = NoSquare
	b.Bfull = 1
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

func DisplayBoard(x Board) {
	fmt.Println(x)
}

// String displays rank 8 first and rank 1 last.
func (b Board) String() string {
	var out strings.Builder
	for rank := 7; rank >= 0; rank-- {
		fmt.Fprintf(&out, "%d ", rank+1)
		for file := 0; file < 8; file++ {
			fmt.Fprintf(&out, "%s ", b.Bpieces[rank*8+file])
		}
		out.WriteByte('\n')
	}
	out.WriteString("  a b c d e f g h")
	return out.String()
}

// String returns the usual letter for a piece, or a dot for an empty square.
func (p PieceType) String() string {
	if p.Ptype < 1 || p.Ptype > 6 {
		return "."
	}
	var letters string
	switch p.Pcolor {
	case "white":
		letters = " PRNBQK"
	case "black":
		letters = " prnbqk"
	default:
		return "."
	}
	return string(letters[p.Ptype])
}

// PrintE prints one piece, retaining the original display helper.
func PrintE(p PieceType) error {
	_, err := fmt.Printf("%s ", p)
	return err
}

// PieceTypeA retains the existing piece-to-string helper used by FEN.
func PieceTypeA(p PieceType) string {
	return p.String()
}
