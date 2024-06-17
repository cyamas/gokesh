package board

import (
	"fmt"
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

var ENEMY = map[string]string{
	WHITE: BLACK,
	BLACK: WHITE,
}

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

type Error struct {
	Message string
}

func NewError(format string, a ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(format, a...),
	}
}

type Square struct {
	Piece  Piece
	Row    int
	Column int
	Name   string
}

type Path map[*Square]bool

func (s *Square) SetPiece(piece Piece) {
	s.Piece = piece
}

func (s *Square) IsEmpty() bool {
	if s.Piece.Type() == NULL {
		return true
	}
	return false
}

func (b *Board) getSquare(row, col int) *Square {
	return b.Squares[row][col]
}

type Board struct {
	Squares     [][]*Square
	Moves       []*Move
	WhitePieces map[Piece]bool
	BlackPieces map[Piece]bool
	Checkmate   bool
	Check       bool
	Checkers    []Piece
	Value       float32
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

func (b *Board) Evaluate(turn string) {
	b.Value = 0.0
	if turn == WHITE {
		b.evaluateWhite()
		king := b.GetKing(BLACK)
		checked, checkers := king.IsInCheck(b)
		b.Check = checked
		b.Checkers = checkers
		b.evaluateBlack()
		if b.CheckmateDetected(BLACK, king) {
			b.Checkmate = true
		}
	} else {
		b.evaluateBlack()
		king := b.GetKing(WHITE)
		checked, checkers := king.IsInCheck(b)
		b.Check = checked
		b.Checkers = checkers
		b.evaluateWhite()
		if b.CheckmateDetected(WHITE, king) {
			b.Checkmate = true
		}
	}
}

func (b *Board) evaluateBlack() {
	for piece := range b.BlackPieces {
		piece.SetActiveSquares(b)
		b.Value += piece.Value()
	}
}

func (b *Board) evaluateWhite() {
	for piece := range b.WhitePieces {
		piece.SetActiveSquares(b)
		b.Value += piece.Value()
	}
}

func (b *Board) GetAllValidMoves(color string) []*Move {
	moves := []*Move{}
	pieces := b.getAllies(color)
	for piece := range pieces {
		for sq, activity := range piece.ActiveSquares() {
			if activity == CAPTURE || activity == FREE {
				move := &Move{
					Turn:  color,
					Piece: piece,
					From:  piece.Square(),
					To:    sq,
				}
				moves = append(moves, move)
			}
		}
	}
	return moves
}

func (b *Board) CheckmateDetected(color string, king *King) bool {
	checked, checkers := king.IsInCheck(b)
	if !checked {
		return false
	}

	kingActives := king.ActiveSquares()

	switch {
	case len(kingActives) == 0 && len(checkers) > 1:
		return true
	case len(kingActives) > 0 && king.CanEvadeCheck(kingActives, b):
		return false
	default:
		return !b.piecePreventsCheckmate(king, color, checkers)
	}
}

func (b *Board) piecePreventsCheckmate(king *King, color string, checkers []Piece) bool {
	checker := checkers[0]
	checkerSq := checker.Square()
	checkPath := b.GetAttackedPath(checkerSq, king.Square())
	allies := b.getAllies(color)
	for piece := range allies {
		if piece.Type() == KING {
			continue
		}
		pieceActives := piece.ActiveSquares()
		for sq := range checkPath {
			if moveType, ok := pieceActives[sq]; ok {
				if moveType == FREE || moveType == CAPTURE {
					return true
				}
			}
		}
	}
	return false
}

func (b *Board) getAllies(color string) map[Piece]bool {
	if color == WHITE {
		return b.WhitePieces
	} else {
		return b.BlackPieces
	}
}

func (b *Board) getEnemies(color string) map[Piece]bool {
	if color == WHITE {
		return b.BlackPieces
	} else {
		return b.WhitePieces
	}
}

func (b *Board) GetAttackedSquares(color string) map[*Square][]Piece {
	enemies := b.getEnemies(color)

	attackedSqs := make(map[*Square][]Piece)

	for piece := range enemies {
		if king, ok := piece.(*King); ok {
			b.addKingAttackedSquares(king, attackedSqs)
			continue
		} else {
			pieceActives := piece.ActiveSquares()
			for sq, sqActivity := range pieceActives {
				if _, ok := attackedSqs[sq]; !ok {
					if piece.Type() == PAWN && sqActivity == FREE {
						continue
					}
				}
				addAttackedSquare(sq, piece, attackedSqs)
			}
		}
	}
	return attackedSqs
}

func (b *Board) addKingAttackedSquares(king *King, attackedSqs map[*Square][]Piece) {
	for _, dir := range KING_DIRS {
		row := king.Square().Row + dir[0]
		col := king.Square().Column + dir[1]
		if squareExists(row, col) {
			kingGuardedSq := b.getSquare(row, col)
			addAttackedSquare(kingGuardedSq, king, attackedSqs)
		}
	}
}

func addAttackedSquare(sq *Square, piece Piece, attackeds map[*Square][]Piece) {
	_, ok := attackeds[sq]
	if ok {
		attackeds[sq] = append(attackeds[sq], piece)
	} else {
		attackeds[sq] = []Piece{piece}
	}
}

func (b *Board) GetAttackedPath(from, to *Square) map[*Square]bool {
	if from.Piece.Type() == KNIGHT {
		return map[*Square]bool{from: true}
	}
	bigRow, smallRow := compareCoords(from.Row, to.Row)
	bigCol, smallCol := compareCoords(from.Column, to.Column)
	switch {
	case from.Row == to.Row:
		return b.horizontalPath(bigCol, smallCol, bigRow)
	case from.Column == to.Column:
		return b.verticalPath(bigRow, smallRow, bigCol)
	default:
		return b.diagonalPath(from, to)
	}
}

func (b *Board) horizontalPath(bigCol, smallCol, row int) map[*Square]bool {
	path := make(map[*Square]bool)

	for col := smallCol; col < bigCol; col++ {
		sq := b.Squares[row][col]
		path[sq] = true
	}
	return path
}

func (b *Board) verticalPath(bigRow, smallRow, col int) map[*Square]bool {
	path := make(map[*Square]bool)

	for row := smallRow; row < bigRow; row++ {
		sq := b.Squares[row][col]
		path[sq] = true
	}
	return path
}

func (b *Board) diagonalPath(from, to *Square) map[*Square]bool {
	path := make(map[*Square]bool)

	rows := orderCoords(from.Row, to.Row)
	cols := orderCoords(from.Column, to.Column)
	for i := range rows {
		sq := b.Squares[rows[i]][cols[i]]
		path[sq] = true
	}

	return path
}

func orderCoords(from, to int) []int {
	var coords []int
	if from-to < 0 {
		for coord := from; coord < to; coord++ {
			coords = append(coords, coord)
		}
	} else {
		for coord := from; coord > to; coord-- {
			coords = append(coords, coord)
		}
	}
	return coords
}

func compareCoords(a, b int) (int, int) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}

