/**
 * Created by PhpStorm.
 * @file   redis.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/5/15 15:43
 * @desc   redis.go
 */

package dirver

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// Redis 配置
type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

// RedisConfig 配置
type RedisConfig struct {
	List map[string]Redis `yaml:"list"`
	Db   int              `yaml:"db"`
	TLS  int              `yaml:"tls"`
}

var password = map[string]string{}

// 设置单个实例的密码
func userPassForAddr(addr string) string {
	if pass, exist := password[addr]; exist {
		return pass
	}
	return ""
}

// InitializeRedis 链接Redis服务
func InitializeRedis(config RedisConfig) *redis.Ring {
	addr := map[string]string{}
	for key, conf := range config.List {
		address := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
		addr[key] = address
		password[address] = conf.Password
	}
	option := &redis.RingOptions{
		Addrs: addr,
		DB:    config.Db,
		NewClient: func(opt *redis.Options) *redis.Client {
			pass := userPassForAddr(opt.Addr)
			opt.Password = pass
			return redis.NewClient(opt)
		},
		DialTimeout:  20 * time.Second, // 设置连接超时
		ReadTimeout:  20 * time.Second, // 设置读取超时
		WriteTimeout: 20 * time.Second, // 设置写入超时
	}
	if config.TLS == 1 { // 使用tls加密通讯
		option.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	rdb := redis.NewRing(option)
	// 验证服务
	err := rdb.ForEachShard(context.Background(), func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	return rdb
}
