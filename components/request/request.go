/**
 * Created by Goland
 * @file   request.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2025/1/7 19:36
 * @desc   request.go
 */

package request

import (
	"fmt"
	"net/http"
	url2 "net/url"

	"github.com/x-module/helper/components/request/plugins"
	retry "gopkg.in/h2non/gentleman-retry.v2"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugin"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
	"gopkg.in/h2non/gentleman.v2/plugins/transport"
)

type Request struct {
	*gentleman.Client
}

func NewRequest() *Request {
	return &Request{
		Client: gentleman.New(),
	}
}

func (r *Request) Use(p plugin.Plugin) *Request {
	r.Client.Use(p)
	return r
}

func (r *Request) URL(url string) *Request {
	r.Client.URL(url)
	return r
}

// Debug 开启调试模式
// 如果传入的参数为false，则关闭调试模式
func (r *Request) Debug(isTrace ...bool) *Request {
	if len(isTrace) == 0 || isTrace[0] {
		customTransport := &plugins.DebugTransport{
			Transport: http.DefaultTransport,
		}
		r.Client.Use(transport.Set(customTransport))
	}
	return r
}

// Retry 开启重试机制
// RetryTimes = 3 重试次数
// RetryWait = 100 * time.Millisecond 重试等待时间
func (r *Request) Retry() *Request {
	r.Client.Use(retry.New(retry.ConstantBackoff))
	return r
}

// GetRequest 发送GET请求
func (r *Request) GetRequest(url string, params map[string]string) (*gentleman.Response, error) {
	response, err := r.Client.URL(url).Get().SetQueryParams(params).Send()
	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, fmt.Errorf("invalid server response: %d", response.StatusCode)
	}
	return response, nil
}

// CommonPostRequest 发送普通POST请求
func (r *Request) CommonPostRequest(url string, params map[string]string) (*gentleman.Response, error) {
	formData := url2.Values{}
	for k, v := range params {
		formData.Set(k, v)
	}
	r.Client.Use(body.String(formData.Encode()))
	r.Client.Use(headers.Set("Content-Type", "application/x-www-form-urlencoded"))
	response, err := r.Client.URL(url).Post().Send()
	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, fmt.Errorf("invalid server response: %d", response.StatusCode)
	}
	return response, nil
}

func (r *Request) JsonPostRequest(url string, params any) (*gentleman.Response, error) {
	response, err := r.Client.Use(body.JSON(params)).URL(url).Post().Send()
	if err != nil {
		return nil, err
	}
	//if !response.Ok {
	//	return nil, fmt.Errorf("invalid server response: %d", response.StatusCode)
	//}
	return response, nil
}
