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
)
