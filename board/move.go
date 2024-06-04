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
	valids := m.Piece.ActiveSquares(board)
	moveType, ok := valids[m.To]
	if ok {
		m.Type = moveType
		return true
	}
	return false
}
