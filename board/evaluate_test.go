package board

import (
	"testing"
)

func TestMiniMax(t *testing.T) {
	board1 := New()

	e1 := board1.Squares[ROW_1][COL_E]
	e8 := board1.Squares[ROW_8][COL_E]

	whitePawn := &Pawn{color: WHITE, moved: true, value: 1.00}
	b7 := board1.Squares[ROW_7][COL_B]

	blackBishop := &Bishop{color: BLACK, value: -3.33}
	c8 := board1.Squares[ROW_8][COL_C]

	whiteBishop := &Bishop{color: WHITE, value: 3.33}
	e6 := board1.Squares[ROW_6][COL_E]

	board1.SetPiece(&King{color: WHITE, moved: true, value: 99.9}, e1)
	board1.SetPiece(&King{color: BLACK, moved: true, value: -99.9}, e8)
	board1.SetPiece(whitePawn, b7)
	board1.SetPiece(blackBishop, c8)
	board1.SetPiece(whiteBishop, e6)

	board2 := New()
	board2.SetupFromFen("r1b1kbnr/2p3pp/p4p2/1p1pp3/1q1n4/2Q3B1/PPP1PPPP/RN2KBNR")

	board3 := New()
	board3.SetupFromFen("rn5k/7p/p2r4/1p1P1pKR/1P5P/P4P2/2B2P2/8")

	board4 := New()
	board4.SetupFromFen("5b1r/2p2ppp/p4k2/1p2p3/1Q3B2/4P3/PPP2PPP/3RK1NR")

	board5 := New()
	board5.SetupFromFen("1k6/4Q3/1K6/8/3p4/3P4/8/8")

	tests := []struct {
		board    *Board
		input    string
		expected string
	}{
		//{board1, BLACK, "BISHOP TAKES PAWN: C8 -> B7"},
		//{board2, WHITE, "QUEEN TAKES QUEEN: C3 -> B4"},
		//{board3, BLACK, "ROOK TAKES PAWN: D6 -> D5"},
		//{board4, WHITE, "QUEEN B4 -> C3"},
		//{board5, BLACK, "KING: B8 -> A8"},
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
