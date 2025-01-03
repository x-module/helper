/**
 * Created by PhpStorm.
 * @file   redis.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/9 20:41
 * @desc   redis.go
 */

package subscribe

import (
	"github.com/x-module/helper/components/handler"
)

type RedisSubscribe struct {
}

func NewRedisSubscribe() {
	SubscribeHandler = new(RedisSubscribe)
}

// Subscribe 订阅消息
func (s *RedisSubscribe) Subscribe(channel string, callback handler.SubscribeCallback) error {
	err := handler.RedisHandler.Subscribe(channel, func(message string) {
		// 处理消息
		callback(message)
	})
	return err
}

// Publish 发布数据
func (s *RedisSubscribe) Publish(channel string, message any) error {
	err := handler.RedisHandler.Publish(channel, message)
	if err != nil {
		return err
	}
	return nil
}
