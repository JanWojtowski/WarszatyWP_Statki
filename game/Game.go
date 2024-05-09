package game

import (
	"StatkiBasic/httpClient"
	"context"
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"time"
)

type PlayerInfo struct {
	nickname   string
	desc       string
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

func PlayGame() {
	playerInfo := PlayerInfo{
		nickname: "Yuki",
		desc:     "Snow",
		coords: []string{"A1", "A3", "B9", "C7", "D1",
			"D2", "D3", "D4", "D7", "E7",
			"F1", "F2", "F3", "F5", "G5",
			"G8", "G9", "I4", "J4", "J8"},
	}

	ownBoard := false
	gameMode := "single"

	gameinfo := GameInfo{gameStarter(gameMode, ownBoard, playerInfo)}

	if !ownBoard {
		playerInfo.coords = httpClient.GetBoard(gameinfo.AuthToken)
	}

	fmt.Println(gameinfo)
	fmt.Println("--------------------------------------------")

	tempInfo := httpClient.GetOpponentInfo(gameinfo.AuthToken)

	opponentInfo := OpponentInfo{
		nickname: tempInfo[0],
		desc:     tempInfo[1],
	}
	fmt.Println(opponentInfo)
	fmt.Println("--------------------------------------------")

	ui := gui.NewGUI(true)

	txt := gui.NewText(1, 1, "Press on any coordinate to shot it.", nil)
	ui.Draw(txt)
	ui.Draw(gui.NewText(1, 2, "Press Ctrl+C to exit", nil))

	opponentBoard := gui.NewBoard(1, 4, nil)
	ui.Draw(opponentBoard)
	playerBoard := gui.NewBoard(50, 4, nil)
	ui.Draw(playerBoard)

	ui.Draw(gui.NewText(1, 26, opponentInfo.desc, nil))
	ui.Draw(gui.NewText(50, 26, playerInfo.desc, nil))

	go func() {
		for {
			boardsUpdater(gameinfo.AuthToken, playerInfo, playerBoard, opponentBoard)
			char := opponentBoard.Listen(context.TODO())
			if canFire(gameinfo.AuthToken) {
				fireResult := fire(char, gameinfo.AuthToken, &playerInfo)
				txt.SetText(fmt.Sprintf("Fired at Coordinates: %s. %s!!!", char, fireResult))
			} else {
				txt.SetText("Wait for your turn!")
			}
		}
	}()

	ui.Start(context.TODO(), nil)
}

func gameStarter(gameMode string, ownBoard bool, playerInfo PlayerInfo) string {
	var token string
	if !ownBoard {
		playerInfo.coords = []string{}
	}

	switch gameMode {
	case "single":
		token = httpClient.StartGameWithBot(playerInfo.nickname, playerInfo.desc, playerInfo.coords)
		time.Sleep(3 * time.Second)
	case "multiplayer":
		token = httpClient.StartGameMulti(playerInfo.nickname, playerInfo.desc, playerInfo.coords, "Ktos")
		time.Sleep(3 * time.Second)
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
	numbers[0] = int(cord[0]) - 65
	numbers[1] = int(cord[1]) - 49

	return numbers
}

func playerBoardStatusUpdate(board *gui.Board, cords []string, oppCords []string) {
	states := [10][10]gui.State{}
	for _, ply := range cords {
		numbers := cordToNumbers(ply)
		states[numbers[0]][numbers[1]] = gui.Ship
	}
	for _, opp := range oppCords {
		if stringInSlice(opp, cords) {
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

func boardsUpdater(authKey string, playerInfo PlayerInfo, plyBoard *gui.Board, oppBoard *gui.Board) {
	status := httpClient.GetStatus(authKey)
	playerBoardStatusUpdate(plyBoard, playerInfo.coords, status.OppShots)
	opponentBoardStatusUpdate(oppBoard, playerInfo.hitCoords, playerInfo.missCoords)
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
