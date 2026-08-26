// MIT License

// Copyright (c) 2026 Aziz Shamuratov

package chess

//Basic type for Pieces in chess logic. (Just basic OOP, nothing too crazy here)
type PieceType struct {
	Pcolor string // "White", "Black", ""
	Ptype  int8   // 1-Pawn 2-Rook 3-Knight 4-Bishop 5-Queen 6-King
}

func newPieceType() *PieceType {
	return &PieceType{
		Pcolor: "", Ptype: 0,
	}
}

//??? idk why but I will need it in future
type Square int

type Board struct {
	Bpieces [64]*PieceType
	Bmove   string // White ,Black
}

// Convert chess notation into number: e4 -> 28 (Implemented manipulating on ASCII conversion)
//I am not sure if I want to return plain int or obj Square, note taken.
func Str2Sqr(str string) (Square, error) {
	runes := []rune(str)
	// I wrote this line without any help I promise
	x := (((int(runes[1]) - 49) * 8) + ((int(runes[0])) - 97))
	sqr := Square(x)
	return sqr, nil
}

//8.22 3:20 Changes to creation of new board.
func NewBoard() (Board, error) {
	i := 0
	var default_pieces []PieceType
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
			default_pieces = append(default_pieces, piece)
		}

		// Create for White: Bishop(5) -> Knight(6) -> Rook(7)
		if (i > 4) && (i < 8) {
			piece.Ptype = j_white
			piece.Pcolor = "White"
			j_white--
			default_pieces = append(default_pieces, piece)
		}

		// Create for White: Set of Pawns (8-15)
		if (i > 7) && (i < 16) {

			default_pieces = append(default_pieces, piece)
		}

		// Create empty pieces (16-47)
		if (i > 15) && (i < 48) {
			default_pieces = append(default_pieces, piece)
		}

		// Create for Black: Set of Pawns (48-55)
		if (i > 47) && (i < 56) {
			default_pieces = append(default_pieces, piece)
		}

		// Create for Black: Rook(56) -> Knight(57) -> Bishop(58) -> Queen(59) -> King(60)
		if (i > 55) && (i < 60) {
			k_black++ //I am either idiot or genious for this line of code
			default_pieces = append(default_pieces, piece)
		}

		// Create for Black: Bishop(61) -> Knight(62) -> Rook (63)
		if (i > 60) && (i < 64) {
			j_black--
			default_pieces == append(default_pieces, piece)
		}

		i++
	}
	board := Board(default_pieces)

}

func White() string {
	return "White"
}

func Black() string {
	return "Black"
}

//Just let it be here I hate empty main
func Sum(x int, y int) (int, error) {
	return x + y, nil
}
