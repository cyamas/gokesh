package board

const (
	PAWN   = "PAWN"
	KNIGHT = "KNIGHT"
	BISHOP = "BISHOP"
	ROOK   = "ROOK"
	QUEEN  = "QUEEN"
	KING   = "KING"
	NULL   = "NULL"
)

type Pin struct {
	Piece Piece
	Path  map[*Square]bool
}

type Piece interface {
	Type() string
	Value() float64
	SetColor(color string)
	Color() string
	SetSquare(square *Square)
	Square() *Square
	SetActiveSquares(board *Board)
	ActiveSquares() map[*Square]SqActivity
	SetPin(piece Piece, path map[*Square]bool)
	ResetPin()
	Pin() *Pin
	IsAlly(color string) bool
	IsEnemy(color string) bool
	SetMoved()
	HasMoved() bool
}

type King struct {
	square        *Square
	color         string
	value         float64
	moved         bool
	Checked       bool
	Checkers      []Piece
	Castled       bool
	activeSquares map[*Square]SqActivity
}

var KING_DIRS = map[string][2]int{
	"N":  {-1, 0},
	"E":  {0, 1},
	"S":  {1, 0},
	"W":  {0, -1},
	"NW": {-1, -1},
	"NE": {-1, 1},
	"SE": {1, 1},
	"SW": {1, -1},
}

func (k *King) Type() string                              { return KING }
func (k *King) Value() float64                            { return k.value }
func (k *King) SetColor(color string)                     { k.color = color }
func (k *King) Color() string                             { return k.color }
func (k *King) SetSquare(square *Square)                  { k.square = square }
func (k *King) Square() *Square                           { return k.square }
func (k *King) SetMoved()                                 { k.moved = true }
func (k *King) HasMoved() bool                            { return k.moved }
func (k *King) ActiveSquares() map[*Square]SqActivity     { return k.activeSquares }
func (k *King) SetPin(piece Piece, path map[*Square]bool) {}
func (k *King) ResetPin()                                 {}
func (k *King) Pin() *Pin                                 { return nil }
func (k *King) IsAlly(color string) bool                  { return k.color == color }
func (k *King) IsEnemy(color string) bool                 { return k.color != color }

func (k *King) SetActiveSquares(board *Board) {
	actives := make(map[*Square]SqActivity)
	unsafes := board.GetAttackedSquares(k.color)
	for _, coords := range KING_DIRS {
		candRow := k.square.Row + coords[0]
		candCol := k.square.Column + coords[1]

		if squareExists(candRow, candCol) {
			cand := board.getSquare(candRow, candCol)
			if !cand.IsEmpty() && cand.Piece.IsAlly(k.color) {
				continue
			}
			_, ok := unsafes[cand]
			if !ok {
				if cand.IsEmpty() {
					actives[cand] = FREE
				} else {
					actives[cand] = CAPTURE
				}
			}
		}
	}

	if !k.moved && !k.Checked {
		k.checkForShortCastle(board, unsafes, actives)
		k.checkForLongCastle(board, unsafes, actives)
	}

	k.activeSquares = actives
}

func (k *King) CanEvadeCheck(actives map[*Square]SqActivity, board *Board) bool {
	kingSq := k.Square()
	kingSq.SetPiece(&Null{})
	attackedSqs := board.GetAttackedSquares(k.Color())
	kingSq.SetPiece(k)
	for sq := range actives {
		if _, ok := attackedSqs[sq]; !ok {
			return true
		}
	}
	return false
}

