package installer

import "errors"

var (

	// ErrStepsNoStepper means the steps does not have any stepper.
	ErrStepsNoStepper = errors.New("Steps has no stepper")
	// ErrStepsExecuted means the steps is already executed.
	ErrStepsExecuted = errors.New("Steps is already executed")
	// ErrStepsNotExecuted means the steps is not executed.
	ErrStepsNotExecuted = errors.New("Steps is not executed")

	// ErrStepNoDoer means the step does not have doer.
	ErrStepNoDoer = errors.New("Step has no doer")
	// ErrStepNoUndoer means the step does not have undoer.
	ErrStepNoUndoer = errors.New("Step has no undoer")
	// ErrStepExecuted means the step is already executed.
	ErrStepExecuted = errors.New("Step is already executed")
	// ErrStepNotExecuted means the step is not executed.
	ErrStepNotExecuted = errors.New("Step is not executed")
)
