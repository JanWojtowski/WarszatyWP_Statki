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
			if len(shipsTab) < 4 {
				if len(currentShip) == 0 {
					if canPlaceNewShip(char, shipsTab) {
						currentShip = append(currentShip, char)
						shipsTab = append(shipsTab, char)
					} else {
						mis.SetText("To close to another ships")
					}
				} else {
					if !checkSides(char, shipsWithoutCurrent(shipsTab, currentShip)) {
						if checkSides(char, currentShip) {
							currentShip = append(currentShip, char)
							shipsTab = append(shipsTab, char)
						} else {
							mis.SetText("Ship parts edges must be connected")
						}
					} else {
						mis.SetText("To close to another ship")
					}
				}
			} else if len(shipsTab) < 7 {
				if len(shipsTab) == 4 {
					currentShip = []string{}
					txt.SetText("1st 3pcs ship")
				}
				if len(currentShip) == 0 {
					if canPlaceNewShip(char, shipsTab) {
						currentShip = append(currentShip, char)
						shipsTab = append(shipsTab, char)
					} else {
						mis.SetText("To close to another ships")
					}
				} else {
					if !checkSides(char, shipsWithoutCurrent(shipsTab, currentShip)) {
						if checkSides(char, currentShip) {
							currentShip = append(currentShip, char)
							shipsTab = append(shipsTab, char)
						} else {
							mis.SetText("Ship parts edges must be connected")
						}
					} else {
						mis.SetText("To close to another ship")
					}
				}
			} else if len(shipsTab) < 10 {
				if len(shipsTab) == 7 {
					currentShip = []string{}
					txt.SetText("2nd 3pcs ship")
				}
				if len(currentShip) == 0 {
					if canPlaceNewShip(char, shipsTab) {
						currentShip = append(currentShip, char)
						shipsTab = append(shipsTab, char)
					} else {
						mis.SetText("To close to another ships")
					}
				} else {
					if !checkSides(char, shipsWithoutCurrent(shipsTab, currentShip)) {
						if checkSides(char, currentShip) {
							currentShip = append(currentShip, char)
							shipsTab = append(shipsTab, char)
						} else {
							mis.SetText("Ship parts edges must be connected")
						}
					} else {
						mis.SetText("To close to another ship")
					}
				}
			} else if len(shipsTab) < 12 {
				if len(shipsTab) == 10 {
					currentShip = []string{}
					txt.SetText("1st 2pcs ship")
				}
				if len(currentShip) == 0 {
					if canPlaceNewShip(char, shipsTab) {
						currentShip = append(currentShip, char)
						shipsTab = append(shipsTab, char)
					} else {
						mis.SetText("To close to another ships")
					}
				} else {
					if !checkSides(char, shipsWithoutCurrent(shipsTab, currentShip)) {
						if checkSides(char, currentShip) {
							currentShip = append(currentShip, char)
							shipsTab = append(shipsTab, char)
						} else {
							mis.SetText("Ship parts edges must be connected")
						}
					} else {
						mis.SetText("To close to another ship")
					}
				}
			} else if len(shipsTab) < 14 {
				if len(shipsTab) == 12 {
					currentShip = []string{}
					txt.SetText("2nd 2pcs ship")
				}
				if len(currentShip) == 0 {
					if canPlaceNewShip(char, shipsTab) {
						currentShip = append(currentShip, char)
						shipsTab = append(shipsTab, char)
					} else {
						mis.SetText("To close to another ships")
					}
				} else {
					if !checkSides(char, shipsWithoutCurrent(shipsTab, currentShip)) {
						if checkSides(char, currentShip) {
							currentShip = append(currentShip, char)
							shipsTab = append(shipsTab, char)
						} else {
							mis.SetText("Ship parts edges must be connected")
						}
					} else {
						mis.SetText("To close to another ship")
					}
				}
			} else if len(shipsTab) < 16 {
				if len(shipsTab) == 14 {
					currentShip = []string{}
					txt.SetText("3rd 2pcs ship")
				}
				if len(currentShip) == 0 {
					if canPlaceNewShip(char, shipsTab) {
						currentShip = append(currentShip, char)
						shipsTab = append(shipsTab, char)
					} else {
						mis.SetText("To close to another ships")
					}
				} else {
					if !checkSides(char, shipsWithoutCurrent(shipsTab, currentShip)) {
						if checkSides(char, currentShip) {
							currentShip = append(currentShip, char)
							shipsTab = append(shipsTab, char)
						} else {
							mis.SetText("Ship parts edges must be connected")
						}
					} else {
						mis.SetText("To close to another ship")
					}
				}
			} else if len(shipsTab) < 17 {
				if len(shipsTab) == 16 {
					currentShip = []string{}
					txt.SetText("1st 1pcs ship")
				}
				if len(currentShip) == 0 {
					if canPlaceNewShip(char, shipsTab) {
						currentShip = append(currentShip, char)
						shipsTab = append(shipsTab, char)
					} else {
						mis.SetText("To close to another ships")
					}
				} else {
					if !checkSides(char, shipsWithoutCurrent(shipsTab, currentShip)) {
						if checkSides(char, currentShip) {
							currentShip = append(currentShip, char)
							shipsTab = append(shipsTab, char)
						} else {
							mis.SetText("Ship parts edges must be connected")
						}
					} else {
						mis.SetText("To close to another ship")
					}
				}
			} else if len(shipsTab) < 18 {
				if len(shipsTab) == 17 {
					currentShip = []string{}
					txt.SetText("2nd 1pcs ship")
				}
				if len(currentShip) == 0 {
					if canPlaceNewShip(char, shipsTab) {
						currentShip = append(currentShip, char)
						shipsTab = append(shipsTab, char)
					} else {
						mis.SetText("To close to another ships")
					}
				} else {
					if !checkSides(char, shipsWithoutCurrent(shipsTab, currentShip)) {
						if checkSides(char, currentShip) {
							currentShip = append(currentShip, char)
							shipsTab = append(shipsTab, char)
						} else {
							mis.SetText("Ship parts edges must be connected")
						}
					} else {
						mis.SetText("To close to another ship")
					}
				}
			} else if len(shipsTab) < 19 {
				if len(shipsTab) == 18 {
					currentShip = []string{}
					txt.SetText("3rd 1pcs ship")
				}
				if len(currentShip) == 0 {
					if canPlaceNewShip(char, shipsTab) {
						currentShip = append(currentShip, char)
						shipsTab = append(shipsTab, char)
					} else {
						mis.SetText("To close to another ships")
					}
				} else {
					if !checkSides(char, shipsWithoutCurrent(shipsTab, currentShip)) {
						if checkSides(char, currentShip) {
							currentShip = append(currentShip, char)
							shipsTab = append(shipsTab, char)
						} else {
							mis.SetText("Ship parts edges must be connected")
						}
					} else {
						mis.SetText("To close to another ship")
					}
				}
			} else if len(shipsTab) < 20 {
				if len(shipsTab) == 19 {
					currentShip = []string{}
					txt.SetText("4st 1pcs ship")
				}
				if len(currentShip) == 0 {
					if canPlaceNewShip(char, shipsTab) {
						currentShip = append(currentShip, char)
						shipsTab = append(shipsTab, char)

						ui.Remove(line1)
						ui.Remove(line2)
						ui.Remove(line3)
						ui.Remove(txt)
						ui.Remove(mis)
						ui.Remove(board)

						game.PlayerInfo.ownBoard = true
						game.PlayerInfo.coords = shipsTab

						gameModeSelect(game)
					} else {
						mis.SetText("To close to another ships")
					}
				} else {
					if !checkSides(char, shipsWithoutCurrent(shipsTab, currentShip)) {
						if checkSides(char, currentShip) {
							currentShip = append(currentShip, char)
							shipsTab = append(shipsTab, char)
						} else {
							mis.SetText("Ship parts edges must be connected")
						}
					} else {
						mis.SetText("To close to another ship")
					}
				}
			}

			states := [10][10]gui.State{}
			for _, ply := range shipsTab {
				numbers := cordToNumbers(ply)
				states[numbers[0]][numbers[1]] = gui.Ship
			}
			board.SetStates(states)
		}
	}()
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
