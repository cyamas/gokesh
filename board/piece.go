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

type Piece interface {
	Type() string
	Value() int
	SetColor(color string)
	Color() string
	SetSquare(square *Square)
	Square() *Square
	ActiveSquares(board *Board) map[*Square]SqActivity
}

type King struct {
	square *Square
	color  string
	value  int
	Moved  bool
}

func (k *King) Type() string             { return KING }
func (k *King) Value() int               { return k.value }
func (k *King) SetColor(color string)    { k.color = color }
func (k *King) Color() string            { return k.color }
func (k *King) SetSquare(square *Square) { k.square = square }
func (k *King) Square() *Square          { return k.square }

func (k *King) ActiveSquares(board *Board) map[*Square]SqActivity {
	actives := make(map[*Square]SqActivity)
	unsafes := board.GetAttackedSquares(k.color)

	if !k.Moved && !k.IsInCheck(unsafes) {
		k.checkForShortCastle(board, unsafes, actives)
		k.checkForLongCastle(board, unsafes, actives)
	}

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

	for _, coords := range dirs {
		candRow := k.square.Row + coords[0]
		candCol := k.square.Column + coords[1]

		if squareExists(candRow, candCol) {
			cand := getCandidateSquare(board, candRow, candCol)
			if cand.Piece.Type() != NULL && cand.Piece.Color() == k.color {
				continue
			}
			if board.SquareIsSafe(k.color, cand) {
				if cand.Piece.Type() == NULL {
					actives[cand] = FREE
				} else {
					actives[cand] = CAPTURE
				}
			}
		}
	}

	return actives
}

