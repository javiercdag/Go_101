package multidimchess

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

type Position struct {
	Coordinates []int
}

type Piece struct {
	Type     string // e.g. Queen, King, Bishop, etc.
	Color    string // White or Black
	Position Position
}

type Board struct {
	Pieces []Piece
}

func copyPosition(p Position) Position {
	return Position{Coordinates: []int{p.Coordinates[0], p.Coordinates[1]}}
}

func copyPiece(p Piece) Piece {
	return Piece{Type: p.Type, Color: p.Color, Position: copyPosition(p.Position)}
}

func copyBoard(b Board) Board {
	newPieces := make([]Piece, len(b.Pieces))

	for i, p := range b.Pieces {
		newPieces[i] = copyPiece(p)
	}

	return Board{Pieces: newPieces}
}
