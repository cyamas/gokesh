package board

import (
	"fmt"
	"math"
)

func (b *Board) Evaluate(turn string) {
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
	if b.DrawDetected() {
		b.Draw = true
		return
	}
}

func (b *Board) evaluateWhite() {
	for piece := range b.WhitePieces {
		piece.SetActiveSquares(b)
		b.Value += piece.Value()
		if king, ok := piece.(*King); ok {
			if king.Castled {
				b.Value += 0.33
			}
		}
		switch {
		case isMinorPiece(piece) && !piece.HasMoved():
			b.Value -= 0.33
		}
	}
}

func (b *Board) evaluateBlack() {
	for piece := range b.BlackPieces {
		piece.SetActiveSquares(b)
		b.Value += piece.Value()
		if king, ok := piece.(*King); ok {
			if king.Castled {
				b.Value -= 0.33
			}
		}
		switch {
		case isMinorPiece(piece) && !piece.HasMoved():
			b.Value += 0.33
		}
	}
}

func (b *Board) BestMove(turn string) *Move {
	move, _ := b.MiniMax(turn, math.Inf(-1), math.Inf(1), 4)
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
	if b.Checkmate {
		fmt.Println("CHECKMATE DETECTED")
		if turn == WHITE {
			return nil, -99.9
		} else {
			return nil, 99.9
		}
	}
	if b.Stalemate || b.Draw {
		fmt.Println("DRAW DETECTED")
		return nil, 0.0
	}

	if turn == WHITE {
		maxEval := math.Inf(-1)
		maxMove := &Move{}

		valids := b.GetAllValidMoves(turn)

		for _, move := range valids {
			b.MovePiece(move)
			_, eval := b.MiniMax(BLACK, alpha, beta, depth-1)
			//fmt.Println("DEPTH: ", depth)
			b.UndoMove()
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

		for _, move := range valids {
			b.MovePiece(move)
			_, eval := b.MiniMax(WHITE, alpha, beta, depth-1)
			b.UndoMove()
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
