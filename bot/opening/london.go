package opening

import "github.com/cyamas/gokesh/board"

var LondonMoves = map[string]MoveFunc{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR":              pawnD2ToD4,
	"rnbqkbnr/ppp1pppp/8/3p4/3P4/8/PPP1PPPP/RNBQKBNR":          bishopC1ToF4,
	"rn1qkbnr/ppp1pppp/8/3p1b2/3P1B2/8/PPP1PPPP/RN1QKBNR":      pawnE2ToE3,
	"rnbqkbnr/pp2pppp/8/2pp4/3P1B2/8/PPP1PPPP/RN1QKBNR":        pawnE2ToE3,
	"rnbqkb1r/ppp1pppp/5n2/3p4/3P1B2/8/PPP1PPPP/RN1QKBNR":      pawnE2ToE3,
	"r1bqkbnr/ppp1pppp/2n5/3p4/3P1B2/8/PPP1PPPP/RN1QKBNR":      pawnE2ToE3,
	"rn1qkbnr/ppp2ppp/4p3/3p1b2/3P1B2/4P3/PPP2PPP/RN1QKBNR":    knightG1ToF3,
	"r2qkbnr/ppp2ppp/2n1p3/3p1b2/3P1B2/4PN2/PPP2PPP/RN1QKB1R":  knightB1ToD2,
	"rn1qkbnr/ppp2ppp/4p3/3p4/3P1B2/3bP3/PPP2PPP/RN1QK1NR":     pawnC2ToD3,
	"rn1qkb1r/ppp1pppp/5n2/3p1b2/3P1B2/4P3/PPP2PPP/RN1QKBNR":   bishopF1ToD3,
	"r2qkbnr/pppnpppp/8/3p1b2/3P1B2/4P3/PPP2PPP/RN1QKBNR":      bishopF1ToD3,
	"r2qkbnr/ppp1pppp/2n5/3p1b2/3P1B2/4P3/PPP2PPP/RN1QKBNR":    knightB1ToD2,
	"r2qkb1r/ppp2ppp/2n1pn2/3p1b2/3P1B2/4PN2/PPPN1PPP/R2QKB1R": pawnC2ToC4,
	"r2qkb1r/ppp1pppp/2n2n2/3p1b2/3P1B2/4P3/PPPN1PPP/R2QKBNR":  pawnC2ToC4,
	"r2qkbnr/1pp1pppp/p1n5/3p1b2/3P1B2/4P3/PPPN1PPP/R2QKBNR":   pawnC2ToC4,
	"r2qkbnr/ppp2ppp/2n1p3/3p1b2/3P1B2/4P3/PPPN1PPP/R2QKBNR":   pawnC2ToC4,
	"r2qk1nr/ppp2ppp/2nbp3/3p1b2/3P1B2/4PN2/PPPN1PPP/R2QKB1R":  bishopF4ToG3,
	"rn1qk1nr/ppp2ppp/3bp3/3p1b2/3P1B2/4PN2/PPP2PPP/RN1QKB1R":  bishopF4ToD6,
	"rn1qk1nr/pp3ppp/3pp3/3p1b2/3P4/4PN2/PPP2PPP/RN1QKB1R":     pawnC2ToC4,
	"rn1qk1nr/pp3ppp/3pp3/5b2/2pP4/4PN2/PP3PPP/RN1QKB1R":       bishopF1ToC4,
	"rn1qk2r/pp3ppp/3ppn2/5b2/2BP4/4PN2/PP3PPP/RN1QK2R":        shortCastleWhite,
}

func pawnD2ToD4(brd *board.Board) *board.Move {
	d2 := brd.Squares[ROW_2][COL_D]
	d4 := brd.Squares[ROW_4][COL_D]

	return &board.Move{
		Turn:  WHITE,
		Piece: d2.Piece,
		From:  d2,
		To:    d4,
	}
}

func bishopC1ToF4(brd *board.Board) *board.Move {
	c1 := brd.Squares[ROW_1][COL_C]
	f4 := brd.Squares[ROW_4][COL_F]

	return &board.Move{
		Turn:  WHITE,
		Piece: c1.Piece,
		From:  c1,
		To:    f4,
	}
}

func pawnE2ToE3(brd *board.Board) *board.Move {
	e2 := brd.Squares[ROW_2][COL_E]
	e3 := brd.Squares[ROW_3][COL_E]

	return &board.Move{
		Turn:  WHITE,
		Piece: e2.Piece,
		From:  e2,
		To:    e3,
	}
}

func knightG1ToF3(brd *board.Board) *board.Move {
	g1 := brd.Squares[ROW_1][COL_G]
	f3 := brd.Squares[ROW_3][COL_F]

	return &board.Move{
		Turn:  WHITE,
		Piece: g1.Piece,
		From:  g1,
		To:    f3,
	}
}

func bishopF1ToD3(brd *board.Board) *board.Move {
	f1 := brd.Squares[ROW_1][COL_F]
	d3 := brd.Squares[ROW_3][COL_D]

	return &board.Move{
		Turn:  WHITE,
		Piece: f1.Piece,
		From:  f1,
		To:    d3,
	}
}

func pawnC2ToD3(brd *board.Board) *board.Move {
	c2 := brd.Squares[ROW_2][COL_C]
	d3 := brd.Squares[ROW_3][COL_D]

	return &board.Move{
		Turn:  WHITE,
		Piece: c2.Piece,
		From:  c2,
		To:    d3,
	}
}

func knightB1ToD2(brd *board.Board) *board.Move {
	b1 := brd.Squares[ROW_1][COL_B]
	d2 := brd.Squares[ROW_2][COL_D]

	return &board.Move{
		Turn:  WHITE,
		Piece: b1.Piece,
		From:  b1,
		To:    d2,
	}
}

func pawnC2ToC4(brd *board.Board) *board.Move {
	c2 := brd.Squares[ROW_2][COL_C]
	c4 := brd.Squares[ROW_4][COL_C]

	return &board.Move{
		Turn:  WHITE,
		Piece: c2.Piece,
		From:  c2,
		To:    c4,
	}
}

func bishopF4ToD6(brd *board.Board) *board.Move {
	f4 := brd.Squares[ROW_4][COL_F]
	d6 := brd.Squares[ROW_6][COL_D]

	return &board.Move{
		Turn:  WHITE,
		Piece: f4.Piece,
		From:  f4,
		To:    d6,
	}
}

func bishopF4ToG3(brd *board.Board) *board.Move {
	f4 := brd.Squares[ROW_4][COL_F]
	g3 := brd.Squares[ROW_3][COL_G]

	return &board.Move{
		Turn:  WHITE,
		Piece: f4.Piece,
		From:  f4,
		To:    g3,
	}
}

func bishopF1ToC4(brd *board.Board) *board.Move {
	f1 := brd.Squares[ROW_1][COL_F]
	c4 := brd.Squares[ROW_4][COL_C]

	return &board.Move{
		Turn:  WHITE,
		Piece: f1.Piece,
		From:  f1,
		To:    c4,
	}
}

func shortCastleWhite(brd *board.Board) *board.Move {
	e1 := brd.Squares[ROW_1][COL_E]
	g1 := brd.Squares[ROW_1][COL_G]

	return &board.Move{
		Turn:  WHITE,
		Piece: e1.Piece,
		From:  e1,
		To:    g1,
	}
}
