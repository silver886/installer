package installer

import "errors"

var (

	// ErrStepsNoStepper means the steps does not have any stepper.
	ErrStepsNoStepper = errors.New("Steps has no stepper")
	// ErrStepsExecuted means the steps is already executed.
	ErrStepsExecuted = errors.New("Steps is already executed")
	// ErrStepsNonDone means the steps is not done.
	ErrStepsNonDone = errors.New("Steps is not done")
	// ErrStepsNonUndone means the steps is not undone.
	ErrStepsNonUndone = errors.New("Steps is not undone")

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
