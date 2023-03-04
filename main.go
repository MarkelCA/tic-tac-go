package main

import (
    "fmt"
    "log"
)

type Board [3][3]rune
type player rune

type Player struct {
    piece rune
    name string
}

type Game struct {
    board      Board
    player1    Player
    player2    Player
    firstMove  Player
    lastMoved  *Player
    winner     *Player
}

func getTrue() bool {
    return true
}

func (game *Game) next() {
    fmt.Println("Write a position:")
    fmt.Println()

    var pos int
    _, err := fmt.Scanln(&pos)
    if err != nil {
        log.Fatal(err)
    }

    var playerMoving Player

    if game.lastMoved == nil {
        playerMoving = game.firstMove
    } else if game.lastMoved.piece == game.player1.piece {
        playerMoving = game.player2
    } else {
        playerMoving = game.player1
    }

    game.board.MakeMove(pos, playerMoving)
    game.lastMoved = &playerMoving
}

func (board Board) isFinished() bool {
    return false
}

func main() {
    board  := NewBoard()
    markel := Player{'x', "Markel"}
    bot    := Player{'o', "Player 2"}

    game := Game {
        board:     board,
        firstMove: markel,
        player1:   markel,
        player2:   bot,
    }

    fmt.Println("TicTacToe\n----------")
    for !game.board.isFinished() {
        game.next()
    }

    return


    board.MakeMove(1, markel)
    board.MakeMove(2, bot)
    board.MakeMove(3, markel)
    board.MakeMove(4, bot)
    board.MakeMove(5, markel)
    board.MakeMove(6, bot)
    board.MakeMove(7, markel)
    board.MakeMove(8, bot)
    board.MakeMove(9, markel)
}

func (board *Board) MakeMove(pos int, player Player) {
    col  := 3
    row  := pos / 3
    rest := pos % 3

    if rest > 0 {
        row += 1
        col = rest
    }

    board[row-1][col-1] = player.piece
    board.Print()
    fmt.Println("---------------")
}


func NewBoard() Board {
    board := Board{}
    for i, row := range board {
        for j, _ := range row {
            board[i][j] = '-'
        }
    }

    return board
}

func (board Board) Print() {
    for _, row := range board {
        for _, cell := range row {
            fmt.Printf("%c ", cell)
        }
        fmt.Println()
    }
}
