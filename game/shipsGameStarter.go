package game

import (
	"context"
	tl "github.com/grupawp/termloop"
)

type ShipsGame struct {
	Game       *tl.Game
	Level      *tl.BaseLevel
	PlayerInfo PlayerInfo
}

func NewGame() {
	shipsGame := ShipsGame{
		Game:  tl.NewGame(),
		Level: tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite}),
	}

	go MainMenu(shipsGame)
	shipsGame.Game.Start(context.TODO())
}