func (k *King) checkForShortCastle(board *Board, unsafes map[*Square][]Piece, actives map[*Square]SqActivity) {
	var hSquare *Square
	var gSquare *Square
	var fSquare *Square
	if k.color == WHITE {
		hSquare = board.Squares[ROW_1][COL_H]
		gSquare = board.Squares[ROW_1][COL_G]
		fSquare = board.Squares[ROW_1][COL_F]
	} else {
		hSquare = board.Squares[ROW_8][COL_H]
		gSquare = board.Squares[ROW_8][COL_G]
		fSquare = board.Squares[ROW_8][COL_F]
	}

	if rookCanCastle(hSquare) && gSquare.IsEmpty() && fSquare.IsEmpty() {
		_, gUnsafe := unsafes[gSquare]
		_, fUnsafe := unsafes[fSquare]
		if !gUnsafe && !fUnsafe {
			actives[gSquare] = CASTLE
		}
	}
}

func (k *King) checkForLongCastle(board *Board, unsafes map[*Square][]Piece, actives map[*Square]SqActivity) {
	var dSquare *Square
	var cSquare *Square
	var aSquare *Square

	if k.color == WHITE {
		dSquare = board.Squares[ROW_1][COL_D]
		cSquare = board.Squares[ROW_1][COL_C]
		aSquare = board.Squares[ROW_1][COL_A]
	} else {
		dSquare = board.Squares[ROW_8][COL_D]
		cSquare = board.Squares[ROW_8][COL_C]
		aSquare = board.Squares[ROW_8][COL_A]
	}

	if rookCanCastle(aSquare) && cSquare.IsEmpty() && dSquare.IsEmpty() {
		_, cUnsafe := unsafes[cSquare]
		_, dUnsafe := unsafes[dSquare]
		if !cUnsafe && !dUnsafe {
			actives[cSquare] = CASTLE
		}
	}
}

func rookCanCastle(square *Square) bool {
	rook, ok := square.Piece.(*Rook)
	if !ok {
		return false
	}
	if rook.moved {
		return false
	}
	return true
}

func (k *King) SetCheck(board *Board) {
	attackedSqs := board.GetAttackedSquares(k.color)
	if checkingPieces, ok := attackedSqs[k.square]; ok {
		k.Checked = true
		k.Checkers = checkingPieces
		return
	}

	k.Checked = false
	k.Checkers = []Piece{}
}

type Queen struct {
	square        *Square
	color         string
	value         float64
	activeSquares map[*Square]SqActivity
	pin           *Pin
	moved         bool
}

func (q *Queen) Type() string   { return QUEEN }
func (q *Queen) Value() float64 { return q.value }
func (q *Queen) SetValue() {
	if q.color == WHITE {
		q.value = 9.5
	} else {
		q.value = -9.5
	}
}
func (q *Queen) SetColor(color string)                 { q.color = color }
func (q *Queen) Color() string                         { return q.color }
func (q *Queen) IsAlly(color string) bool              { return q.color == color }
func (q *Queen) IsEnemy(color string) bool             { return q.color != color }
func (q *Queen) SetSquare(square *Square)              { q.square = square }
func (q *Queen) Square() *Square                       { return q.square }
func (q *Queen) SetMoved()                             { q.moved = true }
func (q *Queen) HasMoved() bool                        { return q.moved }
func (q *Queen) ActiveSquares() map[*Square]SqActivity { return q.activeSquares }
func (q *Queen) SetActiveSquares(board *Board) {
	dirs := map[string][2]int{
		"N":  {-1, 0},
		"E":  {0, 1},
		"S":  {1, 0},
		"W":  {0, -1},
		"NW": {-1, -1},
		"NE": {-1, 1},
		"SE": {1, 1},
		"SW": {1, -1},
	}
	if q.pin != nil {
		q.activeSquares = calcBRQPinnedActives(q, dirs, board)
	} else {
		actives := calcBRQActives(q, dirs, board)
		q.activeSquares = actives
	}
}

func (q *Queen) Pin() *Pin { return q.pin }
func (q *Queen) SetPin(piece Piece, path map[*Square]bool) {
	q.pin = &Pin{piece, path}
}
func (q *Queen) ResetPin() { q.pin = nil }

