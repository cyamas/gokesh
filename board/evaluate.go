package board

import (
	"math"
	"sort"
)

func (b *Board) Evaluate(turn string) {
	if b.DrawDetected() {
		b.Draw = true
		return
	}
	b.Value = 0.0
	b.resetCheck(turn)
	b.resetPins()
	if turn == WHITE {
		b.evaluateWhite()
		blackKing := b.GetKing(BLACK)
		blackKing.SetCheck(b)
		b.evaluateBlack()
		if b.CheckmateDetected(BLACK) {
			b.Checkmate = true
			return
		}
		if b.StalemateDetected(BLACK) {
			b.Stalemate = true
			return
		}
	} else {
		b.evaluateBlack()
		whiteKing := b.GetKing(WHITE)
		whiteKing.SetCheck(b)
		b.evaluateWhite()
		if b.CheckmateDetected(WHITE) {
			b.Checkmate = true
			return
		}
		if b.StalemateDetected(WHITE) {
			b.Stalemate = true
			return
		}
	}
	b.setSquareGuards()
}

func (b *Board) evaluateWhite() {
	for piece := range b.WhitePieces {
		piece.SetActiveSquares(b)
		b.Value += piece.Value()
		if isMinorPiece(piece) && !piece.HasMoved() {
			b.Value -= 0.95
		}
		if king, ok := piece.(*King); ok {
			if !king.Castled {
				b.Value -= 0.95
			}
		}
	}
}

func (b *Board) evaluateBlack() {
	for piece := range b.BlackPieces {
		piece.SetActiveSquares(b)
		b.Value += piece.Value()
		if isMinorPiece(piece) && !piece.HasMoved() {
			b.Value += 0.95
		}
		if king, ok := piece.(*King); ok {
			if !king.Castled {
				b.Value += 0.95
			}
		}
	}
}

func (b *Board) BestMove(turn string) *Move {
	move, _ := b.MiniMax(turn, math.Inf(-1), math.Inf(1), 3)
	return move
}

func (b *Board) SimPosition(move *Move) *Board {
	simBoard := b.Copy()
	simBoard.Evaluate(ENEMY[move.Turn])
	simMove := move.Copy(simBoard)
	simBoard.MovePiece(simMove)
	simBoard.Evaluate(move.Turn)

	return simBoard
}

func (b *Board) MiniMax(turn string, alpha float64, beta float64, depth int) (*Move, float64) {
	if depth == 0 {
		return nil, b.Value
	}
	if b.Draw {
		return nil, 0.0
	}
	if b.Checkmate {
		if turn == WHITE {
			return nil, -99.9
		} else {
			return nil, 99.9
		}
	}
	if b.Stalemate || b.Draw {
		return nil, 0.0
	}

	if turn == WHITE {
		maxEval := math.Inf(-1)
		maxMove := &Move{}

		valids := b.GetAllValidMoves(turn)
		for _, move := range valids {
			sim := b.SimPosition(move)
			_, eval := sim.MiniMax(BLACK, alpha, beta, depth-1)
			if eval > maxEval {
				maxEval = eval
				maxMove = move
			}
			alpha = math.Max(alpha, eval)
			if beta < alpha {
				break
			}
		}
		return maxMove, maxEval

	} else {
		minEval := math.Inf(1)
		minMove := &Move{}
		valids := b.GetAllValidMoves(turn)
		if len(valids) == 0 {
			return nil, 0.0
		}
		for _, move := range valids {
			sim := b.SimPosition(move)
			_, eval := sim.MiniMax(WHITE, alpha, beta, depth-1)
			if eval < minEval {
				minEval = eval
				minMove = move
			}
			beta = math.Min(beta, eval)
			if beta < alpha {
				break
			}
		}
		return minMove, minEval
	}
}

func (b *Board) NextPositions(turn string) map[*Board]*Move {
	positions := make(map[*Board]*Move)
	valids := b.GetAllValidMoves(turn)
	for _, move := range valids {
		positions[b.SimPosition(move)] = move
	}
	return positions
}

type OrderedPositions []*Board

func (op OrderedPositions) Len() int           { return len(op) }
func (op OrderedPositions) Less(i, j int) bool { return op[i].Value < op[j].Value }
func (op OrderedPositions) Swap(i, j int)      { op[i], op[j] = op[j], op[i] }

func orderedPositions(turn string, positions map[*Board]*Move) OrderedPositions {
	ordered := OrderedPositions{}
	for pos := range positions {
		ordered = append(ordered, pos)
	}
	if turn == WHITE {
		sort.Sort(sort.Reverse(OrderedPositions(ordered)))
	} else {
		sort.Sort(OrderedPositions(ordered))
	}
	return ordered
}
