package opening

import (
	"github.com/cyamas/gokesh/board"
)

var CaroKannMoves = map[string]MoveFunc{
	"rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR":       pawnC7ToC6,
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR":       pawnC7ToC6,
	"rnbqkbnr/pp1ppppp/2p5/8/4P3/5N2/PPPP1PPP/RNBQKB1R":   pawnD7ToD5,
	"rnbqkbnr/pp1ppppp/2p5/8/3PP3/8/PPP2PPP/RNBQKBNR":     pawnD7ToD5,
	"rnbqkbnr/pp1ppppp/2p5/8/4P3/2N5/PPPP1PPP/R1BQKBNR":   pawnD7ToD5,
	"rnbqkbnr/pp2pppp/2p5/3p4/3PP3/2N5/PPP2PPP/R1BQKBNR":  pawnD5ToE4,
	"rnbqkbnr/pp2pppp/2p5/3pP3/3P4/8/PPP2PPP/RNBQKBNR":    bishopC8ToF5,
	"rnbqkbnr/pp2pppp/2p5/3P4/3P4/8/PPP2PPP/RNBQKBNR":     pawnC6ToD5,
	"rnbqkbnr/pp2pppp/8/3p4/3P4/5N2/PPP2PPP/RNBQKB1R":     knightB8ToC6,
	"r1bqkbnr/pp2pppp/2n5/3p4/2PP4/5N2/PP3PPP/RNBQKB1R":   knightG8ToF6,
	"rnbqkbnr/pp2pppp/8/3p4/2PP4/8/PP3PPP/RNBQKBNR":       knightG8ToF6,
	"rnbqkbnr/pp2pppp/8/3p4/3P1B2/8/PPP2PPP/RN1QKBNR":     bishopC8ToF5,
	"rn1qkbnr/pp2pppp/2p5/3pPb2/3P4/2N5/PPP2PPP/R1BQKBNR": pawnE7ToE6,
	"rn1qkbnr/pp2pppp/2p5/3pPb2/3P4/5N2/PPP2PPP/RNBQKB1R": pawnE7ToE6,
	"rn1qkbnr/pp2pppp/2p5/3pPb2/2PP4/8/PP3PPP/RNBQKBNR":   pawnE7ToE6,
	"rn1qkbnr/pp2pppp/2p5/3pPb2/3P4/3B4/PPP2PPP/RNBQK1NR": bishopF5ToD3,
	"rnbqkbnr/pp2pppp/2p5/8/3PN3/8/PPP2PPP/R1BQKBNR":      knightG8ToF6,
	"rnbqkb1r/pp2pppp/2p2N2/8/3P4/8/PPP2PPP/R1BQKBNR":     pawnE7ToF6,
	"r1bqkb1r/pp2pppp/2n2n2/3P4/3P4/5N2/PP3PPP/RNBQKB1R":  knightF6ToD5,
	"rnbqkb1r/pp3ppp/2p2p2/8/3P4/5N2/PPP2PPP/R1BQKB1R":    bishopF8ToD6,
	"rnbqkb1r/pp3ppp/2p2p2/8/3P4/2P5/PP3PPP/R1BQKBNR":     bishopF8ToD6,
	"rnbqkb1r/pp3ppp/2p2p2/8/3P4/P7/1PP2PPP/R1BQKBNR":     bishopF8ToD6,
	"rnbqk2r/pp3ppp/2pb1p2/8/3P4/5N2/PPP1BPPP/R1BQK2R":    shortCastleBlack,
	"rnbqk2r/pp3ppp/2pb1p2/8/3P4/3B1N2/PPP2PPP/R1BQK2R":   shortCastleBlack,
	"rnbqk2r/pp3ppp/2pb1p2/8/2BP4/5N2/PPP2PPP/R1BQK2R":    shortCastleBlack,
	"rnbqk2r/pp3ppp/2pb1p2/8/3P4/2PB4/PP3PPP/R1BQK1NR":    shortCastleBlack,
	"rnbqk2r/pp3ppp/2pb1p2/8/3P4/P7/1PP1BPPP/R1BQK1NR":    shortCastleBlack,
	"rnbqk2r/pp3ppp/2pb1p2/8/3P4/P2B4/1PP2PPP/R1BQK1NR":   shortCastleBlack,
	"rnbqk2r/pp3ppp/2pb1p2/8/3P4/P4N2/1PP2PPP/R1BQKB1R":   shortCastleBlack,
	"rnbqk2r/pp3ppp/2pb1p2/8/2BP4/P7/1PP2PPP/R1BQK1NR":    shortCastleBlack,
}

