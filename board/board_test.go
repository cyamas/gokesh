package board

import (
	"fmt"
	"testing"
)

func TestBestMove(t *testing.T) {
	board := New()

	e8 := board.Squares[ROW_8][COL_E]
	blackKing := &King{color: BLACK}

	e1 := board.Squares[ROW_1][COL_E]
	whiteKing := &King{color: WHITE}

	e5 := board.Squares[ROW_5][COL_E]
	whiteEpawn := &Pawn{color: WHITE, Moved: true, value: 1.00}

	d5 := board.Squares[ROW_5][COL_D]
	blackDPawn := &Pawn{color: BLACK, Moved: true, value: -1.00}

	c4 := board.Squares[ROW_4][COL_C]
	whiteBishop := &Bishop{color: WHITE, value: 3.33}

	f6 := board.Squares[ROW_6][COL_F]
	blackBishop := &Bishop{color: BLACK, value: -3.33}

	board.SetPiece(whiteKing, e1)
	board.SetPiece(blackKing, e8)
	board.SetPiece(whiteEpawn, e5)
	board.SetPiece(blackDPawn, d5)
	board.SetPiece(whiteBishop, c4)
	board.SetPiece(blackBishop, f6)

	tests := []struct {
		input    string
		expected string
	}{
		{WHITE, "BISHOP: C4 -> B5"},
		{BLACK, "BISHOP: F6 -> H4"},
	}

	for _, tt := range tests {
		board.Evaluate(ENEMY[tt.input])
		best := board.BestMove(tt.input)
		receipt := fmt.Sprintf("%s: %s -> %s", best.Piece.Type(), best.From.Name, best.To.Name)
		if receipt != tt.expected {
			t.Fatalf("Receipt should be '%s'. Got '%s'", tt.expected, receipt)
		}
	}
}

func TestMoveIsSafe(t *testing.T) {
	board := New()

	e8 := board.Squares[ROW_8][COL_E]
	blackKing := &King{color: BLACK}

	e1 := board.Squares[ROW_1][COL_E]
	whiteKing := &King{color: WHITE}

	e5 := board.Squares[ROW_5][COL_E]
	whiteEpawn := &Pawn{color: WHITE, Moved: true, value: 1.00}

	d5 := board.Squares[ROW_5][COL_D]
	blackDPawn := &Pawn{color: BLACK, Moved: true, value: -1.00}

	f1 := board.Squares[ROW_1][COL_F]
	whiteBishop := &Bishop{color: WHITE, value: 3.33}

	f8 := board.Squares[ROW_8][COL_F]
	blackBishop := &Bishop{color: BLACK, value: -3.33}

	board.SetPiece(whiteKing, e1)
	board.SetPiece(blackKing, e8)
	board.SetPiece(whiteEpawn, e5)
	board.SetPiece(blackDPawn, d5)
	board.SetPiece(whiteBishop, f1)
	board.SetPiece(blackBishop, f8)

	tests := []struct {
		id       int
		input    *Move
		expected bool
	}{
		{
			1,
			&Move{
				Turn:  WHITE,
				Piece: whiteBishop,
				From:  f1,
				To:    board.Squares[ROW_4][COL_C],
			},
			false,
		},
		{
			2,
			&Move{
				Turn:  WHITE,
				Piece: whiteBishop,
				From:  f1,
				To:    board.Squares[ROW_3][COL_D],
			},
			true,
		},
		{
			3,
			&Move{
				Turn:  BLACK,
				Piece: blackBishop,
				From:  f8,
				To:    board.Squares[ROW_6][COL_D],
			},
			false,
		},
		{
			4,
			&Move{
				Turn:  BLACK,
				Piece: blackBishop,
				From:  f8,
				To:    board.Squares[ROW_5][COL_C],
			},
			true,
		},
	}

	for _, tt := range tests {
		board.Evaluate(ENEMY[tt.input.Turn])
		result := tt.input.IsSafe(board)
		if result != tt.expected {
			t.Fatalf("Test %d: Move should be %t. Got %t", tt.id, tt.expected, result)
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
	h1 := board.getSquare(ROW_1, COL_H)
	a1 := board.getSquare(ROW_1, COL_A)
	b2 := board.getSquare(ROW_2, COL_B)
	h8 := board.getSquare(ROW_8, COL_H)
	a8 := board.getSquare(ROW_8, COL_A)

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
