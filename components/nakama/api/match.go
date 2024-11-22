/**
 * Created by Goland.
 * @file   matche.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/14 11:30
 * @desc   matche.go
 */

package api

import (
	"errors"
	"github.com/x-module/helper/components/nakama/common"
	"github.com/x-module/helper/components/request"

	"github.com/gin-gonic/gin"
)

type Match struct {
	common.NakamaApi
}
type MatchList struct {
	Matches []Matches `json:"matches"`
}
type Matches struct {
	MatchID       string `json:"match_id"`
	Authoritative bool   `json:"authoritative"`
	Label         string `json:"label"`
	Size          int    `json:"size"`
	TickRate      int    `json:"tick_rate"`
	HandlerName   string `json:"handler_name"`
}

func NewMatch(token string) *Match {
	match := new(Match)
	match.Token = token
	return match
}

type MatchState struct {
	Presences []Presences `json:"presences"`
	Tick      string      `json:"tick"`
	State     string      `json:"state"`
}
type Presences struct {
	UserID      string `json:"user_id"`
	SessionID   string `json:"session_id"`
	Username    string `json:"username"`
	Persistence bool   `json:"persistence"`
	Status      any    `json:"status"`
}

func (a *Match) GetMatchList(url string, mode string) (MatchList, error) {
	response, err := request.NewRequest().Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(url)
	if err != nil {
		return MatchList{}, err
	}
	defer response.Close()
	if response.StatusCode() != 200 {
		return MatchList{}, errors.New("request nakama server error")
	}
	var matchList MatchList
	err = response.Json(&matchList)
	if err != nil {
		return MatchList{}, err
	}
	return matchList, nil
}

// GetState 比赛状态
func (a *Match) GetState(url string, mode string) (MatchState, error) {
	response, err := request.NewRequest().Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(url)
	if err != nil {
		return MatchState{}, err
	}
	defer response.Close()
	if response.StatusCode() != 200 {
		return MatchState{}, errors.New("request nakama server error")
	}
	var matchState MatchState
	err = response.Json(&matchState)
	if err != nil {
		return MatchState{}, err
	}
	return matchState, nil
}
