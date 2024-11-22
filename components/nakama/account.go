/**
 * Created by Goland.
 * @file   account.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/8 19:32
 * @desc   account.go
 */

package nakama

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/x-module/helper/components/request"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type AccountInfo struct {
	Account     AccountData `json:"account"`
	DisableTime any         `json:"disable_time"`
}
type AccountData struct {
	User        User   `json:"user"`
	Wallet      string `json:"wallet"`
	Email       string `json:"email"`
	Devices     []any  `json:"devices"`
	CustomID    string `json:"custom_id"`
	VerifyTime  any    `json:"verify_time"`
	DisableTime string `json:"disable_time"`
}
type Accounts struct {
	Users      []User `json:"users"`
	TotalCount int    `json:"total_count"`
	NextCursor string `json:"next_cursor"`
}

type BanPlayer struct {
	ID          string     `json:"id"`
	Username    string     `json:"username"`
	DisplayName string     `json:"display_name"`
	AvatarURL   string     `json:"avatar_url"`
	LangTag     string     `json:"lang_tag"`
	Metadata    string     `json:"metadata"`
	EdgeCount   int        `json:"edge_count"`
	CreateTime  CreateTime `json:"create_time"`
	UpdateTime  UpdateTime `json:"update_time"`
	SteamID     string     `json:"steam_id,omitempty"`
}
type CreateTime struct {
	Seconds int `json:"seconds"`
}
type UpdateTime struct {
	Seconds int `json:"seconds"`
}

type User struct {
	ID                    string    `json:"id"`
	Username              string    `json:"username"`
	DisplayName           string    `json:"display_name"`
	AvatarURL             string    `json:"avatar_url"`
	LangTag               string    `json:"lang_tag"`
	Location              string    `json:"location"`
	Timezone              string    `json:"timezone"`
	Metadata              string    `json:"metadata"`
	FacebookID            string    `json:"facebook_id"`
	GoogleID              string    `json:"google_id"`
	GamecenterID          string    `json:"gamecenter_id"`
	SteamID               string    `json:"steam_id"`
	Online                bool      `json:"online"`
	EdgeCount             int       `json:"edge_count"`
	CreateTime            time.Time `json:"create_time"`
	UpdateTime            time.Time `json:"update_time"`
	FacebookInstantGameID string    `json:"facebook_instant_game_id"`
	AppleID               string    `json:"apple_id"`
}

type Encoder struct{}
type Params struct {
	Updates   any     `json:"updates"`
	CloneFrom any     `json:"cloneFrom"`
	Encoder   Encoder `json:"encoder"`
	Map       any     `json:"map"`
}
type NormalizedNames struct{}
type LazyUpdate struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Op    string `json:"op"`
}
type LazyInit struct {
	NormalizedNames NormalizedNames   `json:"normalizedNames"`
	LazyUpdate      any               `json:"lazyUpdate"`
	Headers         map[string]string `json:"headers"`
}
type Headers struct {
	NormalizedNames NormalizedNames   `json:"normalizedNames"`
	LazyUpdate      []LazyUpdate      `json:"lazyUpdate"`
	Headers         map[string]string `json:"headers"`
	LazyInit        LazyInit          `json:"lazyInit"`
}
type FriendResponse struct {
	Friends []Friends `json:"friends"`
	Cursor  string    `json:"cursor"`
}
type Friends struct {
	State      int       `json:"state"`
	UpdateTime time.Time `json:"update_time"`
	User       User      `json:"user,omitempty"`
}
type Payload struct {
	Params  Params  `json:"params"`
	Headers Headers `json:"headers"`
}

// GetAccountList 获取用户列表
func (n *Api) GetAccountList(apiUrl string, filter string, cursor string, mode string) (Accounts, error) {
	apiUrl = apiUrl + "?a=a"
	if filter != "" {
		filter = url.QueryEscape(filter)
		apiUrl = fmt.Sprintf("%s&filter=%s", apiUrl, filter)
	}
	if cursor != "" {
		apiUrl = fmt.Sprintf("%s&cursor=%s", apiUrl, cursor)
	}
	response, err := request.NewRequest().Debug(mode == gin.DebugMode).SetHeaders(n.GetNakamaHeader()).SetTimeout(10).Get(apiUrl)
	if err != nil {
		return Accounts{}, err
	}
	defer response.Close()
	if response.StatusCode() != 200 {
		return Accounts{}, errors.New("request nakama server error")
	}
	var accounts Accounts
	err = response.Json(&accounts)
	if err != nil {
		return Accounts{}, err
	}
	return accounts, nil
}

