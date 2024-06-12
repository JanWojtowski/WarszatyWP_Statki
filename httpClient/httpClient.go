package httpClient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type GameStarter struct {
	Coords     []string `json:"coords"`
	Desc       string   `json:"desc"`
	Nick       string   `json:"nick"`
	TargetNick string   `json:"target_nick"`
	Wpbot      bool     `json:"wpbot"`
}
type Auth struct {
	Auth string `json:"X-Auth-Token"`
}

type Opponent struct {
	Opponent     string `json:"opponent"`
	OpponentDesc string `json:"opp_desc"`
}

type LobbyInfo []struct {
	Nick string `json:"nick"`
}

type PlayerStats struct {
	Nick   string `json:"nick"`
	Games  int    `json:"games"`
	Wins   int    `json:"wins"`
	Rank   int    `json:"rank"`
	Points int    `json:"points"`
}

type Stats struct {
	Stats []PlayerStats `json:"stats"`
}

type Status struct {
	GameStatus     string   `json:"game_status"`
	LastGameStatus string   `json:"last_game_status"`
	Nick           string   `json:"nick"`
	OppShots       []string `json:"opp_shots"`
	Opponent       string   `json:"opponent"`
	ShouldFire     bool     `json:"should_fire"`
	Timer          int      `json:"timer"`
}

type Cord struct {
	Cord string `json:"coord"`
}

func StartGameWithBot(nick string, desc string, coords []string) string {
	posturl := "https://go-pjatk-server.fly.dev/api/game"

	body, err := json.Marshal(GameStarter{
		Coords:     coords,
		Desc:       desc,
		Nick:       nick,
		TargetNick: "",
		Wpbot:      true,
	})

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	authKey := res.Header.Get("x-auth-token")

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusServiceUnavailable {
			time.Sleep(1 * time.Second)
			return StartGameWithBot(nick, desc, coords)
		} else if res.StatusCode == http.StatusTooManyRequests {
			time.Sleep(1 * time.Second)
			return StartGameWithBot(nick, desc, coords)
		} else {
			log.Panicln(res.Status)
		}
	}

	return authKey
}

func StartGameMultiAttack(nick string, desc string, coords []string, target string) string {
	posturl := "https://go-pjatk-server.fly.dev/api/game"

	body, err := json.Marshal(GameStarter{
		Coords:     coords,
		Desc:       desc,
		Nick:       nick,
		TargetNick: target,
		Wpbot:      false,
	})

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	authKey := res.Header.Get("x-auth-token")

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusServiceUnavailable {
			time.Sleep(500 * time.Millisecond)
			return StartGameMultiAttack(nick, desc, coords, target)
		} else if res.StatusCode == http.StatusTooManyRequests {
			time.Sleep(500 * time.Millisecond)
			return StartGameMultiAttack(nick, desc, coords, target)
		} else {
			log.Panicln(res.Status)
		}
	}
	return authKey
}

func StartGameMultiLobby(nick string, desc string, coords []string) string {
	posturl := "https://go-pjatk-server.fly.dev/api/game"

	body, err := json.Marshal(GameStarter{
		Coords: coords,
		Desc:   desc,
		Nick:   nick,
		Wpbot:  false,
	})

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	authKey := res.Header.Get("x-auth-token")

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusServiceUnavailable {
			time.Sleep(500 * time.Millisecond)
			return StartGameMultiLobby(nick, desc, coords)
		} else if res.StatusCode == http.StatusTooManyRequests {
			time.Sleep(500 * time.Millisecond)
			return StartGameMultiLobby(nick, desc, coords)
		} else {
			log.Panicln(res.Status)
		}
	}
	return authKey
}

func GetOpponentInfo(authKey string) []string {
	geturl := "https://go-pjatk-server.fly.dev/api/game/desc"

	r, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		panic(err)
	}

	r.Header.Set("X-Auth-Token", authKey)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusServiceUnavailable {
			time.Sleep(500 * time.Millisecond)
			return GetOpponentInfo(authKey)
		} else if res.StatusCode == http.StatusTooManyRequests {
			time.Sleep(500 * time.Millisecond)
			return GetOpponentInfo(authKey)
		} else {
			log.Panicln(res.Status)
		}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data Opponent
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	if data.Opponent == "" {
		time.Sleep(500 * time.Millisecond)
		return GetOpponentInfo(authKey)
	}

	ret := []string{data.Opponent, data.OpponentDesc}

	return ret
}

