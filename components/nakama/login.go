/**
 * Created by Goland.
 * @file   login.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/11 11:37
 * @desc   login.go
 */

package nakama

import (
	"crypto"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/x-module/helper/components/request"
	"github.com/x-module/helper/xlog"
	"time"
)

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
func (n *Api) parseConsoleToken(hmacSecretByte []byte, tokenString string) (username, email string, role UserRole, exp int64, ok bool) {
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
func (n *Api) testToken() (int, error) {
	token, err := jwt.Parse(n.Token.Token, func(token *jwt.Token) (any, error) {
		if s, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || s.Hash != crypto.SHA256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(n.loginParams.SignKey), nil
	})
	if err != nil {
		return InvalidToken, err
	}
	_, _, _, exp, ok := n.parseConsoleToken([]byte(n.loginParams.SignKey), n.Token.Token)
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
func (n *Api) GetToken() (LoginToken, error) {
	if n.Token.Token == "" {
		token, err := n.Login(n.loginParams)
		if err != nil {
			return LoginToken{}, err
		} else {
			return token, err
		}
	} else {
		_, err := n.testToken()
		if err != nil {
			n.Token = LoginToken{}
			return n.GetToken()
		} else {
			return n.Token, err
		}
	}
}

// Login 登录操作
func (n *Api) Login(params LoginParams) (LoginToken, error) {
	n.loginParams = params
	data := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: params.UserName,
		Password: params.Password,
	}

	url := fmt.Sprintf("%s%s", n.host, AuthenticateApiUrl)
	response, err := request.NewRequest().Debug(n.mode == xlog.DebugMode).Json().SetTimeout(10).Post(url, data)
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

	n.Token = loginToken
	return loginToken, nil
}
