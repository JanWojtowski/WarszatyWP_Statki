package game

import (
	"StatkiBasic/game/utility"
	"context"
	tl "github.com/grupawp/termloop"
	"github.com/nsf/termbox-go"
	"io/ioutil"
)

type channel struct {
	ch chan string
}

func MainMenu(game ShipsGame) {
	game.Level = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite})
	temp := channel{make(chan string)}

	button1 := utility.NewClickableRectangle(
		tl.NewRectangle(25, 20, 30, 3, tl.Attr(termbox.ColorRed)),
		"play",
		temp.ch)
	button2 := utility.NewClickableRectangle(
		tl.NewRectangle(25, 25, 30, 3, tl.Attr(termbox.ColorGreen)),
		"stats",
		temp.ch)

	dat, err := ioutil.ReadFile("game/files/title.txt")
	if err != nil {
		panic(err)
	}

	title := tl.NewEntityFromCanvas(5, 5, tl.CanvasFromString(string(dat)))

	game.Level.AddEntity(title)
	game.Level.AddEntity(button1)
	game.Level.AddEntity(tl.NewText(37, 21, "Start", tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorRed)))
	game.Level.AddEntity(button2)
	game.Level.AddEntity(tl.NewText(35, 26, "Leaderboard", tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorGreen)))

	game.Game.Screen().SetLevel(game.Level)

	go func() {
		for {
			char := button1.Listen(context.TODO())

			if char == "play" {
				if game.PlayerInfo.nickname != "" && game.PlayerInfo.desc != "" {
					BoardSelect(game)
				} else {
					if game.PlayerInfo.nickname == "" {
						changeNickname(game)
					} else if game.PlayerInfo.desc == "" {
						changeDescription(game)
					}
				}
				return
			} else if char == "stats" {
				StatsBoard(game)
			}

		}
	}()

}
