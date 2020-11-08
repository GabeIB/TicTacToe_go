package gameboard

import(
	"testing"
)

func TestCountInARow(t *testing.T){
	g := NewGameBoard(3)
	g.Move(1,1)
	ans := g.countInARow('x', 1, 1, 1, 0)
	if ans != 1 {
		t.Errorf("Incorrect count")
	}
	g.Move(2,2)
	ans = g.countInARow('x', 2, 2, 1, 1)
	if ans != 2 {
		t.Errorf("Incorrect count")
	}
	ans = g.countInARow('x', 2, 2, 0, 1)
	if ans != 1 {
		t.Errorf("Incorrect count")
	}
}

//table based unit test
func TestGameOver(t *testing.T){
	g = NewGameBoard(3)
	var tests = []struct{
		p1_moves [][2]int
		p2_moves [][2]int
	}{
		{[][2]int{{0,0}, {1,1}, {2,2}}, [][2]int{{1,0}, {0,1}}},
		{[][2]int{{1,0}, {2,0}, {0,0}}, [][2]int{{2,2}, {1,2}}},
		{[][2]int{{2,2}, {2,1}, {1,1}}, [][2]int{{0,2}, {0,1}, {0,0}}},
	}
	for i, tt := range tests {
		turn := 0
		for{
			g.Move(tt.p1_moves[turn][0], tt.p1_moves[turn][1])
			ans := g.GameOver()
			if turn == len(tt.p1_moves)-1 && turn == len(tt.p2_moves) {
				//last move
				if ans != true {
					t.Errorf("false negative for game over: test %d, turn %d", i, turn)
				}
				break
			} else if ans != false {
				g.Print()
				t.Errorf("false positive for game over: test %d, turn %d", i, turn)
			}
			g.NextTurn()
			g.Move(tt.p2_moves[turn][0], tt.p2_moves[turn][1])
			ans = g.GameOver()
			if turn == len(tt.p2_moves)-1 && turn == len(tt.p1_moves)-1 {
				//last move
				if ans != true {
					t.Errorf("false negative for game over: test %d, turn %d", i, turn)
				}
				break
			} else if ans != false {
				t.Errorf("false positive for game over: test %d, turn %d", i, turn)
			}
			g.NextTurn()
			turn++
		}
		g = NewGameBoard(3)
	}
}


func TestMove(t *testing.T){
	g := NewGameBoard(3)
	ans := g.Move(0,0)
	if ans != nil{
		t.Errorf("Move(0,0) returned error")
	}
	if g.board[0][0] != 'x'{
		t.Errorf("Move(0,0) didn't change board")
	}
	g.Move(2,1)
	if g.board[2][1] != 'x'{
		t.Errorf("Move(2,1) didn't change board appropriately")
	}
	ans = g.Move(-1,3)
	if ans == nil{
		t.Errorf("Move(-1,3) didn't throw error")
	}
	ans = g.Move(2,-3)
	if ans == nil{
		t.Errorf("Move(2,-3) didn't throw error")
	}
	ans = g.Move(4, 1)
	if ans == nil{
		t.Errorf("Move(4,1) didn't throw error")
	}
	ans = g.Move(1,5)
	if ans == nil{
		t.Errorf("Move(1,5) didn't throw error")
	}
	ans = g.Move(0,0) //trying to move same spot twice
	if ans == nil{
		t.Errorf("duplicate move to (0,0) didn't throw error")
	}
}
