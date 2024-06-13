package game

import (
	"context"
	gui "github.com/JanWojtowski/warships-gui"
	tl "github.com/grupawp/termloop"
)

func BoardBuilder(game ShipsGame) {

	game.Level = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite})
	game.Game.Screen().SetLevel(game.Level)

	ui := gui.NewGUI(true, *game.Game)

	board := gui.NewBoard(20, 11, nil)

	line1 := gui.NewText(15, 3, "Place your ships by clicking on the board", nil)
	ui.Draw(line1)
	line2 := gui.NewText(12, 4, "You will be placing them in order from largest to smallest", nil)
	ui.Draw(line2)
	line3 := gui.NewText(15, 6, "Now placing: ", nil)
	ui.Draw(line3)
	txt := gui.NewText(30, 6, "4pcs ship", nil)
	ui.Draw(txt)

	mis := gui.NewText(15, 8, "Here will be shown any tips if you try place you ship wrongly", nil)
	ui.Draw(mis)

	ui.Draw(board)

	go func() {
		var currentShip []string
		var shipsTab []string
		for {
			char := board.Listen(context.TODO())
			shipPlacer(char, &currentShip, &shipsTab, mis)

			switch len(shipsTab) {
			case 4:
				currentShip = []string{}
				txt.SetText("1st 3pcs ship")
			case 7:
				currentShip = []string{}
				txt.SetText("2nd 3pcs ship")
			case 10:
				currentShip = []string{}
				txt.SetText("1st 2pcs ship")
			case 12:
				currentShip = []string{}
				txt.SetText("2nd 2pcs ship")
			case 14:
				currentShip = []string{}
				txt.SetText("3rd 2pcs ship")
			case 16:
				currentShip = []string{}
				txt.SetText("1st 1pcs ship")
			case 17:
				currentShip = []string{}
				txt.SetText("2nd 1pcs ship")
			case 18:
				currentShip = []string{}
				txt.SetText("3rd 1pcs ship")
			case 19:
				currentShip = []string{}
				txt.SetText("4st 1pcs ship")
			case 20:
				ui.Remove(line1)
				ui.Remove(line2)
				ui.Remove(line3)
				ui.Remove(txt)
				ui.Remove(mis)
				ui.Remove(board)

				game.PlayerInfo.ownBoard = true
				game.PlayerInfo.coords = shipsTab

				gameModeSelect(game)
			}

			states := [10][10]gui.State{}
			for _, ply := range shipsTab {
				numbers := cordToNumbers(ply)
				states[numbers[0]][numbers[1]] = gui.Ship
			}
			paintSides(currentShip, &states, gui.Sunk)
			paintSides(shipsWithoutCurrent(shipsTab, currentShip), &states, gui.Hit)
			paintSlants(shipsWithoutCurrent(shipsTab, currentShip), &states, gui.Hit)
			board.SetStates(states)
		}
	}()
}

func shipPlacer(char string, currentShip *[]string, shipsTab *[]string, mis *gui.Text) {
	if !stringInSlice(char, *shipsTab) {
		if len(*currentShip) == 0 {
			if canPlaceNewShip(char, *shipsTab) {
				*currentShip = append(*currentShip, char)
				*shipsTab = append(*shipsTab, char)
			} else {
				mis.SetText("To close to another ships")
			}
		} else {
			if !checkSides(char, shipsWithoutCurrent(*shipsTab, *currentShip)) {
				if checkSides(char, *currentShip) {
					*currentShip = append(*currentShip, char)
					*shipsTab = append(*shipsTab, char)
				} else {
					mis.SetText("Ship parts edges must be connected!")
				}
			} else {
				mis.SetText("To close to another ship!")
			}
		}
	} else {
		mis.SetText("You already placed ship here!")
	}
}

func canPlaceNewShip(char string, ships []string) bool {
	if !checkSides(char, ships) && !checkSlants(char, ships) {
		return true
	}
	return false
}

