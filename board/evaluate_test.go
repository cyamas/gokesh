package board

import (
	"fmt"
	"testing"
)

func TestMiniMax(t *testing.T) {
	testNum := 1
	brd1 := New()
	brd1.SetupFromFen("2bqk3/5p2/8/1Q6/8/8/3P4/4K3")
	brd1.Evaluate(BLACK)

	brd2 := New()
	brd2.SetupFromFen("rn3rk1/pp2bppp/2p2p1B/8/2Pq4/1P1B2QP/P4PPK/8")
	brd2.Evaluate(WHITE)

	tests := []struct {
		board *Board
		turn  string
	}{
		{brd1, BLACK},
		{brd1, WHITE},
		{brd2, BLACK},
	}

	for _, tt := range tests {
		ogFen := tt.board.Fen()
		tt.board.Evaluate(ENEMY[tt.turn])
		move := tt.board.BestMove(tt.turn)
		fen := tt.board.Fen()

		if ogFen != fen {
			fmt.Println("TESTNUM: ", testNum)
			t.Fatalf("Fen should be %s. Got %s", ogFen, fen)
		}
		tt.board.MovePiece(move)
		testNum++
	}
}
