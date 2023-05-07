package state

type WhitePawnMoveFinder struct{}

func (m *WhitePawnMoveFinder) Find(coor Coordinate, board [8][8]*Piece) (moves map[Coordinate]bool) {
	if coor.Y+1 < 8 {
		if board[coor.Y+1][coor.X] == nil {
			moves[Coordinate{X: coor.X, Y: coor.Y + 1}] = true
		}
		if coor.X+1 < 8 && board[coor.Y+1][coor.X+1].Colour == Black {
			moves[Coordinate{X: coor.X + 1, Y: coor.Y + 1}] = true
		}
		if coor.X-1 > -1 && board[coor.Y+1][coor.X-1].Colour == Black {
			moves[Coordinate{X: coor.X - 1, Y: coor.Y + 1}] = true
		}
	}
	return moves
}