type Rook struct {
	square        *Square
	color         string
	value         float64
	moved         bool
	CastleSq      *Square
	activeSquares map[*Square]SqActivity
	pin           *Pin
}

func (r *Rook) Type() string                          { return ROOK }
func (r *Rook) Value() float64                        { return r.value }
func (r *Rook) SetColor(color string)                 { r.color = color }
func (r *Rook) Color() string                         { return r.color }
func (r *Rook) IsAlly(color string) bool              { return r.color == color }
func (r *Rook) IsEnemy(color string) bool             { return r.color != color }
func (r *Rook) SetSquare(square *Square)              { r.square = square }
func (r *Rook) Square() *Square                       { return r.square }
func (r *Rook) SetMoved()                             { r.moved = true }
func (r *Rook) HasMoved() bool                        { return r.moved }
func (r *Rook) ActiveSquares() map[*Square]SqActivity { return r.activeSquares }
func (r *Rook) Pin() *Pin                             { return r.pin }

func (r *Rook) SetActiveSquares(board *Board) {
	dirs := map[string][2]int{
		"N": {-1, 0},
		"E": {0, 1},
		"S": {1, 0},
		"W": {0, -1},
	}
	if r.pin != nil {
		r.activeSquares = calcBRQPinnedActives(r, dirs, board)
	} else {
		r.activeSquares = calcBRQActives(r, dirs, board)
	}
}

func (r *Rook) SetPin(piece Piece, path map[*Square]bool) {
	r.pin = &Pin{piece, path}
}
func (r *Rook) ResetPin() { r.pin = nil }

type Bishop struct {
	square        *Square
	color         string
	value         float64
	activeSquares map[*Square]SqActivity
	pin           *Pin
	moved         bool
}

func (b *Bishop) Type() string                          { return BISHOP }
func (b *Bishop) Value() float64                        { return b.value }
func (b *Bishop) SetColor(color string)                 { b.color = color }
func (b *Bishop) Color() string                         { return b.color }
func (b *Bishop) IsAlly(color string) bool              { return b.color == color }
func (b *Bishop) IsEnemy(color string) bool             { return b.color != color }
func (b *Bishop) SetSquare(square *Square)              { b.square = square }
func (b *Bishop) Square() *Square                       { return b.square }
func (b *Bishop) SetMoved()                             { b.moved = true }
func (b *Bishop) HasMoved() bool                        { return b.moved }
func (b *Bishop) Pin() *Pin                             { return b.pin }
func (b *Bishop) ActiveSquares() map[*Square]SqActivity { return b.activeSquares }
func (b *Bishop) SetActiveSquares(board *Board) {
	dirs := map[string][2]int{
		"NW": {-1, -1},
		"NE": {-1, 1},
		"SE": {1, 1},
		"SW": {1, -1},
	}
	if b.pin != nil {
		b.activeSquares = calcBRQPinnedActives(b, dirs, board)
	} else {
		b.activeSquares = calcBRQActives(b, dirs, board)
	}
}

func (b *Bishop) SetPin(piece Piece, path map[*Square]bool) {
	b.pin = &Pin{piece, path}
}
func (b *Bishop) ResetPin() { b.pin = nil }

type Knight struct {
	square        *Square
	color         string
	value         float64
	activeSquares map[*Square]SqActivity
	pin           *Pin
	moved         bool
}

var KNIGHT_DIRS = [8][2]int{
	{-2, -1},
	{-2, 1},
	{-1, 2},
	{1, 2},
	{2, 1},
	{2, -1},
	{1, -2},
	{-1, -2},
}

