package main

import (
    "testing"
)

func TestGetTrue(t *testing.T) {
    got := true
    want := true

    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }

}

func TestIsFinished(t *testing.T) {
    //emptyBoard  := Board{ 
        //{'-','-','-'},
        //{'-','-','-'},
        //{'-','-','-'},
    //}

    //got := emptyBoard.isFinished()
    //want := false

    //if result := emptyBoard.isFinished() ; result == true {
        //t.Errorf("got %v, wanted %v", got, want)
    //}

    finishedBoard  := Board{ 
        {'x','x','x'},
        {'x','-','-'},
        {'x','-','-'},
    }

    if result := finishedBoard.isWinnerRow(0) ; result == false {
        t.Errorf("result %v, wanted %v", result, false)
    }
    if result := finishedBoard.isWinnerRow(1) ; result != false {
        t.Errorf("result %v, wanted %v", result, false)
    }
    if result := finishedBoard.isWinnerCol(0) ; result == false {
        t.Errorf("result %v, wanted %v", result, false)
    }
    if result := finishedBoard.isWinnerCol(1) ; result != false {
        t.Errorf("result %v, wanted %v", result, false)
    }
}

func TestIsWinnerDiagonal(t *testing.T) {
    finishedBoard  := Board{ 
        {'x','-','-'},
        {'-','x','-'},
        {'-','-','x'},
    }

    if result := finishedBoard.isWinningDiagonal(1) ; result == false {
        t.Errorf("result %v, wanted %v", result, false)
    }
    if result := finishedBoard.isWinningDiagonal(9) ; result == false {
        t.Errorf("result %v, wanted %v", result, false)
    }

    finishedBoard2  := Board{ 
        {'-','-','o'},
        {'-','o','-'},
        {'o','-','-'},
    }

    if result := finishedBoard2.isWinningDiagonal(3) ; result == false {
        t.Errorf("result %v, wanted %v", result, false)
    }
    if result := finishedBoard2.isWinningDiagonal(7) ; result == false {
        t.Errorf("result %v, wanted %v", result, false)
    }

}

func TestGetColRow(t *testing.T) {
    if col,row,err := GetColRow(1) ; col != 0 || row != 0 && err != nil {
        t.Errorf("Col(expected=%v, received=%v), Row(expected=%v, received=%v)", 0, col, 0, row)
    }

    if col,row,err := GetColRow(2) ; col != 1 || row != 0 && err != nil {
        t.Errorf("Col(expected=%v, received=%v), Row(expected=%v, received=%v)", 0, col, 0, row)
    }
     
    if col,row,err := GetColRow(3) ; col != 2 || row != 0 && err != nil {
        t.Errorf("Col(expected=%v, received=%v), Row(expected=%v, received=%v)", 0, col, 0, row)
    }

    if col,row,err := GetColRow(4) ; col != 0 || row != 1 && err != nil {
        t.Errorf("Col(expected=%v, received=%v), Row(expected=%v, received=%v)", 0, col, 0, row)
    }

    if col,row,err := GetColRow(5) ; col != 1 || row != 1 && err != nil {
        t.Errorf("Col(expected=%v, received=%v), Row(expected=%v, received=%v)", 0, col, 0, row)
    }

    if col,row,err := GetColRow(6) ; col != 2 || row != 1 && err != nil {
        t.Errorf("Col(expected=%v, received=%v), Row(expected=%v, received=%v)", 0, col, 0, row)
    }

    if col,row,err := GetColRow(7) ; col != 0 || row != 2 && err != nil {
        t.Errorf("Col(expected=%v, received=%v), Row(expected=%v, received=%v)", 0, col, 0, row)
    }

    if col,row,err := GetColRow(8) ; col != 1 || row != 2 && err != nil {
        t.Errorf("Col(expected=%v, received=%v), Row(expected=%v, received=%v)", 0, col, 0, row)
    }

    if col,row,err := GetColRow(9) ; col != 2 || row != 2 && err != nil {
        t.Errorf("Col(expected=%v, received=%v), Row(expected=%v, received=%v)", 0, col, 0, row)
    }

    if _,_,err := GetColRow(0) ; err == nil {
        t.Error("Must have thrown an exception")
    }

    if _,_,err := GetColRow(10) ; err == nil {
        t.Error("Must have thrown an exception")
    }
}
