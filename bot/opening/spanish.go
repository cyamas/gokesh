package opening

import "github.com/cyamas/gokesh/board"

func Spanish(brd *board.Board, opening *Opening) *board.Move {
	switch len(brd.Moves) {
	case 0:
		return pawnToE4(brd)
	case 2:
		return knightToF3(brd)
	default:
		return bishopToB5(brd)
	}

}

func pawnToE4(brd *board.Board) *board.Move {
	e2 := brd.Squares[ROW_2][COL_E]
	e4 := brd.Squares[ROW_4][COL_E]
	return &board.Move{
		Turn:  WHITE,
		Piece: e2.Piece,
		From:  e2,
		To:    e4,
	}
}

func knightToF3(brd *board.Board) *board.Move {
	g1 := brd.Squares[ROW_1][COL_G]
	f3 := brd.Squares[ROW_3][COL_F]
	return &board.Move{
		Turn:  WHITE,
		Piece: g1.Piece,
		From:  g1,
		To:    f3,
	}
}

func bishopToB5(brd *board.Board) *board.Move {
	f1 := brd.Squares[ROW_1][COL_F]
	b5 := brd.Squares[ROW_5][COL_B]
	return &board.Move{
		Turn:  WHITE,
		Piece: f1.Piece,
		From:  f1,
		To:    b5,
	}
}
