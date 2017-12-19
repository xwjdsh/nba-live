package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getGames() ([]*Game, error) {
	res, err := http.Get("http://bifen4m.qiumibao.com/json/list.htm")
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	var gameResp struct {
		Code   string  `json:"code"`
		Second string  `json:"second"`
		List   []*Game `json:"list"`
	}
	err = json.Unmarshal(result, &gameResp)
	if err == nil {
		games := []*Game{}
		for _, game := range gameResp.List {
			if game.PeriodCn != "" && game.Type == "basketball" {
				game.PeriodCn = strings.Replace(game.PeriodCn, "\n", " ", -1)
				games = append(games, game)
			}
		}
		return games, nil
	}
	return nil, err
}

func getMaxsid(gameId string) (int, error) {
	res, err := http.Get(fmt.Sprintf("http://dingshi4pc.qiumibao.com/livetext/data/cache/max_sid/%s/0.htm", gameId))
	if err != nil {
		return 0, err
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	res.Body.Close()
	return strconv.Atoi(string(result))
}

func getGameInfo(gameId string) (*GameInfo, error) {
	res, err := http.Get(fmt.Sprintf("http://bifen4pc2.qiumibao.com/json/%s/%s.htm", time.Now().Format("2006-01-02"), gameId))
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	gameInfo := &GameInfo{}
	err = json.Unmarshal(result, gameInfo)
	return gameInfo, err
}

func getLiveRecord(gameId string, maxSid int) ([]*GameLiveRecord, error) {
	res, err := http.Get(fmt.Sprintf("http://dingshi4pc.qiumibao.com/livetext/data/cache/livetext/%s/0/lit_page_2/%d.htm", gameId, maxSid))
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	records := []*GameLiveRecord{}
	err = json.Unmarshal(result, &records)
	return records, err
}
