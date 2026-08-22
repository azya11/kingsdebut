// MIT License

// Copyright (c) 2026 Aziz Shamuratov

package main

import (
	"fmt"

	"kingsdebut.com/xo/internal/chess"
)

func main() {

	//Just let it be here for now, soon I will come back and make it work.
	x, err := chess.Sum(1, 1)
	if err == nil {
		fmt.Printf("1+1==%v", x)
	} else {
		fmt.Printf("failed to executue")
	}

}
