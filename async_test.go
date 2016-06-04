package async

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	ran := 0
	results := Parallel([]Task{
		func(done Done) {
			ran += 1
			done(nil)
		},
		func(done Done) {
			ran += 1
			done(nil)
		},
	})

	if len(results) != 2 {
		t.Errorf("Expected results to have same length as tasks got%d, expected %d\n", len(results), 2)
	}
	if ran != 2 {
		t.Errorf("Expected to execute two func only, executed %d\n", ran)
	}
}

func TestParallelErrored(t *testing.T) {
	results := Parallel([]Task{
		func(done Done) {
			done(errors.New("42"))
		},
		func(done Done) {
			done(nil)
		},
	})

	if results[0].Err == nil {
		t.Errorf("Expected index 0 to have an error\n")
	}
}

func TestParallelInvokeCallbackTwice(t *testing.T) {
	results := Parallel([]Task{
		func(done Done) {
			done(nil)
		},
		func(done Done) {
			done(nil)
			done(nil)
		},
	})

	if results[1].Err != nil {
		t.Errorf("Expected error of task #1 to be nil\n")
	}
}

func TestParallelNoTasks(t *testing.T) {
	results := Parallel([]Task{})

	if len(results) != 0 {
		t.Errorf("Expected results to have same length as tasks got%d, expected %d\n", len(results), 0)
	}
}

func TestParallelLimit(t *testing.T) {
	ran := 0
	results, err := ParallelLimit(1, []Task{
		func(done Done) {
			ran += 1
			done(nil)
		},
		func(done Done) {
			ran += 1
			done(nil)
		},
	})

	if len(results) != 2 {
		t.Errorf("Expected results to have same length as tasks got%d, expected %d\n", len(results), 2)
	}
	if ran != 2 {
		t.Errorf("Expected to execute two func only, executed %d\n", ran)
	}
	if err != nil {
		t.Errorf("Expected err to be nil, got %q", err)
	}
}

func TestParallelLimitErrored(t *testing.T) {
	results, err := ParallelLimit(2, []Task{
		func(done Done) {
			done(errors.New("42"))
		},
		func(done Done) {
			done(nil)
		},
	})

	if results[0].Err == nil {
		t.Errorf("Expected index 0 to have an error\n")
	}
	if err != nil {
		t.Errorf("Expected err to be nil, got %q", err)
	}
}

func TestParallelLimitInvokeCallbackTwice(t *testing.T) {
	results, err := ParallelLimit(2, []Task{
		func(done Done) {
			done(nil)
		},
		func(done Done) {
			done(nil)
			done(nil)
		},
	})

	if results[1].Err != nil {
		t.Errorf("Expected error of task #1 to be nil\n")
	}
	if err != nil {
		t.Errorf("Expected err to be nil, got %q", err)
	}
}

func TestParallelLimitToZero(t *testing.T) {
	_, err := ParallelLimit(0, []Task{
		func(done Done) {
			done(nil)
		},
		func(done Done) {
			done(nil)
		},
	})

	if err == nil {
		t.Errorf("Expected err to be not nil")
	}
}

func TestParallelLimitNoTasks(t *testing.T) {
	results, err := ParallelLimit(1, []Task{})

	if len(results) != 0 {
		t.Errorf("Expected results to have same length as tasks got%d, expected %d\n", len(results), 0)
	}
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
}
