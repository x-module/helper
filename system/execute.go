/**
 * Created by Goland
 * @file   execute.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/11/27 16:23
 * @desc   execute.go
 */

package system

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type OutFunc func(string)
type Execute struct {
	stdOut OutFunc
	errOut OutFunc
}

func NewExecute() *Execute {
	return &Execute{}
}

func (e *Execute) SetStdOut(stdOut OutFunc) {
	e.stdOut = stdOut
}

func (e *Execute) SetErrOut(errOut OutFunc) {
	e.errOut = errOut
}

func (e *Execute) ExecuteCommand(ctx context.Context, command string, args ...string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, command, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error creating stdout pipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("error creating stderr pipe: %v", err)
	}

	if err = cmd.Start(); err != nil {
		return fmt.Errorf("error starting command: %v", err)
	}

	// 修改这里：使用 bufio.Scanner 读取输出并调用回调函数
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if n > 0 && e.stdOut != nil {
				e.stdOut(strings.Replace(string(buf[:n]), "\n", "", -1))
			}
			if err != nil {
				break
			}
		}
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stderr.Read(buf)
			if n > 0 && e.errOut != nil {
				e.stdOut(strings.Replace(string(buf[:n]), "\n", "", -1))
			}
			if err != nil {
				break
			}
		}
	}()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("error waiting for command: %v", err)
	}

	return nil
}
