package board

import (
	"fmt"

	"github.com/cyamas/gokesh/game"
)

const (
	BLACK = "BLACK"
	WHITE = "WHITE"

	COL_A = 0
	COL_B = 1
	COL_C = 2
	COL_D = 3
	COL_E = 4
	COL_F = 5
	COL_G = 6
	COL_H = 7

	ROW_1 = 7
	ROW_2 = 6
	ROW_3 = 5
	ROW_4 = 4
	ROW_5 = 3
	ROW_6 = 2
	ROW_7 = 1
	ROW_8 = 0
)

var whiteStartSquares = map[string]Piece{
	"A2": &Pawn{},
	"B2": &Pawn{},
	"C2": &Pawn{},
	"D2": &Pawn{},
	"E2": &Pawn{},
	"F2": &Pawn{},
	"G2": &Pawn{},
	"H2": &Pawn{},
	"B1": &Knight{},
	"G1": &Knight{},
	"C1": &Bishop{},
	"F1": &Bishop{},
	"A1": &Rook{},
	"H1": &Rook{},
	"D1": &Queen{},
	"E1": &King{},
}

var blackStartSquares = map[string]Piece{
	"A7": &Pawn{},
	"B7": &Pawn{},
	"C7": &Pawn{},
	"D7": &Pawn{},
	"E7": &Pawn{},
	"F7": &Pawn{},
	"G7": &Pawn{},
	"H7": &Pawn{},
	"B8": &Knight{},
	"G8": &Knight{},
	"C8": &Bishop{},
	"F8": &Bishop{},
	"A8": &Rook{},
	"H8": &Rook{},
	"D8": &Queen{},
	"E8": &King{},
}

type Board struct {
	Squares     [][]*Square
	Moves       []*Move
	WhitePieces map[Piece]bool
	BlackPieces map[Piece]bool
}

func New() *Board {
	board := &Board{
		WhitePieces: make(map[Piece]bool),
		BlackPieces: make(map[Piece]bool),
	}

	rows := []string{"8", "7", "6", "5", "4", "3", "2", "1"}
	cols := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	for i := range 8 {
		boardRow := []*Square{}
		rowNum := rows[i]
		for j := range 8 {
			colLetter := cols[j]
			square := &Square{
				Row:    i,
				Column: j,
				Name:   colLetter + rowNum,
				Piece:  &Null{},
			}
			boardRow = append(boardRow, square)
		}
		board.Squares = append(board.Squares, boardRow)
	}

	return board
}

func (b *Board) GetAttackedSquares(color string) map[*Square]bool {
	var enemies map[Piece]bool

	if color == WHITE {
		enemies = b.BlackPieces
	} else {
		enemies = b.WhitePieces
	}

	actives := make(map[*Square]bool)
	for piece := range enemies {
		if piece.Type() == KING {
			continue
		}
		pieceActives := piece.ActiveSquares(b)
		for sq, sqActivity := range pieceActives {
			if _, ok := actives[sq]; !ok {
				if piece.Type() == PAWN && sqActivity == FREE {
					continue
				}
				actives[sq] = true
			}
		}
	}
	return actives
}

func (b *Board) SetPiece(piece Piece, square *Square) {
	piece.SetSquare(square)
	square.SetPiece(piece)
	if piece.Color() == WHITE {
		_, ok := b.WhitePieces[piece]
		if !ok {
			b.WhitePieces[piece] = true
		}

	} else {
		_, ok := b.BlackPieces[piece]
		if !ok {
			b.BlackPieces[piece] = true
		}
	}
}

func (b *Board) SquareIsSafe(color string, cand *Square) bool {
	var enemies map[Piece]bool
	if color == WHITE {
		enemies = b.BlackPieces
	} else {
		enemies = b.WhitePieces
	}

	for piece := range enemies {
		if piece.Type() == KING {
			if b.squareGuardedByEnemyKing(piece, cand) {
				return false
			}
		} else {
			valids := piece.ActiveSquares(b)
			moveType, ok := valids[cand]
			if ok {
				if piece.Type() == PAWN && moveType == FREE {
					continue
				} else {
					return false
				}
			}
		}
	}

	return true
}

func (b *Board) squareGuardedByEnemyKing(king Piece, cand *Square) bool {
	dirs := map[string][2]int{
		"N":  {-1, 0},
		"E":  {0, 1},
		"S":  {1, 0},
		"W":  {0, -1},
		"NW": {-1, -1},
		"NE": {-1, 1},
		"SE": {1, 1},
		"SW": {1, -1},
	}
	for _, coords := range dirs {
		row := king.Square().Row + coords[0]
		col := king.Square().Column + coords[1]
		if squareExists(row, col) && b.Squares[row][col] == cand {
			return true
		}
	}
	return false
}