func checkSides(char string, ships []string) bool {
	coords := cordToNumbers(char)

	if coords[0] > 0 && coords[0] < 9 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[0]+1 == temp[0] && coords[1] == temp[1] {
				return true
			} else if coords[0]-1 == temp[0] && coords[1] == temp[1] {
				return true
			}
		}
	} else if coords[0] == 0 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[0]+1 == temp[0] && coords[1] == temp[1] {
				return true
			}
		}
	} else if coords[0] == 9 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[0]-1 == temp[0] && coords[1] == temp[1] {
				return true
			}
		}
	}

	if coords[1] > 0 && coords[1] < 9 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[1]+1 == temp[1] && coords[0] == temp[0] {
				return true
			} else if coords[1]-1 == temp[1] && coords[0] == temp[0] {
				return true
			}
		}
	} else if coords[1] == 0 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[1]+1 == temp[1] && coords[0] == temp[0] {
				return true
			}
		}
	} else if coords[1] == 9 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[1]-1 == temp[1] && coords[0] == temp[0] {
				return true
			}
		}
	}
	return false
}

func checkSlants(char string, ships []string) bool {
	coords := cordToNumbers(char)

	if coords[0] > 0 && coords[1] > 0 && coords[0] < 9 && coords[1] < 9 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[0]+1 == temp[0] && coords[1]+1 == temp[1] {
				return true
			} else if coords[0]+1 == temp[0] && coords[1]-1 == temp[1] {
				return true
			} else if coords[0]-1 == temp[0] && coords[1]+1 == temp[1] {
				return true
			} else if coords[0]-1 == temp[0] && coords[1]-1 == temp[1] {
				return true
			}
		}
	} else if coords[0] == 0 && coords[1] == 0 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[0]+1 == temp[0] && coords[1]+1 == temp[1] {
				return true
			}
		}
	} else if coords[0] == 9 && coords[1] == 9 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[0]-1 == temp[0] && coords[1]-1 == temp[1] {
				return true
			}
		}
	} else if coords[0] == 0 && coords[1] == 9 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[0]+1 == temp[0] && coords[1]-1 == temp[1] {
				return true
			}
		}
	} else if coords[0] == 9 && coords[1] == 0 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[0]-1 == temp[0] && coords[1]+1 == temp[1] {
				return true
			}
		}
	} else if coords[0] == 0 && coords[1] > 0 && coords[1] < 9 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[0]+1 == temp[0] && coords[1]+1 == temp[1] {
				return true
			} else if coords[0]+1 == temp[0] && coords[1]-1 == temp[1] {
				return true
			}
		}
	} else if coords[0] == 9 && coords[1] > 0 && coords[1] < 9 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[0]-1 == temp[0] && coords[1]+1 == temp[1] {
				return true
			} else if coords[0]-1 == temp[0] && coords[1]-1 == temp[1] {
				return true
			}
		}
	} else if coords[1] == 0 && coords[0] > 0 && coords[0] < 9 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[1]+1 == temp[1] && coords[0]+1 == temp[0] {
				return true
			} else if coords[1]+1 == temp[1] && coords[0]-1 == temp[0] {
				return true
			}
		}
	} else if coords[1] == 9 && coords[0] > 0 && coords[0] < 9 {
		for _, ship := range ships {
			temp := cordToNumbers(ship)
			if coords[1]-1 == temp[1] && coords[0]+1 == temp[0] {
				return true
			} else if coords[1]-1 == temp[1] && coords[0]-1 == temp[0] {
				return true
			}
		}
	}
	return false
}

func shipsWithoutCurrent(ships []string, current []string) []string {
	var temp []string
	for _, ship := range ships {
		if !stringInSlice(ship, current) {
			temp = append(temp, ship)
		}
	}
	return temp
}

