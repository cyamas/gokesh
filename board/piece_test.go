package board

import (
	"fmt"
	"testing"
)

func TestCheckmate(t *testing.T) {
	board1 := New()

	whiteKing1 := &King{color: WHITE}
	d1 := board1.Squares[ROW_1][COL_D]

	blackKing1 := &King{color: BLACK}
	e8 := board1.Squares[ROW_8][COL_E]

	blackQueen1 := &Queen{color: BLACK}
	a1 := board1.Squares[ROW_1][COL_A]

	blackRook1 := &Rook{color: BLACK}
	b2 := board1.Squares[ROW_2][COL_B]

	board1.SetPiece(whiteKing1, d1)
	board1.SetPiece(blackKing1, e8)
	board1.SetPiece(blackQueen1, a1)
	board1.SetPiece(blackRook1, b2)

	board2 := New()

	whiteKing2 := &King{color: WHITE}
	h1 := board2.Squares[ROW_1][COL_H]

	whiteGPawn2 := &Pawn{color: WHITE}
	g2 := board2.Squares[ROW_2][COL_G]

	whiteHPawn2 := &Pawn{color: WHITE}
	h2 := board2.Squares[ROW_2][COL_H]

	blackQueen2 := &Queen{color: BLACK}
	b1 := board2.Squares[ROW_1][COL_B]

	blackKing2 := &King{color: BLACK}
	a8 := board2.Squares[ROW_8][COL_A]

	board2.SetPiece(whiteKing2, h1)
	board2.SetPiece(blackKing2, a8)
	board2.SetPiece(whiteGPawn2, g2)
	board2.SetPiece(whiteHPawn2, h2)
	board2.SetPiece(blackQueen2, b1)

	board3 := New()

	whiteKing3 := &King{color: WHITE}
	f1 := board3.Squares[ROW_1][COL_F]

	blackKing3 := &King{color: BLACK}
	b8 := board3.Squares[ROW_8][COL_B]

	whiteQueen3 := &Queen{color: WHITE}
	b7 := board3.Squares[ROW_7][COL_B]

	whiteKnight3 := &Knight{color: WHITE}
	d6 := board3.Squares[ROW_6][COL_D]

	board3.SetPiece(whiteKing3, f1)
	board3.SetPiece(blackKing3, b8)
	board3.SetPiece(whiteQueen3, b7)
	board3.SetPiece(whiteKnight3, d6)

	board4 := New()
	whiteKing4 := &King{color: WHITE}
	h4 := board4.Squares[ROW_4][COL_H]

	blackKing4 := &King{color: BLACK}
	f4_4 := board4.Squares[ROW_4][COL_F]

	blackRook4 := &Rook{color: BLACK}
	h8 := board4.Squares[ROW_8][COL_H]

	board4.SetPiece(whiteKing4, h4)
	board4.SetPiece(blackKing4, f4_4)
	board4.SetPiece(blackRook4, h8)

	tests := []struct {
		turn     string
		testKing *King
		board    *Board
		expected bool
	}{
		{WHITE, whiteKing1, board1, true},
		{WHITE, whiteKing2, board2, true},
		{BLACK, blackKing3, board3, true},
		{WHITE, whiteKing4, board4, true},
	}

	for _, tt := range tests {
		tt.board.Evaluate(tt.turn)
		if tt.board.CheckmateDetected(tt.testKing.Color(), tt.testKing) != tt.expected {
			t.Fatalf(
				"%s KING on %s: Checkmate should be %t",
				tt.testKing.color,
				tt.testKing.Square().Name,
				tt.expected,
			)
		}
	}
}

func TestPieceBlocksCheck(t *testing.T) {
	board1 := New()
	board2 := New()

	whiteDKing := &King{color: WHITE}
	d1 := board1.Squares[ROW_1][COL_D]

	blackFKing := &King{color: BLACK}
	f8 := board1.Squares[ROW_8][COL_F]

	blackHKing := &King{color: BLACK}
	h5 := board2.Squares[ROW_5][COL_H]

	whiteAKing := &King{color: WHITE}
	a5 := board2.Squares[ROW_5][COL_A]

	whiteRook := &Rook{color: WHITE}
	e1 := board1.Squares[ROW_1][COL_E]

	blackBishop := &Bishop{color: BLACK}
	b3 := board1.Squares[ROW_3][COL_B]

	blackQueen := &Queen{color: BLACK}
	d4 := board1.Squares[ROW_4][COL_D]

	whiteQueen := &Queen{color: WHITE}
	h2 := board1.Squares[ROW_2][COL_H]

	whiteBishop := &Bishop{color: WHITE}
	f3 := board2.Squares[ROW_3][COL_F]

	blackKnight := &Knight{color: BLACK}
	c6 := board2.Squares[ROW_6][COL_C]

	blackPawn := &Pawn{color: BLACK, Moved: true}
	g5 := board2.Squares[ROW_5][COL_G]

	whitePawn := &Pawn{color: WHITE, Moved: true}
	b4 := board2.Squares[ROW_4][COL_B]

	board1.SetPiece(blackFKing, f8)
	board1.SetPiece(whiteDKing, d1)
	board1.SetPiece(whiteQueen, h2)
	board1.SetPiece(blackQueen, d4)
	board1.SetPiece(blackBishop, b3)
	board1.SetPiece(whiteRook, e1)

	board2.SetPiece(blackHKing, h5)
	board2.SetPiece(whiteAKing, a5)
	board2.SetPiece(whiteBishop, f3)
	board2.SetPiece(blackKnight, c6)
	board2.SetPiece(whitePawn, b4)
	board2.SetPiece(blackPawn, g5)

	tests := []struct {
		board    *Board
		input    *Move
		expected string
	}{
		{
			board2,
			&Move{
				Turn:  BLACK,
				Piece: blackPawn,
				From:  g5,
				To:    board2.Squares[ROW_4][COL_G],
			},
			"PAWN: G5 -> G4",
		},
		{
			board2,
			&Move{
				Turn:  WHITE,
				Piece: whitePawn,
				From:  b4,
				To:    board2.Squares[ROW_5][COL_B],
			},
			"PAWN: B4 -> B5 is not a valid move",
		},
		{
			board1,
			&Move{
				Turn:  WHITE,
				Piece: whiteQueen,
				From:  h2,
				To:    board1.Squares[ROW_2][COL_D],
			},
			"QUEEN: H2 -> D2 is not a valid move",
		},
	}

	for _, tt := range tests {
		tt.board.Evaluate(tt.input.Turn)
		receipt, err := tt.board.MovePiece(tt.input)
		if receipt != tt.expected {
			t.Fatalf("Receipt should be %s. Got %s", tt.expected, receipt)
		}
		if err == nil {
			if !testPieceHasMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should have moved", tt.input.Piece.Type())
			}
		} else {
			if !testPieceHasNotMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should not have moved", tt.input.Piece.Type())
			}
		}
	}

}

