package installer

import "errors"

var (

	// ErrStepsNoStepper means the steps does not have any stapper.
	ErrStepsNoStepper = errors.New("Steps has no stapper")
	// ErrStepsNonDone means the steps had not been done.
	ErrStepsNonDone = errors.New("Steps is not done")
	// ErrStepsNonUndone means the steps had not been undone.
	ErrStepsNonUndone = errors.New("Steps is not undone")
	// ErrStepsDone means the steps had already done.
	ErrStepsDone = errors.New("Steps is already done")
	// ErrStepsUndone means the steps had already undone.
	ErrStepsUndone = errors.New("Steps is already undone")

	// ErrStepNoDoer means the step does not have doer.
	ErrStepNoDoer = errors.New("Step has no doer")
	// ErrStepNoUndoer means the step does not have undoer.
	ErrStepNoUndoer = errors.New("Step has no undoer")
	// ErrStepExecuted means the step is already executed.
	ErrStepExecuted = errors.New("Step is already executed")
	// ErrStepNonDone means the step is not done.
	ErrStepNonDone = errors.New("Step is not done")
	// ErrStepNonUndone means the step is not undone.
	ErrStepNonUndone = errors.New("Step is not undone")
)
