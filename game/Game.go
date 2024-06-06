package game

import (
	"StatkiBasic/game/utility"
	"StatkiBasic/httpClient"
	"context"
	"fmt"
	gui "github.com/JanWojtowski/warships-gui"
	tl "github.com/grupawp/termloop"
	"github.com/nsf/termbox-go"
	"time"
)

type PlayerInfo struct {
	nickname   string
	desc       string
	ownBoard   bool
	coords     []string
	hitCoords  []string
	missCoords []string
}

type OpponentInfo struct {
	nickname string
	desc     string
}

type GameInfo struct {
	AuthToken string
}

func StartGame(game ShipsGame, gameMode string, opponent string) {
	playerInfo := game.PlayerInfo

	gameInfo := GameInfo{gameStarter(gameMode, playerInfo, opponent)}

	if !playerInfo.ownBoard {
		playerInfo.coords = httpClient.GetBoard(gameInfo.AuthToken)
	}

	tempInfo := httpClient.GetOpponentInfo(gameInfo.AuthToken)
	opponentInfo := OpponentInfo{tempInfo[0], tempInfo[1]}

	game.Level = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite})
	game.Game.Screen().SetLevel(game.Level)

	surrenderChan := channel{make(chan string)}

	surrenderButton := utility.NewClickableRectangle(
		tl.NewRectangle(38, 40, 20, 3, tl.Attr(termbox.ColorRed)),
		"surrender",
		surrenderChan.ch)

	game.Level.AddEntity(surrenderButton)
	game.Level.AddEntity(tl.NewText(43, 41, "Surrender", tl.Attr(termbox.ColorBlack), tl.Attr(termbox.ColorRed)))

	ui := gui.NewGUI(true, *game.Game)

	timer := gui.NewText(42, 1, "Game will start soon.", nil)
	txt := gui.NewText(30, 3, "Press on any coordinate to shot it.", nil)
	turn := gui.NewText(40, 5, "Waiting for game to start", nil)
	playerNick := gui.NewText(50, 8, playerInfo.nickname, nil)
	opponentNick := gui.NewText(1, 8, opponentInfo.nickname, nil)
	opponentBoard := gui.NewBoard(1, 11, nil)
	playerBoard := gui.NewBoard(50, 11, nil)

	playerDesc := descriptionFormater(52, 33, 40, playerInfo.desc)
	opponentDesc := descriptionFormater(3, 33, 40, opponentInfo.desc)

	ui.Draw(txt)
	ui.Draw(timer)
	ui.Draw(turn)
	ui.Draw(opponentBoard)
	ui.Draw(playerBoard)
	ui.Draw(opponentNick)
	ui.Draw(playerNick)
	for _, element := range playerDesc {
		ui.Draw(&element)
	}
	for _, element := range opponentDesc {
		ui.Draw(&element)
	}

	go func() {
		for {
			char := opponentBoard.Listen(context.TODO())
			if canFire(gameInfo.AuthToken) {
				fireResult := fire(char, gameInfo.AuthToken, &playerInfo)
				txt.SetText(fmt.Sprintf("Fired at Coordinates: %s. %s!!!", char, fireResult))
				boardsUpdater(gameInfo.AuthToken, playerInfo, playerBoard, opponentBoard)
			} else {
				txt.SetText("Wait for your turn!")
			}
		}
	}()

	go func() {
		for {
			char := surrenderButton.Listen(context.TODO())
			if char == "surrender" {
				httpClient.DeleteSurrender(gameInfo.AuthToken)
				ui.Remove(txt)
				ui.Remove(timer)
				ui.Remove(turn)
				ui.Remove(opponentBoard)
				ui.Remove(playerBoard)
				ui.Remove(opponentNick)
				ui.Remove(playerNick)
				for _, element := range opponentDesc {
					ui.Remove(&element)
				}
				for _, element := range playerDesc {
					ui.Remove(&element)
				}
				MainMenu(game)
				return
			}
		}
	}()

	go func() {
		for {
			time.Sleep(300 * time.Millisecond)
			status := boardsUpdater(gameInfo.AuthToken, playerInfo, playerBoard, opponentBoard)
			timer.SetText(fmt.Sprintf("Timer: %d", status.Timer))
			if status.ShouldFire {
				turn.SetText(fmt.Sprintf("Turn of: %s", status.Nick))
			} else {
				turn.SetText(fmt.Sprintf("Turn of: %s", status.Opponent))
			}
			if status.GameStatus == "ended" {
				return
			}
		}
	}()
}

