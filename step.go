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

// NewStep creates step with doer and undoer.
func NewStep(doer func() error, undoer func() error) (*Step, error) {
	s := &Step{
		doer:   doer,
		undoer: undoer,
	}
	if doer == nil {
		return nil, ErrStepNoDoer
	} else if s.undoer == nil {
		return nil, ErrStepNoUndoer
	}
	return s, nil
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

// Done retuen the status of doer.
func (s *Step) Done() bool {
	return s.done
}

// Undone retuen the status of undoer.
func (s *Step) Undone() bool {
	return s.undone
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
