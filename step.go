package installer

import "sync"

// Step is the basic component of a doer.
type Step struct {
	mutex *sync.Mutex

	doer  func() error
	doErr error
	done  bool

	undoer  func() error
	undoErr error
	undone  bool
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
	s.done = false
	s.undoErr = nil
	s.undone = false
}

// Do triggers the doer.
func (s *Step) Do() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.checkDoer(); err != nil {
		return err
	}
	if s.done {
		return ErrStepDone
	}
	defer func() { s.done = true }()
	s.doErr = s.doer()
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
	if s.undone {
		return ErrStepUndone
	}
	defer func() { s.undone = true }()
	s.undoErr = s.undoer()
	if s.undoErr != nil {
		return s.undoErr
	}
	return nil
}

// Done retuen the status of doer.
func (s *Step) Done() bool {
	return s.done
}

// Undone retuen the status of undoer.
func (s *Step) Undone() bool {
	return s.undone
}

// DoneStep retuen the step status of doer.
func (s *Step) DoneStep() int {
	if s.done {
		return 1
	}
	return 0
}

// UndoneStep retuen the step status of undoer.
func (s *Step) UndoneStep() int {
	if s.undone {
		return 1
	}
	return 0
}

// DoneProgress retuen the progress status of doer.
func (s *Step) DoneProgress() float64 {
	if s.done {
		return 1
	}
	return 0
}

// UndoneProgress retuen the progress status of undoer.
func (s *Step) UndoneProgress() float64 {
	if s.undone {
		return 1
	}
	return 0
}

// DoError retuen the error during doing action.
func (s *Step) DoError() error {
	if !s.done {
		return ErrStepNonDone
	}
	return s.doErr
}

// UndoError retuen the error during undoing action.
func (s *Step) UndoError() error {
	if !s.undone {
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
