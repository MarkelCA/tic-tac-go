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
    emptyBoard  := Board{ 
        {'-','-','-'},
        {'-','-','-'},
        {'-','-','-'},
    }

    got := emptyBoard.isFinished()
    want := false

    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }

    finishedBoard  := Board{ 
        {'x','-','-'},
        {'x','-','-'},
        {'x','-','-'},
    }

    got2 := finishedBoard.isFinished()
    want2 := false

    if got2 != want2 {
        t.Errorf("got2 %v, wanted %v", got2, want2)
    }
}