func GetStatus(authKey string) Status {
	geturl := "https://go-pjatk-server.fly.dev/api/game"

	r, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		panic(err)
	}

	r.Header.Set("X-Auth-Token", authKey)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusServiceUnavailable {
			time.Sleep(500 * time.Millisecond)
			return GetStatus(authKey)
		} else if res.StatusCode == http.StatusTooManyRequests {
			time.Sleep(500 * time.Millisecond)
			return GetStatus(authKey)
		} else {
			log.Panicln(res.Status)
		}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data Status
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	return data
}

func PostFire(cord string, authKey string) string {
	posturl := "https://go-pjatk-server.fly.dev/api/game/fire"

	body, err := json.Marshal(Cord{
		Cord: cord,
	})

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Auth-Token", authKey)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode != http.StatusOK {
			if res.StatusCode == http.StatusServiceUnavailable {
				time.Sleep(500 * time.Millisecond)
				return PostFire(cord, authKey)
			} else if res.StatusCode == http.StatusTooManyRequests {
				time.Sleep(500 * time.Millisecond)
				return PostFire(cord, authKey)
			} else {
				log.Panicln(res.Status)
			}
		}
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data struct {
		Result string `json:"result"`
	}

	err = json.Unmarshal(resBody, &data)
	if err != nil {
		panic(err.Error())
	}

	return data.Result
}

func GetBoard(authKey string) []string {
	geturl := "https://go-pjatk-server.fly.dev/api/game/board"

	r, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		panic(err)
	}

	r.Header.Set("X-Auth-Token", authKey)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusServiceUnavailable {
			time.Sleep(500 * time.Millisecond)
			return GetBoard(authKey)
		} else if res.StatusCode == http.StatusTooManyRequests {
			time.Sleep(500 * time.Millisecond)
			return GetBoard(authKey)
		} else {
			log.Panicln(res.Status)
		}
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data struct {
		Board []string `json:"board"`
	}

	err = json.Unmarshal(resBody, &data)
	if err != nil {
		panic(err.Error())
	}

	return data.Board
}

func GetLobby() LobbyInfo {
	geturl := "https://go-pjatk-server.fly.dev/api/lobby"

	r, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusServiceUnavailable {
			time.Sleep(500 + time.Millisecond)
			return GetLobby()
		} else if res.StatusCode == http.StatusTooManyRequests {
			time.Sleep(500 * time.Millisecond)
			return GetLobby()
		} else {
			log.Panicln(res.Status)
		}
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data LobbyInfo
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		panic(err.Error())
	}

	return data
}

func DeleteSurrender(authKey string) bool {
	deleteurl := "https://go-pjatk-server.fly.dev/api/game/abandon"

	r, err := http.NewRequest("DELETE", deleteurl, nil)
	if err != nil {
		panic(err)
	}

	r.Header.Set("X-Auth-Token", authKey)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusServiceUnavailable {
			time.Sleep(500 * time.Millisecond)
			return DeleteSurrender(authKey)
		} else if res.StatusCode == http.StatusTooManyRequests {
			time.Sleep(500 * time.Millisecond)
			return DeleteSurrender(authKey)
		} else {
			log.Panicln(res.Status)
		}
	}

	return true
}

func GetRefresh(authKey string) {
	geturl := "https://go-pjatk-server.fly.dev/api/game/refresh"

	r, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		panic(err)
	}
	r.Header.Set("X-Auth-Token", authKey)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusServiceUnavailable {
			time.Sleep(500 * time.Millisecond)
			GetRefresh(authKey)
		} else if res.StatusCode == http.StatusTooManyRequests {
			time.Sleep(500 * time.Millisecond)
			GetRefresh(authKey)
		} else {
			log.Panicln(res.Status)
		}
	}
}

func GetStats() []PlayerStats {
	geturl := "https://go-pjatk-server.fly.dev/api/stats"

	r, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusServiceUnavailable {
			time.Sleep(500 + time.Millisecond)
			return GetStats()
		} else if res.StatusCode == http.StatusTooManyRequests {
			time.Sleep(500 * time.Millisecond)
			return GetStats()
		} else {
			log.Panicln(res.Status)
		}
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var stats Stats
	err = json.Unmarshal(resBody, &stats)
	if err != nil {
		panic(err.Error())
	}

	return stats.Stats
}
