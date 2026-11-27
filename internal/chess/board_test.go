// MIT License
// Copyright (c) 2026 Aziz Shamuratov

package chess

import (
	"strings"
	"testing"
)

func TestSquareConversions(t *testing.T) {
	cases := []struct {
		notation string
		square   Square
	}{
		{"a1", 0}, {"e4", 28}, {"h8", 63},
	}
	for _, tc := range cases {
		got, err := Str2Sqr(tc.notation)
		if err != nil || got != tc.square {
			t.Errorf("Str2Sqr(%q) = %d, %v; want %d", tc.notation, got, err, tc.square)
		}
	}
	for square := Square(0); square < 64; square++ {
		notation, err := Sqr2Str(square)
		if err != nil {
			t.Fatalf("Sqr2Str(%d): %v", square, err)
		}
		got, err := Str2Sqr(notation)
		if err != nil || got != square {
			t.Errorf("round trip for %d via %q = %d, %v", square, notation, got, err)
		}
	}
	for _, input := range []string{"", "a", "a10", "z9", "A1", "a0", "é1"} {
		if _, err := Str2Sqr(input); err == nil {
			t.Errorf("Str2Sqr(%q) accepted invalid input", input)
		}
	}
	for _, square := range []Square{-1, 64} {
		if _, err := Sqr2Str(square); err == nil {
			t.Errorf("Sqr2Str(%d) accepted invalid input", square)
		}
	}
}

func TestStartingPosition(t *testing.T) {
	board, err := NewBoard()
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		square Square
		piece  PieceType
	}{
		{0, PieceType{"white", 2}}, {4, PieceType{"white", 6}},
		{8, PieceType{"white", 1}}, {28, PieceType{}},
		{48, PieceType{"black", 1}}, {59, PieceType{"black", 5}},
		{60, PieceType{"black", 6}}, {63, PieceType{"black", 2}},
	}
	for _, tc := range cases {
		if got := board.Bpieces[tc.square]; got != tc.piece {
			t.Errorf("square %d = %#v; want %#v", tc.square, got, tc.piece)
		}
	}
	if board.Bmove != "white" {
		t.Errorf("side to move = %q; want white", board.Bmove)
	}
}

func TestBoardString(t *testing.T) {
	board, err := NewBoard()
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(board.String(), "\n")
	if len(lines) != 9 {
		t.Fatalf("got %d lines, want 9: %q", len(lines), board.String())
	}
	if lines[0] != "8 r n b q k b n r " {
		t.Errorf("top rank = %q", lines[0])
	}
	if lines[7] != "1 R N B Q K B N R " {
		t.Errorf("bottom rank = %q", lines[7])
	}
	if lines[8] != "  a b c d e f g h" {
		t.Errorf("file labels = %q", lines[8])
	}
}