func (k *King) checkForShortCastle(board *Board, unsafes map[*Square]bool, actives map[*Square]SqActivity) {
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

func (k *King) checkForLongCastle(board *Board, unsafes map[*Square]bool, actives map[*Square]SqActivity) {
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
	if rook.Moved {
		return false
	}
	return true
}

func (k *King) IsInCheck(attackedSquares map[*Square]bool) bool {
	_, ok := attackedSquares[k.square]
	if ok {
		return true
	}
	return false
}

type Queen struct {
	square *Square
	color  string
	value  int
}

func (q *Queen) Type() string             { return QUEEN }
func (q *Queen) Value() int               { return q.value }
func (q *Queen) SetColor(color string)    { q.color = color }
func (q *Queen) Color() string            { return q.color }
func (q *Queen) SetSquare(square *Square) { q.square = square }
func (q *Queen) Square() *Square          { return q.square }

func (q *Queen) ActiveSquares(board *Board) map[*Square]SqActivity {
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

	return calcBRQActives(q, dirs, board)
}

type Rook struct {
	square   *Square
	color    string
	value    int
	Moved    bool
	CastleSq *Square
}

func (r *Rook) Type() string             { return ROOK }
func (r *Rook) Value() int               { return r.value }
func (r *Rook) SetColor(color string)    { r.color = color }
func (r *Rook) Color() string            { return r.color }
func (r *Rook) SetSquare(square *Square) { r.square = square }
func (r *Rook) Square() *Square          { return r.square }

func (r *Rook) ActiveSquares(board *Board) map[*Square]SqActivity {
	dirs := map[string][2]int{
		"N": {-1, 0},
		"E": {0, 1},
		"S": {1, 0},
		"W": {0, -1},
	}

	return calcBRQActives(r, dirs, board)
}

type Bishop struct {
	square *Square
	color  string
	value  int
}

func (b *Bishop) Type() string             { return BISHOP }
func (b *Bishop) Value() int               { return b.value }
func (b *Bishop) SetColor(color string)    { b.color = color }
func (b *Bishop) Color() string            { return b.color }
func (b *Bishop) SetSquare(square *Square) { b.square = square }
func (b *Bishop) Square() *Square          { return b.square }

func (b *Bishop) ActiveSquares(board *Board) map[*Square]SqActivity {
	dirs := map[string][2]int{
		"NW": {-1, -1},
		"NE": {-1, 1},
		"SE": {1, 1},
		"SW": {1, -1},
	}

	return calcBRQActives(b, dirs, board)
}

type Knight struct {
	square *Square
	color  string
	value  int
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

func (kn *Knight) Type() string             { return KNIGHT }
func (kn *Knight) Value() int               { return kn.value }
func (kn *Knight) SetColor(color string)    { kn.color = color }
func (kn *Knight) Color() string            { return kn.color }
func (kn *Knight) SetSquare(square *Square) { kn.square = square }
func (kn *Knight) Square() *Square          { return kn.square }

func (kn *Knight) ActiveSquares(board *Board) map[*Square]SqActivity {
	actives := make(map[*Square]SqActivity)

	for _, dir := range KNIGHT_DIRS {
		candRow := kn.square.Row + dir[0]
		candCol := kn.square.Column + dir[1]

		if squareExists(candRow, candCol) {
			cand := board.Squares[candRow][candCol]
			switch {
			case cand.Piece.Type() == NULL:
				actives[cand] = FREE
			case cand.Piece.Type() != NULL && cand.Piece.Color() != kn.color:
				actives[cand] = CAPTURE
			case cand.Piece.Type() != NULL && cand.Piece.Color() != kn.color:
				actives[cand] = GUARDED
			}
		}
	}

	return actives
}

type Pawn struct {
	square     *Square
	color      string
	value      int
	PrevSquare *Square
	Moved      bool
}

func (p *Pawn) Type() string             { return PAWN }
func (p *Pawn) Value() int               { return p.value }
func (p *Pawn) SetColor(color string)    { p.color = color }
func (p *Pawn) Color() string            { return p.color }
func (p *Pawn) SetSquare(square *Square) { p.square = square }
func (p *Pawn) Square() *Square          { return p.square }

func (p *Pawn) ActiveSquares(board *Board) map[*Square]SqActivity {
	actives := make(map[*Square]SqActivity)

	p.addFreeMoves(board, actives)
	p.addCapturesAndGuardeds(board, actives)

	enPassantSq, ok := p.checkEnPassant(board)
	if ok {
		actives[enPassantSq] = EN_PASSANT
	}
	return actives
}

func (p *Pawn) checkEnPassant(board *Board) (*Square, bool) {
	lastMove := board.LastMove()
	if lastMove == nil {
		return nil, false
	}
	if lastMove.Piece.Type() == PAWN {
		switch p.color {
		case WHITE:
			if p.square.Row != ROW_5 {
				return nil, false
			}
			if lastMove.From.Row == ROW_7 && lastMove.To.Row == ROW_5 {
				if calcColumnDiff(lastMove.To.Column, p.square.Column) == 1 {
					epRow := p.square.Row - 1
					epCol := lastMove.To.Column
					epSquare := board.Squares[epRow][epCol]
					return epSquare, true
				}
			}
		case BLACK:
			if p.square.Row != 4 {
				return nil, false
			}
			if lastMove.From.Row == ROW_2 && lastMove.To.Row == ROW_4 {
				if calcColumnDiff(lastMove.To.Column, p.square.Column) == 1 {
					epRow := p.square.Row + 1
					epCol := lastMove.To.Column
					epSquare := board.Squares[epRow][epCol]
					return epSquare, true
				}
			}
		}
	}
	return nil, false
}

func (p *Pawn) addFreeMoves(board *Board, actives map[*Square]SqActivity) {
	freeCol := p.square.Column
	freeRow := p.freeMoveRow(1)

	cand := getCandidateSquare(board, freeRow, freeCol)
	if cand.Piece.Type() == NULL {
		actives[cand] = FREE

		if !p.Moved {
			dblRow := p.freeMoveRow(2)
			dblCand := getCandidateSquare(board, dblRow, freeCol)
			if dblCand.Piece.Type() == NULL {
				actives[dblCand] = FREE
			}
		}
	}
}

func (p *Pawn) freeMoveRow(dist int) int {
	currRow := p.square.Row
	if p.color == WHITE {
		return currRow - dist
	}
	return currRow + dist
}

func (p *Pawn) addCapturesAndGuardeds(board *Board, actives map[*Square]SqActivity) {
	row := p.captureRow()
	leftCol, rightCol := p.captureColumns()

	if squareExists(row, leftCol) {
		cand := board.Squares[row][leftCol]
		if cand.Piece.Type() != NULL {
			if cand.Piece.Color() != p.color {
				actives[cand] = CAPTURE
			} else {
				actives[cand] = GUARDED
			}
		}
	}

	if squareExists(row, rightCol) {
		cand := board.Squares[row][rightCol]
		if cand.Piece.Type() != NULL {
			if cand.Piece.Color() != p.color {
				actives[cand] = CAPTURE
			} else {
				actives[cand] = GUARDED
			}
		}
	}
}

func (p *Pawn) captureRow() int {
	currRow := p.square.Row
	if p.color == WHITE {
		return currRow - 1
	}
	return currRow + 1
}

func (p *Pawn) captureColumns() (int, int) {
	currCol := p.square.Column

	if p.color == WHITE {
		return currCol - 1, currCol + 1
	}

	return currCol + 1, currCol - 1

}

func (p *Pawn) checkForValidCapture(board *Board, row, col int) (*Square, bool) {
	if 0 <= col && col <= 7 {
		cand := getCandidateSquare(board, row, col)
		if !cand.IsEmpty() && cand.Piece.Color() != p.color {
			return cand, true
		}
	}
	return nil, false
}

type Null struct {
	square *Square
	color  string
	value  int
}

func (n *Null) Type() string                                      { return NULL }
func (n *Null) Value() int                                        { return n.value }
func (n *Null) SetColor(color string)                             { n.color = color }
func (n *Null) Color() string                                     { return "NULL" }
func (n *Null) SetSquare(square *Square)                          { n.square = square }
func (n *Null) Square() *Square                                   { return n.square }
func (n *Null) ActiveSquares(board *Board) map[*Square]SqActivity { return map[*Square]SqActivity{} }

func getCandidateSquare(board *Board, row, col int) *Square {
	return board.Squares[row][col]
}

func squareExists(row int, col int) bool {
	return 0 <= row && row <= 7 && 0 <= col && col <= 7
}

func calcColumnDiff(col1 int, col2 int) int {
	diff := col1 - col2
	if diff < 0 {
		return -diff
	} else {
		return diff
	}
}

func calcBRQActives(piece Piece, dirs map[string][2]int, board *Board) map[*Square]SqActivity {
	actives := make(map[*Square]SqActivity)

	for dist := 1; dist < 8; dist++ {
		for dir, coords := range dirs {
			candRow := piece.Square().Row + (coords[0] * dist)
			candCol := piece.Square().Column + (coords[1] * dist)
			if squareExists(candRow, candCol) {
				cand := getCandidateSquare(board, candRow, candCol)
				switch {
				case cand.Piece.Type() != NULL && cand.Piece.Color() == piece.Color():
					actives[cand] = GUARDED
					delete(dirs, dir)
				case cand.Piece.Type() == NULL:
					actives[cand] = FREE
				case cand.Piece.Type() != NULL && cand.Piece.Color() != piece.Color():
					actives[cand] = CAPTURE
					delete(dirs, dir)
				}
			} else {
				delete(dirs, dir)
			}
		}
	}
	return actives
}
