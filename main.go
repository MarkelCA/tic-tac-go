package main

import (
    "fmt"
    "strconv"
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

func parseInt(text string) (int, error) {
    return strconv.Atoi(text)
}

func (game *Game) next() {
    fmt.Println("Write a position:")

    var posStr string
    _, scanErr := fmt.Scanln(&posStr)
    if scanErr != nil {
        fmt.Printf("Scan error: %v.\n", scanErr)
        return
    }

    var pos int
    if n,parseErr := parseInt(posStr) ; parseErr != nil {
        fmt.Println("Not valid position, must be a number. Try again")
        return
    } else {
        pos = n
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

    if game.lastMove == nil {
        fmt.Println("Last move not valid, cancelling")
        return false
    }

    pos := game.lastMove.position
    col,row,_ := GetColRow(pos)

    var winConditions []bool

    if pos == 1 || pos == 9 {
        winConditions = append(winConditions, game.board.isWinningDiagonal(1))
    }

    if pos == 3 || pos == 7 {
        winConditions = append(winConditions, game.board.isWinningDiagonal(3))
    }

    winConditions = append(winConditions, game.board.isWinnerRow(row))
    winConditions = append(winConditions, game.board.isWinnerCol(col))

    fmt.Println(winConditions)

    for _,condition := range winConditions {
        if condition == true {
            game.printWinner()
            return true
        }
    }

    return false
}

func (game Game) printWinner() {
    fmt.Printf("Winner: %v!", game.lastMove)
}

func (game Game) PrintTitle() {
    fmt.Println("TicTacToe\n----------")
}
func (game Game) printEnd() {
    fmt.Println("\n--------\nGame finished")
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

    game.next()
    for !game.isFinished() {
        fmt.Println("---------------")
        game.next()
    }
    game.printEnd()
}

func (board *Board) MakeMove(pos int, player Player) error  {
    col, row, err := GetColRow(pos)
    if err != nil {
        return err
    }
    if board[row][col] != '-' {
        return fmt.Errorf("Position already busy. Select another one")
    }
    board[row][col] = player.piece
    board.Print()

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
