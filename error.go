package installer

import "errors"

var (
	// ErrStepNoDoer means the step does not have doer.
	ErrStepNoDoer = errors.New("Step has no doer")
	// ErrStepNoUndoer means the step does not have undoer.
	ErrStepNoUndoer = errors.New("Step has no undoer")
	// ErrStepNonDone means the step had not been done.
	ErrStepNonDone = errors.New("Step is not done")
	// ErrStepNonUndone means the step had not been undone.
	ErrStepNonUndone = errors.New("Step is not undone")
	// ErrStepDone means the step had already done.
	ErrStepDone = errors.New("Step is already done")
	// ErrStepUndone means the step had already undone.
	ErrStepUndone = errors.New("Step is already undone")
)
