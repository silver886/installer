package installer

import (
	"errors"
	"testing"
)

func TestStepssNew(t *testing.T) {
	t.Log("Normally create a steps.")
	var normalTest = [][]Stepper{
		{
			func() *Step {
				s, _ := NewStep(
					func() error { return nil },
					func() error { return nil },
				)
				return s
			}(),
			func() *Step {
				s, _ := NewStep(
					func() error { return nil },
					func() error { return nil },
				)
				return s
			}(),
		},
		{
			func() *Step {
				s, _ := NewStep(
					func() error { return errors.New("") },
					func() error { return nil },
				)
				return s
			}(),
			func() *Step {
				s, _ := NewStep(
					func() error { return nil },
					func() error { return errors.New("") },
				)
				return s
			}(),
		},
		{
			func() *Step {
				s, _ := NewStep(
					func() error { return errors.New("") },
					func() error { return errors.New("") },
				)
				return s
			}(),
			func() *Step {
				s, _ := NewStep(
					func() error { return errors.New("") },
					func() error { return errors.New("") },
				)
				return s
			}(),
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			if s, err := NewSteps(tt); s == nil || s.steppers == nil || len(s.steppers) == 0 || err != nil {
				t.Error("Steps should be able to create.")
			}
		})
	}

	t.Log("Create a steps without steppers.")
	var missingTest = [][]Stepper{
		{},
		nil,
	}
	for _, tt := range missingTest {
		t.Run("Missing", func(t *testing.T) {
			if s, err := NewSteps(tt); s != nil || err != ErrStepsNoStepper {
				t.Error("Steps should not be able to create.")
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
				func() *Step {
					s, _ := NewStep(
						func() error { return nil },
						func() error { return nil },
					)
					return s
				}(),
				func() *Step {
					s, _ := NewStep(
						func() error { return nil },
						func() error { return nil },
					)
					return s
				}(),
			},
			result: nil,
		},
		{
			steppers: []Stepper{
				func() *Step {
					s, _ := NewStep(
						func() error { return nil },
						func() error { return errors.New("") },
					)
					return s
				}(),
				func() *Step {
					s, _ := NewStep(
						func() error { return errors.New("") },
						func() error { return nil },
					)
					return s
				}(),
			},
			result: errors.New(""),
		},
	}

	t.Log("Normally do a steps.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				steppers: tt.steppers,
			}
			if err := s.Do(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Steps should be able to do.")
			}
		})
	}

	t.Log("Do an emtpy steps.")
	t.Run("Emtpy", func(t *testing.T) {
		s := &Steps{}
		if err := s.Do(); err != ErrStepsNoStepper {
			t.Error("Steps should not be able to do.")
		}
	})

	t.Log("Do a done step.")
	for _, tt := range test {
		t.Run("Done", func(t *testing.T) {
			s := &Steps{
				steppers: tt.steppers,
				done:     true,
			}
			if err := s.Do(); err != ErrStepsDone {
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
				func() *Step {
					s, _ := NewStep(
						func() error { return nil },
						func() error { return nil },
					)
					return s
				}(),
				func() *Step {
					s, _ := NewStep(
						func() error { return nil },
						func() error { return nil },
					)
					return s
				}(),
			},
			result: nil,
		},
		{
			steppers: []Stepper{
				func() *Step {
					s, _ := NewStep(
						func() error { return nil },
						func() error { return errors.New("") },
					)
					return s
				}(),
				func() *Step {
					s, _ := NewStep(
						func() error { return errors.New("") },
						func() error { return nil },
					)
					return s
				}(),
			},
			result: errors.New(""),
		},
	}

	t.Log("Normally undo a steps.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				steppers: tt.steppers,
			}
			if err := s.Undo(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Steps should be able to undo.")
			}
		})
	}

	t.Log("Undo an emtpy steps.")
	t.Run("Emtpy", func(t *testing.T) {
		s := &Steps{}
		if err := s.Undo(); err != ErrStepsNoStepper {
			t.Error("Steps should not be able to undo.")
		}
	})

	t.Log("Undo a done step.")
	for _, tt := range test {
		t.Run("Undone", func(t *testing.T) {
			s := &Steps{
				steppers: tt.steppers,
				undone:   true,
			}
			if err := s.Undo(); err != ErrStepsUndone {
				t.Error("Steps should not be able to undo.")
			}
		})
	}
}

func TestStepsDone(t *testing.T) {
	t.Log("Get done status.")
	var normalTest = []bool{
		true,
		false,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				done: tt,
			}
			if s.Done() != tt {
				t.Error("Done status of steps should be the same.")
			}
		})
	}
}

