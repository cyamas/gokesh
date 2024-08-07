package board

import (
	"fmt"
	"strconv"
	"strings"
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
)

var ENEMY = map[string]string{
	WHITE: BLACK,
	BLACK: WHITE,
}

type Board struct {
	Squares        [][]*Square
	Moves          []*Move
	WhitePieces    map[Piece]bool
	BlackPieces    map[Piece]bool
	PromotedPawns  []*Pawn
	CapturedPieces []Piece
	Checkmate      bool
	Stalemate      bool
	Draw           bool
	Value          float64
	Fens           map[string]int
	Receipts       []string
}

func New() *Board {
	board := &Board{
		WhitePieces: make(map[Piece]bool),
		BlackPieces: make(map[Piece]bool),
	}

	rows := []string{"8", "7", "6", "5", "4", "3", "2", "1"}
	cols := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	for i := range 8 {
		boardRow := []*Square{}
		rowNum := rows[i]
		for j := range 8 {
			colLetter := cols[j]
			square := &Square{
				Row:    i,
				Column: j,
				Name:   colLetter + rowNum,
				Piece:  &Null{},
			}
			boardRow = append(boardRow, square)
		}
		board.Squares = append(board.Squares, boardRow)
	}
	board.Fens = map[string]int{}

	return board
}

func (b *Board) Copy() *Board {
	copy := New()
	for i, row := range copy.Squares {
		for j, sq := range row {
			ogSq := b.Squares[i][j]
			switch piece := ogSq.Piece.(type) {
			case *Pawn:
				pawn := &Pawn{color: piece.color, value: piece.value, moveCount: piece.moveCount}
				copy.SetPiece(pawn, sq)
			case *Knight:
				knight := &Knight{color: piece.color, value: piece.value, moveCount: piece.moveCount}
				copy.SetPiece(knight, sq)
			case *Bishop:
				bishop := &Bishop{color: piece.color, value: piece.value, moveCount: piece.moveCount}
				copy.SetPiece(bishop, sq)
			case *Rook:
				rook := &Rook{color: piece.color, value: piece.value, moveCount: piece.moveCount}
				copy.SetPiece(rook, sq)
			case *Queen:
				queen := &Queen{color: piece.color, value: piece.value, moveCount: piece.moveCount}
				copy.SetPiece(queen, sq)
			case *King:
				king := &King{color: piece.color, value: piece.value, moveCount: piece.moveCount}
				copy.SetPiece(king, sq)
			default:
				copy.SetPiece(&Null{}, sq)
			}
		}
	}
	for _, receipt := range b.Receipts {
		copy.Receipts = append(copy.Receipts, receipt)
	}
	for fen, count := range b.Fens {
		copy.Fens[fen] = count
	}
	for _, move := range b.Moves {
		copy.Moves = append(copy.Moves, move)
	}
	return copy
}

type Error struct {
	Message string
}

func NewError(format string, a ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(format, a...),
	}
}

type Square struct {
	Piece       Piece
	Row         int
	Column      int
	Name        string
	WhiteGuards []Piece
	BlackGuards []Piece
}

type Path map[*Square]bool

func (s *Square) GetGuardsAndValue(color string) ([]Piece, float64) {
	var value float64 = 0.0
	var guards []Piece
	if color == WHITE {
		guards = s.WhiteGuards
	} else {
		guards = s.BlackGuards
	}
	for _, piece := range guards {
		value += piece.Value()
	}
	return guards, value
}

func (s *Square) SetPiece(piece Piece) {
	s.Piece = piece
}

func (s *Square) IsEmpty() bool {
	if s.Piece.Type() == NULL {
		return true
	}
	return false
}

func (b *Board) GetSquare(row, col int) *Square {
	return b.Squares[row][col]
}

