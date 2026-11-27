// MIT License
// Copyright (c) 2026 Aziz Shamuratov

package chess

import (
	"fmt"
	"strconv"
	"strings"
)

// FEN returns the six fields of the position in Forsyth-Edwards notation.
func (b Board) FEN() string {
	var out strings.Builder
	for rank := 7; rank >= 0; rank-- {
		if rank != 7 {
			out.WriteByte('/')
		}
		empty := 0
		for file := 0; file < 8; file++ {
			piece := b.Bpieces[rank*8+file]
			if piece.Ptype == 0 {
				empty++
				continue
			}
			if empty > 0 {
				out.WriteByte(byte('0' + empty))
				empty = 0
			}
			out.WriteString(piece.String())
		}
		if empty > 0 {
			out.WriteByte(byte('0' + empty))
		}
	}
	side := "w"
	if b.Bmove == "black" {
		side = "b"
	}
	castling := b.Bcastling
	if castling == "" {
		castling = "-"
	}
	target := "-"
	if b.Benpass != NoSquare {
		if sq, err := Sqr2Str(b.Benpass); err == nil {
			target = sq
		}
	}
	fmt.Fprintf(&out, " %s %s %s %d %d", side, castling, target, b.Bhalf, b.Bfull)
	return out.String()
}

// FEN retains the original function form for callers using it.
func FEN(b Board) string { return b.FEN() }

// ParseFEN parses the six FEN fields. It validates notation and field values,
// but does not decide whether the position can arise in a legal game.
func ParseFEN(fen string) (Board, error) {
	var b Board
	b.Benpass = NoSquare
	fields := strings.Fields(fen)
	if len(fields) != 6 {
		return b, fmt.Errorf("FEN: expected 6 fields, got %d", len(fields))
	}

	ranks := strings.Split(fields[0], "/")
	if len(ranks) != 8 {
		return b, fmt.Errorf("piece placement: expected 8 ranks, got %d", len(ranks))
	}
	for i, rank := range ranks {
		file := 0
		previousDigit := false
		for _, ch := range rank {
			if ch >= '1' && ch <= '8' {
				if previousDigit {
					return b, fmt.Errorf("piece placement: consecutive empty counts on rank %d", 8-i)
				}
				file += int(ch - '0')
				previousDigit = true
			} else {
				previousDigit = false
				color := "white"
				letters := "PRNBQK"
				if ch >= 'a' && ch <= 'z' {
					color = "black"
					letters = "prnbqk"
				}
				kind := strings.IndexRune(letters, ch)
				if kind < 0 {
					return b, fmt.Errorf("piece placement: invalid piece %q on rank %d", ch, 8-i)
				}
				if file >= 8 {
					return b, fmt.Errorf("piece placement: rank %d exceeds 8 squares", 8-i)
				}
				b.Bpieces[(7-i)*8+file] = PieceType{Pcolor: color, Ptype: int8(kind + 1)}
				file++
			}
			if file > 8 {
				return b, fmt.Errorf("piece placement: rank %d exceeds 8 squares", 8-i)
			}
		}
		if file != 8 {
			return b, fmt.Errorf("piece placement: rank %d has %d squares, want 8", 8-i, file)
		}
	}
	switch fields[1] {
	case "w":
		b.Bmove = "white"
	case "b":
		b.Bmove = "black"
	default:
		return b, fmt.Errorf("side to move: expected w or b, got %q", fields[1])
	}
	if fields[2] != "-" {
		last := -1
		for _, c := range fields[2] {
			index := strings.IndexRune("KQkq", c)
			if index <= last || index < 0 {
				return b, fmt.Errorf("castling rights: invalid or out-of-order value %q", fields[2])
			}
			last = index
		}
		b.Bcastling = fields[2]
	}
	if fields[3] != "-" {
		sq, err := Str2Sqr(fields[3])
		if err != nil {
			return b, fmt.Errorf("en passant target: %w", err)
		}
		if (b.Bmove == "white" && fields[3][1] != '6') || (b.Bmove == "black" && fields[3][1] != '3') {
			return b, fmt.Errorf("en passant target: %q is on the wrong rank for %s to move", fields[3], b.Bmove)
		}
		b.Benpass = sq
	}
	var err error
	b.Bhalf, err = parseFENNumber(fields[4], 0)
	if err != nil {
		return b, fmt.Errorf("halfmove clock: %w", err)
	}
	b.Bfull, err = parseFENNumber(fields[5], 1)
	if err != nil {
		return b, fmt.Errorf("fullmove number: %w", err)
	}
	return b, nil
}

func parseFENNumber(s string, minimum int) (int, error) {
	if s == "" || (len(s) > 1 && s[0] == '0') {
		return 0, fmt.Errorf("invalid number %q", s)
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0, fmt.Errorf("invalid number %q", s)
		}
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q: %w", s, err)
	}
	if n < minimum {
		return 0, fmt.Errorf("number %d must be at least %d", n, minimum)
	}
	return n, nil
}
