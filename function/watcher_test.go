/**
 * Created by Goland
 * @file   debug.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 11:40
 * @desc   watcher_test.go
 */

package function

import (
	"github.com/x-module/helper/internal"
	"testing"
)

func TestWatcher(t *testing.T) {
	assert := internal.NewAssert(t, "TestWatcher")

	w := NewWatcher()
	w.Start()

	longRunningTask()

	assert.Equal(true, w.excuting)

	w.Stop()

	usedTime := w.GetElapsedTime().Milliseconds()
	t.Log("Elapsed Time (milSecond)", usedTime)
	assert.Equal(false, w.excuting)
	w.Reset()

	assert.Equal(int64(0), w.startTime)
	assert.Equal(int64(0), w.stopTime)
}

func longRunningTask() []int64 {
	var data []int64
	for i := 0; i < 10000000; i++ {
		data = append(data, int64(i))
	}
	return data
}
