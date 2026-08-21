// MIT License

// Copyright (c) 2026 Aziz Shamuratov

package chess

type PieceType struct {
	Pcolor bool // True-White False-Black
	Ptype  int8 // 1-Pawn 2-Knight 3-Bishop 4-Rook 5-Queen 6-King
}

type Square int

func Str2Sqr(str string) (int, error) {
	runes := []rune(str)

	x := (((int(runes[1]) - 49) * 8) + ((int(runes[0])) - 97))
	// "a":  0,	1: 00
	// "b": 1,	2: 10
	// "c": 2,	3: 20
	// "d": 3,	4: 30
	// "e": 4,	5: 40
	// "f": 5,	6: 50
	// "g": 6,	7: 60
	// "h": 7,	8: 70
	return x, nil
}

//Testing function to ensure package import works
func Sum(a int, b int) (int, error) {
	return a + b, nil
}
