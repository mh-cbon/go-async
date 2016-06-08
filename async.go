// Package async help to execute function in parallel.
package async

import (
	"errors"
	// "fmt"
)

// func indicating task completion
type Done func(error)

// func type of an async task
type Task func(Done)

// Provide information about
// completion status of a task.
// Err is the value returned when Done was invoked.
// If Done is invoked multiple times, it is ignored.
// Index is the index of the task.
type Result struct {
	done  bool
	Err   error
	Index int
}

// Run multiple task with a limit of concurrent execution.
// []Result length is always equal to the number of task to execute.
// []Result returns completion information for each Task.
// if limit < 1, it returns an error.
func ParallelLimit(limit int, handlers []Task) ([]Result, error) {
	results := make([]Result, len(handlers))
	if limit < 1 {
		return results, errors.New("Limit must be greater than 0")
	}
	if len(handlers) < 1 {
		return results, nil
	}
	out := make(chan Result)
	in := make(chan Task, limit)
	for index, t := range handlers {
		go func(index int, t Task) {
			in <- t
			t(func(err error) {
				out <- Result{Index: index, Err: err, done: true}
				<-in
			})
		}(index, t)
	}
	for res := range out {
    if results[res.Index].done == false {
  		results[res.Index] = res
    }
		if hasEnded(results) {
			close(out)
			close(in)
		}
	}
	return results, nil
}

// Run all provided tasks in a parallel fashion.
func Parallel(handlers []Task) []Result {
	results := make([]Result, len(handlers))
	if len(handlers) < 1 {
		return results
	}
	out := make(chan Result)
	for index, t := range handlers {
		go func(index int, t Task) {
			t(func(err error) {
				out <- Result{Index: index, Err: err, done: true}
			})
		}(index, t)
	}
	for res := range out {
    if results[res.Index].done == false {
  		results[res.Index] = res
    }
		if hasEnded(results) {
			close(out)
		}
	}
	return results
}
func hasEnded(results []Result) bool {
	ended := true
	for _, res := range results {
		ended = ended && res.done
	}
	return ended
}

// func helper to tell if any Result failed with an error.
func HasErrors(results []Result) bool {
	h := false
	for _, res := range results {
		h = h || res.Err != nil
	}
	return h
}