func (b *Board) GetKing(color string) *King {
	if color == WHITE {
		for piece := range b.WhitePieces {
			if piece.Type() == KING {
				return piece.(*King)
			}
		}
	}
	for piece := range b.BlackPieces {
		if piece.Type() == KING {
			return piece.(*King)
		}
	}
	return nil
}

func (b *Board) SetPiece(piece Piece, square *Square) {
	piece.SetSquare(square)
	square.SetPiece(piece)
	if piece.Color() == WHITE {
		b.WhitePieces[piece] = true
	} else {
		b.BlackPieces[piece] = true
	}
}

func (b *Board) SetupPieces() {
	var whiteStartSquares = map[string]Piece{
		"A2": &Pawn{value: 1.0},
		"B2": &Pawn{value: 1.0},
		"C2": &Pawn{value: 1.0},
		"D2": &Pawn{value: 1.0},
		"E2": &Pawn{value: 1.0},
		"F2": &Pawn{value: 1.0},
		"G2": &Pawn{value: 1.0},
		"H2": &Pawn{value: 1.0},
		"B1": &Knight{value: 3.05},
		"G1": &Knight{value: 3.05},
		"C1": &Bishop{value: 3.33},
		"F1": &Bishop{value: 3.33},
		"A1": &Rook{value: 5.63, CastleSq: b.Squares[ROW_1][COL_D]},
		"H1": &Rook{value: 5.63, CastleSq: b.Squares[ROW_1][COL_F]},
		"D1": &Queen{value: 9.5},
		"E1": &King{value: 99.9},
	}

	var blackStartSquares = map[string]Piece{
		"A7": &Pawn{value: -1.0},
		"B7": &Pawn{value: -1.0},
		"C7": &Pawn{value: -1.0},
		"D7": &Pawn{value: -1.0},
		"E7": &Pawn{value: -1.0},
		"F7": &Pawn{value: -1.0},
		"G7": &Pawn{value: -1.0},
		"H7": &Pawn{value: -1.0},
		"B8": &Knight{value: -3.05},
		"G8": &Knight{value: -3.05},
		"C8": &Bishop{value: -3.33},
		"F8": &Bishop{value: -3.33},
		"A8": &Rook{value: -5.63, CastleSq: b.Squares[ROW_8][COL_D]},
		"H8": &Rook{value: -5.63, CastleSq: b.Squares[ROW_8][COL_F]},
		"D8": &Queen{value: -9.5},
		"E8": &King{value: -99.9},
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
	b.Evaluate(WHITE)
}

func (b *Board) MovePiece(move *Move) (string, *Error) {
	receipt := ""
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
			receipt = b.executeFreeMove(move)
			b.Evaluate(move.Turn)
			return receipt, nil
		case CAPTURE:
			receipt = b.executeCaptureMove(move)
			b.Evaluate(move.Turn)
			return receipt, nil
		case EN_PASSANT:
			receipt = b.executeEnPassantMove(move)
			b.Evaluate(move.Turn)
			return receipt, nil
		case CASTLE:
			receipt = b.executeCastleMove(move)
			b.Evaluate(move.Turn)
			return receipt, nil
		}
	}

	return b.invalidMove(move)
}

