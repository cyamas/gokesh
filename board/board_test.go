package board

import (
	"testing"
)

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
