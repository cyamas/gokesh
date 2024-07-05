package board

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

const (
	FREE       = "FREE"
	CAPTURE    = "CAPTURE"
	GUARDED    = "GUARDED"
	EN_PASSANT = "ENPASSANT"
	CASTLE     = "CASTLE"
	CHECK      = "CHECK"
)

var centerSquares = make(map[*Square]bool)

type SqActivity string

type Move struct {
	Turn      string
	Type      SqActivity
	Piece     Piece
	From      *Square
	To        *Square
	Promotion Piece
	Value     float64
}

func (b *Board) GetAllValidMoves(color string) []*Move {
	moves := []*Move{}
	pieces := b.getAllies(color)
	for piece := range pieces {
		for sq, activity := range piece.ActiveSquares() {
			if activity != GUARDED {
				move := &Move{
					Turn:  color,
					Piece: piece,
					From:  piece.Square(),
					To:    sq,
				}
				move.Evaluate(activity, b)

				if move.Piece.Type() == PAWN && (move.To.Row == ROW_1 || move.To.Row == ROW_8) {
					move.Promotion = &Queen{color: color}
				}
				moves = append(moves, move)
			}
		}
	}
	sort.Slice(moves, func(i, j int) bool {
		return moves[i].Value > moves[j].Value
	})
	return moves
}

func (m *Move) Evaluate(activity SqActivity, board *Board) {
	switch activity {
	case CASTLE:
		m.Value = float64(0.9)
	case CAPTURE:
		m.Value = math.Abs(m.To.Piece.Value())
	case EN_PASSANT:
		m.Value = float64(1)
	}
	if m.DevelopsMinorPiece() {
		m.Value += float64(0.9)
	}
}

func (m *Move) AddsCenterControl() {
	switch m.Piece.(type) {
	case *Pawn:

	}
}

func (b *Board) getRooks(color string) []*Rook {
	rooks := []*Rook{}
	var pieces map[Piece]bool
	if color == WHITE {
		pieces = b.WhitePieces
	} else {
		pieces = b.BlackPieces
	}
	for piece := range pieces {
		if rook, ok := piece.(*Rook); ok {
			rooks = append(rooks, rook)
		}
	}
	return rooks
}

func (m *Move) DevelopsMinorPiece() bool {
	if isMinorPiece(m.Piece) && !m.Piece.HasMoved() {
		return true
	}
	return false
}

func isMinorPiece(piece Piece) bool {
	if piece.Type() == KNIGHT || piece.Type() == BISHOP {
		return true
	}
	return false
}

func (m *Move) Copy(simBoard *Board) *Move {
	fromRow := m.From.Row
	fromCol := m.From.Column
	toRow := m.To.Row
	toCol := m.To.Column
	simFrom := simBoard.Squares[fromRow][fromCol]
	simTo := simBoard.Squares[toRow][toCol]
	return &Move{
		Turn:  m.Turn,
		Piece: simFrom.Piece,
		From:  simFrom,
		To:    simTo,
	}
}

func (m *Move) IsValid(board *Board) bool {
	valids := m.Piece.ActiveSquares()
	moveType, ok := valids[m.To]
	if ok {
		m.Type = moveType
		return true
	}
	return false
}

func (b *Board) MovePiece(move *Move) (string, *Error) {
	receipt := ""
	if move.IsValid(b) {
		move.Piece.IncrementMoveCount()

		switch move.Type {
		case FREE:
			receipt = b.executeFreeMove(move)
			b.Receipts = append(b.Receipts, receipt)
			b.addFen(b.Fen())
			b.Evaluate(move.Turn)
			return receipt, nil
		case CAPTURE:
			receipt = b.executeCaptureMove(move)
			b.Receipts = append(b.Receipts, receipt)
			b.Evaluate(move.Turn)
			return receipt, nil
		case EN_PASSANT:
			receipt = b.executeEnPassantMove(move)
			b.Receipts = append(b.Receipts, receipt)
			b.Evaluate(move.Turn)
			return receipt, nil
		case CASTLE:
			receipt = b.executeCastleMove(move)
			b.Receipts = append(b.Receipts, receipt)
			b.Evaluate(move.Turn)
			return receipt, nil
		}
	}
	return b.invalidMove(move)
}

func (b *Board) UndoMove() {
	last := b.LastMove()
	prevFen := b.Fen()
	if last.Promotion != nil {
		b.RemovePiece(last.Promotion, last.To)
		b.SetPiece(last.Piece, last.From)
	}
	switch last.Type {
	case FREE:
		b.undoFreeMove(last)
	case CAPTURE:
		b.undoCaptureMove(last)
	case EN_PASSANT:
		b.undoEnPassant(last)
	case CASTLE:
		b.undoCastle(last)
	}
	last.Piece.DecrementMoveCount()
	b.resetCheck(last.Turn)
	b.Checkmate = false
	b.Stalemate = false
	b.Draw = false
	b.removeLastReceipt()
	b.removeLastFen(prevFen)
	b.removeLastMove()
	b.Evaluate(ENEMY[last.Turn])
}

