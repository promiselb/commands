package commands

import (
	"fmt"
	"math/rand"
	"strings"
)

type TicTacToeBoard struct {
	B [3][3]int
}

func NewTicTacToeBoard() *TicTacToeBoard {
	return &TicTacToeBoard{
		B: [3][3]int{},
	}
}

func (board *TicTacToeBoard) Print() string {
	s := ""
	for _, row := range board.B {
		s += "|"
		for _, v := range row {
			if v == 1 {
				s += "X"
			} else if v == 2 {
				s += "O"
			} else {
				s += " "
			}
			s += "|"
		}
		s += "\n"
		s += strings.Repeat("-", 7)
		s += "\n"
	}
	return s
}

func (board *TicTacToeBoard) IsSolved() int {
	var results [8]int

	// -------
	for i, row := range board.B {
		if row[0] == row[1] && row[0] == row[2] {
			results[i] = row[0]
		}
	}

	// |
	for i := range board.B {
		if board.B[0][i] == board.B[1][i] && board.B[0][i] == board.B[2][i] {
			results[i+2] = board.B[0][i]
		}
	}

	// \
	if board.B[0][0] == board.B[1][1] && board.B[0][0] == board.B[2][2] {
		results[6] = board.B[0][0]
	}

	// /
	if board.B[0][2] == board.B[1][1] && board.B[0][2] == board.B[2][0] {
		results[7] = board.B[0][2]
	}

	for _, v := range results {
		if v != 0 {
			return v
		}
	}
	return 0
}

type AIPlayer struct {
	*TicTacToeBoard
	Team int
}

func (bot *AIPlayer) IsAvailable(index int) bool {
	if index < 0 || index > 8 {
		return false
	}
	column, row := getCoordinates(index)
	if bot.TicTacToeBoard.B[column][row] == 1 || bot.TicTacToeBoard.B[column][row] == 2 {
		return false
	}

	return true
}

func (bot *AIPlayer) Capture(index int) {
	if bot.IsAvailable(index) {
		c, r := getCoordinates(index)
		bot.TicTacToeBoard.B[c][r] = bot.Team
	}
}

func getCoordinates(i int) (int, int) {
	return divide(i, 3)
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func CreateRoomAndPlay(team int, commenceFirst bool) (int, error) {
	if team < 1 || team > 2 {
		return 0, errTicTacToeTeam
	}

	oppositeTeam := 0
	if team == 1 {
		oppositeTeam = 2
	} else {
		oppositeTeam = 1
	}
	TheBoard := NewTicTacToeBoard()
	bot := AIPlayer{
		TicTacToeBoard: TheBoard,
		Team:           oppositeTeam,
	}

	var place int
	if commenceFirst {
	thisOne:
		for {
			fmt.Println(TheBoard.Print())
			for {
				fmt.Print("input a place to capture: ")
				fmt.Scanln(&place)
				if IsAvailable(TheBoard, place) {
					c, r := getCoordinates(place)
					TheBoard.B[c][r] = team
					break
				}
			}

			for {
				place = rand.Intn(8)
				if bot.IsAvailable(place) {
					bot.Capture(place)
					break
				}
			}
			x := TheBoard.IsSolved()
			if x != -1 {
				fmt.Printf("\x1bc")
				continue thisOne
			}
			fmt.Println(x)
			fmt.Println(TheBoard.Print())
			return x, nil

			// fmt.Print("\x0c")
		}
	} else {
		for {
			fmt.Println(TheBoard.Print())
			for {
				if bot.IsAvailable(rand.Intn(8)) {
					c, r := getCoordinates(place)
					TheBoard.B[c][r] = bot.Team
					break
				}
			}
			TheBoard.Print()
			fmt.Println()
			for {
				fmt.Print("input a place to capture: ")
				fmt.Scanln(&place)
				if IsAvailable(TheBoard, place) {
					c, r := getCoordinates(place)
					TheBoard.B[c][r] = team
					break
				}
			}

			x := TheBoard.IsSolved()
			if x == 1 || x == 2 {
				return x, nil
			}
			fmt.Print("\x0c")
		}
	}
}

func IsAvailable(board *TicTacToeBoard, index int) bool {
	if index < 0 || index > 8 {
		return false
	}

	column, row := getCoordinates(index)
	if board.B[column][row] == 1 || board.B[column][row] == 2 {
		return false
	}

	return true
}