func TestAbsolutePins(t *testing.T) {
	board := New()

	whiteKing := &King{color: WHITE}
	e1 := board.Squares[ROW_1][COL_E]

	blackKing := &King{color: BLACK}
	e8 := board.Squares[ROW_8][COL_E]

	whiteBishop := &Bishop{color: WHITE}
	a4 := board.Squares[ROW_4][COL_A]

	blackDPawn := &Pawn{color: BLACK}
	d7 := board.Squares[ROW_7][COL_D]

	blackRook := &Rook{color: BLACK}
	c6 := board.Squares[ROW_6][COL_C]

	blackBishop := &Bishop{color: BLACK}
	b4 := board.Squares[ROW_4][COL_B]

	whiteDPawn := &Pawn{color: WHITE}
	d2 := board.Squares[ROW_2][COL_D]

	board.SetPiece(whiteDPawn, d2)
	board.SetPiece(blackBishop, b4)
	board.SetPiece(whiteKing, e1)
	board.SetPiece(blackKing, e8)
	board.SetPiece(whiteBishop, a4)
	board.SetPiece(blackDPawn, d7)
	board.SetPiece(blackRook, c6)

	tests := []struct {
		input    *Move
		expected string
	}{
		{
			&Move{
				Turn:  BLACK,
				Piece: blackRook,
				From:  c6,
				To:    board.Squares[ROW_5][COL_C],
			},
			"ROOK: C6 -> C5",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackDPawn,
				From:  d7,
				To:    board.Squares[ROW_6][COL_D],
			},
			"PAWN: D7 -> D6 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteDPawn,
				From:  d2,
				To:    board.Squares[ROW_4][COL_D],
			},
			"PAWN: D2 -> D4 is not a valid move",
		},
	}

	for _, tt := range tests {
		board.Evaluate(tt.input.Turn)
		receipt, err := board.MovePiece(tt.input)

		if receipt != tt.expected {
			t.Fatalf("Receipt should be '%s'. Got '%s'", tt.expected, receipt)
		}
		if err == nil {
			if !testPieceHasMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should have moved", tt.input.Piece.Type())
			}
		} else {
			if !testPieceHasNotMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should not have moved", tt.input.Piece.Type())
			}
		}
	}

}

func TestPawnPromotion(t *testing.T) {
	board := New()

	c7 := board.Squares[ROW_7][COL_C]
	e7 := board.Squares[ROW_7][COL_E]
	d2 := board.Squares[ROW_2][COL_D]
	f2 := board.Squares[ROW_2][COL_F]
	b8 := board.Squares[ROW_8][COL_B]
	g1 := board.Squares[ROW_1][COL_G]
	h7 := board.Squares[ROW_7][COL_H]
	a2 := board.Squares[ROW_2][COL_A]

	whiteEPawn := &Pawn{color: WHITE, Moved: true}
	whiteCPawn := &Pawn{color: WHITE, Moved: true}
	blackDPawn := &Pawn{color: BLACK, Moved: true}
	blackFPawn := &Pawn{color: BLACK, Moved: true}
	blackQueen := &Queen{color: BLACK}
	whiteBishop := &Bishop{color: WHITE}
	whiteKing := &King{color: WHITE}
	blackKing := &King{color: BLACK}

	board.SetPiece(whiteEPawn, e7)
	board.SetPiece(whiteCPawn, c7)
	board.SetPiece(blackDPawn, d2)
	board.SetPiece(blackFPawn, f2)
	board.SetPiece(blackQueen, b8)
	board.SetPiece(whiteBishop, g1)
	board.SetPiece(whiteKing, h7)
	board.SetPiece(blackKing, a2)

	tests := []struct {
		input             *Move
		expectedReceipt   string
		expectedPromotion string
	}{
		{
			&Move{
				Turn:      WHITE,
				Piece:     whiteEPawn,
				From:      e7,
				To:        board.Squares[ROW_8][COL_E],
				Promotion: &Queen{color: WHITE},
			},
			"PAWN: E7 -> E8 (PROMOTION: QUEEN)",
			QUEEN,
		},
		{
			&Move{
				Turn:      BLACK,
				Piece:     blackDPawn,
				From:      d2,
				To:        board.Squares[ROW_1][COL_D],
				Promotion: &Knight{color: BLACK},
			},
			"PAWN: D2 -> D1 (PROMOTION: KNIGHT)",
			KNIGHT,
		},
		{
			&Move{
				Turn:      WHITE,
				Piece:     whiteCPawn,
				From:      c7,
				To:        b8,
				Promotion: &Bishop{color: WHITE},
			},
			"PAWN TAKES QUEEN: C7 -> B8 (PROMOTION: BISHOP)",
			BISHOP,
		},
		{
			&Move{
				Turn:      BLACK,
				Piece:     blackFPawn,
				From:      f2,
				To:        g1,
				Promotion: &Rook{color: BLACK},
			},
			"PAWN TAKES BISHOP: F2 -> G1 (PROMOTION: ROOK)",
			ROOK,
		},
	}

	for _, tt := range tests {
		board.Evaluate(tt.input.Turn)
		receipt, err := board.MovePiece(tt.input)
		if err != nil {
			t.Fatalf("Test case should not return error")
		}
		if receipt != tt.expectedReceipt {
			t.Fatalf("Receipt should be '%s'. Got '%s'", tt.expectedReceipt, receipt)
		}
		if tt.input.From.Piece.Type() != NULL {
			t.Fatalf("Square %s should be NULL. Got %s", tt.input.From.Name, tt.input.From.Piece.Type())
		}
		if tt.input.To.Piece.Type() != tt.expectedPromotion {
			t.Fatalf("Square %s should now have Queen. Got %s", tt.input.To.Name, tt.input.To.Piece.Type())
		}
	}

}