// GetAccountBanList 获取用用列表
func (n *Api) GetAccountBanList(apiUrl string, UserID string, UserName string, Offset int, Limit int, mode string) ([]BanPlayer, error) {
	params := map[string]any{
		"user_id":   UserID,
		"user_name": UserName,
		"offset":    Offset,
		"limit":     Limit,
	}
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(map[string]string{
		"Accept": "application/json",
	}).SetTimeout(20).Post(apiUrl, params)
	if err != nil {
		return nil, err
	}
	defer response.Close()

	if response.StatusCode() != 200 {
		return nil, errors.New("request nakama server error")
	}
	// c, _ := response.Content()
	// utils.JsonDisplay(c)
	var accounts []BanPlayer
	err = response.Json(&accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// GetAccountDetail 获取用户详情
func (n *Api) GetAccountDetail(id string, url string, mode string) (AccountInfo, error) {
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(n.GetNakamaHeader()).SetTimeout(10).Get(url)
	if err != nil {
		return AccountInfo{}, err
	}
	defer response.Close()
	if response.StatusCode() != 200 {
		return AccountInfo{}, errors.New("request nakama server error")
	}
	var accountInfo AccountInfo
	err = response.Json(&accountInfo)
	if err != nil {
		return AccountInfo{}, err
	}
	return accountInfo, nil
}

func (n *Api) UpdateAccount(id string, params []byte, url string, mode string) (string, error) {
	type Payload struct {
		Username    string `json:"username"`
		DisplayName string `json:"display_name"`
		AvatarURL   string `json:"avatar_url"`
		Location    string `json:"location"`
		Timezone    string `json:"timezone"`
		Metadata    string `json:"metadata"`
	}
	var data Payload
	_ = json.Unmarshal(params, &data)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(n.GetNakamaHeader()).SetTimeout(10).Post(url, data)
	if err != nil {
		return "", err
	}
	defer response.Close()
	type ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details []any  `json:"details"`
	}
	if response.StatusCode() != 200 {
		res, _ := response.Content()
		var errorResp ErrorResponse
		_ = json.Unmarshal([]byte(res), &errorResp)
		return errorResp.Message, errors.New(errorResp.Message)
	}
	return "success", nil
}

// Unlink account unlink
func (n *Api) Unlink(url string, mode string) error {
	data := Payload{
		Params: Params{},
		Headers: Headers{
			NormalizedNames: NormalizedNames{},
			LazyUpdate: []LazyUpdate{
				{
					Name:  "Authorization",
					Value: "Bearer ",
					Op:    "s",
				},
			},
			Headers:  map[string]string{},
			LazyInit: LazyInit{},
		},
	}
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(n.GetNakamaHeader()).Json().SetTimeout(10).Post(url, data)
	if err != nil {
		return err
	}
	defer response.Close()
	type ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details []any  `json:"details"`
	}
	if response.StatusCode() != 200 {
		res, _ := response.Content()
		var errorResp ErrorResponse
		_ = json.Unmarshal([]byte(res), &errorResp)
		return errors.New(errorResp.Message)
	}
	return nil
}

// ChangeAccount 修改邮箱密码
func (n *Api) ChangeAccount(email string, password string, url string, mode string) error {
	type Payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	data := Payload{
		Email:    email,
		Password: password,
	}

	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(n.GetNakamaHeader()).Json().SetTimeout(10).Post(url, data)
	if err != nil {
		return err
	}
	defer response.Close()
	type ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details []any  `json:"details"`
	}
	if response.StatusCode() != 200 {
		res, _ := response.Content()
		var errorResp ErrorResponse
		_ = json.Unmarshal([]byte(res), &errorResp)
		return errors.New(errorResp.Message)
	}
	return nil
}

// GetFriends 获取账户朋友
func (n *Api) GetFriends(url string, mode string) (FriendResponse, error) {
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(n.GetNakamaHeader()).SetTimeout(10).Get(url)
	if err != nil {
		return FriendResponse{}, err
	}
	defer response.Close()
	if response.StatusCode() != 200 {
		errorMsg, _ := response.Content()
		return FriendResponse{}, errors.New(errorMsg)
	}
	var friendResponse FriendResponse
	err = response.Json(&friendResponse)
	if err != nil {
		return FriendResponse{}, err
	}
	return friendResponse, nil
}

// DeleteFriend 删除好友
func (n *Api) DeleteFriend(url string, mode string) error {
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

// DeleteAccount 删除账户
func (n *Api) DeleteAccount(url string, mode string) error {
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

func (n *Api) Enable(url string, mode string) error {
	data := Payload{
		Params: Params{},
		Headers: Headers{
			NormalizedNames: NormalizedNames{},
			LazyUpdate: []LazyUpdate{
				{
					Name:  "Authorization",
					Value: "Bearer ",
					Op:    "s",
				},
			},
			Headers:  map[string]string{},
			LazyInit: LazyInit{},
		},
	}
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(n.GetNakamaHeader()).Json().SetTimeout(10).Post(url, data)
	if err != nil {
		return err
	}
	defer response.Close()
	type ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details []any  `json:"details"`
	}
	if response.StatusCode() != 200 {
		res, _ := response.Content()
		var errorResp ErrorResponse
		_ = json.Unmarshal([]byte(res), &errorResp)
		return errors.New(errorResp.Message)
	}
	return nil
}
func (n *Api) Disable(url string, mode string) error {
	data := Payload{
		Params: Params{},
		Headers: Headers{
			NormalizedNames: NormalizedNames{},
			LazyUpdate: []LazyUpdate{
				{
					Name:  "Authorization",
					Value: "Bearer ",
					Op:    "s",
				},
			},
			Headers:  map[string]string{},
			LazyInit: LazyInit{},
		},
	}
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(n.GetNakamaHeader()).Json().SetTimeout(10).Post(url, data)
	if err != nil {
		return err
	}
	defer response.Close()
	type ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details []any  `json:"details"`
	}
	if response.StatusCode() != 200 {
		res, _ := response.Content()
		var errorResp ErrorResponse
		_ = json.Unmarshal([]byte(res), &errorResp)
		return errors.New(errorResp.Message)
	}
	return nil
}
