/**
 * Created by Goland
 * @file   ssh2.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/11/26 14:02
 * @desc   ssh2.go
 */

package ssh

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/x-module/helper/xlog"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

type HookFunc func(key int, command string)

type SSHClient struct {
	client *ssh.Client
}

type ServerInfo struct {
	Host     string
	Port     int
	Password string
	Username string
	Key      string
}

func NewSSHClient(serverInfo ServerInfo) (*SSHClient, error) {
	config := &ssh.ClientConfig{}
	if serverInfo.Key != "" {
		// 读取私钥
		key, err := os.ReadFile(serverInfo.Key)
		if err != nil {
			return nil, fmt.Errorf("unable to read private key: %v", err)
		}
		// 解析私钥
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, fmt.Errorf("unable to parse private key: %v", err)
		}
		config = &ssh.ClientConfig{
			User: serverInfo.Username,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	} else {
		config = &ssh.ClientConfig{
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
			return nil, err
		}
		defer client.Close()
	}
	host := fmt.Sprintf("%s:%d", serverInfo.Host, serverInfo.Port)
	xlog.Debugf("链接当前服务器: %s", host)
	// 连接到远程服务器
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", serverInfo.Host, serverInfo.Port), config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect: %v", err)
	}
	return &SSHClient{client: client}, nil
}

// ExecuteStreamCommand 执行流式命令并实时获取输出
func (s *SSHClient) ExecuteStreamCommand(command string, stdCallback func(string), errCallback func(string)) error {
	// 创建新会话
	session, err := s.client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	// 获取标准输出管道
	stdout, err := session.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}

	// 获取标准错误管道
	stderr, err := session.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}

	// 启动命令
	if err := session.Start(command); err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}

	// 处理标准输出
	go func() {
		reader := bufio.NewReader(stdout)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					log.Printf("Error reading stdout: %v", err)
				}
				return
			}
			stdCallback(line)
		}
	}()

	// 处理标准错误
	go func() {
		reader := bufio.NewReader(stderr)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					log.Printf("Error reading stderr: %v", err)
				}
				return
			}
			errCallback(line)
		}
	}()
	// 等待命令完成
	if err = session.Wait(); err != nil {
		return fmt.Errorf("command failed: %v", err)
	}
	return nil
}

// ExecuteCommand 执行单个命令
func (s *SSHClient) ExecuteCommand(command string) (resOut string, errOut string, err error) {
	session, err := s.client.NewSession()
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
func (s *SSHClient) ExecuteCommands(serverInfo ServerInfo, commands []string, hooks ...HookFunc) (resOut []string, errOut []string, err error) {
	hook := func(key int, command string) {}
	if len(hooks) > 0 {
		hook = hooks[0]
	}

	for key, cmd := range commands {
		hook(key, cmd)
		// 为每个命令创建一个新的会话
		session, err := s.client.NewSession()
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

// Close 关闭SSH连接
func (s *SSHClient) Close() error {
	return s.client.Close()
}
