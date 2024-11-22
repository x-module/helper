/**
 * Created by Goland
 * @file   ssh.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/11/21 22:42
 * @desc   ssh.go
 */

package ssh

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
)

type HookFunc func(key int, command string)

type ServerInfo struct {
	Host     string
	Port     string
	Password string
}

// ExecuteCommand 执行单个命令
func ExecuteCommand(serverInfo ServerInfo, command string) (resOut string, errOut string, err error) {
	// 配置 SSH 客户端
	config := &ssh.ClientConfig{
		User: "root", // 替换为你的用户名
		Auth: []ssh.AuthMethod{
			ssh.Password(serverInfo.Password), // 替换为你的密码
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 仅用于测试，生产环境中应使用更安全的方式
	}
	// 连接到 SSH 服务器
	server := fmt.Sprintf("%s:22", serverInfo.Host)
	client, err := ssh.Dial("tcp", server, config) // 替换为你的服务器 IP 地址和端口
	if err != nil {
		return
	}
	defer client.Close()
	// 为每个命令创建一个新的会话
	session, err := client.NewSession()
	if err != nil {
		return
	}
	// 使用 bytes.Buffer 捕获标准输出
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	// 使用 bytes.Buffer 捕获标准输出
	var stderrBuf bytes.Buffer
	session.Stderr = &stderrBuf
	defer session.Close()

	if err = session.Run(command); err != nil {
		return
	}
	resOut = stdoutBuf.String()
	errOut = stderrBuf.String()
	return
}

// ExecuteCommands 执行多个命令
func ExecuteCommands(serverInfo ServerInfo, commands []string, hooks ...HookFunc) (resOut []string, errOut []string, err error) {
	// 配置 SSH 客户端
	config := &ssh.ClientConfig{
		User: "root", // 替换为你的用户名
		Auth: []ssh.AuthMethod{
			ssh.Password(serverInfo.Password), // 替换为你的密码
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 仅用于测试，生产环境中应使用更安全的方式
	}
	// 连接到 SSH 服务器
	server := fmt.Sprintf("%s:22", serverInfo.Host)
	client, err := ssh.Dial("tcp", server, config) // 替换为你的服务器 IP 地址和端口
	if err != nil {
		return
	}
	defer client.Close()

	hook := func(key int, command string) {}
	if len(hooks) > 0 {
		hook = hooks[0]
	}

	for key, cmd := range commands {
		hook(key, cmd)
		// 为每个命令创建一个新的会话
		session, err := client.NewSession()
		if err != nil {
			return nil, nil, err
		}
		// 使用 bytes.Buffer 捕获标准输出
		var stdoutBuf bytes.Buffer
		session.Stdout = &stdoutBuf
		// 使用 bytes.Buffer 捕获标准输出
		var stderrBuf bytes.Buffer
		session.Stderr = &stderrBuf
		// 执行命令
		if err = session.Run(cmd); err != nil {
			return nil, nil, err
		}

		resOut = append(resOut, stdoutBuf.String())
		errOut = append(errOut, stderrBuf.String())
		session.Close()
	}
	return
}
