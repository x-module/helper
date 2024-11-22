/**
 * Created by Goland.
 * @file   count.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/14 18:18
 * @desc   count.go
 */

package api

import (
	"errors"
	"github.com/x-module/helper/components/nakama/common"
	"github.com/x-module/helper/components/request"
	"time"
)

type Count struct {
	common.NakamaApi
}
type CountResponse struct {
	Nodes     []Node    `json:"nodes"`
	Timestamp time.Time `json:"timestamp"`
}
type Node struct {
	Name           string  `json:"name"`
	Health         int     `json:"health"`
	SessionCount   int     `json:"session_count"`
	PresenceCount  int     `json:"presence_count"`
	MatchCount     int     `json:"match_count"`
	GoroutineCount int     `json:"goroutine_count"`
	AvgLatencyMs   float64 `json:"avg_latency_ms"`
	AvgRateSec     float64 `json:"avg_rate_sec"`
	AvgInputKbs    float64 `json:"avg_input_kbs"`
	AvgOutputKbs   float64 `json:"avg_output_kbs"`
}

func NewCount(token string) *Count {
	count := new(Count)
	count.Token = token
	return count
}

func (a *Count) GetGameServerInfo(apiUrl string, debug bool) (CountResponse, error) {
	response, err := request.NewRequest().Debug(debug).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(apiUrl)
	if err != nil {
		return CountResponse{}, err
	}
	defer response.Close()
	if response.StatusCode() != 200 {
		return CountResponse{}, errors.New("request nakama server error")
	}
	var countResponse CountResponse
	_, err = response.JsonReturn(&countResponse)
	if err != nil {
		return CountResponse{}, err
	}
	return countResponse, nil
}
