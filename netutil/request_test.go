/**
 * Created by Goland
 * @file   request_test.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 14:23
 * @desc   request_test.go
 */

package netutil

import (
	"fmt"
	"testing"
)

func TestRequest(t *testing.T) {
	response, err := RequestUtils.SetCookies(map[string]string{
		"name":     "124",
		"password": "1234",
	}).SetHeaders(map[string]string{
		"auth": "super man",
		"sign": "sign the request",
	}).Debug().SetTimeout(1).Json().Get("http://127.0.0.1:9090", map[string]interface{}{
		"username": "username",
		"password": "password",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Content())

}
