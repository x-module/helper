/**
 * Created by Goland
 * @file   transport.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2025/1/8 14:00
 * @desc   transport.go
 */

package plugins

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

type DebugTransport struct {
	Transport http.RoundTripper
}

func (t *DebugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// 打印请求信息
	fmt.Println("================================ DebugTrace ================================")
	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		fmt.Println("Failed to dump request:", err)
	} else {
		fmt.Println(strings.TrimSpace(string(dump)))
	}

	// 执行请求
	res, err := t.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// 打印响应信息
	dump, err = httputil.DumpResponse(res, false)
	if err != nil {
		fmt.Println("Failed to dump response:", err)
	} else {
		fmt.Println("-------------------------------- Response ---------------------------------")
		fmt.Println(strings.TrimSpace(string(dump)))
	}
	fmt.Println("================================ DebugTrace ================================")
	return res, nil
}