func (kn *Knight) Type() string                          { return KNIGHT }
func (kn *Knight) Value() float64                        { return kn.value }
func (kn *Knight) SetColor(color string)                 { kn.color = color }
func (kn *Knight) Color() string                         { return kn.color }
func (kn *Knight) IsAlly(color string) bool              { return kn.color == color }
func (kn *Knight) IsEnemy(color string) bool             { return kn.color != color }
func (kn *Knight) SetSquare(square *Square)              { kn.square = square }
func (kn *Knight) Square() *Square                       { return kn.square }
func (kn *Knight) SetMoved()                             { kn.moved = true }
func (kn *Knight) HasMoved() bool                        { return kn.moved }
func (kn *Knight) Pin() *Pin                             { return kn.pin }
func (kn *Knight) ActiveSquares() map[*Square]SqActivity { return kn.activeSquares }

func (kn *Knight) SetPin(piece Piece, path map[*Square]bool) {
	kn.pin = &Pin{piece, path}
}
func (kn *Knight) ResetPin() { kn.pin = nil }

func (kn *Knight) SetActiveSquares(board *Board) {
	king := board.GetKing(kn.color)
	if kn.pin != nil || (king.Checked && len(king.Checkers) > 1) {
		kn.activeSquares = map[*Square]SqActivity{}
		return
	}

	actives := make(map[*Square]SqActivity)

	for _, dir := range KNIGHT_DIRS {
		candRow := kn.square.Row + dir[0]
		candCol := kn.square.Column + dir[1]

		if squareExists(candRow, candCol) {
			cand := board.Squares[candRow][candCol]
			switch {
			case cand.IsEmpty():
				actives[cand] = FREE
			case cand.Piece.IsEnemy(kn.color):
				actives[cand] = CAPTURE
			case cand.Piece.IsAlly(kn.color):
				actives[cand] = GUARDED
			}
		}
	}
	if king.Checked {
		kn.filterForCheck(actives, board, king)
	}

	kn.activeSquares = actives
}

func (kn *Knight) filterForCheck(actives map[*Square]SqActivity, board *Board, king *King) {
	checkerSq := king.Checkers[0].Square()
	kingSq := king.Square()
	checkPath := board.GetAttackedPath(checkerSq, kingSq)
	for sq := range actives {
		if _, ok := checkPath[sq]; !ok {
			delete(actives, sq)
		}
	}
}

type Pawn struct {
	square        *Square
	color         string
	value         float64
	moved         bool
	activeSquares map[*Square]SqActivity
	pin           *Pin
}

func (p *Pawn) Type() string                              { return PAWN }
func (p *Pawn) Value() float64                            { return p.value }
func (p *Pawn) SetColor(color string)                     { p.color = color }
func (p *Pawn) Color() string                             { return p.color }
func (p *Pawn) SetSquare(square *Square)                  { p.square = square }
func (p *Pawn) IsAlly(color string) bool                  { return p.color == color }
func (p *Pawn) IsEnemy(color string) bool                 { return p.color != color }
func (p *Pawn) Square() *Square                           { return p.square }
func (p *Pawn) SetMoved()                                 { p.moved = true }
func (p *Pawn) HasMoved() bool                            { return p.moved }
func (p *Pawn) Pin() *Pin                                 { return p.pin }
func (p *Pawn) SetPin(piece Piece, path map[*Square]bool) { p.pin = &Pin{piece, path} }
func (p *Pawn) ResetPin()                                 { p.pin = nil }
func (p *Pawn) ActiveSquares() map[*Square]SqActivity     { return p.activeSquares }

func (p *Pawn) SetActiveSquares(board *Board) {
	king := board.GetKing(p.color)
	if p.pinnedHorizontally() || (king.Checked && len(king.Checkers) > 1) {
		p.activeSquares = map[*Square]SqActivity{}
		return
	}

	actives := map[*Square]SqActivity{}
	p.addForwardSquares(actives, board)
	p.addDiagonalSquares(actives, board)
	if king.Checked {
		p.filterForCheck(actives, board, king)
	}
	p.activeSquares = actives
}

