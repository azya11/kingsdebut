package src

import "slices"

type Pieces struct {
	id         int
	isCaptured bool
	hasMoved   bool
	legalMoves []string
	position   string
}

func Move(piece *Pieces, target string) bool {
	if slices.Contains(piece.legalMoves, target) {
		//check if the other pieces got captured
		//piece.position = target
		if piece.hasMoved == false {
			piece.hasMoved = true
		}

		return true
	} else {
		return false
	}
}

func UpdateMoves(piece *Pieces) bool {
	//crete logic here later
	return true
}