func TestCastle(t *testing.T) {
	shortBoard := New()
	shortBlackSq := shortBoard.Squares[ROW_8][COL_G]
	shortWhiteSq := shortBoard.Squares[ROW_1][COL_G]
	shortWhiteKing := &King{color: WHITE}
	shortBlackKing := &King{color: BLACK}
	whiteFPawn := &Pawn{color: WHITE}
	blackFPawn := &Pawn{color: BLACK}
	shortWhiteRook := &Rook{color: WHITE, CastleSq: shortBoard.Squares[ROW_1][COL_F]}
	shortBlackRook := &Rook{color: BLACK, CastleSq: shortBoard.Squares[ROW_8][COL_F]}

	shortBoard.SetPiece(shortWhiteKing, shortBoard.Squares[ROW_1][COL_E])
	shortBoard.SetPiece(shortBlackKing, shortBoard.Squares[ROW_8][COL_E])
	shortBoard.SetPiece(shortWhiteRook, shortBoard.Squares[ROW_1][COL_H])
	shortBoard.SetPiece(shortBlackRook, shortBoard.Squares[ROW_8][COL_H])
	shortBoard.SetPiece(whiteFPawn, shortBoard.Squares[ROW_2][COL_F])
	shortBoard.SetPiece(blackFPawn, shortBoard.Squares[ROW_7][COL_F])

	longBoard := New()
	longBlackSq := longBoard.Squares[ROW_8][COL_C]
	longWhiteSq := longBoard.Squares[ROW_1][COL_C]
	longWhiteKing := &King{color: WHITE}
	longBlackKing := &King{color: BLACK}
	whiteDPawn := &Pawn{color: WHITE}
	longWhiteRook := &Rook{color: WHITE, CastleSq: longBoard.Squares[ROW_1][COL_D]}
	longBlackRook := &Rook{color: BLACK, CastleSq: longBoard.Squares[ROW_8][COL_D]}

	longBoard.SetPiece(longWhiteKing, longBoard.Squares[ROW_1][COL_E])
	longBoard.SetPiece(longBlackKing, longBoard.Squares[ROW_8][COL_E])
	longBoard.SetPiece(longWhiteRook, longBoard.Squares[ROW_1][COL_A])
	longBoard.SetPiece(longBlackRook, longBoard.Squares[ROW_8][COL_A])
	longBoard.SetPiece(whiteDPawn, longBoard.Squares[ROW_2][COL_D])

	checkedBoard := New()
	checkedBlackSq := checkedBoard.Squares[ROW_8][COL_C]
	checkedWhiteKing := &King{color: WHITE}
	checkedBlackKing := &King{color: BLACK}
	checkingWhiteFPawn := &Pawn{color: WHITE, Moved: true}
	checkedWhiteRook := &Rook{color: WHITE, CastleSq: checkedBoard.Squares[ROW_1][COL_D]}
	checkedBlackRook := &Rook{color: BLACK, CastleSq: checkedBoard.Squares[ROW_8][COL_D]}

	checkedBoard.SetPiece(checkedWhiteKing, checkedBoard.Squares[ROW_1][COL_E])
	checkedBoard.SetPiece(checkedBlackKing, checkedBoard.Squares[ROW_8][COL_E])
	checkedBoard.SetPiece(checkedWhiteRook, checkedBoard.Squares[ROW_1][COL_A])
	checkedBoard.SetPiece(checkedBlackRook, checkedBoard.Squares[ROW_8][COL_A])
	checkedBoard.SetPiece(checkingWhiteFPawn, checkedBoard.Squares[ROW_7][COL_F])

	blockedBoard := New()
	blockedBlackSq := blockedBoard.Squares[ROW_8][COL_C]
	blockedWhiteSq := blockedBoard.Squares[ROW_1][COL_C]
	blockedWhiteKing := &King{color: WHITE}
	blockedBlackKing := &King{color: BLACK}
	blockedWhiteRook := &Rook{color: WHITE, CastleSq: blockedBoard.Squares[ROW_1][COL_D]}
	blockedBlackRook := &Rook{color: BLACK, CastleSq: blockedBoard.Squares[ROW_8][COL_D]}
	blockingWhiteKnight := &Knight{color: WHITE}
	blockingBlackKnight := &Knight{color: BLACK}

	blockedBoard.SetPiece(blockedWhiteKing, blockedBoard.Squares[ROW_1][COL_E])
	blockedBoard.SetPiece(blockedBlackKing, blockedBoard.Squares[ROW_8][COL_E])
	blockedBoard.SetPiece(blockingWhiteKnight, blockedBoard.Squares[ROW_1][COL_C])
	blockedBoard.SetPiece(blockingBlackKnight, blockedBoard.Squares[ROW_8][COL_D])
	blockedBoard.SetPiece(blockedWhiteRook, blockedBoard.Squares[ROW_1][COL_A])
	blockedBoard.SetPiece(blockedBlackRook, blockedBoard.Squares[ROW_8][COL_A])

	enemyBlockedBoard := New()
	enemyBlockedWhiteSq := enemyBlockedBoard.Squares[ROW_1][COL_C]
	enemyBlockedBlackSq := enemyBlockedBoard.Squares[ROW_8][COL_G]
	enemyBlockedWhiteKing := &King{color: WHITE}
	enemyBlockedBlackKing := &King{color: BLACK}
	enemyBlockedWhiteRook := &Rook{color: WHITE, CastleSq: enemyBlockedBoard.Squares[ROW_1][COL_D]}
	enemyBlockedBlackRook := &Rook{color: BLACK, CastleSq: enemyBlockedBoard.Squares[ROW_8][COL_F]}
	enemyBlackBishop := &Bishop{color: BLACK}
	enemyWhiteBishop := &Bishop{color: WHITE}

	enemyBlockedBoard.SetPiece(enemyBlockedWhiteKing, enemyBlockedBoard.Squares[ROW_1][COL_E])
	enemyBlockedBoard.SetPiece(enemyBlockedBlackKing, enemyBlockedBoard.Squares[ROW_8][COL_E])
	enemyBlockedBoard.SetPiece(enemyBlackBishop, enemyBlockedBoard.Squares[ROW_5][COL_G])
	enemyBlockedBoard.SetPiece(enemyWhiteBishop, enemyBlockedBoard.Squares[ROW_2][COL_A])
	enemyBlockedBoard.SetPiece(enemyBlockedWhiteRook, enemyBlockedBoard.Squares[ROW_1][COL_A])
	enemyBlockedBoard.SetPiece(enemyBlockedBlackRook, enemyBlockedBoard.Squares[ROW_8][COL_H])

	kingMovedBoard := New()
	kingMovedWhiteSq := kingMovedBoard.Squares[ROW_1][COL_C]
	kingMovedBlackSq := kingMovedBoard.Squares[ROW_8][COL_C]
	kingMovedWhiteKing := &King{color: WHITE, Moved: true}
	kingMovedBlackKing := &King{color: BLACK, Moved: true}
	kingMovedWhiteRook := &Rook{color: WHITE, CastleSq: kingMovedBoard.Squares[ROW_1][COL_D]}
	kingMovedBlackRook := &Rook{color: BLACK, CastleSq: kingMovedBoard.Squares[ROW_8][COL_D]}
	kingMovedWhitePawn := &Pawn{color: WHITE}
	kingMovedBlackPawn := &Pawn{color: BLACK}

	kingMovedBoard.SetPiece(kingMovedWhiteKing, kingMovedBoard.Squares[ROW_1][COL_E])
	kingMovedBoard.SetPiece(kingMovedBlackKing, kingMovedBoard.Squares[ROW_8][COL_E])
	kingMovedBoard.SetPiece(kingMovedWhiteRook, kingMovedBoard.Squares[ROW_1][COL_A])
	kingMovedBoard.SetPiece(kingMovedBlackRook, kingMovedBoard.Squares[ROW_8][COL_A])
	kingMovedBoard.SetPiece(kingMovedWhitePawn, kingMovedBoard.Squares[ROW_2][COL_D])
	kingMovedBoard.SetPiece(kingMovedBlackPawn, kingMovedBoard.Squares[ROW_7][COL_D])

	rookMovedBoard := New()
	rookMovedWhiteSq := rookMovedBoard.Squares[ROW_1][COL_G]
	rookMovedBlackSq := rookMovedBoard.Squares[ROW_8][COL_G]
	rookMovedWhiteKing := &King{color: WHITE}
	rookMovedBlackKing := &King{color: BLACK}
	rookMovedWhiteRook := &Rook{color: WHITE, Moved: true, CastleSq: rookMovedBoard.Squares[ROW_1][COL_F]}
	rookMovedBlackRook := &Rook{color: BLACK, Moved: true, CastleSq: rookMovedBoard.Squares[ROW_8][COL_F]}
	rookMovedWhitePawn := &Pawn{color: WHITE}
	rookMovedBlackPawn := &Pawn{color: BLACK}

	rookMovedBoard.SetPiece(rookMovedWhiteKing, rookMovedBoard.Squares[ROW_1][COL_E])
	rookMovedBoard.SetPiece(rookMovedBlackKing, rookMovedBoard.Squares[ROW_8][COL_E])
	rookMovedBoard.SetPiece(rookMovedWhiteRook, rookMovedBoard.Squares[ROW_1][COL_H])
	rookMovedBoard.SetPiece(rookMovedBlackRook, rookMovedBoard.Squares[ROW_8][COL_H])
	rookMovedBoard.SetPiece(rookMovedWhitePawn, rookMovedBoard.Squares[ROW_1][COL_F])
	rookMovedBoard.SetPiece(rookMovedBlackPawn, rookMovedBoard.Squares[ROW_8][COL_F])

	tests := []struct {
		input    *Move
		board    *Board
		rook     *Rook
		expected string
	}{
		{
			&Move{
				Turn:  WHITE,
				Piece: shortWhiteKing,
				From:  shortWhiteKing.Square(),
				To:    shortWhiteSq,
			},
			shortBoard,
			shortWhiteRook,
			"KING CASTLES SHORT",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: shortBlackKing,
				From:  shortBlackKing.Square(),
				To:    shortBlackSq,
			},
			shortBoard,
			shortBlackRook,
			"KING CASTLES SHORT",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: longWhiteKing,
				From:  longWhiteKing.Square(),
				To:    longWhiteSq,
			},
			longBoard,
			longWhiteRook,
			"KING CASTLES LONG",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: longBlackKing,
				From:  longBlackKing.Square(),
				To:    longBlackSq,
			},
			longBoard,
			longBlackRook,
			"KING CASTLES LONG",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: checkedBlackKing,
				From:  checkedBlackKing.Square(),
				To:    checkedBlackSq,
			},
			checkedBoard,
			checkedBlackRook,
			"KING: E8 -> C8 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: blockedWhiteKing,
				From:  blockedWhiteKing.Square(),
				To:    blockedWhiteSq,
			},
			blockedBoard,
			blockedWhiteRook,
			"KING: E1 -> C1 is not a valid move",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blockedBlackKing,
				From:  blockedBlackKing.Square(),
				To:    blockedBlackSq,
			},
			blockedBoard,
			blockedBlackRook,
			"KING: E8 -> C8 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: enemyBlockedWhiteKing,
				From:  enemyBlockedWhiteKing.Square(),
				To:    enemyBlockedWhiteSq,
			},
			enemyBlockedBoard,
			enemyBlockedWhiteRook,
			"KING: E1 -> C1 is not a valid move",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: enemyBlockedBlackKing,
				From:  enemyBlockedBlackKing.Square(),
				To:    enemyBlockedBlackSq,
			},
			enemyBlockedBoard,
			enemyBlockedBlackRook,
			"KING: E8 -> G8 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: kingMovedWhiteKing,
				From:  kingMovedWhiteKing.Square(),
				To:    kingMovedWhiteSq,
			},
			kingMovedBoard,
			kingMovedWhiteRook,
			"KING: E1 -> C1 is not a valid move",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: kingMovedBlackKing,
				From:  kingMovedBlackKing.Square(),
				To:    kingMovedBlackSq,
			},
			kingMovedBoard,
			kingMovedBlackRook,
			"KING: E8 -> C8 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: rookMovedWhiteKing,
				From:  rookMovedWhiteKing.Square(),
				To:    rookMovedWhiteSq,
			},
			rookMovedBoard,
			rookMovedWhiteRook,
			"KING: E1 -> G1 is not a valid move",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: rookMovedBlackKing,
				From:  rookMovedBlackKing.Square(),
				To:    rookMovedBlackSq,
			},
			rookMovedBoard,
			rookMovedBlackRook,
			"KING: E8 -> G8 is not a valid move",
		},
	}

	for _, tt := range tests {
		ogRookSq := tt.rook.Square()
		tt.board.Evaluate(tt.input.Turn)
		receipt, err := tt.board.MovePiece(tt.input)

		if receipt != tt.expected {
			t.Fatalf("receipt should be '%s'. Got '%s'", tt.expected, receipt)
		}
		if err == nil {
			if !testPieceHasMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should have moved", tt.input.Piece.Type())
			}
			if tt.rook.Square() != tt.rook.CastleSq {
				t.Fatalf("ROOK should be on %s. Got %s", tt.rook.CastleSq.Name, tt.rook.Square().Name)
			}
		} else {
			if !testPieceHasNotMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should not have moved", tt.input.Piece.Type())
			}
			if tt.rook.Square() != ogRookSq {
				t.Fatalf("ROOK should not have moved")
			}
		}
	}

}

