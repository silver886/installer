package installer

// Step is the basic component of a installer.
type Step struct {
	installer  func() error
	installErr error
	installed  bool

	uninstaller  func() error
	uninstallErr error
	uninstalled  bool
}

// RegisterInstaller sets the installer.
func (s *Step) RegisterInstaller(installer func() error) error {
	if s.installed {
		return ErrStepInstalled
	}
	s.installer = installer
	return nil
}

// RegisterUninstaller sets the uninstaller.
func (s *Step) RegisterUninstaller(uninstaller func() error) error {
	if s.uninstalled {
		return ErrStepUninstalled
	}
	s.uninstaller = uninstaller
	return nil
}

// Install triggers the installer.
func (s *Step) Install() error {
	if s.installer == nil {
		return ErrStepNoInstaller
	}
	if s.installed {
		return ErrStepInstalled
	}
	s.installErr, s.installed = s.installer(), true
	if s.installErr != nil {
		return s.installErr
	}
	return nil
}

// Uninstall triggers the uninstaller.
func (s *Step) Uninstall() error {
	if s.uninstaller == nil {
		return ErrStepNoUninstaller
	}
	if s.uninstalled {
		return ErrStepUninstalled
	}
	s.uninstallErr, s.uninstalled = s.uninstaller(), true
	if s.uninstallErr != nil {
		return s.uninstallErr
	}
	return nil
}

// InstallError retuen the error during the installation.
func (s *Step) InstallError() error {
	if !s.installed {
		return ErrStepNonInstalled
	}
	return s.installErr
}

// UninstallError retuen the error during the installation.
func (s *Step) UninstallError() error {
	if !s.uninstalled {
		return ErrStepNonUninstalled
	}
	return s.uninstallErr
}
