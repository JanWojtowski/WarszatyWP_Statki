package game

import (
	"StatkiBasic/game/utility"
	"context"
	tl "github.com/grupawp/termloop"
	"github.com/nsf/termbox-go"
	"io/ioutil"
)

func gameModeSelect(game ShipsGame) {
	game.Level = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite})
	gameModeChan := channel{make(chan string)}

	singleButton := utility.NewClickableRectangle(
		tl.NewRectangle(10, 15, 30, 3, tl.Attr(termbox.ColorRed)),
		"single",
		gameModeChan.ch)
	multiButton := utility.NewClickableRectangle(
		tl.NewRectangle(45, 15, 30, 3, tl.Attr(termbox.ColorGreen)),
		"multi",
		gameModeChan.ch)
	backButton := utility.NewClickableRectangle(
		tl.NewRectangle(27, 25, 30, 3, tl.Attr(termbox.ColorWhite)),
		"back",
		gameModeChan.ch)

	dat, err := ioutil.ReadFile("game/files/gamemode.txt")
	if err != nil {
		panic(err)
	}

	title := tl.NewEntityFromCanvas(5, 5, tl.CanvasFromString(string(dat)))

	game.Level.AddEntity(title)
	game.Level.AddEntity(singleButton)
	game.Level.AddEntity(tl.NewText(22, 16, "Single", tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorRed)))
	game.Level.AddEntity(multiButton)
	game.Level.AddEntity(tl.NewText(57, 16, "Multi", tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorGreen)))
	game.Level.AddEntity(backButton)
	game.Level.AddEntity(tl.NewText(36, 26, "Back to menu", tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorWhite)))

	game.Game.Screen().SetLevel(game.Level)

	go func() {
		for {
			char := singleButton.Listen(context.TODO())

			if char == "back" {
				MainMenu(game)
				return
			} else if char == "single" {
				StartGame(game, char)
				return
			} else if char == "multi" {
				OpponentSelect(game)
				return
			}
		}
	}()

}