func (b *Board) CreatePiece(color string, name string) Piece {
	switch name {
	case PAWN:
		p := &Pawn{color: color, value: 1}
		if color == BLACK {
			p.value *= float64(-1)
		}
		return p
	case KNIGHT:
		kn := &Knight{color: color, value: 3.05}
		if color == BLACK {
			kn.value *= float64(-1)
		}
		return kn
	case BISHOP:
		bish := &Bishop{color: color, value: 3.33}
		if color == BLACK {
			bish.value *= float64(-1)
		}
		return bish
	case ROOK:
		r := &Rook{color: color}
		if color == BLACK {
			r.value *= float64(-1)
		}
		return r
	case QUEEN:
		q := &Queen{color: color}
		if color == BLACK {
			q.value *= float64(-1)
		}
		return q
	case KING:
		k := &King{color: color, value: 99.9}
		if color == BLACK {
			k.value *= float64(-1)
		}
		return k
	default:
		return &Null{}
	}
}

func (b *Board) resetCheck(color string) {
	king := b.GetKing(color)
	king.Checked = false
	king.Checkers = []Piece{}
}

func (b *Board) setSquareGuards() {
	blackAttackedSqs := b.GetAttackedSquares(WHITE)
	whiteAttackedSqs := b.GetAttackedSquares(BLACK)
	for _, row := range b.Squares {
		for _, sq := range row {
			if blackPieces, ok := blackAttackedSqs[sq]; ok {
				sq.BlackGuards = blackPieces
			} else {
				sq.BlackGuards = []Piece{}
			}
			if whitePieces, ok := whiteAttackedSqs[sq]; ok {
				sq.WhiteGuards = whitePieces
			} else {
				sq.WhiteGuards = []Piece{}
			}
		}
	}
}

func (b *Board) Fen() string {
	fen := ""

	for i, row := range b.Squares {
		emptySqs := 0
		for _, sq := range row {
			if sq.Piece.Type() != NULL && emptySqs > 0 {
				fen += strconv.Itoa(emptySqs)
				emptySqs = 0
			}

			switch sq.Piece.Type() {
			case NULL:
				emptySqs += 1
			case PAWN:
				if sq.Piece.Color() == WHITE {
					fen += "P"
				} else {
					fen += "p"
				}
			case KNIGHT:
				if sq.Piece.Color() == WHITE {
					fen += "N"
				} else {
					fen += "n"
				}
			case BISHOP:
				if sq.Piece.Color() == WHITE {
					fen += "B"
				} else {
					fen += "b"
				}
			case ROOK:
				if sq.Piece.Color() == WHITE {
					fen += "R"
				} else {
					fen += "r"
				}
			case QUEEN:
				if sq.Piece.Color() == WHITE {
					fen += "Q"
				} else {
					fen += "q"
				}
			case KING:
				if sq.Piece.Color() == WHITE {
					fen += "K"
				} else {
					fen += "k"
				}
			}
		}
		if emptySqs > 0 {
			fen += strconv.Itoa(emptySqs)
		}
		if i < 7 {
			fen += "/"
		}
	}
	return fen
}

