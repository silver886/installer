package installer

import (
	"errors"
	"math"
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

func TestStepsFin(t *testing.T) {
	t.Log("Get fin status.")
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
			steppers: []Stepper{nil, nil, nil, nil, nil},
			step:     -5,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
				step:     tt.step,
			}
			if s.Fin() != (int(math.Abs(float64(tt.step))) == len(tt.steppers)) {
				t.Error("Fin status of steps should be the same.")
			}
		})
	}
}

func TestStepsStep(t *testing.T) {
	t.Log("Get step status.")
	var normalTest = []int{
		0,
		3,
		5,
		-3,
		-5,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				mutex: &sync.Mutex{},
				step:  tt,
			}
			if s.Step() != int(math.Abs(float64(tt))) {
				t.Error("Done step status should be the same.")
			}
			if s.Step() < 0 {
				t.Error("Done step should be non-negative.")
			}
		})
	}
}

func TestStepsProgress(t *testing.T) {
	t.Log("Get progress status.")
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
			if s.checkSteppers() != nil && s.Progress() > 0 ||
				s.checkSteppers() == nil && s.Progress() != math.Abs(float64(tt.step))/float64(len(tt.steppers)) {
				t.Error("Progress status of steps should be the same.")
			}
			if s.Progress() < 0 || s.Progress() > 1 {
				t.Error("Progress exceeds the range of 0~1.")
			}
		})
	}
}

func TestStepsError(t *testing.T) {
	var test = []struct {
		steppers []Stepper
		doStep   int
		undoStep int
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
			doStep:   1,
			undoStep: 1,
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
			doStep:   1,
			undoStep: 0,
		},
	}

	t.Log("Get do error from an executed steps.")
	for _, tt := range test {
		t.Run("Normal do", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
			}
			s.Do()
			if err := s.Error(); err != tt.steppers[tt.doStep].Error() &&
				err.Error() != tt.steppers[tt.doStep].Error().Error() {
				t.Error("Do error should be able to get.")
			}
		})
	}

	t.Log("Get do error from a non-executed step.")
	for _, tt := range test {
		t.Run("Non-executed do", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
			}
			if err := s.Error(); err != ErrStepsNotExecuted {
				t.Error("Do error should not be able to get.")
			}
		})
	}
	t.Log("Get undo error from an executed steps.")
	for _, tt := range test {
		t.Run("Normal undo", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
			}
			s.Undo()
			if err := s.Error(); err != tt.steppers[tt.undoStep].Error() &&
				err.Error() != tt.steppers[tt.undoStep].Error().Error() {
				t.Error("Undo error should be able to get.")
			}
		})
	}

	t.Log("Get undo error from a non-executed step.")
	for _, tt := range test {
		t.Run("Non-executed undo", func(t *testing.T) {
			s := &Steps{
				mutex:    &sync.Mutex{},
				steppers: tt.steppers,
			}
			if err := s.Error(); err != ErrStepsNotExecuted {
				t.Error("Undo error should not be able to get.")
			}
		})
	}
}
