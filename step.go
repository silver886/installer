package installer

import (
	"math"
	"sync"
)

// Step is the basic component of a doer.
type Step struct {
	mutex *sync.Mutex
	step  int
	err   error

	doer   func() error
	undoer func() error
}

// NewStep creates step with doer and undoer.
func NewStep(doer func() error, undoer func() error) *Step {
	return &Step{
		mutex:  &sync.Mutex{},
		doer:   doer,
		undoer: undoer,
	}
}

// Do triggers the doer.
func (s *Step) Do() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.doer == nil {
		return ErrStepNoDoer
	}
	if s.step != 0 {
		return ErrStepExecuted
	}
	s.err = s.doer()
	s.step++
	if s.err != nil {
		return s.err
	}
	return nil
}

// Undo triggers the undoer.
func (s *Step) Undo() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.undoer == nil {
		return ErrStepNoUndoer
	}
	if s.step != 0 {
		return ErrStepExecuted
	}
	s.err = s.undoer()
	s.step--
	if s.err != nil {
		return s.err
	}
	return nil
}

// Error return the error during executing action.
func (s *Step) Error() error {
	if s.step == 0 {
		return ErrStepNotExecuted
	}
	return s.err
}

// Action return current action of step.
func (s *Step) Action() int {
	if s.step > 0 {
		return 1
	} else if s.step < 0 {
		return -1
	}
	return 0
}

// Fin return the status of step.
func (s *Step) Fin() bool {
	return math.Abs(float64(s.step)) == 1
}

// Step return the step status of step.
func (s *Step) Step() int {
	return int(math.Abs(float64(s.step)))
}

// Progress return the progress status of step.
func (s *Step) Progress() float64 {
	return math.Abs(float64(s.step))
}

// Reset clears the status.
func (s *Step) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.err = nil
	s.step = 0
}