func gameStarter(gameMode string, playerInfo PlayerInfo, opponent string) string {
	var token string

	if !playerInfo.ownBoard {
		playerInfo.coords = []string{}
	}

	switch gameMode {
	case "single":
		token = httpClient.StartGameWithBot(playerInfo.nickname, playerInfo.desc, playerInfo.coords)
		time.Sleep(1 * time.Second)
	case "multiplayerAttack":
		token = httpClient.StartGameMultiAttack(playerInfo.nickname, playerInfo.desc, playerInfo.coords, opponent)
		time.Sleep(1 * time.Second)
	case "multiplayerLobby":
		token = httpClient.StartGameMultiLobby(playerInfo.nickname, playerInfo.desc, playerInfo.coords)
		time.Sleep(1 * time.Second)
		for {
			if httpClient.GetStatus(token).GameStatus != "waiting" {
				break
			} else {
				httpClient.GetRefresh(token)
				time.Sleep(1 * time.Second)
			}
		}
	default:
		token = ""
	}
	if token != "" {
		return token
	} else {
		panic("brak tokena gry")
	}
}

func cordToNumbers(cord string) []int {
	numbers := []int{0, 0}

	if len(cord) == 3 {
		numbers[0] = int(cord[0]) - 65
		numbers[1] = 9
	} else {
		numbers[0] = int(cord[0]) - 65
		numbers[1] = int(cord[1]) - 49
	}

	return numbers
}

func playerBoardStatusUpdate(board *gui.Board, plyCords []string, oppCords []string) {
	states := [10][10]gui.State{}
	for _, ply := range plyCords {
		numbers := cordToNumbers(ply)
		states[numbers[0]][numbers[1]] = gui.Ship
	}
	for _, opp := range oppCords {
		if stringInSlice(opp, plyCords) {
			numbers := cordToNumbers(opp)
			states[numbers[0]][numbers[1]] = gui.Hit
		} else {
			numbers := cordToNumbers(opp)
			states[numbers[0]][numbers[1]] = gui.Miss
		}
	}
	board.SetStates(states)
}

func opponentBoardStatusUpdate(board *gui.Board, hitCords []string, missCords []string) {
	states := [10][10]gui.State{}
	for _, hit := range hitCords {
		numbers := cordToNumbers(hit)
		states[numbers[0]][numbers[1]] = gui.Hit
	}
	for _, miss := range missCords {
		numbers := cordToNumbers(miss)
		states[numbers[0]][numbers[1]] = gui.Miss
	}
	board.SetStates(states)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func boardsUpdater(authKey string, playerInfo PlayerInfo, plyBoard *gui.Board, oppBoard *gui.Board) httpClient.Status {
	status := httpClient.GetStatus(authKey)
	playerBoardStatusUpdate(plyBoard, playerInfo.coords, status.OppShots)
	opponentBoardStatusUpdate(oppBoard, playerInfo.hitCoords, playerInfo.missCoords)
	return status
}

func fire(cord string, authKey string, playerInfo *PlayerInfo) string {
	fireResult := httpClient.PostFire(cord, authKey)
	switch fireResult {
	case "miss":
		if stringInSlice(cord, playerInfo.hitCoords) {
			return "You have already shoted at that space"
		} else {
			playerInfo.missCoords = append(playerInfo.missCoords, cord)
			return "miss"
		}
	case "hit":
		playerInfo.hitCoords = append(playerInfo.hitCoords, cord)
		return "hit"
	case "sunk":
		playerInfo.hitCoords = append(playerInfo.hitCoords, cord)
		return "SUNKEN"
	default:
		fmt.Println("fire fail")
		panic("fire fail")
	}
}

func canFire(authKey string) bool {
	return httpClient.GetStatus(authKey).ShouldFire
}

func descriptionFormater(x, y, maxlinelength int, desc string) []gui.Text {
	var texts []gui.Text
	line := ""
	var lines []string
	i := 0
	for _, ch := range desc {
		line = line + fmt.Sprintf("%c", ch)
		i++
		if i%maxlinelength == 0 {
			lines = append(lines, line)
			line = ""
		}
	}
	lines = append(lines, line)
	for _, line := range lines {
		texts = append(texts, *gui.NewText(x, y, line, nil))
		y++
	}

	return texts
}
