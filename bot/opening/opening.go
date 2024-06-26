package opening

import (
	"github.com/cyamas/gokesh/board"
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

	PAWN   = "PAWN"
	KNIGHT = "KNIGHT"
	BISHOP = "BISHOP"
	ROOK   = "ROOK"
	QUEEN  = "QUEEN"
	KING   = "KING"
	NULL   = "NULL"

	CARO_KANN = "CAROKANN"
	LONDON    = "LONDON"

	CLASSICAL = "CLASSICAL"
	EXCHANGE  = "EXCHANGE"
	ADVANCE   = "ADVANCE"
	MAIN      = "MAIN"
)

var ENEMY = map[string]string{
	WHITE: BLACK,
	BLACK: WHITE,
}

type MoveFunc func(*board.Board) *board.Move

type Opening struct {
	Name      string
	Variation string
	Moves     map[string]MoveFunc
}

func Play(opening string, brd *board.Board) *Opening {
	o := &Opening{
		Name:      opening,
		Variation: MAIN,
	}
	switch opening {
	case CARO_KANN:
		o.Moves = CaroKannMoves
	case LONDON:
		o.Moves = LondonMoves
	}
	return o
}

func (o *Opening) NextMove(brd *board.Board) *board.Move {
	fen := brd.Fen()
	if move, ok := o.Moves[fen]; ok {
		return move(brd)
	}
	return nil
}

func (o *Opening) SetVariation(fen string) {
	switch o.Name {
	case CARO_KANN:
		switch fen {
		case "rnbqkbnr/pp2pppp/2p5/3P4/3P4/8/PPP2PPP/RNBQKBNR":
			o.Variation = EXCHANGE
		case "rnbqkbnr/pp2pppp/2p5/3pP3/3P4/8/PPP2PPP/RNBQKBNR":
			o.Variation = ADVANCE
		}
	}
}
