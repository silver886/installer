package installer

import (
	"errors"
	"sync"
	"testing"
)

func TestStepsNew(t *testing.T) {
	t.Log("Normally create a steps.")
	var normalTest = [][]Stepper{
		{
			NewStep(
				func() error { return nil },
				func() error { return nil },
			),
			NewStep(
				func() error { return nil },
				func() error { return nil },
			),
		},
		{
			NewStep(
				func() error { return errors.New("") },
				func() error { return nil },
			),
			NewStep(
				func() error { return nil },
				func() error { return errors.New("") },
			),
		},
		{
			NewStep(
				func() error { return errors.New("") },
				func() error { return errors.New("") },
			),
			NewStep(
				func() error { return errors.New("") },
				func() error { return errors.New("") },
			),
		},
		{},
		nil,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			if s := NewSteps(tt); s == nil {
				t.Error("Steps should be able to create.")
			}
		})
	}
}

func TestStepsReset(t *testing.T) {
	var test = []struct {
		steppers []Stepper
		result   error
	}{
		{
			steppers: []Stepper{
				NewStep(
					func() error { return nil },
					func() error { return nil },
				),
				NewStep(
					func() error { return nil },
					func() error { return nil },
				),
			},
			result: nil,
		},
		{
			steppers: []Stepper{
				NewStep(
					func() error { return nil },
					func() error { return errors.New("") },
				),
				NewStep(
					func() error { return errors.New("") },
					func() error { return nil },
				),
			},
			result: errors.New(""),
		},
	}

	t.Log("Normally do a steps.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
			}
			s.Do()
			s.Reset()
			if err := s.Do(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Steps should be able to do.")
			}
		})
	}
}

func TestStepsDo(t *testing.T) {
	var test = []struct {
		steppers []Stepper
		result   error
	}{
		{
			steppers: []Stepper{
				NewStep(
					func() error { return nil },
					func() error { return nil },
				),
				NewStep(
					func() error { return nil },
					func() error { return nil },
				),
			},
			result: nil,
		},
		{
			steppers: []Stepper{
				NewStep(
					func() error { return nil },
					func() error { return errors.New("") },
				),
				NewStep(
					func() error { return errors.New("") },
					func() error { return nil },
				),
			},
			result: errors.New(""),
		},
	}

	t.Log("Normally do a steps.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
			}
			if err := s.Do(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Steps should be able to do.")
			}
		})
	}

	t.Log("Do an emtpy steps.")
	t.Run("Emtpy", func(t *testing.T) {
		s := &Steps{
			mutex: &sync.Mutex{},
		}
		if err := s.Do(); err != ErrStepsNoStepper {
			t.Error("Steps should not be able to do.")
		}
	})

	t.Log("Do a done step.")
	for _, tt := range test {
		t.Run("Done", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
				step:     1,
			}
			if err := s.Do(); err != ErrStepsExecuted {
				t.Error("Steps should not be able to do.")
			}
		})
	}
}

func TestStepsUndo(t *testing.T) {
	var test = []struct {
		steppers []Stepper
		result   error
	}{
		{
			steppers: []Stepper{
				NewStep(
					func() error { return nil },
					func() error { return nil },
				),
				NewStep(
					func() error { return nil },
					func() error { return nil },
				),
			},
			result: nil,
		},
		{
			steppers: []Stepper{
				NewStep(
					func() error { return nil },
					func() error { return errors.New("") },
				),
				NewStep(
					func() error { return errors.New("") },
					func() error { return nil },
				),
			},
			result: errors.New(""),
		},
	}

	t.Log("Normally undo a steps.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
			}
			if err := s.Undo(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Steps should be able to undo.")
			}
		})
	}

	t.Log("Undo an emtpy steps.")
	t.Run("Emtpy", func(t *testing.T) {
		s := &Steps{
			mutex: &sync.Mutex{},
		}
		if err := s.Undo(); err != ErrStepsNoStepper {
			t.Error("Steps should not be able to undo.")
		}
	})

	t.Log("Undo a done step.")
	for _, tt := range test {
		t.Run("Undone", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
				step:     -1,
			}
			if err := s.Undo(); err != ErrStepsExecuted {
				t.Error("Steps should not be able to undo.")
			}
		})
	}
}

func TestStepsDone(t *testing.T) {
	t.Log("Get done status.")
	var normalTest = []struct {
		steppers []Stepper
		step     int
	}{
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     0,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     3,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     5,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     -3,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
				step:     tt.step,
			}
			if s.Done() != (tt.step == len(tt.steppers)) {
				t.Error("Done status of steps should be the same.")
			}
		})
	}
}

func TestStepsUndone(t *testing.T) {
	t.Log("Get undone status.")
	var normalTest = []struct {
		steppers []Stepper
		step     int
	}{
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     0,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     -3,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     -5,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     3,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
				step:     tt.step,
			}
			if s.Undone() != (tt.step == -len(tt.steppers)) {
				t.Error("Undone status of steps should be the same.")
			}
		})
	}
}