func (p *Pawn) addForwardSquares(actives map[*Square]SqActivity, board *Board) {
	if p.pinnedDiagonally() {
		return
	}

	row := p.candidateRow(1)
	col := p.square.Column

	if !p.moved {
		dblCand := board.getSquare(p.candidateRow(2), col)
		betweenSq := board.getSquare(row, col)
		if dblCand.IsEmpty() && betweenSq.IsEmpty() {
			actives[dblCand] = FREE
		}
	}

	if squareExists(row, col) {
		cand := board.getSquare(row, col)
		if cand.IsEmpty() {
			actives[cand] = FREE
		}
	}
}

func (p *Pawn) addDiagonalSquares(actives map[*Square]SqActivity, board *Board) {
	if p.pinnedVertically() {
		return
	}
	row := p.candidateRow(1)
	cols := p.candidateCols()

loop:
	for _, col := range cols {
		if squareExists(row, col) {
			cand := board.getSquare(row, col)
			switch {
			case p.pinnedDiagonally():
				pinner := p.pin.Piece
				if cand.Piece == pinner {
					actives[cand] = CAPTURE
				}
				continue loop
			case !cand.IsEmpty() && cand.Piece.IsEnemy(p.color):
				actives[cand] = CAPTURE
			case cand.IsEmpty() && p.canEnPassant(cand, board):
				actives[cand] = EN_PASSANT
			default:
				actives[cand] = GUARDED
			}
		}
	}
}

func (p *Pawn) canEnPassant(cand *Square, board *Board) bool {
	if len(board.Moves) == 0 {
		return false
	}
	lastMove := board.LastMove()

	if lastMove.Piece.Type() == PAWN {
		currRow := ROW_5
		fromRow := ROW_7
		if p.color == BLACK {
			currRow = ROW_4
			fromRow = ROW_2
		}

		if p.square.Row == currRow &&
			lastMove.To.Row == currRow &&
			lastMove.From.Row == fromRow &&
			lastMove.From.Column == cand.Column {

			return true
		}
	}

	return false
}

func (p *Pawn) filterForCheck(actives map[*Square]SqActivity, board *Board, king *King) {
	checker := king.Checkers[0]
	checkPath := board.GetAttackedPath(checker.Square(), king.Square())
	for sq := range actives {
		if _, ok := checkPath[sq]; !ok {
			delete(actives, sq)
		}
	}
}
func (p *Pawn) candidateCols() [2]int {
	col := p.square.Column
	return [2]int{col - 1, col + 1}
}

func (p *Pawn) candidateRow(dist int) int {
	row := p.square.Row
	if p.color == WHITE {
		return row - dist
	}
	return row + dist
}

func (p *Pawn) pinnedHorizontally() bool {
	if p.pin == nil {
		return false
	}
	pinnerRow := p.pin.Piece.Square().Row

	return pinnerRow == p.square.Row
}

func (p *Pawn) pinnedVertically() bool {
	if p.pin == nil {
		return false
	}
	pinnerCol := p.pin.Piece.Square().Column

	return pinnerCol == p.square.Column
}

func (p *Pawn) pinnedDiagonally() bool {
	if p.pin == nil {
		return false
	}
	pinnerRow := p.pin.Piece.Square().Row
	pinnerCol := p.pin.Piece.Square().Column

	return pinnerRow != p.square.Row && pinnerCol != p.square.Column
}

type Null struct {
	square *Square
	color  string
	value  float64
}

func (n *Null) Type() string                              { return NULL }
func (n *Null) Value() float64                            { return n.value }
func (n *Null) SetColor(color string)                     { n.color = color }
func (n *Null) Color() string                             { return "NULL" }
func (n *Null) IsAlly(color string) bool                  { return false }
func (n *Null) IsEnemy(color string) bool                 { return false }
func (n *Null) SetSquare(square *Square)                  { n.square = square }
func (n *Null) Square() *Square                           { return n.square }
func (n *Null) SetMoved()                                 {}
func (n *Null) HasMoved() bool                            { return false }
func (n *Null) ActiveSquares() map[*Square]SqActivity     { return nil }
func (n *Null) SetActiveSquares(board *Board)             {}
func (n *Null) SetPin(piece Piece, path map[*Square]bool) {}
func (n *Null) ResetPin()                                 {}
func (n *Null) Pin() *Pin                                 { return nil }