func (b *Board) SetupFromFen(fen string) {
	fen = strings.ReplaceAll(fen, "/", "")
	runes := []rune(fen)
	idx := 0
	runeDigits := map[rune]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8}
	empties := 0
	for _, row := range b.Squares {
		for _, sq := range row {
			if empties > 0 {
				empties--
			} else {
				ch := runes[idx]
				switch ch {
				case 'p':
					p := &Pawn{color: BLACK, value: -1.00}
					if sq.Row != ROW_7 {
						p.moveCount = 1
					}
					b.SetPiece(p, sq)
				case 'P':
					p := &Pawn{color: WHITE, value: 1.00}
					if sq.Row != ROW_2 {
						p.moveCount = 1
					}
					b.SetPiece(p, sq)
				case 'n':
					kn := &Knight{color: BLACK, value: -3.05}
					b.SetPiece(kn, sq)
				case 'N':
					kn := &Knight{color: WHITE, value: 3.05}
					b.SetPiece(kn, sq)
				case 'b':
					bish := &Bishop{color: BLACK, value: -3.33}
					b.SetPiece(bish, sq)
				case 'B':
					bish := &Bishop{color: WHITE, value: 3.33}
					b.SetPiece(bish, sq)
				case 'r':
					r := &Rook{color: BLACK, value: -5.63, moveCount: 1}
					b.SetPiece(r, sq)
				case 'R':
					r := &Rook{color: WHITE, value: 5.63, moveCount: 1}
					b.SetPiece(r, sq)
				case 'q':
					q := &Queen{color: BLACK, value: -9.5}
					b.SetPiece(q, sq)
				case 'Q':
					q := &Queen{color: WHITE, value: 9.5}
					b.SetPiece(q, sq)
				case 'k':
					k := &King{color: BLACK, value: -99.9}
					b.SetPiece(k, sq)
				case 'K':
					k := &King{color: WHITE, value: 99.9}
					b.SetPiece(k, sq)
				default:
					empties = runeDigits[ch] - 1
				}
				idx++
			}
		}
	}
}

func (b *Board) DrawDetected() bool {
	if b.DrawByRepetition() {
		return true
	}
	if b.DrawByInsufficientMaterial() {
		return true
	}
	return false
}

func (b *Board) DrawByRepetition() bool {
	currFen := b.Fen()
	if count, ok := b.Fens[currFen]; ok {
		if count > 2 {
			return true
		}
	}
	return false
}

func (b *Board) DrawByInsufficientMaterial() bool {
	whiteValue := 0.0
	whitePawns := 0
	blackValue := 0.0
	blackPawns := 0

	for piece := range b.WhitePieces {
		if piece.Type() == PAWN {
			whitePawns++
		}
		whiteValue += piece.Value()
	}
	for piece := range b.BlackPieces {
		if piece.Type() == PAWN {
			blackPawns++
		}
		blackValue -= piece.Value()
	}

	if whiteValue < 104.0 && blackValue < 104.0 && whitePawns == 0 && blackPawns == 0 {
		return true
	}

	return false
}

func (b *Board) resetPins() {
	for wp := range b.WhitePieces {
		wp.ResetPin()
	}
	for bp := range b.BlackPieces {
		bp.ResetPin()
	}
}

func (b *Board) CheckmateDetected(color string) bool {
	king := b.GetKing(color)
	if !king.Checked {
		return false
	}

	kingActives := king.ActiveSquares()

	switch {
	case len(kingActives) == 0 && len(king.Checkers) > 1:
		return true
	case len(kingActives) > 0 && king.CanEvadeCheck(kingActives, b):
		return false
	default:
		return !b.piecePreventsCheckmate(king)
	}
}

func (b *Board) StalemateDetected(color string) bool {
	allies := b.getAllies(color)
	for ally := range allies {
		for _, activity := range ally.ActiveSquares() {
			if activity != GUARDED {
				return false
			}
		}
	}
	return true
}

func (b *Board) piecePreventsCheckmate(king *King) bool {
	checker := king.Checkers[0]
	checkerSq := checker.Square()
	checkPath := b.GetAttackedPath(checkerSq, king.Square())
	allies := b.getAllies(king.Color())
	for piece := range allies {
		if piece.Type() == KING {
			continue
		}
		pieceActives := piece.ActiveSquares()
		for sq := range checkPath {
			if moveType, ok := pieceActives[sq]; ok {
				if moveType == FREE || moveType == CAPTURE {
					return true
				}
			}
		}
	}
	return false
}

func (b *Board) getAllies(color string) map[Piece]bool {
	if color == WHITE {
		return b.WhitePieces
	} else {
		return b.BlackPieces
	}
}

func (b *Board) getEnemies(color string) map[Piece]bool {
	if color == WHITE {
		return b.BlackPieces
	} else {
		return b.WhitePieces
	}
}