func paintSides(ships []string, shipsState *[10][10]gui.State, state gui.State) {
	for _, ship := range ships {
		coords := cordToNumbers(ship)

		if coords[0] > 0 && coords[0] < 9 {
			if shipsState[coords[0]+1][coords[1]] != gui.Ship {
				shipsState[coords[0]+1][coords[1]] = state
			}
			if shipsState[coords[0]-1][coords[1]] != gui.Ship {
				shipsState[coords[0]-1][coords[1]] = state
			}
		} else if coords[0] == 0 {
			if shipsState[coords[0]+1][coords[1]] != gui.Ship {
				shipsState[coords[0]+1][coords[1]] = state
			}
		} else if coords[0] == 9 {
			if shipsState[coords[0]-1][coords[1]] != gui.Ship {
				shipsState[coords[0]-1][coords[1]] = state
			}
		}

		if coords[1] > 0 && coords[1] < 9 {
			if shipsState[coords[0]][coords[1]+1] != gui.Ship {
				shipsState[coords[0]][coords[1]+1] = state
			}
			if shipsState[coords[0]][coords[1]-1] != gui.Ship {
				shipsState[coords[0]][coords[1]-1] = state
			}
		} else if coords[1] == 0 {
			if shipsState[coords[0]][coords[1]+1] != gui.Ship {
				shipsState[coords[0]][coords[1]+1] = state
			}
		} else if coords[1] == 9 {
			if shipsState[coords[0]][coords[1]-1] != gui.Ship {
				shipsState[coords[0]][coords[1]-1] = state
			}
		}
	}
}

func paintSlants(ships []string, shipsState *[10][10]gui.State, state gui.State) {
	for _, ship := range ships {
		coords := cordToNumbers(ship)

		if coords[0] > 0 && coords[1] > 0 && coords[0] < 9 && coords[1] < 9 {
			if shipsState[coords[0]+1][coords[1]+1] != gui.Ship {
				shipsState[coords[0]+1][coords[1]+1] = state
			}
			if shipsState[coords[0]+1][coords[1]-1] != gui.Ship {
				shipsState[coords[0]+1][coords[1]-1] = state
			}
			if shipsState[coords[0]-1][coords[1]+1] != gui.Ship {
				shipsState[coords[0]-1][coords[1]+1] = state
			}
			if shipsState[coords[0]-1][coords[1]-1] != gui.Ship {
				shipsState[coords[0]-1][coords[1]-1] = state
			}

		} else if coords[0] == 0 && coords[1] == 0 {
			if shipsState[coords[0]+1][coords[1]+1] != gui.Ship {
				shipsState[coords[0]+1][coords[1]+1] = state
			}
		} else if coords[0] == 9 && coords[1] == 9 {
			if shipsState[coords[0]-1][coords[1]-1] != gui.Ship {
				shipsState[coords[0]-1][coords[1]-1] = state
			}
		} else if coords[0] == 0 && coords[1] == 9 {
			if shipsState[coords[0]+1][coords[1]-1] != gui.Ship {
				shipsState[coords[0]+1][coords[1]-1] = state
			}
		} else if coords[0] == 9 && coords[1] == 0 {
			if shipsState[coords[0]-1][coords[1]+1] != gui.Ship {
				shipsState[coords[0]-1][coords[1]+1] = state
			}
		} else if coords[0] == 0 && coords[1] > 0 && coords[1] < 9 {
			if shipsState[coords[0]+1][coords[1]+1] != gui.Ship {
				shipsState[coords[0]+1][coords[1]+1] = state
			}
			if shipsState[coords[0]+1][coords[1]-1] != gui.Ship {
				shipsState[coords[0]+1][coords[1]-1] = state
			}
		} else if coords[0] == 9 && coords[1] > 0 && coords[1] < 9 {
			if shipsState[coords[0]-1][coords[1]+1] != gui.Ship {
				shipsState[coords[0]-1][coords[1]+1] = state
			}
			if shipsState[coords[0]-1][coords[1]-1] != gui.Ship {
				shipsState[coords[0]-1][coords[1]-1] = state
			}
		} else if coords[1] == 0 && coords[0] > 0 && coords[0] < 9 {
			if shipsState[coords[0]+1][coords[1]+1] != gui.Ship {
				shipsState[coords[0]+1][coords[1]+1] = state
			}
			if shipsState[coords[0]-1][coords[1]+1] != gui.Ship {
				shipsState[coords[0]-1][coords[1]+1] = state
			}
		} else if coords[1] == 9 && coords[0] > 0 && coords[0] < 9 {
			if shipsState[coords[0]+1][coords[1]-1] != gui.Ship {
				shipsState[coords[0]+1][coords[1]-1] = state
			}
			if shipsState[coords[0]-1][coords[1]-1] != gui.Ship {
				shipsState[coords[0]-1][coords[1]-1] = state
			}

		}
	}
}