func (b *Board) pawnPromoted(last *Move) bool {
	return last.Piece.Type() == PAWN && (last.To.Row == ROW_1 || last.To.Row == ROW_8)
}

func (b *Board) undoFreeMove(move *Move) {
	move.From.SetPiece(move.Piece)
	move.Piece.SetSquare(move.From)
	move.To.SetPiece(&Null{})
}

func (b *Board) undoCaptureMove(move *Move) {
	move.From.SetPiece(move.Piece)
	move.Piece.SetSquare(move.From)

	captured := b.CapturedPieces[len(b.CapturedPieces)-1]
	b.SetPiece(captured, move.To)

	if len(b.CapturedPieces) == 1 {
		b.CapturedPieces = []Piece{}
		return
	}
	b.CapturedPieces = b.CapturedPieces[:len(b.CapturedPieces)-1]
}

func (b *Board) undoEnPassant(move *Move) {
	move.To.SetPiece(&Null{})
	move.From.SetPiece(move.Piece)
	move.Piece.SetSquare(move.From)

	captured := b.CapturedPieces[len(b.CapturedPieces)-1]
	b.SetPiece(captured, captured.Square())

	if len(b.CapturedPieces) == 1 {
		b.CapturedPieces = []Piece{}
		return
	}
	b.CapturedPieces = b.CapturedPieces[:len(b.CapturedPieces)-1]
}

func (b *Board) undoCastle(move *Move) {
	move.From.SetPiece(move.Piece)
	move.Piece.SetSquare(move.From)
	move.To.SetPiece(&Null{})

	b.resetCastledRook(move)
}

func (b *Board) removeLastMove() {
	if len(b.Moves) == 1 {
		b.Moves = []*Move{}
		return
	}
	b.Moves = b.Moves[:len(b.Moves)-1]
}

func (b *Board) removeLastFen(prevFen string) {
	if b.Fens[prevFen] < 2 {
		delete(b.Fens, prevFen)
	} else {
		b.Fens[prevFen]--
	}
}

func (b *Board) removeLastReceipt() {
	if len(b.Receipts) == 1 {
		b.Receipts = []string{}
		return
	}
	b.Receipts = b.Receipts[:len(b.Receipts)-1]
}

func (b *Board) resetCastledRook(last *Move) {
	switch last.To.Name {
	case "G1":
		f1 := b.GetSquare(ROW_1, COL_F)
		rook := f1.Piece
		rook.SetMoveCount(0)
		h1 := b.GetSquare(ROW_1, COL_H)
		h1.SetPiece(rook)
		rook.SetSquare(h1)
		f1.SetPiece(&Null{})
	case "C1":
		d1 := b.GetSquare(ROW_1, COL_D)
		rook := d1.Piece
		rook.SetMoveCount(0)
		a1 := b.GetSquare(ROW_1, COL_A)
		a1.SetPiece(rook)
		rook.SetSquare(a1)
		d1.SetPiece(&Null{})
	case "G8":
		f8 := b.GetSquare(ROW_8, COL_F)
		rook := f8.Piece
		rook.SetMoveCount(0)
		h8 := b.GetSquare(ROW_8, COL_H)
		h8.SetPiece(rook)
		rook.SetSquare(h8)
		f8.SetPiece(&Null{})
	case "C8":
		d8 := b.GetSquare(ROW_8, COL_D)
		rook := d8.Piece
		rook.SetMoveCount(0)
		a8 := b.GetSquare(ROW_8, COL_A)
		a8.SetPiece(rook)
		rook.SetSquare(a8)
		d8.SetPiece(&Null{})
	}
}

func (b *Board) addFen(fen string) {
	if _, ok := b.Fens[fen]; ok {
		b.Fens[fen] += 1
		return
	}
	b.Fens[fen] = 1
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

	b.CapturedPieces = append(b.CapturedPieces, capturedPiece)
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
	queen := &Queen{color: move.Turn}
	queen.SetValue()
	b.SetPiece(queen, move.To)
	move.Promotion = queen
	b.Moves = append(b.Moves, move)

	receipt += fmt.Sprintf(" (PROMOTION: QUEEN)")
	return receipt

}

func (b *Board) executeEnPassantMove(move *Move) string {
	captureSq := b.Squares[move.Piece.Square().Row][move.To.Column]
	capturedPiece := captureSq.Piece
	b.CapturedPieces = append(b.CapturedPieces, capturedPiece)
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
	king := b.GetKing(move.Turn)
	b.SetPiece(king, move.To)
	move.From.SetPiece(&Null{})
	b.Moves = append(b.Moves, move)
	king.Castled = true

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

func (b *Board) RandomMove(color string) *Move {
	valids := b.GetAllValidMoves(color)
	for {
		randIdx := rand.Intn(len(valids))
		return valids[randIdx]

	}
}
