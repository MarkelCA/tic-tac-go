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

type Move struct {
    player   Player
    position int
}

type Game struct {
    board      Board
    player1    Player
    player2    Player
    firstMove  Player
    lastMove  *Move
    winner     *Player
    turnCount  int
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

    if game.lastMove == nil {
        playerMoving = game.firstMove
    } else if game.lastMove.player.piece == game.player1.piece {
        playerMoving = game.player2
    } else {
        playerMoving = game.player1
    }

    moveErr := game.board.MakeMove(pos, playerMoving)
    if moveErr != nil {
        fmt.Println(moveErr)
    } else {
        move := Move{playerMoving,pos}
        game.lastMove = &move
        game.turnCount += 1
    }
}

func contains(slice []rune, element rune) bool {
    for _,x := range slice {
        if x == element {
            return true
        }
    }
    return false
}

func (board Board) isWinnerRow(row int) bool {
    fullRow := board[row]
    return !contains(board[row][:], '-') && fullRow[0] == fullRow[1] && fullRow[1] == fullRow[2]

}

func (board Board) isWinnerCol(col int) bool {
    fullCol := [3]rune{board[0][col], board[1][col], board[2][col]}
    return !contains(fullCol[:], '-') && fullCol[0] == fullCol[1] && fullCol[1] == fullCol[2]

}

func (board Board) isWinningDiagonal(pos int) bool {
    var fullDiagonal [3]rune
    if pos == 1 || pos == 9 {
        fullDiagonal = [3]rune{board[0][0], board[1][1], board[2][2]}
    } else if pos == 3 || pos == 7 {
        fullDiagonal = [3]rune{board[0][2], board[1][1], board[2][0]}
    }

    return !contains(fullDiagonal[:], '-') && fullDiagonal[0] == fullDiagonal[1] && fullDiagonal[1] == fullDiagonal[2]
}

func (game Game) isFinished() bool {
    if game.turnCount == 9 {
        return true
    }

    fmt.Println(game.lastMove)


    row := 0
    return game.board.isWinnerRow(row)


    //return board[0][0] == board[0][1] && board[0][1h == board[0][2]


    //fmt.Println(board)

    //for _,piece := range row {

        //if piece == '-' {
            //return false
        //}
        //fmt.Printf("%c %c\n", piece, previous)

        //if piece != previous {
            //return false
        //}
        //previous = piece
    //}

    //return true
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

    for !game.isFinished() {
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
