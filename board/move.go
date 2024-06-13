package board

const (
	FREE       = "FREE"
	CAPTURE    = "CAPTURE"
	GUARDED    = "GUARDED"
	EN_PASSANT = "ENPASSANT"
	CASTLE     = "CASTLE"
	CHECK      = "CHECK"
)

type SqActivity string

type Move struct {
	Turn      string
	Type      SqActivity
	Piece     Piece
	From      *Square
	To        *Square
	Promotion Piece
}

func (m *Move) IsValid(board *Board) bool {
	board.Evaluate(m.Turn)
	valids := m.Piece.ActiveSquares()
	moveType, ok := valids[m.To]
	if ok {
		m.Type = moveType
		return true
	}
	return false
}
