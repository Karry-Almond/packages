package favoriteDB

import "testing"


func TestInit(t *testing.T) {
    Init()
}

func TestNewFavorite(t *testing.T) {
    Init()
    NewFavorite(1,1)
}

func TestCancelFavorite(t *testing.T) {
    Init()
    CancelFavorite(1,1)
}