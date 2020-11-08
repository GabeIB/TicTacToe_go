package main

import(
	"fmt"
	"github.com/GabeIB/TicTacToe_go/gameboard"
)

func getTurn(g *gameboard.GameBoard){
	var x, y int
	fmt.Printf("Player %d's Turn:\n", g.GetTurn())
	validTurn := false
	for !validTurn {
		fmt.Printf("Pick a row: ")
		fmt.Scanln(&x)
		fmt.Printf("Pick a column: ")
		fmt.Scanln(&y)
		err := g.Move(x,y)
		if err == nil {
			validTurn = true
		} else {
			fmt.Println("Invalid Turn! try again.")
		}
	}
}

func printWinner(g *gameboard.GameBoard){
	fmt.Printf("Player %d is the Winner!!!!\n", g.GetTurn())
}

func main(){
	g := gameboard.NewGameBoard(3)
	gameOver := false
	for !gameOver{
		g.Print()
		getTurn(&g)
		if g.GameOver() {
			printWinner(&g)
			gameOver = true
		}
		g.NextTurn()
	}
	fmt.Println("See you later!")
}
