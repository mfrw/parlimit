// Package parlimit limits concurecy of goroutines, by creating
// a buffered channel with the desired concurency limit.
// Any goroutine that starts off, sends an empty struct to the
// channel, if the channel is full, the send blocks.

package palrimit

type limiter chan struct{}

// NewLimiter returns a new limiter which has the capacity to
// keep running 'size' go routines.
func NewLimiter(size int) limit {
	return make(chan struct{}, size)
}

// Start pushes a value to the bufferd channel. It blocks if
// the channel is full
func (l limiter) Start() {
	l <- struct{}{}
}

// Finish should mark the end of the goroutine. It gets a value from
// the channel and discards it.
func (l limiter) Finish() {
	<-l
}