func TestStepsUndone(t *testing.T) {
	t.Log("Get undone status.")
	var normalTest = []bool{
		true,
		false,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				undone: tt,
			}
			if s.Undone() != tt {
				t.Error("Undone status of steps should be the same.")
			}
		})
	}
}

func TestStepsDoneStep(t *testing.T) {
	t.Log("Get done step status.")
	var normalTest = []struct {
		done bool
		step int
	}{
		{
			done: true,
			step: 3,
		},
		{
			done: false,
			step: 3,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				done:     tt.done,
				doneStep: tt.step,
			}
			if !tt.done && s.DoneStep() != 0 ||
				tt.done && s.DoneStep() != tt.step {
				t.Error("Done step status should be the same.")
			}
		})
	}
}

func TestStepsUndoneStep(t *testing.T) {
	t.Log("Get undone step status.")
	var normalTest = []struct {
		undone bool
		step   int
	}{
		{
			undone: true,
			step:   3,
		},
		{
			undone: true,
			step:   5,
		},
		{
			undone: false,
			step:   3,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				undone:     tt.undone,
				undoneStep: tt.step,
			}
			if !tt.undone && s.UndoneStep() != 0 ||
				tt.undone && s.UndoneStep() != tt.step {
				t.Error("Undone step status should be the same.")
			}
		})
	}
}

func TestStepsDoneProgress(t *testing.T) {
	t.Log("Get done progress status.")
	var normalTest = []struct {
		steppers []Stepper
		done     bool
		step     int
	}{
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			done:     true,
			step:     3,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			done:     true,
			step:     5,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			done:     false,
			step:     3,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				steppers: tt.steppers,
				done:     tt.done,
				doneStep: tt.step,
			}
			if !tt.done && s.DoneProgress() != 0 ||
				tt.done && s.DoneProgress() != float64(tt.step)/float64(len(tt.steppers)) {
				t.Error("Done progress status of steps should be the same.")
			}
		})
	}
}

func TestStepsUndoneProgress(t *testing.T) {
	t.Log("Get undone progress status.")
	var normalTest = []struct {
		steppers []Stepper
		undone   bool
		step     int
	}{
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			undone:   true,
			step:     3,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			undone:   true,
			step:     5,
		},
		{
			steppers: []Stepper{nil, nil, nil, nil, nil},
			undone:   false,
			step:     3,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				steppers:   tt.steppers,
				undone:     tt.undone,
				undoneStep: tt.step,
			}
			if !tt.undone && s.UndoneProgress() != 0 ||
				tt.undone && s.UndoneProgress() != float64(tt.step)/float64(len(tt.steppers)) {
				t.Error("Undone progress status of steps should be the same.")
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
				func() *Step {
					s, _ := NewStep(
						func() error { return nil },
						func() error { return nil },
					)
					return s
				}(),
				func() *Step {
					s, _ := NewStep(
						func() error { return nil },
						func() error { return nil },
					)
					return s
				}(),
			},
			step:   0,
			result: nil,
		},
		{
			steppers: []Stepper{
				func() *Step {
					s, _ := NewStep(
						func() error { return nil },
						func() error { return errors.New("") },
					)
					return s
				}(),
				func() *Step {
					s, _ := NewStep(
						func() error { return errors.New("") },
						func() error { return nil },
					)
					return s
				}(),
			},
			step:   1,
			result: errors.New(""),
		},
	}

	t.Log("Get do error from an done steps.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				steppers: tt.steppers,
				doneStep: tt.step,
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
				steppers: tt.steppers,
				doneStep: tt.step,
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
				func() *Step {
					s, _ := NewStep(
						func() error { return nil },
						func() error { return nil },
					)
					return s
				}(),
				func() *Step {
					s, _ := NewStep(
						func() error { return nil },
						func() error { return nil },
					)
					return s
				}(),
			},
			step:   0,
			result: nil,
		},
		{
			steppers: []Stepper{
				func() *Step {
					s, _ := NewStep(
						func() error { return nil },
						func() error { return errors.New("") },
					)
					return s
				}(),
				func() *Step {
					s, _ := NewStep(
						func() error { return errors.New("") },
						func() error { return nil },
					)
					return s
				}(),
			},
			step:   0,
			result: errors.New(""),
		},
	}

	t.Log("Get undo error from an done steps.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Steps{
				steppers:   tt.steppers,
				undoneStep: tt.step,
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
				steppers:   tt.steppers,
				undoneStep: tt.step,
			}
			if err := s.UndoError(); err != ErrStepsNonUndone {
				t.Error("Undo error should not be able to get.")
			}
		})
	}
}