func squareExists(row int, col int) bool {
	return 0 <= row && row <= 7 && 0 <= col && col <= 7
}

func calcOffset(x int, y int) int {
	diff := x - y
	if diff < 0 {
		return -diff
	} else {
		return diff
	}
}

func calcBRQActives(piece Piece, dirs map[string][2]int, board *Board) map[*Square]SqActivity {
	actives := make(map[*Square]SqActivity)
	king := board.GetKing(piece.Color())
	if king.Checked && len(king.Checkers) > 1 {
		return actives
	}

	for _, coords := range dirs {

		var pinnedCand Piece = &Null{}
		xRayed := false
		path := map[*Square]bool{piece.Square(): true}

	distLoop:
		for dist := 1; dist < 8; dist++ {
			row := piece.Square().Row + (coords[0] * dist)
			col := piece.Square().Column + (coords[1] * dist)

			if squareExists(row, col) {
				cand := board.getSquare(row, col)
				path[cand] = true

				switch {

				case xRayed && !cand.IsEmpty():
					if cand.Piece.IsEnemy(piece.Color()) {
						if cand.Piece.Type() == KING {
							pinnedCand.SetPin(piece, path)
							pinnedCand.SetActiveSquares(board)
						}
					}
					break distLoop

				case !xRayed && cand.IsEmpty():
					actives[cand] = FREE

				case !xRayed && !cand.IsEmpty():
					if cand.Piece.IsEnemy(piece.Color()) {
						if cand.Piece.Type() == KING {
							actives[cand] = CAPTURE
							if sq, ok := board.checkSquarePastKing(row, col, coords); ok {
								actives[sq] = GUARDED
							}
							break distLoop
						}

						pinnedCand = cand.Piece
						actives[cand] = CAPTURE
					} else {
						actives[cand] = GUARDED
					}
					xRayed = true
					continue distLoop

				case xRayed && cand.IsEmpty():
					continue distLoop
				}
			}
		}
	}

	if king.Checked {
		return filterBRQActivesForCheck(actives, board, king)
	}
	return actives
}

func filterBRQActivesForCheck(actives map[*Square]SqActivity, board *Board, king *King) map[*Square]SqActivity {
	filtered := map[*Square]SqActivity{}
	checkerSq := king.Checkers[0].Square()
	checkPath := board.GetAttackedPath(checkerSq, king.Square())
	for sq, activity := range actives {
		if _, ok := checkPath[sq]; ok {
			filtered[sq] = activity
		}
	}
	return filtered
}

func calcBRQPinnedActives(piece Piece, dirs map[string][2]int, board *Board) map[*Square]SqActivity {
	actives := map[*Square]SqActivity{}
	king := board.GetKing(piece.Color())
	for _, coords := range dirs {
		for dist := 1; dist < 8; dist++ {
			candRow := piece.Square().Row + (coords[0] * dist)
			candCol := piece.Square().Column + (coords[1] * dist)
			if squareExists(candRow, candCol) {
				cand := board.getSquare(candRow, candCol)
				_, ok := piece.Pin().Path[cand]
				if !ok {
					break
				}

				switch {
				case cand.IsEmpty():
					actives[cand] = FREE
				case cand == piece.Pin().Piece.Square():
					actives[cand] = CAPTURE
					break
				case cand.Piece.Type() == KING && cand.Piece.IsAlly(piece.Color()):
					break
				}
			}
		}
	}
	if king.Checked {
		return filterBRQActivesForCheck(actives, board, king)
	}
	return actives
}
