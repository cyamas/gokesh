package board

import (
	"testing"
)

func TestUndoMove(t *testing.T) {
	testNumber := 1
	// FREE MOVE
	board1 := New()
	board1.SetupPieces()
	e2_1 := board1.GetSquare(ROW_2, COL_E)
	e4_1 := board1.GetSquare(ROW_4, COL_E)

	//PAWN PROMOTION
	board2 := New()
	board2.SetupFromFen("8/pp2P3/8/3k2p1/8/2P3P1/P5P1/5K2")
	e7_2 := board2.GetSquare(ROW_7, COL_E)
	e8_2 := board2.GetSquare(ROW_8, COL_E)

	//CAPTURE
	board3 := New()
	board3.SetupFromFen("rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR")
	e4_3 := board3.GetSquare(ROW_4, COL_E)
	d5_3 := board3.GetSquare(ROW_5, COL_D)

	//CASTLE
	board4 := New()
	board4.SetupFromFen("rnbqkbnr/pp3ppp/2p1p3/3p4/4P3/5N2/PPPPBPPP/RNBQK2R")
	e1_4 := board4.GetSquare(ROW_1, COL_E)
	g1_4 := board4.GetSquare(ROW_1, COL_G)

	whiteKing4 := board4.GetKing(WHITE)
	whiteHRook4 := board4.GetSquare(ROW_1, COL_H).Piece
	whiteKing4.SetMoveCount(0)
	whiteHRook4.SetMoveCount(0)

	//EN_PASSANT
	board5 := New()
	board5.SetupFromFen("rnbqkbnr/pp1ppppp/2p5/4P3/8/8/PPPP1PPP/RNBQKBNR")
	e5_5 := board5.GetSquare(ROW_5, COL_E)
	d6_5 := board5.GetSquare(ROW_6, COL_D)
	d7_5 := board5.GetSquare(ROW_7, COL_D)
	d5_5 := board5.GetSquare(ROW_5, COL_D)
	board5.Evaluate(WHITE)

	enPassantSetupMove := &Move{
		Turn:  BLACK,
		Piece: d7_5.Piece,
		From:  d7_5,
		To:    d5_5,
	}
	board5.MovePiece(enPassantSetupMove)

	tests := []struct {
		board    *Board
		input    *Move
		expected string
	}{
		{
			board1,
			&Move{
				Turn:  WHITE,
				Piece: e2_1.Piece,
				From:  e2_1,
				To:    e4_1,
			},
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
		},
		{
			board2,
			&Move{
				Turn:  WHITE,
				Piece: e7_2.Piece,
				From:  e7_2,
				To:    e8_2,
			},
			"8/pp2P3/8/3k2p1/8/2P3P1/P5P1/5K2",
		},
		{
			board3,
			&Move{
				Turn:  WHITE,
				Piece: e4_3.Piece,
				From:  e4_3,
				To:    d5_3,
			},
			"rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR",
		},
		{
			board4,
			&Move{
				Turn:  WHITE,
				Piece: e1_4.Piece,
				From:  e1_4,
				To:    g1_4,
			},
			"rnbqkbnr/pp3ppp/2p1p3/3p4/4P3/5N2/PPPPBPPP/RNBQK2R",
		},
		{
			board5,
			&Move{
				Turn:  WHITE,
				Piece: e5_5.Piece,
				From:  e5_5,
				To:    d6_5,
			},
			"rnbqkbnr/pp2pppp/2p5/3pP3/8/8/PPPP1PPP/RNBQKBNR",
		},
	}
	for _, tt := range tests {
		tt.board.Evaluate(ENEMY[tt.input.Turn])
		ogValids := tt.board.GetAllValidMoves(tt.input.Turn)
		lenOGWhitePieces := len(tt.board.WhitePieces)
		lenOGBlackPieces := len(tt.board.BlackPieces)
		ogBoard := tt.board.Copy()
		tt.board.MovePiece(tt.input)
		tt.board.UndoMove()
		valids := tt.board.GetAllValidMoves(tt.input.Turn)
		for r, row := range tt.board.Squares {
			for c, sq := range row {
				ogPiece := ogBoard.Squares[r][c].Piece.Type()
				piece := tt.board.Squares[r][c].Piece.Type()
				if ogPiece != piece {
					t.Fatalf("Square %s should have piece %s. Got %s", sq.Name, ogPiece, piece)
				}
			}
		}
		if len(ogValids) != len(valids) {
			t.Fatalf("Len Valid moves should be %d. Got %d", len(ogValids), len(valids))
		}
		fen := tt.board.Fen()
		if fen != tt.expected {
			t.Fatalf("fen should be %s. got %s", tt.expected, fen)
		}
		if len(ogBoard.Moves) != len(tt.board.Moves) {
			t.Fatalf("len board.Moves should be %d. Got %d", len(ogBoard.Moves), len(tt.board.Moves))
		}
		if len(ogBoard.Fens) != len(tt.board.Fens) {
			t.Fatalf("len Fens should be %d. Got %d", len(ogBoard.Fens), len(tt.board.Fens))
		}
		if len(tt.board.WhitePieces) != lenOGWhitePieces {
			t.Fatalf("len WhitePieces should be %d. Got %d", lenOGWhitePieces, len(tt.board.WhitePieces))
		}
		if len(tt.board.BlackPieces) != lenOGBlackPieces {
			t.Fatalf("len BlackPieces should be %d. Got %d", lenOGBlackPieces, len(tt.board.BlackPieces))
		}
		testNumber++
	}
}

