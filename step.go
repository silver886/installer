package installer

// Step is the basic component of a doer.
type Step struct {
	doer  func() error
	doErr error
	done  bool

	undoer  func() error
	undoErr error
	undone  bool
}

// SetDoer sets the doer.
func (s *Step) SetDoer(doer func() error) error {
	if s.done {
		return ErrStepDone
	}
	s.doer = doer
	return nil
}

// SetUndoer sets the undoer.
func (s *Step) SetUndoer(undoer func() error) error {
	if s.undone {
		return ErrStepUndone
	}
	s.undoer = undoer
	return nil
}

// Do triggers the doer.
func (s *Step) Do() error {
	if s.doer == nil {
		return ErrStepNoDoer
	}
	if s.done {
		return ErrStepDone
	}
	s.doErr, s.done = s.doer(), true
	if s.doErr != nil {
		return s.doErr
	}
	return nil
}

// Undo triggers the undoer.
func (s *Step) Undo() error {
	if s.undoer == nil {
		return ErrStepNoUndoer
	}
	if s.undone {
		return ErrStepUndone
	}
	s.undoErr, s.undone = s.undoer(), true
	if s.undoErr != nil {
		return s.undoErr
	}
	return nil
}

// DoError retuen the error during the doation.
func (s *Step) DoError() error {
	if !s.done {
		return ErrStepNonDone
	}
	return s.doErr
}

// UndoError retuen the error during the doation.
func (s *Step) UndoError() error {
	if !s.undone {
		return ErrStepNonUndone
	}
	return s.undoErr
}
