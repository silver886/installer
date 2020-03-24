package installer

import "sync"

// Stepper implements methods that would used by installer steps.
type Stepper interface {
	Do() error
	DoError() error
	Done() bool
	DoneStep() int
	DoneProgress() float64

	Undo() error
	UndoError() error
	Undone() bool
	UndoneStep() int
	UndoneProgress() float64

	Reset()
}

// Steps is the set of steppers.
type Steps struct {
	mutex *sync.Mutex

	steppers []Stepper

	done     bool
	doneStep int

	undone     bool
	undoneStep int
}

// NewSteps creates a set of steppers with given steppers.
func NewSteps(steppers []Stepper) *Steps {
	return &Steps{
		mutex:    &sync.Mutex{},
		steppers: steppers,
	}
}

// Reset clears the status.
func (s *Steps) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, ss := range s.steppers {
		ss.Reset()
	}
	s.done = false
	s.doneStep = 0
	s.undone = false
	s.undoneStep = 0
}

// Do triggers each steppers' doer.
func (s *Steps) Do() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.checkSteppers(); err != nil {
		return err
	}
	if s.done {
		return ErrStepsDone
	}
	s.done = true
	for i, ss := range s.steppers {
		s.doneStep = i
		if err := ss.Do(); err != nil {
			return err
		}
	}
	return nil
}

// Undo triggers each steppers' undoer.
func (s *Steps) Undo() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.checkSteppers(); err != nil {
		return err
	}
	if s.undone {
		return ErrStepsUndone
	}
	s.undone = true
	for i, ss := range s.steppers {
		s.undoneStep = i
		if err := ss.Undo(); err != nil {
			return err
		}
	}
	return nil
}

// Done retuen the status of doer.
func (s *Steps) Done() bool {
	return s.done
}

// Undone retuen the status of undoer.
func (s *Steps) Undone() bool {
	return s.undone
}

// DoneStep retuen the step status of doer.
func (s *Steps) DoneStep() int {
	if s.done {
		return s.doneStep
	}
	return 0
}

// UndoneStep retuen the step status of undoer.
func (s *Steps) UndoneStep() int {
	if s.undone {
		return s.undoneStep
	}
	return 0
}

// DoneProgress retuen the progress status of doer.
func (s *Steps) DoneProgress() float64 {
	if err := s.checkSteppers(); err != nil || !s.done {
		return 0
	}
	return float64(s.doneStep) / float64(len(s.steppers))
}

// UndoneProgress retuen the progress status of undoer.
func (s *Steps) UndoneProgress() float64 {
	if err := s.checkSteppers(); err != nil || !s.undone {
		return 0
	}
	return float64(s.undoneStep) / float64(len(s.steppers))
}

// DoError retuen the error during doing steppers.
func (s *Steps) DoError() error {
	if !s.done {
		return ErrStepsNonDone
	}
	return s.steppers[s.doneStep].DoError()
}

// UndoError retuen the error during undoing steppers.
func (s *Steps) UndoError() error {
	if !s.undone {
		return ErrStepsNonUndone
	}
	return s.steppers[s.undoneStep].UndoError()
}

func (s *Steps) checkSteppers() error {
	if s.steppers == nil || len(s.steppers) == 0 {
		return ErrStepsNoStepper
	}
	return nil
}