func bishopF5ToD3(brd *board.Board) *board.Move {
	f5 := brd.Squares[ROW_5][COL_F]
	d3 := brd.Squares[ROW_3][COL_D]
	return &board.Move{
		Turn:  BLACK,
		Piece: f5.Piece,
		From:  f5,
		To:    d3,
	}
}

func pawnC7ToC6(brd *board.Board) *board.Move {
	c7 := brd.Squares[ROW_7][COL_C]
	c6 := brd.Squares[ROW_6][COL_C]
	return &board.Move{
		Turn:  BLACK,
		Piece: c7.Piece,
		From:  c7,
		To:    c6,
	}
}

func pawnD7ToD5(brd *board.Board) *board.Move {
	d7 := brd.Squares[ROW_7][COL_D]
	d5 := brd.Squares[ROW_5][COL_D]
	return &board.Move{
		Turn:  BLACK,
		Piece: d7.Piece,
		From:  d7,
		To:    d5,
	}
}

func pawnC6ToD5(brd *board.Board) *board.Move {
	c6 := brd.Squares[ROW_6][COL_C]
	d5 := brd.Squares[ROW_5][COL_D]
	return &board.Move{
		Turn:  BLACK,
		Piece: c6.Piece,
		From:  c6,
		To:    d5,
	}
}

func pawnE7ToF6(brd *board.Board) *board.Move {
	e7 := brd.Squares[ROW_7][COL_E]
	f6 := brd.Squares[ROW_6][COL_F]
	return &board.Move{
		Turn:  BLACK,
		Piece: e7.Piece,
		From:  e7,
		To:    f6,
	}
}

func bishopC8ToF5(brd *board.Board) *board.Move {
	c8 := brd.Squares[ROW_8][COL_C]
	f5 := brd.Squares[ROW_5][COL_F]
	return &board.Move{
		Turn:  BLACK,
		Piece: c8.Piece,
		From:  c8,
		To:    f5,
	}
}

func bishopG4ToH5(brd *board.Board) *board.Move {
	g4 := brd.Squares[ROW_4][COL_G]
	h5 := brd.Squares[ROW_5][COL_H]
	return &board.Move{
		Turn:  BLACK,
		Piece: g4.Piece,
		From:  g4,
		To:    h5,
	}
}

func bishopC8ToG4(brd *board.Board) *board.Move {
	c8 := brd.Squares[ROW_8][COL_C]
	g4 := brd.Squares[ROW_4][COL_G]
	return &board.Move{
		Turn:  BLACK,
		Piece: c8.Piece,
		From:  c8,
		To:    g4,
	}
}

func pawnE7ToE6(brd *board.Board) *board.Move {
	e7 := brd.Squares[ROW_7][COL_E]
	e6 := brd.Squares[ROW_6][COL_E]
	return &board.Move{
		Turn:  BLACK,
		Piece: e7.Piece,
		From:  e7,
		To:    e6,
	}
}

func bishopF8ToE7(brd *board.Board) *board.Move {
	f8 := brd.Squares[ROW_8][COL_F]
	e7 := brd.Squares[ROW_7][COL_E]
	return &board.Move{
		Turn:  BLACK,
		Piece: f8.Piece,
		From:  f8,
		To:    e7,
	}
}

