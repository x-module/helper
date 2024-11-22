/**
 * Created by Goland.
 * @file   login.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/11 11:37
 * @desc   login.go
 */

package auth

import (
	"crypto"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/x-module/helper/components/nakama/common"
	"github.com/x-module/helper/components/request"
	"github.com/x-module/helper/xlog"
	"time"
)

// LoginToken 身份验证token
type LoginToken struct {
	Token string   `json:"token"`
	Uname string   `json:"uname"`
	Email string   `json:"email"`
	Role  UserRole `json:"role"`
}

type ConsoleTokenClaims struct {
	Username  string   `json:"usn,omitempty"`
	Email     string   `json:"ema,omitempty"`
	Role      UserRole `json:"rol,omitempty"`
	ExpiresAt int64    `json:"exp,omitempty"`
	Cookie    string   `json:"cki,omitempty"`
}

// InvalidToken 无效token
const InvalidToken = 2

// EffectiveToken 有效token
const EffectiveToken = 1

// ExpireToken 过期token
const ExpireToken = 3

type UserRole int32

type Auth struct {
	common.NakamaApi
	userName string
	password string
	host     string
	model    xlog.LogMode
	signKey  string
}

func NewAuth(userName string, password string, host string, signKey string, model xlog.LogMode) *Auth {
	auth := new(Auth)
	auth.userName = userName
	auth.password = password
	auth.host = host
	auth.model = model
	auth.signKey = signKey
	return auth
}

// Valid 校验
func (stc *ConsoleTokenClaims) Valid() error {
	// Verify expiry.
	if stc.ExpiresAt <= time.Now().UTC().Unix() {
		vErr := new(jwt.ValidationError)
		vErr.Inner = errors.New("token is expired")
		vErr.Errors |= jwt.ValidationErrorExpired
		return vErr
	}
	return nil
}

// 解析token
func (a *Auth) parseConsoleToken(hmacSecretByte []byte, tokenString string) (username, email string, role UserRole, exp int64, ok bool) {
	token, err := jwt.ParseWithClaims(tokenString, &ConsoleTokenClaims{}, func(token *jwt.Token) (any, error) {
		if s, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || s.Hash != crypto.SHA256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecretByte, nil
	})
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*ConsoleTokenClaims)
	if !ok || !token.Valid {
		return
	}
	return claims.Username, claims.Email, claims.Role, claims.ExpiresAt, true
}

// token 检测
func (a *Auth) testToken(loginToken LoginToken) (int, error) {
	token, err := jwt.Parse(loginToken.Token, func(token *jwt.Token) (any, error) {
		if s, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || s.Hash != crypto.SHA256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.signKey), nil
	})
	if err != nil {
		return InvalidToken, err
	}
	_, _, _, exp, ok := a.parseConsoleToken([]byte(a.signKey), loginToken.Token)
	if !ok || !token.Valid {
		// The token or its claims are invalid.
		return InvalidToken, err
	}
	if exp <= time.Now().UTC().Unix() {
		// Token expired.
		return ExpireToken, err
	}
	return EffectiveToken, nil
}

// GetToken 获取身份token
func (a *Auth) GetToken(loginToken LoginToken) (LoginToken, error) {
	if loginToken.Token == "" {
		token, err := a.login()
		if err != nil {
			return LoginToken{}, err
		} else {
			return token, err
		}
	} else {
		_, err := a.testToken(loginToken)
		if err != nil {
			// if checkResult == ExpireToken { // token过期
			return a.GetToken(LoginToken{})
			// }
			// return LoginToken{}, err
		} else {
			return loginToken, err
		}
	}
}

// 登录操作
func (a *Auth) login() (LoginToken, error) {
	data := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: a.userName,
		Password: a.password,
	}

	url := fmt.Sprintf("%s/%s", a.host, common.AuthenticateApiUrl)
	response, err := request.NewRequest().Debug(a.model == xlog.DebugMode).Json().SetTimeout(10).Post(url, data)
	if err != nil {
		return LoginToken{}, err
	}
	defer response.Close()
	if response.StatusCode() != 200 {
		return LoginToken{}, errors.New("request nakama server error")
	}
	var loginToken LoginToken
	err = response.Json(&loginToken)
	if err != nil {
		return LoginToken{}, err
	}
	return loginToken, nil
}