func (b *Board) GetAttackedSquares(color string) map[*Square][]Piece {
	enemies := b.getEnemies(color)

	attackedSqs := make(map[*Square][]Piece)

	for piece := range enemies {
		if king, ok := piece.(*King); ok {
			b.addKingAttackedSquares(king, attackedSqs)
			continue
		} else {
			pieceActives := piece.ActiveSquares()
			for sq, sqActivity := range pieceActives {
				if _, ok := attackedSqs[sq]; !ok {
					if piece.Type() == PAWN && sqActivity == FREE {
						continue
					}
				}
				addAttackedSquare(sq, piece, attackedSqs)
			}
		}
	}
	return attackedSqs
}

func (b *Board) addKingAttackedSquares(king *King, attackedSqs map[*Square][]Piece) {
	for _, dir := range KING_DIRS {
		row := king.Square().Row + dir[0]
		col := king.Square().Column + dir[1]

		if cand, ok := b.GetSquareIfExists(row, col); ok {
			addAttackedSquare(cand, king, attackedSqs)
		}
	}
}

func (b *Board) checkSquarePastKing(row, col int, coords [2]int) (*Square, bool) {
	candRow := row + coords[0]
	candCol := col + coords[1]
	if cand, ok := b.GetSquareIfExists(candRow, candCol); ok {
		return cand, true
	}
	return nil, false
}

func addAttackedSquare(sq *Square, piece Piece, attackeds map[*Square][]Piece) {
	_, ok := attackeds[sq]
	if ok {
		attackeds[sq] = append(attackeds[sq], piece)
	} else {
		attackeds[sq] = []Piece{piece}
	}
}

func (b *Board) GetAttackedPath(from, to *Square) map[*Square]bool {
	if from.Piece.Type() == KNIGHT {
		return map[*Square]bool{from: true}
	}
	bigRow, smallRow := compareCoords(from.Row, to.Row)
	bigCol, smallCol := compareCoords(from.Column, to.Column)
	switch {
	case from.Row == to.Row:
		return b.horizontalPath(bigCol, smallCol, bigRow)
	case from.Column == to.Column:
		return b.verticalPath(bigRow, smallRow, bigCol)
	default:
		return b.diagonalPath(from, to)
	}
}

func (b *Board) horizontalPath(bigCol, smallCol, row int) map[*Square]bool {
	path := make(map[*Square]bool)

	for col := smallCol; col <= bigCol; col++ {
		sq := b.Squares[row][col]
		if sq.Piece.Type() == KING {
			continue
		}
		path[sq] = true
	}
	return path
}

func (b *Board) verticalPath(bigRow, smallRow, col int) map[*Square]bool {
	path := make(map[*Square]bool)
	for row := smallRow; row <= bigRow; row++ {
		sq := b.Squares[row][col]
		if sq.Piece.Type() == KING {
			continue
		}
		path[sq] = true
	}
	return path
}

func (b *Board) diagonalPath(from, to *Square) map[*Square]bool {
	path := make(map[*Square]bool)
	rows := orderCoords(from.Row, to.Row)
	cols := orderCoords(from.Column, to.Column)
	for i := range rows {
		sq := b.Squares[rows[i]][cols[i]]
		path[sq] = true
	}

	return path
}

func orderCoords(from, to int) []int {
	var coords []int
	if from-to < 0 {
		for coord := from; coord < to; coord++ {
			coords = append(coords, coord)
		}
	} else {
		for coord := from; coord > to; coord-- {
			coords = append(coords, coord)
		}
	}
	return coords
}

func compareCoords(a, b int) (int, int) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}

func (b *Board) GetKing(color string) *King {
	if color == WHITE {
		for piece := range b.WhitePieces {
			if king, ok := piece.(*King); ok {
				return king
			}
		}
	} else {
		for piece := range b.BlackPieces {
			if king, ok := piece.(*King); ok {
				return king
			}
		}
	}
	return nil
}

