package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func gameList() ([]*Game, error) {
	res, err := http.Get("https://bifen4m.qiumibao.com/json/list.htm")
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
	return gameResp.List, err
}
