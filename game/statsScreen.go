package game

import (
	"StatkiBasic/game/utility"
	"StatkiBasic/httpClient"
	"context"
	tl "github.com/grupawp/termloop"
	"github.com/nsf/termbox-go"
)

func StatsBoard(game ShipsGame) {
	game.Level = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite})
	gameModeChan := channel{make(chan string)}

	backButton := utility.NewClickableRectangle(
		tl.NewRectangle(1, 1, 15, 3, tl.Attr(termbox.ColorWhite)),
		"back",
		gameModeChan.ch)

	stats := httpClient.GetStats()

	if len(stats) == 0 {
		game.Level.AddEntity(tl.NewText(30, 7, "No players found", tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorWhite)))
	} else if len(stats) > 0 {
		for i := 1; i <= len(stats); i++ {
			if i == 1 {
				game.Level.AddEntity(utility.NewRectangle(
					tl.NewRectangle(20, 2+4*i, 30, 3, tl.Attr(termbox.ColorYellow))))
				game.Level.AddEntity(tl.NewText(30, 3+4*i, stats[i-1].Nick,
					tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorYellow)))
			} else {
				game.Level.AddEntity(utility.NewRectangle(
					tl.NewRectangle(20, 2+4*i, 30, 3, tl.Attr(termbox.ColorWhite))))
				game.Level.AddEntity(tl.NewText(30, 3+4*i, stats[i-1].Nick,
					tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorWhite)))
			}
		}
	}
	game.Level.AddEntity(backButton)
	game.Level.AddEntity(tl.NewText(3, 2, "<--- Back", tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorWhite)))
	game.Game.Screen().SetLevel(game.Level)

	go func() {
		for {
			char := backButton.Listen(context.TODO())

			if char == "back" {
				MainMenu(game)
				return
			}
		}
	}()

}
