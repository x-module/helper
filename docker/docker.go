/**
 * Created by Goland
 * @file   docker.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/11/22 16:15
 * @desc   docker.go
 */

package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// GetContainers 获取 Docker 容器列表
func GetContainers(host string, port int, version string) (result []types.ContainerJSON, err error) {
	// 创建 Docker 客户端
	hostAddr := fmt.Sprintf("tcp://%s:%d", host, port)
	cli, err := client.NewClientWithOpts(
		client.WithHost(hostAddr),
		client.WithVersion(version), // 指定版本
	)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	// 获取容器列表
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, err
	}
	// 遍历容器
	for _, cont := range containers {
		// 获取容器详细信息
		inspect, err := cli.ContainerInspect(context.Background(), cont.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, inspect)
	}
	return result, nil
}
