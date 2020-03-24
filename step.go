package installer

import "sync"

// Step is the basic component of a doer.
type Step struct {
	mutex *sync.Mutex
	step  int

	doer  func() error
	doErr error

	undoer  func() error
	undoErr error
}

// NewStep creates step with doer and undoer.
func NewStep(doer func() error, undoer func() error) *Step {
	return &Step{
		mutex:  &sync.Mutex{},
		doer:   doer,
		undoer: undoer,
	}
}

// Reset clears the status.
func (s *Step) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.doErr = nil
	s.undoErr = nil
	s.step = 0
}

// Do triggers the doer.
func (s *Step) Do() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.checkDoer(); err != nil {
		return err
	}
	if s.step != 0 {
		return ErrStepExecuted
	}
	s.doErr = s.doer()
	s.step++
	if s.doErr != nil {
		return s.doErr
	}
	return nil
}

// Undo triggers the undoer.
func (s *Step) Undo() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.checkUndoer(); err != nil {
		return err
	}
	if s.step != 0 {
		return ErrStepExecuted
	}
	s.undoErr = s.undoer()
	s.step--
	if s.undoErr != nil {
		return s.undoErr
	}
	return nil
}

// Done retuen the status of doer.
func (s *Step) Done() bool {
	return s.step == 1
}

// Undone retuen the status of undoer.
func (s *Step) Undone() bool {
	return s.step == -1
}

// DoneStep retuen the step status of doer.
func (s *Step) DoneStep() int {
	if s.step <= 0 {
		return 0
	}
	return s.step
}

// UndoneStep retuen the step status of undoer.
func (s *Step) UndoneStep() int {
	if s.step >= 0 {
		return 0
	}
	return -s.step
}

// DoneProgress retuen the progress status of doer.
func (s *Step) DoneProgress() float64 {
	if s.step <= 0 {
		return 0
	}
	return 1
}

// UndoneProgress retuen the progress status of undoer.
func (s *Step) UndoneProgress() float64 {
	if s.step >= 0 {
		return 0
	}
	return 1
}

// DoError retuen the error during doing action.
func (s *Step) DoError() error {
	if s.step <= 0 {
		return ErrStepNonDone
	}
	return s.doErr
}

// UndoError retuen the error during undoing action.
func (s *Step) UndoError() error {
	if s.step >= 0 {
		return ErrStepNonUndone
	}
	return s.undoErr
}

func (s *Step) checkDoer() error {
	if s.doer == nil {
		return ErrStepNoDoer
	}
	return nil
}

func (s *Step) checkUndoer() error {
	if s.undoer == nil {
		return ErrStepNoUndoer
	}
	return nil
}
