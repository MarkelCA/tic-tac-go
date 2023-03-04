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
    _, scanErr := fmt.Scanln(&pos)
    if scanErr != nil {
        log.Fatal(scanErr)
    }

    var playerMoving Player

    if game.lastMoved == nil {
        playerMoving = game.firstMove
    } else if game.lastMoved.piece == game.player1.piece {
        playerMoving = game.player2
    } else {
        playerMoving = game.player1
    }

    moveErr := game.board.MakeMove(pos, playerMoving)
    if moveErr != nil {
        fmt.Println(moveErr)
    } else {
        game.lastMoved = &playerMoving
    }
}

func (board Board) isFinished() bool {
    previous := board[0][0]
    row := board[0]

    fmt.Println(board)

    for _,piece := range row {

        if piece == '-' {
            return false
        }
        fmt.Printf("%c %c\n", piece, previous)

        if piece != previous {
            return false
        }
        previous = piece
    }

    return true
}

func (game Game) PrintTitle() {
    fmt.Println("TicTacToe\n----------")
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

    game.PrintTitle()

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

func (board *Board) MakeMove(pos int, player Player) error  {
    col, row, err := GetColRow(pos)
    if err != nil {
        return err
    }
    board[row][col] = player.piece
    board.Print()
    fmt.Println("---------------")

    return nil
}

func GetColRow(pos int) (int, int, error) {
    if pos < 1 || pos > 9 {
        return 0, 0, fmt.Errorf("Argument \"pos\" is not valid (%v). It must be between 1 and 9", pos)
    }

    col  := 3
    row  := pos / 3
    rest := pos % 3

    if rest > 0 {
        row += 1
        col = rest
    }

    return col-1, row-1, nil
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
