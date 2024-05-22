package game

import (
	"StatkiBasic/game/utility"
	"context"
	tl "github.com/grupawp/termloop"
)

func changeNickname(game ShipsGame) {
	game.Level = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite})
	temp := channel{make(chan string)}

	inputField := utility.NewInputField(10, 10, temp.ch)
	game.Level.AddEntity(inputField)

	game.Game.Screen().SetLevel(game.Level)

	go func() {
		for {
			char := inputField.Listen(context.TODO())
			game.PlayerInfo.nickname = char
			changeDescription(game)

		}
	}()
}
