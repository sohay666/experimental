package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	itg "toktok/integration"
	"toktok/utils"
)

type reqVideo struct {
	Keyword string `json:"keyword"`
	Type    string `json:"type"`
}

type respVideos struct {
	Code   int               `json:"code"`
	Status string            `json:"status"`
	Data   itg.PaxelResponse `json:"data"`
}

func VideosHandler(w http.ResponseWriter, r *http.Request) {
	s := utils.HTTPSetup{Writer: w, Request: r}

	// 1. get input query from client
	queryValues := s.Request.URL.Query()
	input := reqVideo{
		Keyword: queryValues.Get("keyword"),
		Type:    queryValues.Get("type"),
	}

	// 2. validate the request make sure keywords is filled
	if err := input.validate(); err != nil {
		s.ErrorResp(err.Error())
		return
	}
	resBody := itg.PaxelResponse{}
	uniqueKeyword := utils.Hash(input.Keyword + input.Type)
	// 3. check is keyword is already have an cache , baseon hash keyword
	cache := utils.GetInstanceRedis(utils.CACHE_VIDEO).Get(context.Background(), uniqueKeyword).Val()
	if len(cache) == 0 {
		// 3.a) if cache not found, request to 3party
		paxel := itg.CfgPaxel{Keyword: input.Keyword}
		resp, err := paxel.Search()
		if err != nil {
			s.ErrorResp(err.Error())
			return
		}
		resBody = resp
		// 4. save cache the response from 3party
		if err = saveCache(resBody, uniqueKeyword); err != nil {
			s.ErrorResp(err.Error())
			return
		}
	} else {
		//3.b) if cache is found, load it and convert to object
		resp, err := getCache(cache, resBody)
		if err != nil {
			s.ErrorResp(err.Error())
			return
		}
		resBody = resp
	}

	res := respVideos{
		Code:   http.StatusOK,
		Status: "success",
		Data:   resBody,
	}
	s.SuccessResp(res)
}

func (r reqVideo) validate() error {
	if r.Type == "search" && len(r.Keyword) == 0 {
		return fmt.Errorf("Keyword is required")
	}
	if r.Type != "popular" && r.Type != "search" {
		return fmt.Errorf("type only accept popular/search")
	}
	return nil
}

func saveCache(input itg.PaxelResponse, key string) (err error) {
	value, err := json.Marshal(input)
	if err != nil {
		return
	}
	days := time.Duration(utils.Config.App.ExpiredVideo)
	utils.GetInstanceRedis(utils.CACHE_VIDEO).Set(context.Background(), key, string(value), time.Duration((24*time.Hour)*days))
	return
}

func getCache(str string, input itg.PaxelResponse) (out itg.PaxelResponse, err error) {
	json.Unmarshal([]byte(str), &out)
	if err == nil {
		return
	}
	return
}
