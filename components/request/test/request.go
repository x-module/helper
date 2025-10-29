package main

import (
	"fmt"
	"github.com/x-module/helper/components/request"
)

func main() {
	response, err := request.NewRequest().SetCookies(map[string]string{
		"name":     "124",
		"password": "1234",
	}).SetHeaders(map[string]string{
		"auth": "super man",
		"sign": "sign the request",
	}).Debug().SetTimeout(1).Json().Get("http://127.0.0.1:9090", map[string]any{
		"username": "username",
		"password": "password",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Content())

}
