/**
 * Created by Goland.
 * @file   leaderboard.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/13 17:56
 * @desc   leaderboard.go
 */

package nakama

import (
	"errors"
	"fmt"
	"github.com/x-module/helper/components/request"
	"time"

	"github.com/gin-gonic/gin"
)

type LeaderboardList struct {
	Leaderboards []LeaderboardInfo `json:"leaderboards"`
}
type LeaderboardInfo struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Category      int    `json:"category"`
	SortOrder     int    `json:"sort_order"`
	Size          int    `json:"size"`
	MaxSize       int    `json:"max_size"`
	MaxNumScore   int    `json:"max_num_score"`
	Operator      int    `json:"operator"`
	EndActive     int    `json:"end_active"`
	ResetSchedule string `json:"reset_schedule"`
	Metadata      string `json:"metadata"`
	CreateTime    any    `json:"create_time"`
	StartTime     any    `json:"start_time"`
	EndTime       any    `json:"end_time"`
	Duration      int    `json:"duration"`
	StartActive   int    `json:"start_active"`
	JoinRequired  bool   `json:"join_required"`
	Authoritative bool   `json:"authoritative"`
	Tournament    bool   `json:"tournament"`
}

type LeaderboardRecord struct {
	Records      []Records `json:"records"`
	OwnerRecords []any     `json:"owner_records"`
	NextCursor   string    `json:"next_cursor"`
	PrevCursor   string    `json:"prev_cursor"`
}

type Records struct {
	LeaderboardID string    `json:"leaderboard_id"`
	OwnerID       string    `json:"owner_id"`
	Username      string    `json:"username"`
	Score         string    `json:"score"`
	Subscore      string    `json:"subscore"`
	NumScore      int       `json:"num_score"`
	Metadata      string    `json:"metadata"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
	ExpiryTime    any       `json:"expiry_time"`
	Rank          string    `json:"rank"`
	MaxNumScore   int       `json:"max_num_score"`
}

// GetLeaderboardList 获取排行榜列表
func (n *Api) GetLeaderboardList(url string, mode string) (LeaderboardList, error) {
	response, err := request.NewRequest().Debug(mode == gin.DebugMode).SetHeaders(n.GetNakamaHeader()).SetTimeout(10).Get(url)
	if err != nil {
		return LeaderboardList{}, err
	}
	defer response.Close()

	if response.StatusCode() != 200 {
		return LeaderboardList{}, errors.New("request nakama server error")
	}
	var leaderboardList LeaderboardList
	err = response.Json(&leaderboardList)
	if err != nil {
		return LeaderboardList{}, err
	}
	return leaderboardList, nil
}

// DeleteLeaderboard 删除排行榜
func (n *Api) DeleteLeaderboard(url string, mode string) error {
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(n.GetNakamaHeader()).SetTimeout(10).Delete(url)
	if err != nil {
		return err
	}
	defer response.Close()
	if response.StatusCode() != 200 {
		errorMsg, _ := response.Content()
		return errors.New(errorMsg)
	}
	return nil
}

// GetLeaderboardDetail 获取排行榜详情
func (n *Api) GetLeaderboardDetail(url string, mode string) (LeaderboardInfo, error) {
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(n.GetNakamaHeader()).SetTimeout(10).Get(url)
	if err != nil {
		return LeaderboardInfo{}, err
	}
	defer response.Close()
	if response.StatusCode() != 200 {
		return LeaderboardInfo{}, fmt.Errorf("get leaderboard detail err")
	}
	var leaderboardInfo LeaderboardInfo
	err = response.Json(&leaderboardInfo)
	if err != nil {
		return LeaderboardInfo{}, err
	}
	return leaderboardInfo, nil
}

// GetLeaderboardRecord 获取排行榜记录
func (n *Api) GetLeaderboardRecord(url string, mode string) (LeaderboardRecord, error) {
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(n.GetNakamaHeader()).SetTimeout(10).Get(url)
	if err != nil {

		return LeaderboardRecord{}, err
	}
	defer response.Close()
	if response.StatusCode() != 200 {
		return LeaderboardRecord{}, err
	}
	var leaderboardRecord LeaderboardRecord
	err = response.Json(&leaderboardRecord)
	if err != nil {
		return LeaderboardRecord{}, err
	}
	return leaderboardRecord, nil
}
