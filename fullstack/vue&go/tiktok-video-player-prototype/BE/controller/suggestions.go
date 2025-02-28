package controller

import (
	"fmt"
	"net/http"
	itg "toktok/integration"
	"toktok/utils"
)

type reqSuggestions struct {
	Keyword string `json:"keyword"`
}

type DataSuggestions struct {
	Prefix      string   `json:"prefix"`
	Suggestions []string `json:"suggestions"`
}
type respSuggestions struct {
	Code   int             `json:"code"`
	Status string          `json:"status"`
	Data   DataSuggestions `json:"data"`
}

func SuggestionsHandler(w http.ResponseWriter, r *http.Request) {
	s := utils.HTTPSetup{Writer: w, Request: r}

	// 1. get input query from client
	queryValues := s.Request.URL.Query()
	input := reqVideo{
		Keyword: queryValues.Get("keyword"),
	}

	paxel := itg.CfgPaxel{Keyword: input.Keyword}
	resp, err := paxel.Suggestions()
	if err != nil {
		s.ErrorResp(err.Error())
		return
	}
	res := respSuggestions{
		Code:   http.StatusOK,
		Status: "success",
		Data: DataSuggestions{
			Prefix:      resp.Data.Attributes.Prefix,
			Suggestions: resp.Data.Attributes.Suggestions,
		},
	}
	s.SuccessResp(res)
}

func (r reqSuggestions) validate() error {
	if len(r.Keyword) == 0 {
		return fmt.Errorf("Keyword is required")
	}
	return nil
}
