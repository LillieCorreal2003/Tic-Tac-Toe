package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Set up board with 9 tiles
//Player 1 choose one tile and mark "X"
//Player 2 then chooses a  different tile and marks "O"
//Each player then alternates turns until someone wins
//Check if player wins or both players tie
//Print out congrats

func drawBoard(b []string) {
	fmt.Println(b[0], "|", b[1], "|", b[2])
	fmt.Println("---------")
	fmt.Println(b[3], "|", b[4], "|", b[5])
	fmt.Println("---------")
	fmt.Println(b[6], "|", b[7], "|", b[8])
}

type player interface{
  getPosition() int
	String() string
}

type human string

func (h human) getPosition() int {
	var position int
	fmt.Print("Enter position number: ")
	fmt.Scanln(&position)
	return position
}

func (h human) String() string {
	return string(h)
}

type AI string

func (a AI) getPosition() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(9)
}

func (a AI) String() string {
	return string(a)
}

func hasWinner(board []string) string {
	if board[0] == board[1] && board[1] == board[2]{
		return board[0]
	}
	if board[3] == board[4] && board[4] == board[5]{
		return board[3]
	}
	if board[6] == board[7] && board[7] == board[8]{
		return board[6]
	}
	if board[0] == board[3] && board[3] == board[6]{
		return board[0]
	}
	if board[1] == board[4] && board[4] == board[7]{
		return board[1]
	}
	if board[2] == board[5] && board[5] == board[8]{
		return board[2]
	}
	if board[0] == board[4] && board[4] == board[8]{
		return board[0]
	}
	if board[2] == board[4] && board[4] == board[6]{
		return board[2]
	}
	return " "

}

func isEmpty(board []string) bool {
	for _,pos := range board {
		if pos == " "{
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Welcome to Tic Tac Toe")

	board := []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}

	fmt.Println("1. Human vs Human")
	fmt.Println("2. Human vs Computer")
	fmt.Println("3. Computer vs Computer")
	var choice int
	fmt.Println("Choose which gameplay you want")
	fmt.Scanln(&choice)

	var players []player
	if choice == 1 {
		players = append(players, human("X"), human("O"))
	}
	if choice == 2 {
		players = append(players, human("X"), AI("O"))
	}
	if choice == 3 {
		players = append(players, AI("X"), AI("O"))
	}


  turn := 0

	for hasWinner(board) == " " && isEmpty(board) {
		drawBoard(board)

    player := players[turn%2]
		pos1 := player.getPosition()
		for board[pos1] != " " {
			pos1 = player.getPosition()
		}
		board[pos1] = player.String()
    turn+= 1
	}

	drawBoard(board)
	winner := hasWinner(board)

	if winner != " " {
		fmt.Println("Player", winner, "won")
		return 
	}

	if !isEmpty(board) {
		fmt.Println("Game is tied")
	}
}
