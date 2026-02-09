# Multi-dimensional Chess

In Zedonia the people usually play N-dimensional chess, however in Hungary the people are only capable of playing 2-dimensional chess. Implement the neccessary `copyPosition`, `copyPiece` and `copyBoard` functions that convert a Zedonian chess `Board` to a Hungarian chess `Board` by copying the first `2` coordinates and omitting the remaining ones. Make sure that if the Hungarians develop a method to play in higher dimensions, they are able to do that.

You are given the following structs:

```go
type Position struct {
    Coordinates []int
}

type Piece struct {
    Type string // e.g. Queen, King, Bishop, etc.
    Color string // White or Black
    Position Position
}

type Board struct {
    Pieces []Piece
}
```

Your task is to implement the following functions that perform the expected behavior:

``` go
func copyPosition(p Position) Position
```
``` go
func copyPiece(p Piece) Piece
```
``` go
func copyBoard(b Board) Board
```

Insert the code into the file `exercise.go` at the placeholder `// INSERT YOUR CODE HERE`.
