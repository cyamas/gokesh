package board

import "testing"

func TestMiniMax(t *testing.T) {
	board1 := New()

	e1 := board1.Squares[ROW_1][COL_E]
	e8 := board1.Squares[ROW_8][COL_E]

	whitePawn := &Pawn{color: WHITE, Moved: true, value: 1.00}
	b7 := board1.Squares[ROW_7][COL_B]

	blackBishop := &Bishop{color: BLACK, value: -3.33}
	c8 := board1.Squares[ROW_8][COL_C]

	whiteBishop := &Bishop{color: WHITE, value: 3.33}
	e6 := board1.Squares[ROW_6][COL_E]

	board1.SetPiece(&King{color: WHITE, Moved: true, value: 99.9}, e1)
	board1.SetPiece(&King{color: BLACK, Moved: true, value: -99.9}, e8)
	board1.SetPiece(whitePawn, b7)
	board1.SetPiece(blackBishop, c8)
	board1.SetPiece(whiteBishop, e6)

	tests := []struct {
		board    *Board
		input    string
		expected string
	}{
		{board1, BLACK, "BISHOP TAKES PAWN: C8 -> B7"},
	}

	for _, tt := range tests {
		tt.board.Evaluate(ENEMY[tt.input])
		move := tt.board.BestMove(tt.input)
		receipt, _ := tt.board.MovePiece(move)
		if receipt != tt.expected {
			t.Fatalf("Receipt shoud be %s. Got %s", tt.expected, receipt)
		}
	}
}
