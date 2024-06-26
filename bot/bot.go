package bot

import (
	"github.com/cyamas/gokesh/board"
	"github.com/cyamas/gokesh/bot/opening"
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

	LONDON    = "LONDON"
	CARO_KANN = "CAROKANN"
)

var ENEMY = map[string]string{
	WHITE: BLACK,
	BLACK: WHITE,
}

type Bot struct {
	Name    string
	Color   string
	Opening *opening.Opening
}

func (b *Bot) Move(board *board.Board) *board.Move {
	if len(board.Moves) <= 15 {
		return b.handleOpening(board)
	}
	return board.BestMove(b.Color)
}

func (b *Bot) handleThreats(threats map[*board.Threat]bool, brd *board.Board) *board.Move {
	for threat := range threats {
		ally := threat.Ally
		guards, _ := ally.Square().GetGuardsAndValue(b.Color)
		attackers, _ := ally.Square().GetGuardsAndValue(ENEMY[b.Color])

		attackerSq := threat.Attacker.Square()
		attackerAttackers, allyValue := attackerSq.GetGuardsAndValue(b.Color)
		attackerGuards, enemyValue := attackerSq.GetGuardsAndValue(ENEMY[b.Color])

		switch {
		case len(attackerAttackers) == 0 || len(guards) < len(attackers):
			retreatSq := threat.Ally.BestRetreatSquare(brd)
			if retreatSq == nil {
				return brd.RandomMove(b.Color)
			}
			return &board.Move{
				Turn:  b.Color,
				Piece: threat.Ally,
				From:  threat.Ally.Square(),
				To:    retreatSq,
			}
		case len(attackerGuards) == len(attackerAttackers):
			switch {
			case allyValue <= enemyValue:
				if b.Color == WHITE {
					return &board.Move{
						Turn:  b.Color,
						Piece: attackerSq.WhiteGuards[0],
						From:  attackerSq.WhiteGuards[0].Square(),
						To:    attackerSq,
					}
				} else {
					return &board.Move{
						Turn:  b.Color,
						Piece: attackerSq.BlackGuards[0],
						From:  attackerSq.BlackGuards[0].Square(),
						To:    attackerSq,
					}
				}
			}
		}
	}
	return brd.RandomMove(b.Color)
}

func (b *Bot) handleOpening(brd *board.Board) *board.Move {
	if b.Color == WHITE {
		if b.Opening == nil {
			b.Opening = opening.Play(LONDON, brd)
		}
		move := b.Opening.NextMove(brd)
		if move != nil {
			return move
		}
	} else {
		if b.Opening == nil {
			b.Opening = opening.Play(CARO_KANN, brd)
		}
		move := b.Opening.NextMove(brd)
		if move != nil {
			return move
		}
	}
	return brd.BestMove(b.Color)
}
