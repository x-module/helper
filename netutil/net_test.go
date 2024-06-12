/**
 * Created by Goland
 * @file   net_test.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 14:12
 * @desc   net_test.go
 */

package netutil

import (
	"github.com/x-module/helper/debug"
	"github.com/x-module/helper/internal"
	"net/http"
	"testing"
)

func TestGetPublicIpInfo(t *testing.T) {
	info, err := GetPublicIpInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("GetPublicIpInfo success,info:")
	debug.Display(info)
}

func TestGetInternalIp(t *testing.T) {
	ip := GetInternalIp()
	t.Log("GetInternalIp success,ip:", ip)
}

func TestGetIps(t *testing.T) {
	ip := GetIps()
	t.Log("GetIps success,ip:", ip)
}

func TestGetMac(t *testing.T) {
	mac := GetMacAddress()
	t.Log("GetMac success,mac:", mac)
}

func TestGetRequestPublicIp(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetPublicIpInfo")
	ip := "36.112.24.10"
	request := http.Request{
		Method: "GET",
		Header: http.Header{
			"X-Forwarded-For": {ip},
		},
	}
	publicIp := GetRequestPublicIp(&request)
	assert.Equal(publicIp, ip)
	request = http.Request{
		Method: "GET",
		Header: http.Header{
			"X-Real-Ip": {ip},
		},
	}
	publicIp = GetRequestPublicIp(&request)
	assert.Equal(publicIp, ip)
}

func TestEncodeUrl(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsInternalIP")

	urlAddr := "http://www.lancet.com?a=1&b=[2]"
	encodedUrl, err := EncodeUrl(urlAddr)
	if err != nil {
		t.Log(err)
	}

	expected := "http://www.lancet.com?a=1&b=%5B2%5D"
	assert.Equal(expected, encodedUrl)
}
