package game

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
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

type Bot struct {
	name  string
	Color string
}

func (b *Bot) Name() string { return b.name }

func (b *Bot) CreateMove(board *board.Board) *board.Move {
	valids := board.GetAllValidMoves(b.Color)
	randIdx := rand.Intn(len(valids))
	return valids[randIdx]
}

type Game struct {
	Board *board.Board
	Bot   *Bot
	Turn  string
}

func New(b *board.Board) *Game {
	colors := []string{WHITE, BLACK}
	botColor := colors[rand.Intn(2)]
	return &Game{
		Board: b,
		Bot:   &Bot{name: "GOKESH", Color: botColor},
		Turn:  WHITE,
	}
}

/*func (g *Game) Run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		var moveMsg string
		turnPrompt := fmt.Sprintf("%s's MOVE: ", g.Turn)
		fmt.Fprint(out, turnPrompt)

		if g.Bot.Color == g.Turn {
			moveMsg = g.Bot.CreateMove(g.Board)
		} else {
			scanned := scanner.Scan()
			if !scanned {
				return
			}
			moveMsg = scanner.Text()
		}

		move, err := g.createMove(moveMsg)
		if err != nil {
			fmt.Fprint(out, err.Message+"\n")
			continue
		}
		if move.Piece.Type() == PAWN && (move.To.Row == ROW_1 || move.To.Row == ROW_8) {
			g.handlePawnPromotion(move, out, scanner)
		}
		receipt := g.ExecuteTurn(move)
		fmt.Fprint(out, receipt+"\n")
		if g.Board.Checkmate {
			break
		}
	}
}*/

func (g *Game) createMove(msg string) (*board.Move, *Error) {

	moveParts := strings.Split(msg, " ")
	if len(moveParts) != 3 {
		err := NewError("ERROR: INVALID INPUT. ENTER A VALID MOVE")
		return nil, err
	}

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

	move := &board.Move{
		Turn:  g.Turn,
		Piece: piece,
		From:  fromSq,
		To:    toSq,
	}

	return move, nil

}

func (g *Game) ExecuteTurn(move *board.Move) (string, *Error) {
	moveColor := move.Piece.Color()

	receipt, err := g.Board.MovePiece(move)
	if err != nil {
		if g.Board.Check {
			receipt += " (KING IN CHECK)"
		}
		if move.Piece.Pin() != nil {
			receipt += " (PIECE IS PINNED)"
		}
		boardErr := NewError("BOARD ERROR %s: %s", moveColor, receipt)
		return boardErr.Message, boardErr
	}
	receipt = g.Turn + " " + receipt
	if g.Board.Checkmate {
		return fmt.Sprintf("%s\nCHECKMATE: %s has won", receipt, g.Turn), nil
	}

	if g.Board.Check {
		receipt += fmt.Sprintf("\n%s IN CHECK", ENEMY[g.Turn])
	}

	receipt += fmt.Sprintf("\nBOARD VALUE: %f", g.Board.Value)
	g.nextTurn()
	return receipt, nil
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
			fmt.Println("ERROR: INVALID PIECE FOR PROMOTION")
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
