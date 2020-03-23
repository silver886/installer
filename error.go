package installer

import "errors"

var (
	// ErrStepNoDoer means the step does not have doer seted.
	ErrStepNoDoer = errors.New("No doer")
	// ErrStepNoUndoer means the step does not have undoer seted.
	ErrStepNoUndoer = errors.New("No undoer")
	// ErrStepNonDone means the step had not been done.
	ErrStepNonDone = errors.New("Non done")
	// ErrStepNonUndone means the step had not been undone.
	ErrStepNonUndone = errors.New("Non undone")
	// ErrStepDone means the step had already done.
	ErrStepDone = errors.New("Already done")
	// ErrStepUndone means the step had already undone.
	ErrStepUndone = errors.New("Already undone")
)
