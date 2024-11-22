/**
 * Created by goland.
 * @file   log.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/1 19:10
 * @desc   log.go
 */

package xlog

import "testing"

func TestLog(t *testing.T) {
	log := InitLogger(LogConfig{
		LogPath:  "debug",
		LogModel: "debug",
		LogFile:  "./log",
	})
	log.WithField("debug", "debug").Debug("debug")
	log.Info("info")
	log.WithField("a", 34).Warn("warn")
	log.Error("error")
	log.Fatal("fatal")
}
