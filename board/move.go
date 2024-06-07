package board

const (
	FREE       = "FREE"
	CAPTURE    = "CAPTURE"
	GUARDED    = "GUARDED"
	EN_PASSANT = "ENPASSANT"
	CASTLE     = "CASTLE"
)

type SqActivity string

type Move struct {
	Type      SqActivity
	Piece     Piece
	From      *Square
	To        *Square
	Promotion Piece
}

func (m *Move) IsValid(board *Board) bool {
	if m.Piece.Type() != KING {
		color := m.Piece.Color()
		unsafes := board.GetAttackedSquares(color)

		king := board.GetKing(color)
		if king != nil {
			check, checkingPieces := king.IsInCheck(unsafes)
			if check {
				return m.stopsCheck(king, checkingPieces, board)
			}
			isPinned := m.pieceIsPinned(king, board)
			if isPinned {
				return false
			}
		}
	}

	valids := m.Piece.ActiveSquares(board)
	moveType, ok := valids[m.To]
	if ok {
		m.Type = moveType
		return true
	}
	return false
}

func (m *Move) stopsCheck(king *King, checkingPieces []Piece, board *Board) bool {
	if len(checkingPieces) > 1 {
		return false
	}
	checker := checkingPieces[0]
	checkerSq := checker.Square()

	board.removePiece(checker)
	piecePinned := m.pieceIsPinned(king, board)
	board.SetPiece(checker, checkerSq)

	switch {
	case piecePinned:
		return false
	case checker.Square() == m.To:
		return true
	case checker.Type() == KNIGHT:
		return false
	}

	kingSq := king.Square()
	checkPath := board.GetAttackedPath(checkerSq, kingSq)
	if checkPath[m.To] {
		m.Type = FREE
		return true
	}
	return false
}

func (m *Move) pieceIsPinned(king *King, board *Board) bool {
	m.From.SetPiece(&Null{})
	unsafes := board.GetAttackedSquares(king.Color())
	result, _ := king.IsInCheck(unsafes)
	m.From.SetPiece(m.Piece)

	return result
}
