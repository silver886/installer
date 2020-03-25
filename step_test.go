package installer

import (
	"errors"
	"math"
	"sync"
	"testing"
)

func TestStepNew(t *testing.T) {
	t.Log("Create a step.")
	var normalTest = []struct {
		a func() error
		b func() error
	}{
		{
			func() error { return nil },
			func() error { return nil },
		},
		{
			func() error { return nil },
			func() error { return errors.New("") },
		},
		{
			func() error { return errors.New("") },
			func() error { return nil },
		},
		{
			func() error { return errors.New("") },
			func() error { return errors.New("") },
		},
		{
			nil,
			func() error { return nil },
		},
		{
			nil,
			func() error { return errors.New("") },
		},
		{
			func() error { return nil },
			nil,
		},
		{
			func() error { return errors.New("") },
			nil,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			if s := NewStep(tt.a, tt.b); s == nil {
				t.Error("Step should be able to create.")
			}
		})
	}
}

func TestStepReset(t *testing.T) {
	var test = []struct {
		doer   func() error
		result error
	}{
		{
			doer:   func() error { return nil },
			result: nil,
		},
		{
			doer:   func() error { return errors.New("") },
			result: errors.New(""),
		},
	}

	t.Log("Redo a step.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				doer:  tt.doer,
			}
			s.Do()
			s.Reset()
			if err := s.Do(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Step should be able to redo.")
			}
		})
	}
}

func TestStepDo(t *testing.T) {
	var test = []struct {
		doer   func() error
		result error
	}{
		{
			doer:   func() error { return nil },
			result: nil,
		},
		{
			doer:   func() error { return errors.New("") },
			result: errors.New(""),
		},
	}

	t.Log("Normally do a step.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				doer:  tt.doer,
			}
			if err := s.Do(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Step should be able to do.")
			} else if err != nil && s.err.Error() != tt.result.Error() {
				t.Error("Step should be have a same error internally.")
			}
		})
	}

	t.Log("Do an emtpy step.")
	t.Run("Emtpy", func(t *testing.T) {
		s := &Step{
			mutex: &sync.Mutex{},
		}
		if err := s.Do(); err != ErrStepNoDoer {
			t.Error("Step should not be able to do.")
		}
	})

	t.Log("Do a executed step.")
	for _, tt := range test {
		t.Run("Executed", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				doer:  tt.doer,
				step:  1,
			}
			if err := s.Do(); err != ErrStepExecuted {
				t.Error("Step should not be able to do.")
			}
		})
	}
}

func TestStepUndo(t *testing.T) {
	var test = []struct {
		undoer func() error
		result error
	}{
		{
			undoer: func() error { return nil },
			result: nil,
		},
		{
			undoer: func() error { return errors.New("") },
			result: errors.New(""),
		},
	}

	t.Log("Normally undo a step.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex:  &sync.Mutex{},
				undoer: tt.undoer,
			}
			if err := s.Undo(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Step should be able to undo.")
			} else if err != nil && s.err.Error() != tt.result.Error() {
				t.Error("Step should be have a same error internally.")
			}
		})
	}

	t.Log("Undo an emtpy step.")
	t.Run("Emtpy", func(t *testing.T) {
		s := &Step{
			mutex: &sync.Mutex{},
		}
		if err := s.Undo(); err != ErrStepNoUndoer {
			t.Error("Step should not be able to undo.")
		}
	})

	t.Log("Undo an executed step.")
	for _, tt := range test {
		t.Run("Executed", func(t *testing.T) {
			s := &Step{
				mutex:  &sync.Mutex{},
				undoer: tt.undoer,
				step:   -1,
			}
			if err := s.Undo(); err != ErrStepExecuted {
				t.Error("Step should not be able to undo.")
			}
		})
	}
}

func TestStepFin(t *testing.T) {
	t.Log("Get fin status.")
	var normalTest = []int{
		0,
		1,
		-1,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				step:  tt,
			}
			if s.Fin() != (math.Abs(float64(tt)) == 1) {
				t.Error("Fin status of step should be the same.")
			}
		})
	}
}

func TestStepStep(t *testing.T) {
	t.Log("Get step status.")
	var normalTest = []int{
		0,
		1,
		-1,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				step:  tt,
			}
			if s.Step() != int(math.Abs(float64(tt))) {
				t.Error("Step status should be the same.")
			}
			if s.Step() < 0 {
				t.Error("Step should be non-negative.")
			}
		})
	}
}

func TestStepProgress(t *testing.T) {
	t.Log("Get progress status.")
	var normalTest = []int{
		0,
		1,
		-1,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				step:  tt,
			}
			if s.Progress() != math.Abs(float64(tt)) {
				t.Error("Progress status of step should be the same.")
			}
			if s.Progress() < 0 || s.Progress() > 1 {
				t.Error("Progress exceeds the range of 0~1.")
			}
		})
	}
}

func TestStepError(t *testing.T) {
	var test = []struct {
		a func() error
		b func() error
	}{
		{
			a: func() error { return nil },
			b: func() error { return nil },
		},
		{
			a: func() error { return nil },
			b: func() error { return errors.New("") },
		},
		{
			a: func() error { return errors.New("") },
			b: func() error { return nil },
		},
		{
			a: func() error { return errors.New("") },
			b: func() error { return errors.New("") },
		},
	}

	t.Log("Get do error from an executed step.")
	for _, tt := range test {
		t.Run("Normal do", func(t *testing.T) {
			s := &Step{
				mutex:  &sync.Mutex{},
				doer:   tt.a,
				undoer: tt.b,
			}
			s.Do()
			if err := s.Error(); err != tt.a() && err.Error() != tt.a().Error() {
				t.Error("Do error should be able to get.")
			}
		})
	}

	t.Log("Get do error from a non-executed step.")
	for _, tt := range test {
		t.Run("Non-executed do", func(t *testing.T) {
			s := &Step{
				mutex:  &sync.Mutex{},
				doer:   tt.a,
				undoer: tt.b,
			}
			if err := s.Error(); err != ErrStepNotExecuted {
				t.Error("Do error should not be able to get.")
			}
		})
	}

	t.Log("Get undo error from an executed step.")
	for _, tt := range test {
		t.Run("Normal - undo", func(t *testing.T) {
			s := &Step{
				mutex:  &sync.Mutex{},
				doer:   tt.a,
				undoer: tt.b,
			}
			s.Undo()
			if err := s.Error(); err != tt.b() && err.Error() != tt.b().Error() {
				t.Error("Undo error should be able to get.")
			}
		})
	}

	t.Log("Get undo error from a non-executed step.")
	for _, tt := range test {
		t.Run("Non-executed - undo", func(t *testing.T) {
			s := &Step{
				mutex:  &sync.Mutex{},
				doer:   tt.a,
				undoer: tt.b,
			}
			if err := s.Error(); err != ErrStepNotExecuted {
				t.Error("Undo error should not be able to get.")
			}
		})
	}
}
