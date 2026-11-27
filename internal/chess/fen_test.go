// MIT License
// Copyright (c) 2026 Aziz Shamuratov

package chess

import (
	"strings"
	"testing"
)

const startingFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func TestFENRoundTrip(t *testing.T) {
	cases := []struct{ name, fen string }{
		{"start", startingFEN},
		{"after e4", "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"},
		{"after e4 e5", "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2"},
		{"after Nf3", "rnbqkbnr/pppp1ppp/8/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2"},
		{"kiwipete", "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"},
		{"perft position 3", "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			board, err := ParseFEN(tc.fen)
			if err != nil {
				t.Fatal(err)
			}
			if got := board.FEN(); got != tc.fen {
				t.Errorf("FEN() = %q, want %q", got, tc.fen)
			}
		})
	}
}

func TestFENStartingPositionMatchesNewBoard(t *testing.T) {
	got, err := ParseFEN(startingFEN)
	if err != nil {
		t.Fatal(err)
	}
	want, err := NewBoard()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("parsed board differs from NewBoard: got %#v, want %#v", got, want)
	}
	if got := want.FEN(); got != startingFEN {
		t.Errorf("starting FEN = %q, want %q", got, startingFEN)
	}
}

func TestParseFENInvalid(t *testing.T) {
	cases := []struct{ name, fen, message string }{
		{"field count", "8/8/8/8/8/8/8/8 w - - 0", "expected 6 fields"},
		{"rank count", "8/8/8/8/8/8/8 w - - 0 1", "expected 8 ranks"},
		{"rank width", "9/8/8/8/8/8/8/8 w - - 0 1", "piece placement"},
		{"invalid piece", "7x/8/8/8/8/8/8/8 w - - 0 1", "invalid piece"},
		{"side", "8/8/8/8/8/8/8/8 x - - 0 1", "side to move"},
		{"castling", "8/8/8/8/8/8/8/8 w KK - 0 1", "castling rights"},
		{"target", "8/8/8/8/8/8/8/8 w - e3 0 1", "en passant target"},
		{"halfmove", "8/8/8/8/8/8/8/8 w - - -1 1", "halfmove clock"},
		{"fullmove", "8/8/8/8/8/8/8/8 w - - 0 0", "fullmove number"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParseFEN(tc.fen)
			if err == nil || !strings.Contains(err.Error(), tc.message) {
				t.Errorf("ParseFEN error = %v, want containing %q", err, tc.message)
			}
		})
	}
}