func TestStepsDoneStep(t *testing.T) {
	t.Log("Get done step status.")
	var normalTest = []struct {
		steppers []Stepper
		step     int
	}{
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     0,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     3,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     5,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     -3,
		},
		{
			steppers: []Stepper{},
			step:     3,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
				step:     tt.step,
			}
			if (tt.step <= 0 || s.checkSteppers() != nil) && s.DoneStep() > 0 ||
				tt.step > 0 && s.checkSteppers() == nil && s.DoneStep() != tt.step {
				t.Error("Done step status should be the same.")
			}
			if s.DoneStep() < 0 {
				t.Error("Done step should be non-negative.")
			}
		})
	}
}

func TestStepsUndoneStep(t *testing.T) {
	t.Log("Get undone step status.")
	var normalTest = []struct {
		steppers []Stepper
		step     int
	}{
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     0,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     -3,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     -5,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     3,
		},
		{
			steppers: []Stepper{},
			step:     -3,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
				step:     tt.step,
			}
			if (tt.step >= 0 || s.checkSteppers() != nil) && s.UndoneStep() > 0 ||
				tt.step < 0 && s.checkSteppers() == nil && s.UndoneStep() != -tt.step {
				t.Error("Undone step status should be the same.")
			}
			if s.UndoneStep() < 0 {
				t.Error("Undone step should be non-negative.")
			}
		})
	}
}

func TestStepsDoneProgress(t *testing.T) {
	t.Log("Get done progress status.")
	var normalTest = []struct {
		steppers []Stepper
		step     int
	}{
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     0,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     3,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     5,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     -3,
		},
		{
			steppers: []Stepper{},
			step:     3,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
				step:     tt.step,
			}
			if (tt.step <= 0 || s.checkSteppers() != nil) && s.DoneProgress() > 0 ||
				tt.step > 0 && s.checkSteppers() == nil && s.DoneProgress() != float64(tt.step)/float64(len(tt.steppers)) {
				t.Error("Done progress status of steps should be the same.")
			}
			if s.DoneProgress() < 0 || s.DoneProgress() > 1 {
				t.Error("Done progress exceeds the range of 0~1.")
			}
		})
	}
}

func TestStepsUndoneProgress(t *testing.T) {
	t.Log("Get undone progress status.")
	var normalTest = []struct {
		steppers []Stepper
		step     int
	}{
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     0,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     -3,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     -5,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     3,
		},
		{
			steppers: []Stepper{},
			step:     -3,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
				step:     tt.step,
			}
			if (tt.step >= 0 || s.checkSteppers() != nil) && s.UndoneProgress() > 0 ||
				tt.step < 0 && s.checkSteppers() == nil && s.UndoneProgress() != float64(-tt.step)/float64(len(tt.steppers)) {
				t.Error("Undone progress status of steps should be the same.")
			}
			if s.UndoneProgress() < 0 || s.UndoneProgress() > 1 {
				t.Error("Undone progress exceeds the range of 0~1.")
			}
		})
	}
}

func TestStepsDoError(t *testing.T) {
	var test = []struct {
		steppers []Stepper
		step     int
		result   error
	}{
		{
			steppers: []Stepper{
				NewStep(
					func() error { return nil },
					func() error { return nil },
				),
				NewStep(
					func() error { return nil },
					func() error { return nil },
				),
			},
			step:   0,
			result: nil,
		},
		{
			steppers: []Stepper{
				NewStep(
					func() error { return nil },
					func() error { return errors.New("") },
				),
				NewStep(
					func() error { return errors.New("") },
					func() error { return nil },
				),
			},
			step:   1,
			result: errors.New(""),
		},
	}

	t.Log("Get do error from an done steps.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
			}
			s.Do()
			if err := s.DoError(); err != tt.steppers[tt.step].DoError() &&
				err.Error() != tt.steppers[tt.step].DoError().Error() {
				t.Error("Do error should be able to get.")
			}
		})
	}

	t.Log("Get do error from a non-done step.")
	for _, tt := range test {
		t.Run("Non-done", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
			}
			if err := s.DoError(); err != ErrStepsNonDone {
				t.Error("Do error should not be able to get.")
			}
		})
	}
}

func TestStepsUndoError(t *testing.T) {
	var test = []struct {
		steppers []Stepper
		step     int
		result   error
	}{
		{
			steppers: []Stepper{
				NewStep(
					func() error { return nil },
					func() error { return nil },
				),
				NewStep(
					func() error { return nil },
					func() error { return nil },
				),
			},
			step:   0,
			result: nil,
		},
		{
			steppers: []Stepper{
				NewStep(
					func() error { return nil },
					func() error { return errors.New("") },
				),
				NewStep(
					func() error { return errors.New("") },
					func() error { return nil },
				),
			},
			step:   0,
			result: errors.New(""),
		},
	}

	t.Log("Get undo error from an done steps.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
			}
			s.Undo()
			if err := s.UndoError(); err != tt.steppers[tt.step].UndoError() &&
				err.Error() != tt.steppers[tt.step].UndoError().Error() {
				t.Error("Undo error should be able to get.")
			}
		})
	}

	t.Log("Get duno error from a non-done step.")
	for _, tt := range test {
		t.Run("Non-undone", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
			}
			if err := s.UndoError(); err != ErrStepsNonUndone {
				t.Error("Undo error should not be able to get.")
			}
		})
	}
}