func TestKingInCheck(t *testing.T) {
	board1 := New()

	whiteKing1 := &King{color: WHITE}
	e3 := board1.Squares[ROW_3][COL_E]

	blackBishop1 := &Bishop{color: BLACK}
	a7 := board1.Squares[ROW_7][COL_A]

	blackKing1 := &King{color: BLACK}
	e8 := board1.Squares[ROW_8][COL_E]

	board1.SetPiece(whiteKing1, e3)
	board1.SetPiece(blackKing1, e8)
	board1.SetPiece(blackBishop1, a7)

	board2 := New()

	whiteKing2 := &King{color: WHITE}
	e4 := board2.Squares[ROW_4][COL_E]

	blackKing2 := &King{color: BLACK}
	d8 := board2.Squares[ROW_8][COL_D]

	whiteRook := &Rook{color: WHITE}
	a8 := board2.Squares[ROW_8][COL_A]

	board2.SetPiece(whiteKing2, e4)
	board2.SetPiece(blackKing2, d8)
	board2.SetPiece(whiteRook, a8)

	board3 := New()

	whiteKing3 := &King{color: WHITE}
	a1 := board3.Squares[ROW_1][COL_A]

	blackKing3 := &King{color: BLACK}
	h8 := board3.Squares[ROW_8][COL_H]

	blackPawn3 := &Pawn{color: BLACK, Moved: true}
	b2 := board3.Squares[ROW_2][COL_B]

	blackQueen3 := &Queen{color: BLACK}
	a7_3 := board3.Squares[ROW_7][COL_A]

	board3.SetPiece(whiteKing3, a1)
	board3.SetPiece(blackKing3, h8)
	board3.SetPiece(blackPawn3, b2)
	board3.SetPiece(blackQueen3, a7_3)

	board4 := New()

	whiteKing4 := &King{color: WHITE}
	c1 := board4.Squares[ROW_1][COL_C]

	blackKing4 := &King{color: BLACK}
	c8 := board4.Squares[ROW_8][COL_C]

	blackPawn4 := &Pawn{color: BLACK}
	b7 := board4.Squares[ROW_7][COL_B]

	whiteKnight4 := &Knight{color: WHITE}
	d6 := board4.Squares[ROW_6][COL_D]

	board4.SetPiece(whiteKing4, c1)
	board4.SetPiece(blackKing4, c8)
	board4.SetPiece(blackPawn4, b7)
	board4.SetPiece(whiteKnight4, d6)

	board5 := New()

	whiteKing5 := &King{color: WHITE}
	g1 := board5.Squares[ROW_1][COL_G]

	blackKing5 := &King{color: BLACK}
	f8 := board5.Squares[ROW_8][COL_F]

	whitePawn5 := &Pawn{color: WHITE}
	f2 := board5.Squares[ROW_2][COL_F]

	blackBishop5 := &Bishop{color: BLACK}
	b6 := board5.Squares[ROW_6][COL_B]

	board5.SetPiece(whiteKing5, g1)
	board5.SetPiece(blackKing5, f8)
	board5.SetPiece(whitePawn5, f2)
	board5.SetPiece(blackBishop5, b6)

	tests := []struct {
		turn              string
		input             *King
		board             *Board
		expectedCheck     bool
		expectednumChecks int
	}{
		{WHITE, whiteKing1, board1, true, 1},
		{BLACK, blackKing2, board2, true, 1},
		{WHITE, whiteKing3, board3, true, 2},
		{BLACK, blackKing4, board4, true, 1},
		{WHITE, whiteKing5, board5, false, 0},
	}
	for _, tt := range tests {
		tt.board.Evaluate(tt.turn)
		unsafes := tt.board.GetAttackedSquares(tt.input.color)
		checked, checkingPieces := tt.input.IsInCheck(unsafes)
		if checked != tt.expectedCheck {
			t.Fatalf("King is in check should be %t (%s)", tt.expectedCheck, tt.input.Square().Name)
		}
		if len(checkingPieces) != tt.expectednumChecks {
			for _, piece := range checkingPieces {
				fmt.Println("CHECKING PIECE: ", piece.Type())
				for sq, activity := range piece.ActiveSquares() {
					fmt.Printf("ACTIVE SQ: %s ACTIVITY: %s\n", sq.Name, activity)
				}
			}
			t.Fatalf("King should have %d checking pieces. Got %d", tt.expectednumChecks, len(checkingPieces))
		}

	}
}

