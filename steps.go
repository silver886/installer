package installer

import (
	"math"
	"sync"
)

// Stepper implements methods that would used by installer steps.
type Stepper interface {
	Do() error
	Undo() error
	Error() error

	Action() int
	Fin() bool
	Step() int
	Progress() float64

	Reset()
}

// Steps is the set of steppers.
type Steps struct {
	mutex    *sync.Mutex
	step     int
	steppers []Stepper
}

// NewSteps creates a set of steppers with given steppers.
func NewSteps(steppers []Stepper) *Steps {
	return &Steps{
		mutex:    &sync.Mutex{},
		steppers: steppers,
	}
}

// Do triggers each steppers' doer.
func (s *Steps) Do() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.checkSteppers(); err != nil {
		return err
	}
	if s.step != 0 {
		return ErrStepsExecuted
	}
	for _, ss := range s.steppers {
		s.step++
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
	if s.step != 0 {
		return ErrStepsExecuted
	}
	for _, ss := range s.steppers {
		s.step--
		if err := ss.Undo(); err != nil {
			return err
		}
	}
	return nil
}

// Error return the error during executing steppers.
func (s *Steps) Error() error {
	if s.step == 0 {
		return ErrStepsNotExecuted
	}
	return s.steppers[s.Step()-1].Error()
}

// Action return current action of steps.
func (s *Steps) Action() int {
	if s.step > 0 {
		return 1
	} else if s.step < 0 {
		return -1
	}
	return 0
}

// Fin return the status of steps.
func (s *Steps) Fin() bool {
	return int(math.Abs(float64(s.step))) == len(s.steppers)
}

// Step return the step status of steps.
func (s *Steps) Step() int {
	return int(math.Abs(float64(s.step)))
}

// Progress return the progress status of steps.
func (s *Steps) Progress() float64 {
	if err := s.checkSteppers(); err != nil {
		return 0
	}
	return math.Abs(float64(s.step)) / float64(len(s.steppers))
}

// Reset clears the status.
func (s *Steps) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, ss := range s.steppers {
		ss.Reset()
	}
	s.step = 0
}

func (s *Steps) checkSteppers() error {
	if s.steppers == nil || len(s.steppers) == 0 {
		return ErrStepsNoStepper
	}
	return nil
}