func (b *Board) SetPiece(piece Piece, square *Square) {
	piece.SetSquare(square)
	square.SetPiece(piece)
	if piece.Color() == WHITE {
		b.WhitePieces[piece] = true
	} else {
		b.BlackPieces[piece] = true
	}
}

func (b *Board) SetupPieces() {
	var whiteStartSquares = map[string]Piece{
		"A2": &Pawn{value: 1.0, moveCount: 0},
		"B2": &Pawn{value: 1.0, moveCount: 0},
		"C2": &Pawn{value: 1.0, moveCount: 0},
		"D2": &Pawn{value: 1.0, moveCount: 0},
		"E2": &Pawn{value: 1.0, moveCount: 0},
		"F2": &Pawn{value: 1.0, moveCount: 0},
		"G2": &Pawn{value: 1.0, moveCount: 0},
		"H2": &Pawn{value: 1.0, moveCount: 0},
		"B1": &Knight{value: 3.05, moveCount: 0},
		"G1": &Knight{value: 3.05, moveCount: 0},
		"C1": &Bishop{value: 3.33, moveCount: 0},
		"F1": &Bishop{value: 3.33, moveCount: 0},
		"A1": &Rook{value: 5.63, CastleSq: b.Squares[ROW_1][COL_D], moveCount: 0},
		"H1": &Rook{value: 5.63, CastleSq: b.Squares[ROW_1][COL_F], moveCount: 0},
		"D1": &Queen{value: 9.5, moveCount: 0},
		"E1": &King{value: 99.9, moveCount: 0},
	}

	var blackStartSquares = map[string]Piece{
		"A7": &Pawn{value: -1.0, moveCount: 0},
		"B7": &Pawn{value: -1.0, moveCount: 0},
		"C7": &Pawn{value: -1.0, moveCount: 0},
		"D7": &Pawn{value: -1.0, moveCount: 0},
		"E7": &Pawn{value: -1.0, moveCount: 0},
		"F7": &Pawn{value: -1.0, moveCount: 0},
		"G7": &Pawn{value: -1.0, moveCount: 0},
		"H7": &Pawn{value: -1.0, moveCount: 0},
		"B8": &Knight{value: -3.05, moveCount: 0},
		"G8": &Knight{value: -3.05, moveCount: 0},
		"C8": &Bishop{value: -3.33, moveCount: 0},
		"F8": &Bishop{value: -3.33, moveCount: 0},
		"A8": &Rook{value: -5.63, CastleSq: b.Squares[ROW_8][COL_D], moveCount: 0},
		"H8": &Rook{value: -5.63, CastleSq: b.Squares[ROW_8][COL_F], moveCount: 0},
		"D8": &Queen{value: -9.5, moveCount: 0},
		"E8": &King{value: -99.9, moveCount: 0},
	}
	whiteRows := [2]int{6, 7}
	blackRows := [2]int{0, 1}

	for _, row := range whiteRows {
		for _, sq := range b.Squares[row] {
			piece, ok := whiteStartSquares[sq.Name]
			if !ok {
				continue
			} else {
				piece.SetColor(WHITE)
				piece.SetSquare(sq)
				sq.Piece = piece
				b.WhitePieces[piece] = true
			}
		}
	}

	for _, row := range blackRows {
		for _, sq := range b.Squares[row] {
			piece, ok := blackStartSquares[sq.Name]
			if !ok {
				continue
			} else {
				piece.SetColor(BLACK)
				piece.SetSquare(sq)
				sq.Piece = piece
				b.BlackPieces[piece] = true
			}
		}
	}
	b.Evaluate(WHITE)
}

func (b *Board) RemovePiece(piece Piece, sq *Square) {
	sq.SetPiece(&Null{})
	if piece.Color() == WHITE {
		delete(b.WhitePieces, piece)
	} else {
		delete(b.BlackPieces, piece)
	}
}

func (b *Board) LastMove() *Move {
	if len(b.Moves) > 0 {
		return b.Moves[len(b.Moves)-1]
	}
	return nil
}
