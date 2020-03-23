package installer

import "errors"

var (
	// ErrStepNoInstaller means the step does not have installer registered.
	ErrStepNoInstaller = errors.New("No installer")
	// ErrStepNoUninstaller means the step does not have uninstaller registered.
	ErrStepNoUninstaller = errors.New("No uninstaller")
	// ErrStepNonInstalled means the step had not been installed.
	ErrStepNonInstalled = errors.New("Non installed")
	// ErrStepNonUninstalled means the step had not been uninstalled.
	ErrStepNonUninstalled = errors.New("Non uninstalled")
	// ErrStepInstalled means the step had already installed.
	ErrStepInstalled = errors.New("Already installed")
	// ErrStepUninstalled means the step had already uninstalled.
	ErrStepUninstalled = errors.New("Already uninstalled")
)
