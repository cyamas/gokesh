package game

import (
	"testing"

	"github.com/cyamas/gokesh/board"
)

func TestTurn(t *testing.T) {
	b := board.New()
	b.SetupPieces()

	game := New(b)

	blackEPawn := b.Squares[ROW_7][COL_E].Piece
	whiteEPawn := b.Squares[ROW_2][COL_E].Piece
	whiteRook := b.Squares[ROW_1][COL_A].Piece

	tests := []struct {
		input    *board.Move
		expected string
	}{
		{
			&board.Move{
				Piece: blackEPawn,
				From:  b.Squares[ROW_7][COL_E],
				To:    b.Squares[ROW_5][COL_E],
			},
			"BLACK ERROR: It is WHITE's turn to move",
		},
		{
			&board.Move{
				Piece: whiteEPawn,
				From:  b.Squares[ROW_2][COL_E],
				To:    b.Squares[ROW_4][COL_E],
			},
			"WHITE PAWN: E2 -> E4",
		},
		{
			&board.Move{
				Piece: whiteEPawn,
				From:  b.Squares[ROW_4][COL_E],
				To:    b.Squares[ROW_5][COL_E],
			},
			"WHITE ERROR: It is BLACK's turn to move",
		},
		{
			&board.Move{
				Piece: blackEPawn,
				From:  b.Squares[ROW_7][COL_E],
				To:    b.Squares[ROW_5][COL_E],
			},
			"BLACK PAWN: E7 -> E5",
		},
		{
			&board.Move{
				Piece: whiteRook,
				From:  b.Squares[ROW_1][COL_A],
				To:    b.Squares[ROW_4][COL_A],
			},
			"BOARD ERROR WHITE: ROOK: A1 -> A4 is not a valid move",
		},
	}

	for _, tt := range tests {
		receipt, _ := game.ExecuteTurn(tt.input)
		if receipt != tt.expected {
			t.Fatalf("receipt should be '%s'. Got '%s'", tt.expected, receipt)
		}
	}
}
