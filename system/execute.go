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
	"runtime"
	"strings"
	"time"
)

const (
	Gray = uint8(iota + 90)
	Red
	Green
	Yellow
	Blue
	Magenta
	EndColor = "\033[0m"
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

func (e *Execute) ColorLogS(format string, a ...interface{}) string {
	log := fmt.Sprintf(format, a...)

	var clog string

	if runtime.GOOS != "windows" {
		// Level.
		i := strings.Index(log, "]")
		if log[0] == '[' && i > -1 {
			clog += "[" + e.getColorLevel(log[1:i]) + "]"
		}

		log = log[i+1:]

		// Error.
		log = strings.Replace(log, "[ ", fmt.Sprintf("[\033[%dm", Red), -1)
		log = strings.Replace(log, " ]", EndColor+"]", -1)

		// Path.
		log = strings.Replace(log, "( ", fmt.Sprintf("(\033[%dm", Yellow), -1)
		log = strings.Replace(log, " )", EndColor+")", -1)

		// Highlights.
		log = strings.Replace(log, "# ", fmt.Sprintf("\033[%dm", Gray), -1)
		log = strings.Replace(log, " #", EndColor, -1)

	} else {
		// Level.
		i := strings.Index(log, "]")
		if log[0] == '[' && i > -1 {
			clog += "[" + log[1:i] + "]"
		}

		log = log[i+1:]

		// Error.
		log = strings.Replace(log, "[ ", "[", -1)
		log = strings.Replace(log, " ]", "]", -1)

		// Path.
		log = strings.Replace(log, "( ", "(", -1)
		log = strings.Replace(log, " )", ")", -1)

		// Highlights.
		log = strings.Replace(log, "# ", "", -1)
		log = strings.Replace(log, " #", "", -1)
	}
	return clog + log
}

// ColorLog prints colored log to stdout.
// See color rules in function 'ColorLogS'.
func (e *Execute) ColorLog(format string, a ...interface{}) {
	fmt.Print(e.ColorLogS(format, a...))
}

// getColorLevel returns colored level string by given level.
func (e *Execute) getColorLevel(level string) string {
	level = strings.ToUpper(level)
	switch level {
	case "TRAC":
		return fmt.Sprintf("\033[%dm%s\033[0m", Blue, level)
	case "ERRO":
		return fmt.Sprintf("\033[%dm%s\033[0m", Red, level)
	case "WARN":
		return fmt.Sprintf("\033[%dm%s\033[0m", Magenta, level)
	case "SUCC":
		return fmt.Sprintf("\033[%dm%s\033[0m", Green, level)
	default:
		return level
	}
}
