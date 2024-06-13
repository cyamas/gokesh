package game

import (
	"bufio"
	"fmt"
	"io"
	"strings"

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
	Kings map[string]*board.King
	Turn  string
}

func New(b *board.Board) *Game {
	return &Game{
		Board: b,
		Kings: map[string]*board.King{WHITE: b.GetKing(WHITE), BLACK: b.GetKing(BLACK)},
		Turn:  WHITE,
	}
}

func (g *Game) Run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		turnPrompt := fmt.Sprintf("%s's MOVE: ", g.Turn)
		fmt.Fprint(out, turnPrompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		moveMsg := scanner.Text()
		move, err := g.generateMove(moveMsg)
		if err != nil {
			fmt.Fprint(out, err.Message+"\n")
			continue
		}
		receipt := g.ExecuteTurn(move)
		fmt.Fprint(out, receipt+"\n")
		if g.Board.Checkmate {
			break
		}
	}
}

func (g *Game) generateMove(msg string) (*board.Move, *Error) {

	moveParts := strings.Split(msg, " ")

	pieceStr := strings.TrimSpace(moveParts[0])
	fromStr := strings.TrimSpace(moveParts[1])
	toStr := strings.TrimSpace(moveParts[2])

	fromCoords := squareMap[fromStr]
	toCoords := squareMap[toStr]

	fromSq := g.Board.Squares[fromCoords[0]][fromCoords[1]]
	toSq := g.Board.Squares[toCoords[0]][toCoords[1]]
	piece := fromSq.Piece
	if piece.Type() == NULL {
		err := NewError("ERROR: %s HAS NO PIECE", fromSq.Name)
		return nil, err
	}
	if pieceStr != piece.Type() {
		err := NewError("ERROR: SQUARE DOES NOT HAVE SPECIFIED PIECE")
		return nil, err
	}

	return &board.Move{Turn: g.Turn, Piece: piece, From: fromSq, To: toSq}, nil

}

func (g *Game) ExecuteTurn(move *board.Move) string {
	moveColor := move.Piece.Color()
	if moveColor != g.Turn {
		msg := fmt.Sprintf("%s ERROR: It is %s's turn to move", move.Piece.Color(), g.Turn)
		return msg
	}

	receipt, err := g.Board.MovePiece(move)
	if err != nil {
		boardErr := NewError("BOARD ERROR %s: %s", moveColor, receipt)
		return boardErr.Message
	}
	receipt = g.Turn + " " + receipt
	if g.Board.Checkmate {
		return fmt.Sprintf("%s\nCHECKMATE: %s has won", receipt, g.Turn)
	}

	g.nextTurn()
	return receipt
}

func (g *Game) nextTurn() {
	if g.Turn == WHITE {
		g.Turn = BLACK
	} else {
		g.Turn = WHITE
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
