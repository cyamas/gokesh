package main

import (
	"os"

	"github.com/cyamas/gokesh/board"
	"github.com/cyamas/gokesh/game"
)

func main() {
	board := board.New()
	board.SetupPieces()
	game := game.New(board)
	game.Run(os.Stdin, os.Stdout)
}
