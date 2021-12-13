package week5task

import (
	"sync"
	"time"
)

// Window 滑动窗口接受体
type Window struct {
	Buckets map[int64]*BucketCount
	Mutex   *sync.RWMutex
}

// BucketCount 桶计数
type BucketCount struct {
	Value float64
}

// windowSize 滑动窗口大小
const windowSize = 10

// NewWindow 初始化滑动窗口
func NewWindow() *Window {
	w := &Window{
		Buckets: make(map[int64]*BucketCount),
		Mutex:   &sync.RWMutex{},
	}
	return w
}

// Increment 增加滑动窗口某一秒计数
func (w *Window) Increment(i float64) {
	if i == 0 {
		return
	}

	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	b := w.getCurrentBucket()
	b.Value += 1
	w.removeOldBucket()
}

// getCurrentBucket 获取当前时间桶中的计数
func (w *Window) getCurrentBucket() *BucketCount {
	now := time.Now().Unix()
	if bucket, ok := w.Buckets[now]; ok {
		return bucket
	}
	var bucket *BucketCount
	w.Buckets[now] = bucket
	return bucket
}

// removeOldBucket 删除过期的滑动窗口桶
func (w *Window) removeOldBucket() {
	now := time.Now().Unix() - windowSize
	for timestamp := range w.Buckets {
		if timestamp <= now {
			delete(w.Buckets, timestamp)
		}
	}
}
