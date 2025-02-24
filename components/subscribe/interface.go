/**
 * Created by PhpStorm.
 * @file   interface.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/9 20:42
 * @desc   interface.go
 */

package subscribe

import "github.com/x-module/helper/components/handler"

var SubscribeHandler SubPub

// SubPub 消息发布定义
type SubPub interface {
	// Subscribe 订阅消息
	Subscribe(channel string, callback handler.SubscribeCallback) error
	// Publish 发布消息
	Publish(channel string, message any) error
}
