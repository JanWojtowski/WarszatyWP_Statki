package game

import (
	"StatkiBasic/game/utility"
	"context"
	tl "github.com/grupawp/termloop"
	"github.com/nsf/termbox-go"
)

func changeDescription(game ShipsGame) {
	game.Level = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite})
	temp := channel{make(chan string)}

	game.Level.AddEntity(tl.NewText(10, 8, "Enter description: (max: 100 characters long)", tl.Attr(termbox.ColorWhite), tl.Attr(termbox.ColorBlack)))
	inputField := utility.NewInputField(12, 10, temp.ch, 100)
	game.Level.AddEntity(inputField)
	game.Level.AddEntity(tl.NewText(10, 12, "To confirm press \"ENTER\"", tl.Attr(termbox.ColorWhite), tl.Attr(termbox.ColorBlack)))
	game.Game.Screen().SetLevel(game.Level)

	go func() {
		for {
			char := inputField.Listen(context.TODO())
			game.PlayerInfo.desc = char
			BoardSelect(game)
			return
		}
	}()
}