func TestKingMove(t *testing.T) {
	var board = New()
	d3 := board.Squares[ROW_3][COL_D]
	e5 := board.Squares[ROW_5][COL_E]
	c5 := board.Squares[ROW_5][COL_C]
	e3 := board.Squares[ROW_3][COL_E]
	f3 := board.Squares[ROW_3][COL_F]
	e7 := board.Squares[ROW_7][COL_E]

	whiteKing := &King{color: WHITE}
	blackKing := &King{color: BLACK}
	whitePawn := &Pawn{color: WHITE}
	blackPawn := &Pawn{color: BLACK}
	whiteRook := &Rook{color: WHITE}
	blackBishop := &Bishop{color: BLACK}

	board.SetPiece(whiteKing, d3)
	board.SetPiece(blackKing, e5)
	board.SetPiece(blackPawn, c5)
	board.SetPiece(whitePawn, e3)
	board.SetPiece(whiteRook, f3)
	board.SetPiece(blackBishop, e7)

	tests := []struct {
		input    *Move
		expected string
	}{
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteKing,
				From:  d3,
				To:    board.Squares[ROW_4][COL_E],
			},
			"KING: D3 -> E4 is not a valid move",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackKing,
				From:  e5,
				To:    board.Squares[ROW_4][COL_E],
			},
			"KING: E5 -> E4 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteKing,
				From:  d3,
				To:    board.Squares[ROW_4][COL_C],
			},
			"KING: D3 -> C4",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackKing,
				From:  e5,
				To:    board.Squares[ROW_4][COL_E],
			},
			"KING: E5 -> E4",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackKing,
				From:  board.Squares[ROW_4][COL_E],
				To:    f3,
			},
			"KING TAKES ROOK: E4 -> F3",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteKing,
				From:  board.Squares[ROW_4][COL_C],
				To:    c5,
			},
			"KING: C4 -> C5 is not a valid move",
		},
	}

	for _, tt := range tests {
		board.Evaluate(tt.input.Turn)
		if tt.input.Piece.Type() != KING {
			t.Fatalf("Piece should be a %s. Got %s", KING, tt.input.Piece.Type())
		}

		receipt, err := board.MovePiece(tt.input)

		if receipt != tt.expected {
			t.Fatalf("receipt should be '%s'. Got '%s'", tt.expected, receipt)
		}
		if err == nil {
			if !testPieceHasMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should have moved", tt.input.Piece.Type())
			}
		} else {
			if !testPieceHasNotMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should not have moved", tt.input.Piece.Type())
			}
		}
	}
}

func TestQueenMove(t *testing.T) {
	var board = New()
	e3 := board.Squares[ROW_3][COL_E]
	b6 := board.Squares[ROW_6][COL_B]
	f4 := board.Squares[ROW_4][COL_F]
	d6 := board.Squares[ROW_6][COL_D]
	d7 := board.Squares[ROW_7][COL_D]
	a1 := board.Squares[ROW_1][COL_A]

	whiteQueen := &Queen{color: WHITE}
	blackQueen := &Queen{color: BLACK}
	whitePawn := &Pawn{color: WHITE}
	blackPawn := &Pawn{color: BLACK}
	blackKing := &King{color: BLACK}
	whiteKing := &King{color: WHITE}

	board.SetPiece(whiteQueen, e3)
	board.SetPiece(blackQueen, b6)
	board.SetPiece(whitePawn, f4)
	board.SetPiece(blackPawn, d6)
	board.SetPiece(blackKing, d7)
	board.SetPiece(whiteKing, a1)

	tests := []struct {
		input    *Move
		expected string
	}{
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteQueen,
				From:  e3,
				To:    board.Squares[ROW_6][COL_H],
			},
			"QUEEN: E3 -> H6 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteQueen,
				From:  e3,
				To:    board.Squares[ROW_3][COL_B],
			},
			"QUEEN: E3 -> B3",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackQueen,
				From:  b6,
				To:    e3,
			},
			"QUEEN: B6 -> E3",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteQueen,
				From:  board.Squares[ROW_3][COL_B],
				To:    b6,
			},
			"QUEEN: B3 -> B6",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackQueen,
				From:  e3,
				To:    f4,
			},
			"QUEEN TAKES PAWN: E3 -> F4",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteQueen,
				From:  b6,
				To:    d6,
			},
			"QUEEN TAKES PAWN: B6 -> D6",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackQueen,
				From:  f4,
				To:    d6,
			},
			"QUEEN TAKES QUEEN: F4 -> D6",
		},
	}

	for _, tt := range tests {
		board.Evaluate(tt.input.Turn)
		if tt.input.Piece.Type() != QUEEN {
			t.Fatalf("Piece should be a %s. Got %s", QUEEN, tt.input.Piece.Type())
		}

		receipt, err := board.MovePiece(tt.input)

		if receipt != tt.expected {
			t.Fatalf("receipt should be '%s'. Got '%s'", tt.expected, receipt)
		}
		if err == nil {
			if !testPieceHasMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should have moved", tt.input.Piece.Type())
			}
		} else {
			if !testPieceHasNotMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should not have moved", tt.input.Piece.Type())
			}
		}
	}
}