func (b *Board) SetupPieces() {
	var whiteStartSquares = map[string]Piece{
		"A2": &Pawn{},
		"B2": &Pawn{},
		"C2": &Pawn{},
		"D2": &Pawn{},
		"E2": &Pawn{},
		"F2": &Pawn{},
		"G2": &Pawn{},
		"H2": &Pawn{},
		"B1": &Knight{},
		"G1": &Knight{},
		"C1": &Bishop{},
		"F1": &Bishop{},
		"A1": &Rook{CastleSq: b.Squares[ROW_1][COL_D]},
		"H1": &Rook{CastleSq: b.Squares[ROW_1][COL_F]},
		"D1": &Queen{},
		"E1": &King{},
	}

	var blackStartSquares = map[string]Piece{
		"A7": &Pawn{},
		"B7": &Pawn{},
		"C7": &Pawn{},
		"D7": &Pawn{},
		"E7": &Pawn{},
		"F7": &Pawn{},
		"G7": &Pawn{},
		"H7": &Pawn{},
		"B8": &Knight{},
		"G8": &Knight{},
		"C8": &Bishop{},
		"F8": &Bishop{},
		"A8": &Rook{CastleSq: b.Squares[ROW_8][COL_D]},
		"H8": &Rook{CastleSq: b.Squares[ROW_8][COL_F]},
		"D8": &Queen{},
		"E8": &King{},
	}
	whiteRows := [2]int{6, 7}
	blackRows := [2]int{0, 1}

	for _, row := range whiteRows {
		for _, sq := range b.Squares[row] {
			piece, ok := whiteStartSquares[sq.Name]
			if !ok {
				continue
			} else {
				piece.SetColor(WHITE)
				piece.SetSquare(sq)
				sq.Piece = piece
				b.WhitePieces[piece] = true
			}
		}
	}

	for _, row := range blackRows {
		for _, sq := range b.Squares[row] {
			piece, ok := blackStartSquares[sq.Name]
			if !ok {
				continue
			} else {
				piece.SetColor(BLACK)
				piece.SetSquare(sq)
				sq.Piece = piece
				b.BlackPieces[piece] = true
			}
		}
	}

}

func (b *Board) Move(move *Move) (string, *game.Error) {
	if move.IsValid(b) {
		switch piece := move.Piece.(type) {
		case *Pawn:
			piece.Moved = true
		case *King:
			piece.Moved = true
		case *Rook:
			piece.Moved = true
		}

		switch move.Type {
		case FREE:
			return b.executeFreeMove(move), nil
		case CAPTURE:
			return b.executeCaptureMove(move), nil
		case EN_PASSANT:
			return b.executeEnPassantMove(move), nil
		case CASTLE:
			return b.executeCastleMove(move), nil
		}
	}
	gameError := game.NewError(
		"%s: %s -> %s is not a valid move",
		move.Piece.Type(),
		move.From.Name,
		move.To.Name,
	)
	return gameError.Message, gameError
}

func (b *Board) executeFreeMove(move *Move) string {
	move.Piece.SetSquare(move.To)
	move.To.SetPiece(move.Piece)
	move.From.SetPiece(&Null{})
	b.Moves = append(b.Moves, move)
	receipt := fmt.Sprintf("%s: %s -> %s", move.Piece.Type(), move.From.Name, move.Piece.Square().Name)

	return receipt
}

func (b *Board) executeCaptureMove(move *Move) string {
	capturedPiece := move.To.Piece
	b.removePiece(capturedPiece)

	move.To.SetPiece(move.Piece)
	move.From.SetPiece(&Null{})
	move.Piece.SetSquare(move.To)
	b.Moves = append(b.Moves, move)

	receipt := fmt.Sprintf(
		"%s TAKES %s: %s -> %s",
		move.Piece.Type(),
		capturedPiece.Type(),
		move.From.Name,
		move.To.Name,
	)
	return receipt
}

func (b *Board) executeEnPassantMove(move *Move) string {
	capturedPiece := b.Squares[move.Piece.Square().Row][move.To.Column].Piece
	b.removePiece(capturedPiece)

	move.To.SetPiece(move.Piece)
	move.From.SetPiece(&Null{})
	move.Piece.SetSquare(move.To)
	b.Moves = append(b.Moves, move)

	receipt := fmt.Sprintf(
		"PAWN TAKES PAWN (EN PASSANT): %s -> %s",
		move.From.Name,
		move.To.Name,
	)
	return receipt
}

func (b *Board) executeCastleMove(move *Move) string {
	king := move.Piece
	b.SetPiece(king, move.To)
	move.From.SetPiece(&Null{})
	b.Moves = append(b.Moves, move)

	switch king.Square().Name {
	case "G8":
		h8 := b.Squares[ROW_8][COL_H]
		f8 := b.Squares[ROW_8][COL_F]
		b.castleRook(h8, f8)

		return "KING CASTLES SHORT"

	case "G1":
		h1 := b.Squares[ROW_1][COL_H]
		f1 := b.Squares[ROW_1][COL_F]
		b.castleRook(h1, f1)
		return "KING CASTLES SHORT"

	case "C8":
		a8 := b.Squares[ROW_8][COL_A]
		d8 := b.Squares[ROW_8][COL_D]
		b.castleRook(a8, d8)

		return "KING CASTLES LONG"
	case "C1":
		a1 := b.Squares[ROW_1][COL_A]
		d1 := b.Squares[ROW_1][COL_D]
		b.castleRook(a1, d1)
		return "KING CASTLES LONG"
	default:
		return ""
	}
}

func (b *Board) castleRook(from *Square, to *Square) {
	rook := from.Piece
	b.SetPiece(rook, to)
	from.Piece = &Null{}
}

func (b *Board) removePiece(piece Piece) {
	if piece.Color() == WHITE {
		delete(b.WhitePieces, piece)
	} else {
		delete(b.BlackPieces, piece)
	}
}

func (b *Board) LastMove() *Move {
	if len(b.Moves) > 0 {
		return b.Moves[len(b.Moves)-1]
	}
	return nil
}

type Square struct {
	Piece  Piece
	Row    int
	Column int
	Name   string
}

func (s *Square) SetPiece(piece Piece) {
	s.Piece = piece
}

func (s *Square) IsEmpty() bool {
	if s.Piece.Type() == NULL {
		return true
	}
	return false
}
