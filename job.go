package cron

import (
	"time"

	"github.com/hatchify/atoms"
)

// New creates a new job
func New(fn func()) (jp *Job) {
	var j Job
	// Callback returns false if job was canceled
	j.callback = func() bool {
		// Ignore firing job if canceled
		if !j.Canceled.Get() {
			fn()
			return true
		}
		return false
	}
	return &j
}

// Job represents a job entry, also allows cancelation
type Job struct {
	// Function to be called when job is ran
	callback func() bool

	// Prevents execution if set true before target time
	Canceled atoms.Bool
}

// runAfter will run a function after waiting for a given duration
func (j *Job) runAfter(duration time.Duration) bool {
	time.Sleep(duration)
	return j.callback()
}

// runEvery will run a function continuously with the given duration as a delay
func (j *Job) runEvery(duration time.Duration) {
	for j.runAfter(duration) {
	}
}

// runAt will run a function after waiting until a target time
func (j *Job) runAt(target time.Time) bool {
	return j.runAfter(getDelay(target))
}

// runEveryAt will run a function continuously with the target time every day
func (j *Job) runEveryAt(target time.Time) {
	for j.runAt(target) {
	}
}

// After will run a function after waiting for a given duration
func (j *Job) After(duration time.Duration) *Job {
	go j.runAfter(duration)
	return j
}

// Every will run a function continuously with the given duration as a delay
func (j *Job) Every(duration time.Duration) *Job {
	go j.runEvery(duration)
	return j
}

// At will run a function after waiting until a target time
func (j *Job) At(target time.Time) *Job {
	go j.runAt(target)
	return j
}

// EveryAt will run a function continuously with the target time every day
func (j *Job) EveryAt(target time.Time) *Job {
	go j.runEveryAt(target)
	return j
}