func TestMoveRook(t *testing.T) {
	var board = New()
	e4 := board.Squares[ROW_4][COL_E]
	b4 := board.Squares[ROW_4][COL_B]
	h4 := board.Squares[ROW_4][COL_H]
	b7 := board.Squares[ROW_7][COL_B]
	f1 := board.Squares[ROW_1][COL_F]
	d8 := board.Squares[ROW_8][COL_D]

	whiteRook1 := &Rook{color: WHITE}
	whiteRook2 := &Rook{color: WHITE}
	blackRook1 := &Rook{color: BLACK}
	blackRook2 := &Rook{color: BLACK}

	whiteKing := &King{color: WHITE}
	blackKing := &King{color: BLACK}

	board.SetPiece(whiteRook1, e4)
	board.SetPiece(whiteRook2, b4)
	board.SetPiece(blackRook1, h4)
	board.SetPiece(blackRook2, b7)
	board.SetPiece(whiteKing, f1)
	board.SetPiece(blackKing, d8)

	tests := []struct {
		input    *Move
		expected string
	}{
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteRook1,
				From:  e4,
				To:    board.Squares[ROW_4][COL_A],
			},
			"ROOK: E4 -> A4 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteRook1,
				From:  e4,
				To:    h4,
			},
			"ROOK TAKES ROOK: E4 -> H4",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackRook2,
				From:  b7,
				To:    board.Squares[ROW_3][COL_B],
			},
			"ROOK: B7 -> B3 is not a valid move",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackRook2,
				From:  b7,
				To:    board.Squares[ROW_5][COL_B],
			},
			"ROOK: B7 -> B5",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackRook2,
				From:  board.Squares[ROW_5][COL_B],
				To:    b4,
			},
			"ROOK TAKES ROOK: B5 -> B4",
		},
	}

	for _, tt := range tests {
		board.Evaluate(tt.input.Turn)
		if tt.input.Piece.Type() != ROOK {
			t.Fatalf("Piece should be a %s. Got %s", ROOK, tt.input.Piece.Type())
		}

		receipt, err := board.MovePiece(tt.input)

		if receipt != tt.expected {
			t.Fatalf("receipt should be '%s'. Got '%s'", tt.expected, receipt)
		}
		if err == nil {
			if !testPieceHasMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should have moved", tt.input.Piece.Type())
			}
		} else {
			if !testPieceHasNotMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should not have moved", tt.input.Piece.Type())
			}
		}
	}
}

func TestMoveBishop(t *testing.T) {
	var board = New()
	board.SetupPieces()

	blackQueensBishop := board.Squares[ROW_8][COL_C].Piece
	blackKingsBishop := board.Squares[ROW_8][COL_F].Piece
	whiteQueensBishop := board.Squares[ROW_1][COL_C].Piece
	//whiteKingsBishop := board.Squares[ROW_1][COL_F].Piece

	whiteCPawn := board.Squares[ROW_2][COL_C].Piece
	whiteDPawn := board.Squares[ROW_2][COL_D].Piece
	whiteEPawn := board.Squares[ROW_2][COL_E].Piece
	whiteFPawn := board.Squares[ROW_2][COL_F].Piece
	blackDPawn := board.Squares[ROW_7][COL_D].Piece
	blackEPawn := board.Squares[ROW_7][COL_E].Piece

	moveNonTargetPiece(blackDPawn, blackDPawn.Square(), board.Squares[ROW_5][COL_D], board)
	moveNonTargetPiece(blackEPawn, blackEPawn.Square(), board.Squares[ROW_5][COL_E], board)
	moveNonTargetPiece(whiteCPawn, whiteCPawn.Square(), board.Squares[ROW_4][COL_C], board)
	moveNonTargetPiece(whiteDPawn, whiteDPawn.Square(), board.Squares[ROW_3][COL_D], board)
	moveNonTargetPiece(whiteEPawn, whiteEPawn.Square(), board.Squares[ROW_4][COL_E], board)
	moveNonTargetPiece(whiteFPawn, whiteFPawn.Square(), board.Squares[ROW_4][COL_F], board)

	tests := []struct {
		input    *Move
		expected string
	}{
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteQueensBishop,
				From:  whiteQueensBishop.Square(),
				To:    board.Squares[ROW_3][COL_A],
			},
			"BISHOP: C1 -> A3 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteQueensBishop,
				From:  whiteQueensBishop.Square(),
				To:    board.Squares[ROW_4][COL_F],
			},
			"BISHOP: C1 -> F4 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteQueensBishop,
				From:  whiteQueensBishop.Square(),
				To:    board.Squares[ROW_3][COL_E],
			},
			"BISHOP: C1 -> E3",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackKingsBishop,
				From:  blackKingsBishop.Square(),
				To:    board.Squares[ROW_5][COL_C],
			},
			"BISHOP: F8 -> C5",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteQueensBishop,
				From:  board.Squares[ROW_3][COL_E],
				To:    board.Squares[ROW_5][COL_C],
			},
			"BISHOP TAKES BISHOP: E3 -> C5",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackQueensBishop,
				From:  blackQueensBishop.Square(),
				To:    board.Squares[ROW_5][COL_F],
			},
			"BISHOP: C8 -> F5",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackQueensBishop,
				From:  board.Squares[ROW_5][COL_F],
				To:    board.Squares[ROW_4][COL_E],
			},
			"BISHOP TAKES PAWN: F5 -> E4",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackQueensBishop,
				From:  board.Squares[ROW_4][COL_E],
				To:    board.Squares[ROW_6][COL_G],
			},
			"BISHOP: E4 -> G6",
		},
	}

	for _, tt := range tests {
		board.Evaluate(tt.input.Turn)
		if tt.input.Piece.Type() != BISHOP {
			t.Fatalf("Piece should be a %s. Got %s", BISHOP, tt.input.Piece.Type())
		}

		receipt, err := board.MovePiece(tt.input)

		if receipt != tt.expected {
			t.Fatalf("receipt should be '%s'. Got '%s'", tt.expected, receipt)
		}
		if err == nil {
			if !testPieceHasMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should have moved", tt.input.Piece.Type())
			}
		} else {
			if !testPieceHasNotMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should not have moved", tt.input.Piece.Type())
			}
		}
	}
}