func TestSetupFromFen(t *testing.T) {
	board1 := New()
	board2 := New()
	board3 := New()

	tests := []struct {
		board    *Board
		input    string
		expected string
	}{
		{
			board1,
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
		},
		{
			board2,
			"r1b1k1nr/ppp2ppp/2n5/3Pp3/1b6/P1P2N2/1P1qQPPP/RNB1KB1R",
			"r1b1k1nr/ppp2ppp/2n5/3Pp3/1b6/P1P2N2/1P1qQPPP/RNB1KB1R",
		},
		{
			board3,
			"r3k1nr/pp3pQp/2p5/3Pp3/1P6/P1N2N1b/5PPP/n1BK1B1R",
			"r3k1nr/pp3pQp/2p5/3Pp3/1P6/P1N2N1b/5PPP/n1BK1B1R",
		},
	}

	for _, tt := range tests {
		tt.board.SetupFromFen(tt.input)
		fen := tt.board.Fen()
		if fen != tt.expected {
			t.Fatalf("Fen should be %s. Got %s", tt.expected, fen)
		}
	}
}

func TestBoardCopy(t *testing.T) {
	board1 := New()
	board1.SetupPieces()

	board2 := New()
	e7 := board2.Squares[ROW_7][COL_E]
	whiteKing := &King{color: WHITE}

	f4 := board2.Squares[ROW_4][COL_F]
	blackKing := &King{color: BLACK}

	b4 := board2.Squares[ROW_4][COL_B]
	whiteBishop := &Bishop{color: WHITE}

	h4 := board2.Squares[ROW_4][COL_H]
	blackPawn := &Pawn{color: BLACK}

	b5 := board2.Squares[ROW_5][COL_B]
	whitePawn := &Pawn{color: WHITE}

	c5 := board2.Squares[ROW_5][COL_C]
	whiteRook := &Rook{color: WHITE}

	b8 := board2.Squares[ROW_8][COL_B]
	blackBishop := &Bishop{color: BLACK}

	board2.SetPiece(whiteKing, e7)
	board2.SetPiece(blackKing, f4)
	board2.SetPiece(whiteBishop, b4)
	board2.SetPiece(blackBishop, b8)
	board2.SetPiece(whitePawn, b5)
	board2.SetPiece(blackPawn, h4)
	board2.SetPiece(whiteRook, c5)

	tests := []struct {
		board    *Board
		expected string
	}{
		{board1, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"},
		{board2, "1b6/4K3/8/1PR5/1B3k1p/8/8/8"},
	}

	for _, tt := range tests {
		copy := tt.board.Copy()
		copyFen := copy.Fen()
		if copyFen != tt.expected {
			t.Fatalf("copyFen should be %s. Got %s", tt.expected, copyFen)
		}
	}
}

func TestFen(t *testing.T) {
	board1 := New()
	board1.SetupPieces()

	board2 := New()
	e7 := board2.Squares[ROW_7][COL_E]
	whiteKing := &King{color: WHITE}

	f4 := board2.Squares[ROW_4][COL_F]
	blackKing := &King{color: BLACK}

	b4 := board2.Squares[ROW_4][COL_B]
	whiteBishop := &Bishop{color: WHITE}

	h4 := board2.Squares[ROW_4][COL_H]
	blackPawn := &Pawn{color: BLACK}

	b5 := board2.Squares[ROW_5][COL_B]
	whitePawn := &Pawn{color: WHITE}

	c5 := board2.Squares[ROW_5][COL_C]
	whiteRook := &Rook{color: WHITE}

	b8 := board2.Squares[ROW_8][COL_B]
	blackBishop := &Bishop{color: BLACK}

	board2.SetPiece(whiteKing, e7)
	board2.SetPiece(blackKing, f4)
	board2.SetPiece(whiteBishop, b4)
	board2.SetPiece(blackBishop, b8)
	board2.SetPiece(whitePawn, b5)
	board2.SetPiece(blackPawn, h4)
	board2.SetPiece(whiteRook, c5)

	tests := []struct {
		board    *Board
		expected string
	}{
		{board1, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"},
		{board2, "1b6/4K3/8/1PR5/1B3k1p/8/8/8"},
	}
	for _, tt := range tests {
		fen := tt.board.Fen()
		if fen != tt.expected {
			t.Fatalf("Fen notation should be %s. Got %s", tt.expected, fen)
		}
	}
}

func TestDiagonalPath(t *testing.T) {
	board := New()
	h1 := board.GetSquare(ROW_1, COL_H)
	a1 := board.GetSquare(ROW_1, COL_A)
	b2 := board.GetSquare(ROW_2, COL_B)
	h8 := board.GetSquare(ROW_8, COL_H)
	a8 := board.GetSquare(ROW_8, COL_A)

	tests := []struct {
		from           *Square
		to             *Square
		expectedLength int
	}{
		{
			h1,
			a8,
			7,
		},
		{
			a1,
			h8,
			7,
		},
		{
			a8,
			h1,
			7,
		},
		{
			h8,
			a1,
			7,
		},
		{
			a1,
			b2,
			1,
		},
		{
			b2,
			a1,
			1,
		},
	}
	for _, tt := range tests {
		path := board.diagonalPath(tt.from, tt.to)
		if len(path) != tt.expectedLength {
			t.Fatalf("Path should have length %d. Got %d", tt.expectedLength, len(path))
		}
	}
}

func TestCreateNewBoard(t *testing.T) {
	var board = New()
	if len(board.Squares) != 8 {
		t.Fatalf("Board should have 8 rows. Got %d", len(board.Squares))
	}
	for i, row := range board.Squares {
		if len(row) != 8 {
			t.Fatalf("Row %d should have 8 rows. Got %d", i, len(row))
		}
	}
	tests := []struct {
		input    *Square
		expected string
	}{
		{board.Squares[7][0], "A1"},
		{board.Squares[7][1], "B1"},
		{board.Squares[6][0], "A2"},
	}
	for _, tt := range tests {
		if tt.input.Name != tt.expected {
			t.Fatalf("expected square %s. Got %s", tt.expected, tt.input.Name)
		}
	}
}

func TestSetupPieces(t *testing.T) {
	var board = New()
	tests := []struct {
		input         *Square
		expectedPiece string
		expectedColor string
	}{
		{board.Squares[6][0], "PAWN", WHITE},
		{board.Squares[6][1], "PAWN", WHITE},
		{board.Squares[6][2], "PAWN", WHITE},
		{board.Squares[6][3], "PAWN", WHITE},
		{board.Squares[6][4], "PAWN", WHITE},
		{board.Squares[6][5], "PAWN", WHITE},
		{board.Squares[6][6], "PAWN", WHITE},
		{board.Squares[6][7], "PAWN", WHITE},
		{board.Squares[7][1], "KNIGHT", WHITE},
		{board.Squares[7][6], "KNIGHT", WHITE},
		{board.Squares[7][2], "BISHOP", WHITE},
		{board.Squares[7][5], "BISHOP", WHITE},
		{board.Squares[7][0], "ROOK", WHITE},
		{board.Squares[7][7], "ROOK", WHITE},
		{board.Squares[7][3], "QUEEN", WHITE},
		{board.Squares[7][4], "KING", WHITE},

		{board.Squares[1][0], "PAWN", BLACK},
		{board.Squares[1][1], "PAWN", BLACK},
		{board.Squares[1][2], "PAWN", BLACK},
		{board.Squares[1][3], "PAWN", BLACK},
		{board.Squares[1][4], "PAWN", BLACK},
		{board.Squares[1][5], "PAWN", BLACK},
		{board.Squares[1][6], "PAWN", BLACK},
		{board.Squares[1][7], "PAWN", BLACK},
		{board.Squares[0][1], "KNIGHT", BLACK},
		{board.Squares[0][6], "KNIGHT", BLACK},
		{board.Squares[0][2], "BISHOP", BLACK},
		{board.Squares[0][5], "BISHOP", BLACK},
		{board.Squares[0][0], "ROOK", BLACK},
		{board.Squares[0][7], "ROOK", BLACK},
		{board.Squares[0][3], "QUEEN", BLACK},
		{board.Squares[0][4], "KING", BLACK},
	}

	board.SetupPieces()

	for _, tt := range tests {
		if tt.input.Piece.Type() != tt.expectedPiece {
			t.Fatalf("Expected %s. Got %s", tt.expectedPiece, tt.input.Piece.Type())
		}
		if tt.input.Piece.Color() != tt.expectedColor {
			t.Fatalf("Expected color: %s. Got %s", tt.expectedColor, tt.input.Piece.Color())
		}
	}
}
