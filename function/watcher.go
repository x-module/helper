/**
 * Created by Goland
 * @file   debug.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 11:40
 * @desc   watcher.go
 */

package function

import "time"

type Watcher struct {
	startTime int64
	stopTime  int64
	excuting  bool
}

// NewWatcher Start the watch timer.
func NewWatcher() *Watcher {
	return &Watcher{}
}

// Start the watch timer.
func (w *Watcher) Start() {
	w.startTime = time.Now().UnixNano()
	w.excuting = true
}

// Stop the watch timer.
func (w *Watcher) Stop() {
	w.stopTime = time.Now().UnixNano()
	w.excuting = false
}

// GetElapsedTime get execute elapsed time.
func (w *Watcher) GetElapsedTime() time.Duration {
	if w.excuting {
		return time.Duration(time.Now().UnixNano() - w.startTime)
	}
	return time.Duration(w.stopTime - w.startTime)
}

// Reset the watch timer.
func (w *Watcher) Reset() {
	w.startTime = 0
	w.stopTime = 0
	w.excuting = false
}
