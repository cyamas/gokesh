package board

import (
	"testing"
)

func TestPawnPromotion(t *testing.T) {
	board := New()

	c7 := board.Squares[ROW_7][COL_C]
	e7 := board.Squares[ROW_7][COL_E]
	d2 := board.Squares[ROW_2][COL_D]
	f2 := board.Squares[ROW_2][COL_F]
	b8 := board.Squares[ROW_8][COL_B]
	g1 := board.Squares[ROW_1][COL_G]

	whiteEPawn := &Pawn{color: WHITE, Moved: true}
	whiteCPawn := &Pawn{color: WHITE, Moved: true}
	blackDPawn := &Pawn{color: BLACK, Moved: true}
	blackFPawn := &Pawn{color: BLACK, Moved: true}
	blackQueen := &Queen{color: BLACK}
	whiteBishop := &Bishop{color: WHITE}

	board.SetPiece(whiteEPawn, e7)
	board.SetPiece(whiteCPawn, c7)
	board.SetPiece(blackDPawn, d2)
	board.SetPiece(blackFPawn, f2)
	board.SetPiece(blackQueen, b8)
	board.SetPiece(whiteBishop, g1)

	tests := []struct {
		input             *Move
		expectedReceipt   string
		expectedPromotion string
	}{
		{
			&Move{
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
		receipt, err := board.Move(tt.input)
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
	whiteFPawn := &Pawn{color: WHITE}
	shortWhiteRook := &Rook{color: WHITE, CastleSq: shortBoard.Squares[ROW_1][COL_F]}
	shortBlackRook := &Rook{color: BLACK, CastleSq: shortBoard.Squares[ROW_8][COL_F]}
	shortBoard.SetPiece(shortWhiteRook, shortBoard.Squares[ROW_1][COL_H])
	shortBoard.SetPiece(shortBlackRook, shortBoard.Squares[ROW_8][COL_H])
	shortBoard.SetPiece(whiteFPawn, shortBoard.Squares[ROW_2][COL_F])

	longBoard := New()
	longBlackSq := longBoard.Squares[ROW_8][COL_C]
	longWhiteSq := longBoard.Squares[ROW_1][COL_C]
	whiteDPawn := &Pawn{color: WHITE}
	longWhiteRook := &Rook{color: WHITE, CastleSq: longBoard.Squares[ROW_1][COL_D]}
	longBlackRook := &Rook{color: BLACK, CastleSq: longBoard.Squares[ROW_8][COL_D]}
	longBoard.SetPiece(longWhiteRook, longBoard.Squares[ROW_1][COL_A])
	longBoard.SetPiece(longBlackRook, longBoard.Squares[ROW_8][COL_A])
	longBoard.SetPiece(whiteDPawn, longBoard.Squares[ROW_2][COL_D])

	checkedBoard := New()
	checkedBlackSq := checkedBoard.Squares[ROW_8][COL_C]
	checkedWhiteSq := checkedBoard.Squares[ROW_1][COL_C]
	checkingWhiteFPawn := &Pawn{color: WHITE, Moved: true}
	checkingBlackDKnight := &Knight{color: BLACK}
	checkedWhiteRook := &Rook{color: WHITE, CastleSq: checkedBoard.Squares[ROW_1][COL_D]}
	checkedBlackRook := &Rook{color: BLACK, CastleSq: checkedBoard.Squares[ROW_8][COL_D]}
	checkedBoard.SetPiece(checkedWhiteRook, checkedBoard.Squares[ROW_1][COL_A])
	checkedBoard.SetPiece(checkedBlackRook, checkedBoard.Squares[ROW_8][COL_A])
	checkedBoard.SetPiece(checkingWhiteFPawn, checkedBoard.Squares[ROW_7][COL_F])
	checkedBoard.SetPiece(checkingBlackDKnight, checkedBoard.Squares[ROW_3][COL_F])

	blockedBoard := New()
	blockedBlackSq := blockedBoard.Squares[ROW_8][COL_C]
	blockedWhiteSq := blockedBoard.Squares[ROW_1][COL_C]
	blockedWhiteRook := &Rook{color: WHITE, CastleSq: blockedBoard.Squares[ROW_1][COL_D]}
	blockedBlackRook := &Rook{color: BLACK, CastleSq: blockedBoard.Squares[ROW_8][COL_D]}
	blockedBoard.SetPiece(blockedWhiteRook, blockedBoard.Squares[ROW_1][COL_A])
	blockedBoard.SetPiece(blockedBlackRook, blockedBoard.Squares[ROW_8][COL_A])
	blockingWhiteKnight := &Knight{color: WHITE}
	blockingBlackKnight := &Knight{color: BLACK}
	blockedBoard.SetPiece(blockingWhiteKnight, blockedBoard.Squares[ROW_1][COL_C])
	blockedBoard.SetPiece(blockingBlackKnight, blockedBoard.Squares[ROW_8][COL_D])

	enemyBlockedBoard := New()
	enemyBlockedWhiteSq := enemyBlockedBoard.Squares[ROW_1][COL_C]
	enemyBlockedBlackSq := enemyBlockedBoard.Squares[ROW_8][COL_G]
	enemyBlockedWhiteRook := &Rook{color: WHITE, CastleSq: enemyBlockedBoard.Squares[ROW_1][COL_D]}
	enemyBlockedBlackRook := &Rook{color: BLACK, CastleSq: enemyBlockedBoard.Squares[ROW_8][COL_F]}
	enemyBlockedBoard.SetPiece(enemyBlockedWhiteRook, enemyBlockedBoard.Squares[ROW_1][COL_A])
	enemyBlockedBoard.SetPiece(enemyBlockedBlackRook, enemyBlockedBoard.Squares[ROW_8][COL_H])
	enemyBlackBishop := &Bishop{color: BLACK}
	enemyWhiteBishop := &Bishop{color: WHITE}
	enemyBlockedBoard.SetPiece(enemyBlackBishop, enemyBlockedBoard.Squares[ROW_5][COL_G])
	enemyBlockedBoard.SetPiece(enemyWhiteBishop, enemyBlockedBoard.Squares[ROW_2][COL_A])

	kingMovedBoard := New()
	kingMovedWhiteSq := kingMovedBoard.Squares[ROW_1][COL_C]
	kingMovedBlackSq := kingMovedBoard.Squares[ROW_8][COL_C]
	kingMovedWhiteRook := &Rook{color: WHITE, CastleSq: kingMovedBoard.Squares[ROW_1][COL_D]}
	kingMovedBlackRook := &Rook{color: BLACK, CastleSq: kingMovedBoard.Squares[ROW_8][COL_D]}
	kingMovedBoard.SetPiece(kingMovedWhiteRook, kingMovedBoard.Squares[ROW_1][COL_A])
	kingMovedBoard.SetPiece(kingMovedBlackRook, kingMovedBoard.Squares[ROW_8][COL_A])

	rookMovedBoard := New()
	rookMovedWhiteSq := rookMovedBoard.Squares[ROW_1][COL_G]
	rookMovedBlackSq := rookMovedBoard.Squares[ROW_8][COL_G]
	rookMovedWhiteRook := &Rook{color: WHITE, Moved: true, CastleSq: rookMovedBoard.Squares[ROW_1][COL_F]}
	rookMovedBlackRook := &Rook{color: BLACK, Moved: true, CastleSq: rookMovedBoard.Squares[ROW_8][COL_F]}
	shortBoard.SetPiece(rookMovedWhiteRook, rookMovedBoard.Squares[ROW_1][COL_H])
	shortBoard.SetPiece(rookMovedBlackRook, rookMovedBoard.Squares[ROW_8][COL_H])

	tests := []struct {
		board    *Board
		king     *King
		rook     *Rook
		castleSq *Square
		expected string
	}{
		{shortBoard, &King{color: BLACK}, shortBlackRook, shortBlackSq, "KING CASTLES SHORT"},
		{shortBoard, &King{color: WHITE}, shortWhiteRook, shortWhiteSq, "KING CASTLES SHORT"},
		{longBoard, &King{color: BLACK}, longBlackRook, longBlackSq, "KING CASTLES LONG"},
		{longBoard, &King{color: WHITE}, longWhiteRook, longWhiteSq, "KING CASTLES LONG"},
		{checkedBoard, &King{color: BLACK}, checkedBlackRook, checkedBlackSq, "KING: E8 -> C8 is not a valid move"},
		{checkedBoard, &King{color: WHITE}, checkedWhiteRook, checkedWhiteSq, "KING: E1 -> C1 is not a valid move"},
		{blockedBoard, &King{color: BLACK}, blockedBlackRook, blockedBlackSq, "KING: E8 -> C8 is not a valid move"},
		{blockedBoard, &King{color: WHITE}, blockedWhiteRook, blockedWhiteSq, "KING: E1 -> C1 is not a valid move"},
		{enemyBlockedBoard, &King{color: BLACK}, enemyBlockedBlackRook, enemyBlockedBlackSq, "KING: E8 -> G8 is not a valid move"},
		{enemyBlockedBoard, &King{color: WHITE}, enemyBlockedWhiteRook, enemyBlockedWhiteSq, "KING: E1 -> C1 is not a valid move"},
		{kingMovedBoard, &King{color: BLACK, Moved: true}, kingMovedBlackRook, kingMovedBlackSq, "KING: E8 -> C8 is not a valid move"},
		{kingMovedBoard, &King{color: WHITE, Moved: true}, kingMovedWhiteRook, kingMovedWhiteSq, "KING: E1 -> C1 is not a valid move"},
		{rookMovedBoard, &King{color: BLACK}, rookMovedBlackRook, rookMovedBlackSq, "KING: E8 -> G8 is not a valid move"},
		{rookMovedBoard, &King{color: WHITE}, rookMovedWhiteRook, rookMovedWhiteSq, "KING: E1 -> G1 is not a valid move"},
	}
	for _, tt := range tests {
		ogRookSq := tt.rook.Square()
		var ogKingSq *Square
		if tt.king.color == BLACK {
			tt.board.SetPiece(tt.king, tt.board.Squares[ROW_8][COL_E])
			ogKingSq = tt.board.Squares[ROW_8][COL_E]
		} else {
			tt.board.SetPiece(tt.king, tt.board.Squares[ROW_1][COL_E])
			ogKingSq = tt.board.Squares[ROW_1][COL_E]
		}
		receipt, err := tt.board.Move(&Move{Piece: tt.king, From: tt.king.Square(), To: tt.castleSq})
		if receipt != tt.expected {
			t.Fatalf("Receipt should be '%s'. Got '%s'", tt.expected, receipt)
		}
		if err == nil {
			if !testPieceHasMoved(tt.king, ogKingSq, tt.castleSq) {
				t.Fatalf(
					"KING should have moved to %s. Got %s",
					tt.castleSq.Name,
					tt.king.Square().Name,
				)
			}
			if !testPieceHasMoved(tt.rook, ogRookSq, tt.rook.CastleSq) {
				t.Fatalf(
					"ROOK should have moved to %s. Got %s",
					tt.rook.CastleSq.Name,
					tt.rook.Square().Name,
				)
			}
		} else {
			if tt.king.Square() != ogKingSq {
				t.Fatalf("KING should not have moved")
			}
			if tt.rook.Square() != ogRookSq {
				t.Fatalf("ROOK should not have moved")
			}
		}
	}
}

func TestKingInCheck(t *testing.T) {
	var board = New()

	whiteEKing := &King{color: WHITE}
	e3 := board.Squares[ROW_3][COL_E]
	board.SetPiece(whiteEKing, e3)

	whiteBKing := &King{color: WHITE}
	b2 := board.Squares[ROW_2][COL_B]
	board.SetPiece(whiteBKing, b2)

	whiteHKing := &King{color: WHITE}
	h1 := board.Squares[ROW_1][COL_H]
	board.SetPiece(whiteHKing, h1)

	blackEKing := &King{color: BLACK}
	e8 := board.Squares[ROW_8][COL_E]
	board.SetPiece(blackEKing, e8)

	blackBKing := &King{color: BLACK}
	b7 := board.Squares[ROW_7][COL_B]
	board.SetPiece(blackBKing, b7)

	blackHKing := &King{color: BLACK}
	h6 := board.Squares[ROW_6][COL_H]
	board.SetPiece(blackHKing, h6)

	whiteKnight := &Knight{color: WHITE}
	f5 := board.Squares[ROW_5][COL_F]
	board.SetPiece(whiteKnight, f5)

	blackBishop := &Bishop{color: BLACK}
	d5 := board.Squares[ROW_5][COL_D]
	board.SetPiece(blackBishop, d5)

	whiteRook := &Rook{color: WHITE}
	a8 := board.Squares[ROW_8][COL_A]
	board.SetPiece(whiteRook, a8)

	blackQueen := &Queen{color: BLACK}
	h8 := board.Squares[ROW_8][COL_H]
	board.SetPiece(blackQueen, h8)

	whiteAPawn := &Pawn{color: WHITE}
	a6 := board.Squares[ROW_6][COL_A]
	board.SetPiece(whiteAPawn, a6)

	whiteDPawn := &Pawn{color: WHITE}
	d4 := board.Squares[ROW_4][COL_D]
	board.SetPiece(whiteDPawn, d4)

	tests := []struct {
		input    *King
		expected bool
	}{
		{whiteEKing, false},
		{whiteBKing, false},
		{blackEKing, true},
		{blackBKing, true},
		{whiteHKing, true},
		{blackHKing, true},
	}
	for _, tt := range tests {
		unsafes := board.GetAttackedSquares(tt.input.color)
		if tt.input.IsInCheck(unsafes) != tt.expected {
			t.Fatalf("King is in check should be %t", tt.expected)
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
				Piece: whiteKing,
				From:  d3,
				To:    board.Squares[ROW_4][COL_E],
			},
			"KING: D3 -> E4 is not a valid move",
		},
		{
			&Move{
				Piece: blackKing,
				From:  e5,
				To:    board.Squares[ROW_4][COL_E],
			},
			"KING: E5 -> E4 is not a valid move",
		},
		{
			&Move{
				Piece: whiteKing,
				From:  d3,
				To:    board.Squares[ROW_4][COL_C],
			},
			"KING: D3 -> C4",
		},
		{
			&Move{
				Piece: blackKing,
				From:  e5,
				To:    board.Squares[ROW_4][COL_E],
			},
			"KING: E5 -> E4",
		},
		{
			&Move{
				Piece: blackKing,
				From:  board.Squares[ROW_4][COL_E],
				To:    f3,
			},
			"KING TAKES ROOK: E4 -> F3",
		},
		{
			&Move{
				Piece: whiteKing,
				From:  board.Squares[ROW_4][COL_C],
				To:    c5,
			},
			"KING: C4 -> C5 is not a valid move",
		},
	}

	for _, tt := range tests {
		if tt.input.Piece.Type() != KING {
			t.Fatalf("Piece should be a %s. Got %s", KING, tt.input.Piece.Type())
		}

		receipt, err := board.Move(tt.input)

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

	whiteQueen := &Queen{color: WHITE, square: e3}
	e3.SetPiece(whiteQueen)

	blackQueen := &Queen{color: BLACK, square: b6}
	b6.SetPiece(blackQueen)

	whitePawn := &Pawn{color: WHITE, square: f4}
	f4.SetPiece(whitePawn)

	blackPawn := &Pawn{color: BLACK, square: d6}
	d6.SetPiece(blackPawn)

	tests := []struct {
		input    *Move
		expected string
	}{
		{
			&Move{
				Piece: whiteQueen,
				From:  e3,
				To:    board.Squares[ROW_6][COL_H],
			},
			"QUEEN: E3 -> H6 is not a valid move",
		},
		{
			&Move{
				Piece: whiteQueen,
				From:  e3,
				To:    board.Squares[ROW_3][COL_B],
			},
			"QUEEN: E3 -> B3",
		},
		{
			&Move{
				Piece: blackQueen,
				From:  b6,
				To:    e3,
			},
			"QUEEN: B6 -> E3",
		},
		{
			&Move{
				Piece: whiteQueen,
				From:  board.Squares[ROW_3][COL_B],
				To:    b6,
			},
			"QUEEN: B3 -> B6",
		},
		{
			&Move{
				Piece: blackQueen,
				From:  e3,
				To:    f4,
			},
			"QUEEN TAKES PAWN: E3 -> F4",
		},
		{
			&Move{
				Piece: whiteQueen,
				From:  b6,
				To:    d6,
			},
			"QUEEN TAKES PAWN: B6 -> D6",
		},
		{
			&Move{
				Piece: blackQueen,
				From:  f4,
				To:    d6,
			},
			"QUEEN TAKES QUEEN: F4 -> D6",
		},
	}

	for _, tt := range tests {
		if tt.input.Piece.Type() != QUEEN {
			t.Fatalf("Piece should be a %s. Got %s", QUEEN, tt.input.Piece.Type())
		}

		receipt, err := board.Move(tt.input)

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

	whiteRook1 := &Rook{square: e4, color: WHITE}
	e4.SetPiece(whiteRook1)

	whiteRook2 := &Rook{square: b4, color: WHITE}
	b4.SetPiece(whiteRook2)

	blackRook1 := &Rook{square: h4, color: BLACK}
	h4.SetPiece(blackRook1)

	blackRook2 := &Rook{square: b7, color: BLACK}
	b7.SetPiece(blackRook2)

	tests := []struct {
		input    *Move
		expected string
	}{
		{
			&Move{
				Piece: whiteRook1,
				From:  e4,
				To:    board.Squares[ROW_4][COL_A],
			},
			"ROOK: E4 -> A4 is not a valid move",
		},
		{
			&Move{
				Piece: whiteRook1,
				From:  e4,
				To:    h4,
			},
			"ROOK TAKES ROOK: E4 -> H4",
		},
		{
			&Move{
				Piece: blackRook2,
				From:  b7,
				To:    board.Squares[ROW_3][COL_B],
			},
			"ROOK: B7 -> B3 is not a valid move",
		},
		{
			&Move{
				Piece: blackRook2,
				From:  b7,
				To:    board.Squares[ROW_5][COL_B],
			},
			"ROOK: B7 -> B5",
		},
		{
			&Move{
				Piece: blackRook2,
				From:  board.Squares[ROW_5][COL_B],
				To:    b4,
			},
			"ROOK TAKES ROOK: B5 -> B4",
		},
	}

	for _, tt := range tests {
		if tt.input.Piece.Type() != ROOK {
			t.Fatalf("Piece should be a %s. Got %s", ROOK, tt.input.Piece.Type())
		}

		receipt, err := board.Move(tt.input)

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
				Piece: whiteQueensBishop,
				From:  whiteQueensBishop.Square(),
				To:    board.Squares[ROW_3][COL_A],
			},
			"BISHOP: C1 -> A3 is not a valid move",
		},
		{
			&Move{
				Piece: whiteQueensBishop,
				From:  whiteQueensBishop.Square(),
				To:    board.Squares[ROW_4][COL_F],
			},
			"BISHOP: C1 -> F4 is not a valid move",
		},
		{
			&Move{
				Piece: whiteQueensBishop,
				From:  whiteQueensBishop.Square(),
				To:    board.Squares[ROW_3][COL_E],
			},
			"BISHOP: C1 -> E3",
		},
		{
			&Move{
				Piece: blackKingsBishop,
				From:  blackKingsBishop.Square(),
				To:    board.Squares[ROW_5][COL_C],
			},
			"BISHOP: F8 -> C5",
		},
		{
			&Move{
				Piece: whiteQueensBishop,
				From:  board.Squares[ROW_3][COL_E],
				To:    board.Squares[ROW_5][COL_C],
			},
			"BISHOP TAKES BISHOP: E3 -> C5",
		},
		{
			&Move{
				Piece: blackQueensBishop,
				From:  blackQueensBishop.Square(),
				To:    board.Squares[ROW_5][COL_F],
			},
			"BISHOP: C8 -> F5",
		},
		{
			&Move{
				Piece: blackQueensBishop,
				From:  board.Squares[ROW_5][COL_F],
				To:    board.Squares[ROW_4][COL_E],
			},
			"BISHOP TAKES PAWN: F5 -> E4",
		},
		{
			&Move{
				Piece: blackQueensBishop,
				From:  board.Squares[ROW_4][COL_E],
				To:    board.Squares[ROW_6][COL_G],
			},
			"BISHOP: E4 -> G6",
		},
	}

	for _, tt := range tests {
		if tt.input.Piece.Type() != BISHOP {
			t.Fatalf("Piece should be a %s. Got %s", BISHOP, tt.input.Piece.Type())
		}

		receipt, err := board.Move(tt.input)

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
				Piece: blackQueensKnight,
				From:  blackQueensKnight.Square(),
				To:    board.Squares[ROW_7][COL_D],
			},
			"KNIGHT: B8 -> D7 is not a valid move",
		},
		{
			&Move{
				Piece: blackQueensKnight,
				From:  blackQueensKnight.Square(),
				To:    board.Squares[ROW_1][COL_H],
			},
			"KNIGHT: B8 -> H1 is not a valid move",
		},
		{
			&Move{
				Piece: blackQueensKnight,
				From:  blackQueensKnight.Square(),
				To:    board.Squares[ROW_6][COL_C],
			},
			"KNIGHT: B8 -> C6",
		},
		{
			&Move{
				Piece: whiteKingsKnight,
				From:  whiteKingsKnight.Square(),
				To:    board.Squares[ROW_3][COL_F],
			},
			"KNIGHT: G1 -> F3",
		},
		{
			&Move{
				Piece: blackQueensKnight,
				From:  board.Squares[ROW_6][COL_C],
				To:    board.Squares[ROW_4][COL_D],
			},
			"KNIGHT: C6 -> D4",
		},
		{
			&Move{
				Piece: whiteKingsKnight,
				From:  board.Squares[ROW_3][COL_F],
				To:    board.Squares[ROW_4][COL_D],
			},
			"KNIGHT TAKES KNIGHT: F3 -> D4",
		},
		{
			&Move{
				Piece: blackKingsKnight,
				From:  blackKingsKnight.Square(),
				To:    whiteQueensKnight.Square(),
			},
			"KNIGHT TAKES KNIGHT: E4 -> C3",
		},
	}

	for _, tt := range tests {
		if tt.input.Piece.Type() != KNIGHT {
			t.Fatalf("Piece should be a %s. Got %s", KNIGHT, tt.input.Piece.Type())
		}

		receipt, err := board.Move(tt.input)

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
				Piece: whiteAPawn,
				From:  board.Squares[ROW_2][COL_A],
				To:    board.Squares[ROW_3][COL_A],
			},
			"PAWN: A2 -> A3",
		},
		{
			&Move{
				Piece: whiteAPawn,
				From:  board.Squares[ROW_3][COL_A],
				To:    board.Squares[ROW_4][COL_A],
			},
			"PAWN: A3 -> A4",
		},
		{
			&Move{
				Piece: whiteBPawn,
				From:  board.Squares[ROW_2][COL_B],
				To:    board.Squares[ROW_4][COL_B],
			},
			"PAWN: B2 -> B4",
		},
		{
			&Move{
				Piece: whiteBPawn,
				From:  board.Squares[ROW_4][COL_B],
				To:    board.Squares[ROW_6][COL_B],
			},
			"PAWN: B4 -> B6 is not a valid move",
		},
		{
			&Move{
				Piece: whiteBPawn,
				From:  board.Squares[ROW_4][COL_B],
				To:    board.Squares[ROW_5][COL_B],
			},
			"PAWN: B4 -> B5",
		},
		{
			&Move{
				Piece: blackBPawn,
				From:  board.Squares[ROW_7][COL_B],
				To:    board.Squares[ROW_6][COL_B],
			},
			"PAWN: B7 -> B6",
		},
		{
			&Move{
				Piece: blackBPawn,
				From:  board.Squares[ROW_6][COL_B],
				To:    board.Squares[ROW_5][COL_B],
			},
			"PAWN: B6 -> B5 is not a valid move",
		},
		{
			&Move{
				Piece: blackAPawn,
				From:  board.Squares[ROW_7][COL_A],
				To:    board.Squares[ROW_5][COL_A],
			},
			"PAWN: A7 -> A5",
		},
	}

	for _, tt := range tests {
		if tt.input.Piece.Type() != PAWN {
			t.Fatalf("Piece should be a %s. Got %s", PAWN, tt.input.Piece.Type())
		}
		receipt, err := board.Move(tt.input)

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
				Piece: whiteAPawn,
				From:  board.Squares[ROW_2][COL_A],
				To:    board.Squares[ROW_3][COL_B],
			},
			"PAWN: A2 -> B3 is not a valid move",
		},
		{
			&Move{
				Piece: whiteBPawn,
				From:  board.Squares[ROW_2][COL_B],
				To:    board.Squares[ROW_3][COL_A],
			},
			"PAWN: B2 -> A3 is not a valid move",
		},
		{
			&Move{
				Piece: blackBPawn,
				From:  board.Squares[ROW_7][COL_B],
				To:    board.Squares[ROW_6][COL_A],
			},
			"PAWN: B7 -> A6 is not a valid move",
		},
		{
			&Move{
				Piece: blackAPawn,
				From:  board.Squares[ROW_7][COL_A],
				To:    board.Squares[ROW_6][COL_B],
			},
			"PAWN: A7 -> B6 is not a valid move",
		},
		{
			&Move{
				Piece: whiteEPawn,
				From:  whiteEPawn.Square(),
				To:    blackDPawn.Square(),
			},
			"PAWN TAKES PAWN: E4 -> D5",
		},
		{
			&Move{
				Piece: blackEPawn,
				From:  blackEPawn.Square(),
				To:    whiteDPawn.Square(),
			},
			"PAWN TAKES PAWN: E5 -> D4",
		},
	}
	for _, tt := range tests {
		if tt.input.Piece.Type() != PAWN {
			t.Fatalf("Piece should be a %s. Got %s", PAWN, tt.input.Piece.Type())
		}
		receipt, err := board.Move(tt.input)

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
	g4 := board.Squares[ROW_4][COL_G]
	h2 := board.Squares[ROW_2][COL_H]
	h4 := board.Squares[ROW_4][COL_H]
	f4 := board.Squares[ROW_4][COL_F]
	f2 := board.Squares[ROW_2][COL_F]

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
				Piece: blackGPawn,
				From:  g4,
				To:    board.Squares[ROW_3][COL_F],
			},
			"PAWN: G4 -> F3 is not a valid move",
		},
		{
			&Move{
				Piece: blackGPawn,
				From:  g4,
				To:    board.Squares[ROW_3][COL_H],
			},
			"PAWN TAKES PAWN (EN PASSANT): G4 -> H3",
		},
	}

	for _, tt := range tests {
		if tt.input.Piece.Type() != PAWN {
			t.Fatalf("Piece should be a %s. Got %s", PAWN, tt.input.Piece.Type())
		}
		receipt, err := board.Move(tt.input)

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

	d5 := board.Squares[ROW_5][COL_D]
	e5 := board.Squares[ROW_5][COL_E]
	e7 := board.Squares[ROW_7][COL_E]
	c5 := board.Squares[ROW_5][COL_C]
	c7 := board.Squares[ROW_7][COL_C]

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
				Piece: whiteDPawn,
				From:  d5,
				To:    board.Squares[ROW_6][COL_E],
			},
			"PAWN: D5 -> E6 is not a valid move",
		},
		{
			&Move{
				Piece: whiteDPawn,
				From:  d5,
				To:    board.Squares[ROW_6][COL_C],
			},
			"PAWN TAKES PAWN (EN PASSANT): D5 -> C6",
		},
	}

	for _, tt := range tests {
		if tt.input.Piece.Type() != PAWN {
			t.Fatalf("Piece should be a %s. Got %s", PAWN, tt.input.Piece.Type())
		}
		receipt, err := board.Move(tt.input)

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
	board.Move(move)
}
