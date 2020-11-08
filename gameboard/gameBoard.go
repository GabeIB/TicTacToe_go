package gameboard

import(
	"fmt"
	"errors"
)

type GameBoard struct{
	board [][]rune
	turn int
	lastMove [2]int
}

func (g *GameBoard) GetTurn() int{
	return g.turn
}

func (g *GameBoard) GetLastMove() [2]int{
	return g.lastMove
}

func NewGameBoard(boardSize int) GameBoard{
	var g GameBoard
	g.board = make([][]rune, boardSize)
	for i := range g.board{
		g.board[i] = make([]rune, boardSize)
		for j := range g.board[i]{
			g.board[i][j] = '.'
		}
	}
	g.turn = 1
	return g
}

func (g *GameBoard) Print() {
	for _, line := range g.board{
		for _, ch := range line{
			fmt.Printf("%c", ch)
		}
		fmt.Printf("\n");
	}
}

func (g *GameBoard) Move(x, y int) error {
	if x<0 || x >= len(g.board) || y<0 || y>= len(g.board) {
		return errors.New("Invalid Move")
	}else if g.board[x][y] != '.'{
		return errors.New("Invalid Move")
	}else{
		var userChar rune
		if g.turn == 1 {
			userChar = 'x'
		} else {
			userChar = 'o'
		}
		g.board[x][y] = userChar
		g.lastMove = [2]int{x,y}
		return nil
	}
}

func (g *GameBoard) NextTurn() {
	if g.turn == 1 {
		g.turn = 2
	} else {
		g.turn = 1
	}
}

func _countInARow(playerToken rune, s string) int{
	max := 0
	localCount := 0
	for _, c := range s{
		if c == playerToken {
			localCount++
		}else{
			if localCount > max {
				max = localCount
			}
			localCount = 0
		}
	}
	if localCount > max {
		max = localCount
	}
	return max
}

func (g *GameBoard) countInARow(playerToken rune, startX, startY, itrX, itrY int) int {
	x := startX
	y := startY
	for x < len(g.board) && y < len(g.board) && x >= 0 && y >= 0 {
		x -= itrX
		y -= itrY
	}
	x += itrX
	y += itrY
	s := ""
	for x < len(g.board) && y < len(g.board) && x >= 0 && y >= 0 {
		c := g.board[x][y]
		s += string(c)
		x += itrX
		y += itrY
	}
	return _countInARow(playerToken, s)
}

func (g *GameBoard) GameOver() bool {
	var userChar rune
	if g.turn == 1{
		userChar = 'x'
	}else{
		userChar = 'o'
	}
	if g.countInARow(userChar, g.lastMove[0], g.lastMove[1], 1, 0) == len(g.board) ||
	g.countInARow(userChar, g.lastMove[0], g.lastMove[1], 0, 1) == len(g.board) ||
	g.countInARow(userChar, g.lastMove[0], g.lastMove[1], 1, 1) == len(g.board) ||
	g.countInARow(userChar, g.lastMove[0], g.lastMove[1], -1, 1) == len(g.board) {
		return true
	} else{
		return false
	}
}