func (b *Board) invalidMove(move *Move) (string, *Error) {
	gameError := NewError(
		"%s: %s -> %s is not a valid move",
		move.Piece.Type(),
		move.From.Name,
		move.To.Name,
	)
	return gameError.Message, gameError
}

func (b *Board) executeFreeMove(move *Move) string {
	receipt := fmt.Sprintf("%s: %s -> %s", move.Piece.Type(), move.From.Name, move.To.Name)

	if move.Piece.Type() == PAWN && (move.To.Row == ROW_8 || move.To.Row == ROW_1) {
		return b.executePawnPromotion(move, receipt)
	} else {
		move.From.SetPiece(&Null{})
		b.SetPiece(move.Piece, move.To)
		b.Moves = append(b.Moves, move)

		return receipt
	}
}

func (b *Board) executeCaptureMove(move *Move) string {
	capturedPiece := move.To.Piece

	receipt := fmt.Sprintf(
		"%s TAKES %s: %s -> %s",
		move.Piece.Type(),
		capturedPiece.Type(),
		move.From.Name,
		move.To.Name,
	)

	b.RemovePiece(capturedPiece, move.To)

	if move.Piece.Type() == PAWN && (move.To.Row == ROW_8 || move.To.Row == ROW_1) {
		return b.executePawnPromotion(move, receipt)

	} else {
		move.From.Piece = &Null{}
		b.SetPiece(move.Piece, move.To)
		b.Moves = append(b.Moves, move)

		return receipt
	}

}

func (b *Board) executePawnPromotion(move *Move, receipt string) string {
	b.RemovePiece(move.From.Piece, move.From)
	b.SetPiece(move.Promotion, move.To)
	b.Moves = append(b.Moves, move)

	receipt += fmt.Sprintf(" (PROMOTION: %s)", move.Promotion.Type())
	return receipt

}

func (b *Board) executeEnPassantMove(move *Move) string {
	captureSq := b.Squares[move.Piece.Square().Row][move.To.Column]
	capturedPiece := captureSq.Piece
	b.RemovePiece(capturedPiece, captureSq)

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

func (b *Board) RemovePiece(piece Piece, sq *Square) {
	if piece.Color() == WHITE {
		b.SetPiece(&Null{}, sq)
		delete(b.WhitePieces, piece)
	} else {
		b.SetPiece(&Null{}, sq)
		delete(b.BlackPieces, piece)
	}
}

func (b *Board) LastMove() *Move {
	if len(b.Moves) > 0 {
		return b.Moves[len(b.Moves)-1]
	}
	return nil
}