func TestMoveKnight(t *testing.T) {
	var board = New()
	board.SetupPieces()

	blackQueensKnight := board.Squares[ROW_8][COL_B].Piece
	blackKingsKnight := board.Squares[ROW_8][COL_G].Piece
	whiteQueensKnight := board.Squares[ROW_1][COL_B].Piece
	whiteKingsKnight := board.Squares[ROW_1][COL_G].Piece

	moveNonTargetPiece(blackKingsKnight, blackKingsKnight.Square(), board.Squares[ROW_6][COL_F], board)
	moveNonTargetPiece(blackKingsKnight, blackKingsKnight.Square(), board.Squares[ROW_4][COL_E], board)
	moveNonTargetPiece(whiteQueensKnight, whiteQueensKnight.Square(), board.Squares[ROW_3][COL_C], board)

	tests := []struct {
		input    *Move
		expected string
	}{
		{
			&Move{
				Turn:  BLACK,
				Piece: blackQueensKnight,
				From:  blackQueensKnight.Square(),
				To:    board.Squares[ROW_7][COL_D],
			},
			"KNIGHT: B8 -> D7 is not a valid move",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackQueensKnight,
				From:  blackQueensKnight.Square(),
				To:    board.Squares[ROW_1][COL_H],
			},
			"KNIGHT: B8 -> H1 is not a valid move",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackQueensKnight,
				From:  blackQueensKnight.Square(),
				To:    board.Squares[ROW_6][COL_C],
			},
			"KNIGHT: B8 -> C6",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteKingsKnight,
				From:  whiteKingsKnight.Square(),
				To:    board.Squares[ROW_3][COL_F],
			},
			"KNIGHT: G1 -> F3",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackQueensKnight,
				From:  board.Squares[ROW_6][COL_C],
				To:    board.Squares[ROW_4][COL_D],
			},
			"KNIGHT: C6 -> D4",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteKingsKnight,
				From:  board.Squares[ROW_3][COL_F],
				To:    board.Squares[ROW_4][COL_D],
			},
			"KNIGHT TAKES KNIGHT: F3 -> D4",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackKingsKnight,
				From:  blackKingsKnight.Square(),
				To:    whiteQueensKnight.Square(),
			},
			"KNIGHT TAKES KNIGHT: E4 -> C3",
		},
	}

	for _, tt := range tests {
		board.Evaluate(tt.input.Turn)
		if tt.input.Piece.Type() != KNIGHT {
			t.Fatalf("Piece should be a %s. Got %s", KNIGHT, tt.input.Piece.Type())
		}

		receipt, err := board.MovePiece(tt.input)

		if receipt != tt.expected {
			t.Fatalf("receipt should be '%s'. Got '%s'", tt.expected, receipt)
		}
		if err == nil {
			if !testPieceHasMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should have moved", tt.input.Piece.Type())
			}
		} else {
			if !testPieceHasNotMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should not have moved", tt.input.Piece.Type())
			}
		}
	}

}

func TestFreeMovePawn(t *testing.T) {
	var board = New()
	board.SetupPieces()

	whiteAPawn := board.Squares[ROW_2][COL_A].Piece
	whiteBPawn := board.Squares[ROW_2][COL_B].Piece
	blackAPawn := board.Squares[ROW_7][COL_A].Piece
	blackBPawn := board.Squares[ROW_7][COL_B].Piece

	tests := []struct {
		input    *Move
		expected string
	}{
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteAPawn,
				From:  board.Squares[ROW_2][COL_A],
				To:    board.Squares[ROW_3][COL_A],
			},
			"PAWN: A2 -> A3",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteAPawn,
				From:  board.Squares[ROW_3][COL_A],
				To:    board.Squares[ROW_4][COL_A],
			},
			"PAWN: A3 -> A4",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteBPawn,
				From:  board.Squares[ROW_2][COL_B],
				To:    board.Squares[ROW_4][COL_B],
			},
			"PAWN: B2 -> B4",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteBPawn,
				From:  board.Squares[ROW_4][COL_B],
				To:    board.Squares[ROW_6][COL_B],
			},
			"PAWN: B4 -> B6 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteBPawn,
				From:  board.Squares[ROW_4][COL_B],
				To:    board.Squares[ROW_5][COL_B],
			},
			"PAWN: B4 -> B5",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackBPawn,
				From:  board.Squares[ROW_7][COL_B],
				To:    board.Squares[ROW_6][COL_B],
			},
			"PAWN: B7 -> B6",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackBPawn,
				From:  board.Squares[ROW_6][COL_B],
				To:    board.Squares[ROW_5][COL_B],
			},
			"PAWN: B6 -> B5 is not a valid move",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackAPawn,
				From:  board.Squares[ROW_7][COL_A],
				To:    board.Squares[ROW_5][COL_A],
			},
			"PAWN: A7 -> A5",
		},
	}

	for _, tt := range tests {
		board.Evaluate(tt.input.Turn)
		if tt.input.Piece.Type() != PAWN {
			t.Fatalf("Piece should be a %s. Got %s", PAWN, tt.input.Piece.Type())
		}
		receipt, err := board.MovePiece(tt.input)

		if receipt != tt.expected {
			t.Fatalf("receipt should be: '%s'. Got %s", tt.expected, receipt)
		}
		if err == nil {
			if !testPieceHasMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should have moved", tt.input.Piece.Type())
			}
		} else {
			if !testPieceHasNotMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should not have moved", tt.input.Piece.Type())
			}
		}
	}
}

func TestCaptureMovePawn(t *testing.T) {
	var board = New()
	board.SetupPieces()

	whiteAPawn := board.Squares[ROW_2][COL_A].Piece.(*Pawn)
	whiteBPawn := board.Squares[ROW_2][COL_B].Piece.(*Pawn)
	whiteDPawn := board.Squares[ROW_2][COL_D].Piece.(*Pawn)
	whiteEPawn := board.Squares[ROW_2][COL_E].Piece.(*Pawn)

	blackAPawn := board.Squares[ROW_7][COL_A].Piece.(*Pawn)
	blackBPawn := board.Squares[ROW_7][COL_B].Piece.(*Pawn)
	blackDPawn := board.Squares[ROW_7][COL_D].Piece.(*Pawn)
	blackEPawn := board.Squares[ROW_7][COL_E].Piece.(*Pawn)

	moveNonTargetPiece(blackEPawn, blackEPawn.Square(), board.Squares[ROW_5][COL_E], board)
	moveNonTargetPiece(whiteDPawn, whiteDPawn.Square(), board.Squares[ROW_4][COL_D], board)
	moveNonTargetPiece(whiteEPawn, whiteEPawn.Square(), board.Squares[ROW_4][COL_E], board)
	moveNonTargetPiece(blackDPawn, blackDPawn.Square(), board.Squares[ROW_5][COL_D], board)

	tests := []struct {
		input    *Move
		expected string
	}{
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteAPawn,
				From:  board.Squares[ROW_2][COL_A],
				To:    board.Squares[ROW_3][COL_B],
			},
			"PAWN: A2 -> B3 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteBPawn,
				From:  board.Squares[ROW_2][COL_B],
				To:    board.Squares[ROW_3][COL_A],
			},
			"PAWN: B2 -> A3 is not a valid move",
		},
		{
			&Move{Turn: BLACK,
				Piece: blackBPawn,
				From:  board.Squares[ROW_7][COL_B],
				To:    board.Squares[ROW_6][COL_A],
			},
			"PAWN: B7 -> A6 is not a valid move",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackAPawn,
				From:  board.Squares[ROW_7][COL_A],
				To:    board.Squares[ROW_6][COL_B],
			},
			"PAWN: A7 -> B6 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteEPawn,
				From:  whiteEPawn.Square(),
				To:    blackDPawn.Square(),
			},
			"PAWN TAKES PAWN: E4 -> D5",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackEPawn,
				From:  blackEPawn.Square(),
				To:    whiteDPawn.Square(),
			},
			"PAWN TAKES PAWN: E5 -> D4",
		},
	}
	for _, tt := range tests {
		board.Evaluate(tt.input.Turn)
		if tt.input.Piece.Type() != PAWN {
			t.Fatalf("Piece should be a %s. Got %s", PAWN, tt.input.Piece.Type())
		}
		receipt, err := board.MovePiece(tt.input)

		if receipt != tt.expected {
			t.Fatalf("receipt should be: '%s'. Got %s", tt.expected, receipt)
		}
		if err == nil {
			if tt.input.Type != CAPTURE {
				t.Fatalf("Move Type should be CAPTURE. Got %s", tt.input.Type)
			}
			if !testPieceHasMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should have moved %s -> %s", tt.input.Piece.Type(), tt.input.From.Name, tt.input.To.Name)
			}
		} else {
			if tt.input.Type != "" {
				t.Fatalf("Move Type should be ''. Got '%s'", tt.input.Type)
			}
			if !testPieceHasNotMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should not have moved", tt.input.Piece.Type())
			}
		}
	}
}

