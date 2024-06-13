/**
 * Created by Goland
 * @file   net.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 14:07
 * @desc   net.go
 */

package cmd

import "os"
import "os/exec"
import "bytes"
import "io"

// Run 运行shell命令，返回标准输出、标准错误的内容。
func Run(name string, arg ...string) (sout, serr []byte, err error) {
	c := exec.Command(name, arg...)
	var so bytes.Buffer
	var se bytes.Buffer
	c.Stdout = &so
	c.Stderr = &se
	err = c.Run()
	sout = so.Bytes()
	serr = se.Bytes()
	return
}

// Call 运行shell命令并将结果打印到stdout和stderr。
func Call(name string, arg ...string) error {
	sout, serr, err := Run(name, arg...)
	io.Copy(os.Stdout, bytes.NewBuffer(sout))
	io.Copy(os.Stderr, bytes.NewBuffer(serr))
	return err
}
