package dataWindows

import "time"

type TimedValue interface {
	Timestamp() time.Time
}

type Window struct {
	buffer []TimedValue
	start int
	end int
	Keep time.Duration
}

func NewWindow(keep time.Duration) *Window {
	return &Window{}
}

func (w *Window)trim() {

}

func (w *Window)Insert(x TimedValue) {
	w.trim()
	if w.Len() == len(w.buffer) {
		bigger := make([]TimedValue, 2*len(w.buffer))
		copy(bigger, w.buffer)
	}
}

func (w *Window)Len() int {
	n := len(w.buffer)
	return (((w.end - w.start) + n) % n) + 1
}

func (w *Window)Latest() TimedValue {
	if w.end == w.start {
		return nil
	}
	return w.buffer[w.end]
}
