package lyric

import (
	"sync"
	"time"
)

type Listener func(startTimeMs int64, content string, last bool, index int)

type LRCTimer struct {
	file      *LRCFile
	timer     chan time.Duration
	stop      chan struct{}
	listeners []Listener

	curIndex int
	indexL   sync.Mutex
}

func NewLRCTimer(file *LRCFile) *LRCTimer {
	return &LRCTimer{
		file:  file,
		timer: make(chan time.Duration),
	}
}

func (t *LRCTimer) Timer() chan<- time.Duration {
	return t.timer
}

func (t *LRCTimer) AddListener(l Listener) {
	t.listeners = append(t.listeners, l)
}

func (t *LRCTimer) Start() {
	fragments := t.file.fragments

	if len(fragments) < 1 {
		return
	}

	t.Rewind()
	current := fragments[0]
	t.stop = make(chan struct{})
	for {
		select {
		case duration := <-t.timer:
			if duration < time.Duration(fragments[t.curIndex].StartTimeMs)*time.Millisecond {
				continue
			}

			// Rewind后快速定位
			for t.curIndex < len(fragments)-1 && duration >= time.Duration(fragments[t.curIndex+1].StartTimeMs)*time.Millisecond {
				t.indexL.Lock()
				t.curIndex++
				t.indexL.Unlock()
			}

			last := t.curIndex >= len(fragments)-1

			for _, l := range t.listeners {
				go l(current.StartTimeMs, current.Content, last, t.curIndex)
			}

			if last {
				break
			}

			t.indexL.Lock()
			t.curIndex++
			t.indexL.Unlock()
			current = fragments[t.curIndex]
		case <-t.stop:
			return
		}

	}
}

func (t *LRCTimer) IsStarted() bool {
	return t.timer != nil
}

func (t *LRCTimer) Stop() {
	t.stop <- struct{}{}
	t.timer = nil
}

func (t *LRCTimer) Rewind() {
	t.indexL.Lock()
	defer t.indexL.Unlock()
	t.curIndex = 0
}

func (t *LRCTimer) GetLRCFragment(index int) *LRCFragment {
	if nil == t.file || index >= len(t.file.fragments) || index < 0 {
		return nil
	}

	return &t.file.fragments[index]
}

func (t *LRCTimer) IsEmpty() bool {
	return nil == t.file || len(t.file.fragments) == 0
}