func bishopF8ToD6(brd *board.Board) *board.Move {
	f8 := brd.Squares[ROW_8][COL_F]
	d6 := brd.Squares[ROW_6][COL_D]
	return &board.Move{
		Turn:  BLACK,
		Piece: f8.Piece,
		From:  f8,
		To:    d6,
	}
}

func pawnD5ToE4(brd *board.Board) *board.Move {
	d5 := brd.Squares[ROW_5][COL_D]
	e4 := brd.Squares[ROW_4][COL_E]
	return &board.Move{
		Turn:  BLACK,
		Piece: d5.Piece,
		From:  d5,
		To:    e4,
	}
}

func knightB8ToC6(brd *board.Board) *board.Move {
	b8 := brd.Squares[ROW_8][COL_B]
	c6 := brd.Squares[ROW_6][COL_C]
	return &board.Move{
		Turn:  BLACK,
		Piece: b8.Piece,
		From:  b8,
		To:    c6,
	}
}

func knightG8ToF6(brd *board.Board) *board.Move {
	g8 := brd.Squares[ROW_8][COL_G]
	f6 := brd.Squares[ROW_6][COL_F]
	return &board.Move{
		Turn:  BLACK,
		Piece: g8.Piece,
		From:  g8,
		To:    f6,
	}
}

func knightG8ToE7(brd *board.Board) *board.Move {
	g8 := brd.Squares[ROW_8][COL_G]
	e7 := brd.Squares[ROW_7][COL_E]
	return &board.Move{
		Turn:  BLACK,
		Piece: g8.Piece,
		From:  g8,
		To:    e7,
	}
}

func knightE7ToG6(brd *board.Board) *board.Move {
	e7 := brd.Squares[ROW_7][COL_E]
	g6 := brd.Squares[ROW_6][COL_G]
	return &board.Move{
		Turn:  BLACK,
		Piece: e7.Piece,
		From:  e7,
		To:    g6,
	}
}

func knightB8ToD7(brd *board.Board) *board.Move {
	b8 := brd.Squares[ROW_8][COL_B]
	d7 := brd.Squares[ROW_7][COL_D]
	return &board.Move{
		Turn:  BLACK,
		Piece: b8.Piece,
		From:  b8,
		To:    d7,
	}
}

func knightToB6(brd *board.Board) *board.Move {
	d7 := brd.Squares[ROW_7][COL_D]
	b6 := brd.Squares[ROW_6][COL_B]
	return &board.Move{
		Turn:  BLACK,
		Piece: d7.Piece,
		From:  d7,
		To:    b6,
	}
}

func pawnToH6(brd *board.Board) *board.Move {
	h7 := brd.Squares[ROW_7][COL_H]
	h6 := brd.Squares[ROW_6][COL_H]
	return &board.Move{
		Turn:  BLACK,
		Piece: h7.Piece,
		From:  h7,
		To:    h6,
	}
}

func pawnToH5(brd *board.Board) *board.Move {
	h7 := brd.Squares[ROW_7][COL_H]
	h5 := brd.Squares[ROW_5][COL_H]
	return &board.Move{
		Turn:  BLACK,
		Piece: h7.Piece,
		From:  h7,
		To:    h5,
	}
}

func knightF6ToD5(brd *board.Board) *board.Move {
	f6 := brd.Squares[ROW_6][COL_F]
	d5 := brd.Squares[ROW_5][COL_D]
	return &board.Move{
		Turn:  BLACK,
		Piece: f6.Piece,
		From:  f6,
		To:    d5,
	}
}

func shortCastleBlack(brd *board.Board) *board.Move {
	e8 := brd.Squares[ROW_8][COL_E]
	g8 := brd.Squares[ROW_8][COL_G]
	return &board.Move{
		Turn:  BLACK,
		Piece: e8.Piece,
		From:  e8,
		To:    g8,
	}
}