func TestEnPassantAsBlack(t *testing.T) {
	board := New()
	a6 := board.Squares[ROW_6][COL_A]
	b2 := board.Squares[ROW_2][COL_B]
	g4 := board.Squares[ROW_4][COL_G]
	h2 := board.Squares[ROW_2][COL_H]
	h4 := board.Squares[ROW_4][COL_H]
	f4 := board.Squares[ROW_4][COL_F]
	f2 := board.Squares[ROW_2][COL_F]

	blackKing := &King{color: BLACK}
	board.SetPiece(blackKing, a6)

	whiteKing := &King{color: WHITE}
	board.SetPiece(whiteKing, b2)

	blackGPawn := &Pawn{color: BLACK}
	board.SetPiece(blackGPawn, g4)

	whiteFPawn := &Pawn{color: WHITE}
	board.SetPiece(whiteFPawn, f2)

	whiteHPawn := &Pawn{color: WHITE}
	board.SetPiece(whiteHPawn, h2)

	moveNonTargetPiece(whiteFPawn, f2, f4, board)
	moveNonTargetPiece(whiteHPawn, h2, h4, board)

	tests := []struct {
		input    *Move
		expected string
	}{
		{
			&Move{
				Turn:  BLACK,
				Piece: blackGPawn,
				From:  g4,
				To:    board.Squares[ROW_3][COL_F],
			},
			"PAWN: G4 -> F3 is not a valid move",
		},
		{
			&Move{
				Turn:  BLACK,
				Piece: blackGPawn,
				From:  g4,
				To:    board.Squares[ROW_3][COL_H],
			},
			"PAWN TAKES PAWN (EN PASSANT): G4 -> H3",
		},
	}

	for _, tt := range tests {
		board.Evaluate(tt.input.Turn)
		if tt.input.Piece.Type() != PAWN {
			t.Fatalf("Piece should be a %s. Got %s", PAWN, tt.input.Piece.Type())
		}
		receipt, err := board.MovePiece(tt.input)

		if receipt != tt.expected {
			t.Fatalf("receipt should be: '%s'. Got %s", tt.expected, receipt)
		}
		if err == nil {
			if tt.input.Type != EN_PASSANT {
				t.Fatalf("Move Type should be EN_PASSANT. Got %s", tt.input.Type)
			}
			if !testPieceHasMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should have moved %s -> %s", tt.input.Piece.Type(), tt.input.From.Name, tt.input.To.Name)
			}
		} else {
			if tt.input.Type != "" {
				t.Fatalf("Move Type should be ''. Got '%s'", tt.input.Type)
			}
			if !testPieceHasNotMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should not have moved", tt.input.Piece.Type())
			}
		}
	}
}

func TestEnPassantAsWhite(t *testing.T) {
	var board = New()

	a1 := board.Squares[ROW_1][COL_A]
	h8 := board.Squares[ROW_8][COL_H]
	d5 := board.Squares[ROW_5][COL_D]
	e5 := board.Squares[ROW_5][COL_E]
	e7 := board.Squares[ROW_7][COL_E]
	c5 := board.Squares[ROW_5][COL_C]
	c7 := board.Squares[ROW_7][COL_C]

	whiteKing := &King{color: WHITE}
	board.SetPiece(whiteKing, a1)

	blackKing := &King{color: BLACK}
	board.SetPiece(blackKing, h8)

	whiteDPawn := &Pawn{color: WHITE}
	board.SetPiece(whiteDPawn, d5)

	blackEPawn := &Pawn{color: BLACK}
	board.SetPiece(blackEPawn, e7)

	blackCPawn := &Pawn{color: BLACK}
	board.SetPiece(blackCPawn, c7)

	moveNonTargetPiece(blackEPawn, e7, e5, board)
	moveNonTargetPiece(blackCPawn, c7, c5, board)

	tests := []struct {
		input    *Move
		expected string
	}{
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteDPawn,
				From:  d5,
				To:    board.Squares[ROW_6][COL_E],
			},
			"PAWN: D5 -> E6 is not a valid move",
		},
		{
			&Move{
				Turn:  WHITE,
				Piece: whiteDPawn,
				From:  d5,
				To:    board.Squares[ROW_6][COL_C],
			},
			"PAWN TAKES PAWN (EN PASSANT): D5 -> C6",
		},
	}

	for _, tt := range tests {
		board.Evaluate(tt.input.Turn)
		if tt.input.Piece.Type() != PAWN {
			t.Fatalf("Piece should be a %s. Got %s", PAWN, tt.input.Piece.Type())
		}
		receipt, err := board.MovePiece(tt.input)

		if receipt != tt.expected {
			t.Fatalf("receipt should be: '%s'. Got %s", tt.expected, receipt)
		}
		if err == nil {
			if tt.input.Type != EN_PASSANT {
				t.Fatalf("Move Type should be EN_PASSANT. Got %s", tt.input.Type)
			}
			if !testPieceHasMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should have moved %s -> %s", tt.input.Piece.Type(), tt.input.From.Name, tt.input.To.Name)
			}
		} else {
			if tt.input.Type != "" {
				t.Fatalf("Move Type should be ''. Got '%s'", tt.input.Type)
			}
			if !testPieceHasNotMoved(tt.input.Piece, tt.input.From, tt.input.To) {
				t.Fatalf("%s should not have moved", tt.input.Piece.Type())
			}
		}
	}
}

func testPieceHasMoved(piece Piece, fromSquare *Square, toSquare *Square) bool {
	return fromSquare.IsEmpty() && toSquare.Piece == piece && piece.Square().Name == toSquare.Name
}

func testPieceHasNotMoved(piece Piece, fromSquare *Square, toSquare *Square) bool {
	return toSquare.Piece != piece && fromSquare.Piece == piece && piece.Square() == fromSquare
}

func testEmptySquare(square *Square) bool {
	return square.IsEmpty()
}

func testSquareOccupiedByPiece(square *Square, piece Piece) bool {
	return square.Piece == piece
}

func moveNonTargetPiece(piece Piece, from *Square, to *Square, board *Board) {
	move := &Move{Piece: piece, From: from, To: to}
	board.MovePiece(move)
	board.Evaluate(piece.Color())
}
