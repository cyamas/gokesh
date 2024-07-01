package game

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"

	"github.com/cyamas/gokesh/board"
	"github.com/cyamas/gokesh/bot"
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
)

var squareMap = map[string][2]int{
	"A1": {ROW_1, COL_A},
	"A2": {ROW_2, COL_A},
	"A3": {ROW_3, COL_A},
	"A4": {ROW_4, COL_A},
	"A5": {ROW_5, COL_A},
	"A6": {ROW_6, COL_A},
	"A7": {ROW_7, COL_A},
	"A8": {ROW_8, COL_A},
	"B1": {ROW_1, COL_B},
	"B2": {ROW_2, COL_B},
	"B3": {ROW_3, COL_B},
	"B4": {ROW_4, COL_B},
	"B5": {ROW_5, COL_B},
	"B6": {ROW_6, COL_B},
	"B7": {ROW_7, COL_B},
	"B8": {ROW_8, COL_B},
	"C1": {ROW_1, COL_C},
	"C2": {ROW_2, COL_C},
	"C3": {ROW_3, COL_C},
	"C4": {ROW_4, COL_C},
	"C5": {ROW_5, COL_C},
	"C6": {ROW_6, COL_C},
	"C7": {ROW_7, COL_C},
	"C8": {ROW_8, COL_C},
	"D1": {ROW_1, COL_D},
	"D2": {ROW_2, COL_D},
	"D3": {ROW_3, COL_D},
	"D4": {ROW_4, COL_D},
	"D5": {ROW_5, COL_D},
	"D6": {ROW_6, COL_D},
	"D7": {ROW_7, COL_D},
	"D8": {ROW_8, COL_D},
	"E1": {ROW_1, COL_E},
	"E2": {ROW_2, COL_E},
	"E3": {ROW_3, COL_E},
	"E4": {ROW_4, COL_E},
	"E5": {ROW_5, COL_E},
	"E6": {ROW_6, COL_E},
	"E7": {ROW_7, COL_E},
	"E8": {ROW_8, COL_E},
	"F1": {ROW_1, COL_F},
	"F2": {ROW_2, COL_F},
	"F3": {ROW_3, COL_F},
	"F4": {ROW_4, COL_F},
	"F5": {ROW_5, COL_F},
	"F6": {ROW_6, COL_F},
	"F7": {ROW_7, COL_F},
	"F8": {ROW_8, COL_F},
	"G1": {ROW_1, COL_G},
	"G2": {ROW_2, COL_G},
	"G3": {ROW_3, COL_G},
	"G4": {ROW_4, COL_G},
	"G5": {ROW_5, COL_G},
	"G6": {ROW_6, COL_G},
	"G7": {ROW_7, COL_G},
	"G8": {ROW_8, COL_G},
	"H1": {ROW_1, COL_H},
	"H2": {ROW_2, COL_H},
	"H3": {ROW_3, COL_H},
	"H4": {ROW_4, COL_H},
	"H5": {ROW_5, COL_H},
	"H6": {ROW_6, COL_H},
	"H7": {ROW_7, COL_H},
	"H8": {ROW_8, COL_H},
}

var ENEMY = map[string]string{
	WHITE: BLACK,
	BLACK: WHITE,
}

type Game struct {
	Board *board.Board
	Bot   *bot.Bot
	Turn  string
}

func New(b *board.Board) *Game {
	colors := []string{WHITE, BLACK}
	botColor := colors[rand.Intn(2)]
	return &Game{
		Board: b,
		Bot:   &bot.Bot{Name: "Gokesh", Color: botColor},
		Turn:  WHITE,
	}
}

func (g *Game) ExecuteTurn(move *board.Move) (string, *Error) {
	receipt, err := g.Board.MovePiece(move)
	if err != nil {
		return g.handleBoardError(receipt, move)
	}

	receipt = fmt.Sprintf("%s %s\n", g.Turn, receipt)
	if g.Board.Checkmate {
		return fmt.Sprintf("%s\nCHECKMATE: %s has won", receipt, g.Turn), nil
	}
	if g.Board.Stalemate {
		return fmt.Sprintf("%s\nSTALEMATE: GAME IS A DRAW", receipt), nil
	}
	if g.Board.Draw {
		return fmt.Sprintf("%s\nDRAW: By repetition", receipt), nil
	}
	if g.Board.GetKing(g.Turn).Checked {
		receipt += fmt.Sprintf("\n%s IN CHECK", ENEMY[g.Turn])
	}

	receipt += fmt.Sprintf("\nBOARD VALUE: %f", g.Board.Value)
	g.nextTurn()
	return receipt, nil
}

func (g *Game) handleBoardError(receipt string, move *board.Move) (string, *Error) {
	if g.Board.GetKing(g.Turn).Checked {
		receipt += " (KING IN CHECK)"
	}
	if move.Piece.Pin() != nil {
		receipt += " (PIECE IS PINNED)"
	}
	boardErr := NewError("BOARD ERROR %s: %s", g.Turn, receipt)
	return boardErr.Message, boardErr
}

func (g *Game) nextTurn() {
	if g.Turn == WHITE {
		g.Turn = BLACK
	} else {
		g.Turn = WHITE
	}
}

func (g *Game) handlePawnPromotion(move *board.Move, out io.Writer, scanner *bufio.Scanner) {
	for {
		promotePrompt := "PROMOTE TO: "
		fmt.Fprint(out, promotePrompt)
		scanned := scanner.Scan()
		if !scanned {
			continue
		}
		promoteMsg := scanner.Text()
		switch promoteMsg {
		case QUEEN:
			queen := &board.Queen{}
			queen.SetColor(g.Turn)
			move.Promotion = queen
			return
		case ROOK:
			rook := &board.Rook{}
			rook.SetColor(g.Turn)
			move.Promotion = rook
			return
		case BISHOP:
			bishop := &board.Bishop{}
			bishop.SetColor(g.Turn)
			move.Promotion = bishop
			return
		case KNIGHT:
			knight := &board.Knight{}
			knight.SetColor(g.Turn)
			move.Promotion = knight
			return
		default:
			continue
		}
	}

}

type Error struct {
	Message string
}

func NewError(format string, a ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(format, a...),
	}
}
