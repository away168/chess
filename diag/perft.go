// Package diag provides diagnostic tools for chess engines.
package diag

import (
	"github.com/andrewbackes/chess"
	"github.com/andrewbackes/chess/board"
)

/*******************************************************************************

	Divide:

*******************************************************************************/

// Divide is a diagnostic tool used for figuring out which moves are not
// being generated by an engine. It returns a list of moves and a count
// of how many moves are in that tree of moves with the given depth.
func Divide(G *chess.Game, depth int) map[board.Move]uint64 {
	div := make(map[board.Move]uint64)
	//fmt.Println("Depth", depth)
	var nodes, moveCount uint64
	ml := G.LegalMoves()
	toMove := G.ActiveColor()
	for mv := range ml {
		temp := *G
		temp.MakeMove(mv)

		if temp.Check(toMove) == false {
			//Count it for mate:
			moveCount++
			n := Perft(&temp, depth-1)
			div[mv] = n
			nodes += n
		}
	}
	return div
}

/*******************************************************************************

	Perft:

*******************************************************************************/

// Perft retuns the number of possible moves from the given board position and chess.Game
// state at the given depth.
func Perft(g *chess.Game, depth int) uint64 {
	if depth == 0 {
		return 1
	}
	toMove := g.ActiveColor()
	var nodes uint64
	ml := g.LegalMoves()
	for mv := range ml {
		temp := *g
		temp.QuickMove(mv)
		if temp.Check(toMove) == false {
			nodes += Perft(&temp, depth-1)
		}
	}
	return nodes
}

/*
func perftBreakdown(G *chess.Game, depth int) (nodes, checks, castles, mates, captures, promotions, enpassant uint64) {
	var moveCount uint64

	if depth == 0 {
		return 1, 0, 0, 0, 0, 0, 0
	}

	toMove := G.ActiveColor()
	notToMove := []piece.Color{game.Black, game.White}[toMove]

	isChecked := G.Check(toMove)
	ml := G.LegalMoves()

	for mv := range ml {
		temp := *G
		temp.QuickMove(mv)
		if temp.Check(toMove) == false {
			//Count it for mate:
			moveCount++
			n, c, cstl, m, cap, p, enp := perftBreakdown(&temp, depth-1)
			nodes += n
			checks += c + toInt(temp.Check(notToMove))
			castles += cstl + toInt(isCastle(G, mv))
			mates += m
			captures += cap + toInt(isCapture(G, mv))
			promotions += p + toInt(isPromotion(G, mv))
			enpassant += enp + toInt(isEnPassant(G, mv))
		}
	}
	if moveCount == 0 && isChecked {
		mates++
	}
	return nodes, checks, castles, mates, captures, promotions, enpassant
}
*/

/*******************************************************************************

	Helpers:

*******************************************************************************/

/*
func isCastle(G *chess.Game, m board.Move) bool {
	from, _ := game.SquaresOf(m)
	p := G.Board().OnSquare(from)
	if p.Type == game.King {
		if (m == "e1g1") || (m == "e1c1") || (m == "e8g8") || (m == "e8c8") {
			return true
		}
	}
	return false
}

func isCapture(G *chess.Game, m board.Move) bool {
	_, to := game.SquaresOf(m)
	capPiece := G.Board().OnSquare(to)
	return (capPiece.Type != game.None)
}

func isPromotion(G *chess.Game, m board.Move) bool {
	// TODO: will not work when more notation is added
	return (len(m) > 4)
}

func isEnPassant(G *chess.Game, m board.Move) bool {
	if G.EnPassant() == nil {
		return false
	}
	from, to := game.SquaresOf(m)
	p := G.Board().OnSquare(from)
	return (p.Type == game.Pawn) && (to == *G.EnPassant()) && ((from-to)%8 != 0)
}

func toInt(b bool) uint64 {
	if b == true {
		return 1
	}
	return 0
}
*/