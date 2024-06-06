package game

import (
	"StatkiBasic/game/utility"
	"StatkiBasic/httpClient"
	"context"
	tl "github.com/grupawp/termloop"
	"github.com/nsf/termbox-go"
)

func OpponentSelect(game ShipsGame) {
	game.Level = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite})
	gameModeChan := channel{make(chan string)}

	backButton := utility.NewClickableRectangle(
		tl.NewRectangle(1, 1, 15, 3, tl.Attr(termbox.ColorWhite)),
		"back",
		gameModeChan.ch)

	createLobbyButton := utility.NewClickableRectangle(
		tl.NewRectangle(30, 40, 20, 3, tl.Attr(termbox.ColorCyan)),
		"newLobby",
		gameModeChan.ch)

	lobby := httpClient.GetLobby()

	if len(lobby) == 0 {
		game.Level.AddEntity(tl.NewText(30, 7, "No players found", tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorWhite)))
	} else if len(lobby) > 0 && len(lobby) < 5 {
		for i := 1; i <= len(lobby); i++ {
			game.Level.AddEntity(utility.NewClickableRectangle(
				tl.NewRectangle(20, 2+5*i, 30, 3, tl.Attr(termbox.ColorWhite)),
				lobby[i-1].Nick,
				gameModeChan.ch))
			game.Level.AddEntity(tl.NewText(30, 3+5*i, lobby[i-1].Nick,
				tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorWhite)))
		}
	}

	game.Level.AddEntity(createLobbyButton)
	game.Level.AddEntity(tl.NewText(34, 41, "Create Lobby", tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorCyan)))
	game.Level.AddEntity(backButton)
	game.Level.AddEntity(tl.NewText(3, 2, "<--- Back", tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorWhite)))
	game.Game.Screen().SetLevel(game.Level)

	go func() {
		for {
			char := backButton.Listen(context.TODO())

			if char == "back" {
				gameModeSelect(game)
				return
			} else if char == "newLobby" {
				StartGame(game, "multiplayerLobby", "")
			} else {
				StartGame(game, "multiplayerAttack", char)
			}
		}
	}()

}
