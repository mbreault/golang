package main

import (
	"fmt"
	"math/rand"
	"time"
)

const boardSize int = 3

var gameStates = map[string]string{
	"IN_PROGESS": "IN_PROGESS",
	"DRAW":       "DRAW",
	"X_WINS":     "X_WINS",
	"O_WINS":     "O_WINS",
}

// Move defintion
type Move struct {
	x int
	y int
}

// LegalMoves defintion
type LegalMoves struct {
	Moves []Move
}

func transform(row []int) []string {
	result := make([]string, 0)
	cell := " "
	for i := range row {
		switch row[i] {
		case -1:
			cell = "O"
		case 1:
			cell = "X"
		case 0:
			cell = " "
		}
		result = append(result, cell)
	}
	return result
}

func printBoard(board [][]int) {
	fmt.Println("")
	for _, row := range board {
		fmt.Println(transform(row))
	}
	fmt.Println("")
}

func makeMove(board [][]int, move Move, player int) [][]int {
	board[move.x][move.y] = player
	return board
}

func getMove(i int, j int) Move {
	return Move{
		x: i,
		y: j,
	}
}
func legalMoves(board [][]int) LegalMoves {
	result := LegalMoves{}
	moves := make([]Move, 0)
	for i, row := range board {
		for j := range row {
			if board[i][j] == 0 {
				move := Move{
					x: i,
					y: j,
				}
				moves = append(moves, move)
			}
		}
	}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(moves), func(i, j int) { moves[i], moves[j] = moves[j], moves[i] })
	result.Moves = moves
	return result
}

func sum(row []int) int {
	result := 0
	for i := range row {
		result += row[i]
	}
	return result
}

func playerWins(player int) string {
	result := ""
	if player == 1 {
		result = "X_WINS"
	} else {
		result = "O_WINS"
	}
	return result
}
func getGameState(board [][]int, player int) string {
	result := gameStates["IN_PROGRESS"]
	moves := legalMoves(board)

	// horizontal win
	for _, row := range board {
		if sum(row) == player*boardSize {
			return gameStates[playerWins(player)]
		}
	}

	// vertical win
	for j := range board {
		col := make([]int, 0)
		for i := range board {
			col = append(col, board[i][j])
		}
		if sum(col) == player*boardSize {
			return gameStates[playerWins(player)]
		}
	}

	// forward diagnoal win
	forwardDiag := make([]int, 0)
	for i := range board {
		forwardDiag = append(forwardDiag, board[i][i])
	}
	if sum(forwardDiag) == player*boardSize {
		return gameStates[playerWins(player)]
	}

	// backword diagnoal win
	backwardDiag := make([]int, 0)
	for i := range board {
		backwardDiag = append(backwardDiag, board[i][(boardSize-1)-i])
	}
	if sum(backwardDiag) == player*boardSize {
		return gameStates[playerWins(player)]
	}

	if len(moves.Moves) == 0 {
		return gameStates["DRAW"]
	}

	return result
}

func initBoard() [][]int {
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}
	return board
}

func play() {
	gameState := gameStates["IN_PROGRESS"]

	board := initBoard()
	player := 1

	for gameState == gameStates["IN_PROGRESS"] {
		moves := legalMoves(board)
		nextMove := moves.Moves[0]
		fmt.Println(nextMove)
		board = makeMove(board, nextMove, player)
		printBoard(board)
		gameState = getGameState(board, player)
		player = -player
	}
	print(gameState)
	printBoard(board)
}

func main() {
	play()
}
