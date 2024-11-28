/**
 * Created by Goland
 * @file   log.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/11/26 14:13
 * @desc   log.go
 */

package others

import (
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Message   string
	Level     string
}

type LogContainer struct {
	logs    []LogEntry
	maxSize int
	mu      sync.RWMutex
}

func NewLogContainer(maxSize int) *LogContainer {
	return &LogContainer{
		logs:    make([]LogEntry, 0, maxSize),
		maxSize: maxSize,
	}
}

func (lc *LogContainer) Add(entry LogEntry) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if len(lc.logs) >= lc.maxSize {
		// 删除最早的日志
		lc.logs = lc.logs[1:]
	}
	lc.logs = append(lc.logs, entry)
}

func (lc *LogContainer) GetLogs() []LogEntry {
	lc.mu.RLock()
	defer lc.mu.RUnlock()
	// 返回副本
	logs := make([]LogEntry, len(lc.logs))
	copy(logs, lc.logs)
	return logs
}
