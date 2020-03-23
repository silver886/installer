package installer

import (
	"errors"
	"testing"
)

func TestStepRegisterInstaller(t *testing.T) {
	t.Log("Normally register an installer to a step.")
	var normalTest = []func() error{
		func() error { return nil },
		func() error { return errors.New("") },
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{}
			if err := s.RegisterInstaller(tt); err != nil {
				t.Error("Installer should be able to register.")
			}
		})
	}

	t.Log("Register multiple installers to a step.")
	var mutiTest = []struct {
		a func() error
		b func() error
	}{
		{
			func() error { return nil },
			func() error { return nil },
		},
		{
			func() error { return nil },
			func() error { return errors.New("") },
		},
		{
			func() error { return errors.New("") },
			func() error { return nil },
		},
		{
			func() error { return errors.New("") },
			func() error { return errors.New("") },
		},
	}
	for _, tt := range mutiTest {
		t.Run("Multiple", func(t *testing.T) {
			s := &Step{}
			if err := s.RegisterInstaller(tt.a); err != nil {
				t.Error("The first installer should be able to register.")
			}
			if err := s.RegisterInstaller(tt.b); err != nil {
				t.Error("The second installer should be able to register.")
			}
		})
	}

	t.Log("Register after installed to a step.")
	var installedTest = []func() error{
		func() error { return nil },
		func() error { return errors.New("") },
	}
	for _, tt := range installedTest {
		t.Run("Installed", func(t *testing.T) {
			s := &Step{
				installed: true,
			}
			if err := s.RegisterInstaller(tt); err != ErrStepInstalled {
				t.Error("Installer should not be able to register.")
			}
		})
	}
}

func TestStepRegisterUninstaller(t *testing.T) {
	t.Log("Normally register an uninstaller to a step.")
	var normalTest = []func() error{
		func() error { return nil },
		func() error { return errors.New("") },
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{}
			if err := s.RegisterUninstaller(tt); err != nil {
				t.Error("Uninstaller should be able to register.")
			}
		})
	}

	t.Log("Register multiple uninstallers to a step.")
	var mutiTest = []struct {
		a func() error
		b func() error
	}{
		{
			func() error { return nil },
			func() error { return nil },
		},
		{
			func() error { return nil },
			func() error { return errors.New("") },
		},
		{
			func() error { return errors.New("") },
			func() error { return nil },
		},
		{
			func() error { return errors.New("") },
			func() error { return errors.New("") },
		},
	}
	for _, tt := range mutiTest {
		t.Run("Multiple", func(t *testing.T) {
			s := &Step{}
			if err := s.RegisterUninstaller(tt.a); err != nil {
				t.Error("The first uninstaller should be able to register.")
			}
			if err := s.RegisterUninstaller(tt.b); err != nil {
				t.Error("The second uninstaller should be able to register.")
			}
		})
	}

	t.Log("Register after installed to a step.")
	var installedTest = []func() error{
		func() error { return nil },
		func() error { return errors.New("") },
	}
	for _, tt := range installedTest {
		t.Run("Uninstalled", func(t *testing.T) {
			s := &Step{
				uninstalled: true,
			}
			if err := s.RegisterUninstaller(tt); err != ErrStepUninstalled {
				t.Error("Uninstaller should not be able to register.")
			}
		})
	}
}

func TestStepInstall(t *testing.T) {
	t.Log("Normally install a step.")
	var normalTest = []struct {
		installer func() error
		result    error
	}{
		{
			installer: func() error { return nil },
			result:    nil,
		},
		{
			installer: func() error { return errors.New("") },
			result:    errors.New(""),
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				installer: tt.installer,
			}
			if err := s.Install(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Step should be able to install.")
			} else if err != nil && s.installErr.Error() != tt.result.Error() {
				t.Error("Step should be have a same error internally.")
			}
		})
	}

	t.Log("Install an emtpy step.")
	t.Run("Emtpy", func(t *testing.T) {
		s := &Step{}
		if err := s.Install(); err != ErrStepNoInstaller {
			t.Error("Step should not be able to install.")
		}
	})

	t.Log("Install an insttalled step.")
	var installedTest = []struct {
		installer func() error
		result    error
	}{
		{
			installer: func() error { return nil },
			result:    nil,
		},
		{
			installer: func() error { return errors.New("") },
			result:    errors.New(""),
		},
	}
	for _, tt := range installedTest {
		t.Run("Installed", func(t *testing.T) {
			s := &Step{
				installer: tt.installer,
				installed: true,
			}
			if err := s.Install(); err != ErrStepInstalled {
				t.Error("Step should not be able to install.")
			}
		})
	}
}

func TestStepUninstall(t *testing.T) {
	t.Log("Normally uninstall a step.")
	var normalTest = []struct {
		uninstaller func() error
		result      error
	}{
		{
			uninstaller: func() error { return nil },
			result:      nil,
		},
		{
			uninstaller: func() error { return errors.New("") },
			result:      errors.New(""),
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				uninstaller: tt.uninstaller,
			}
			if err := s.Uninstall(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Step should be able to uninstall.")
			} else if err != nil && s.uninstallErr.Error() != tt.result.Error() {
				t.Error("Step should be have a same error internally.")
			}
		})
	}

	t.Log("Uninstall an emtpy step.")
	t.Run("Emtpy", func(t *testing.T) {
		s := &Step{}
		if err := s.Uninstall(); err != ErrStepNoUninstaller {
			t.Error("Step should not be able to uninstall.")
		}
	})

	t.Log("Uninstall an insttalled step.")
	var uninstalledTest = []struct {
		uninstaller func() error
		result      error
	}{
		{
			uninstaller: func() error { return nil },
			result:      nil,
		},
		{
			uninstaller: func() error { return errors.New("") },
			result:      errors.New(""),
		},
	}
	for _, tt := range uninstalledTest {
		t.Run("Uninstalled", func(t *testing.T) {
			s := &Step{
				uninstaller: tt.uninstaller,
				uninstalled: true,
			}
			if err := s.Uninstall(); err != ErrStepUninstalled {
				t.Error("Step should not be able to uninstall.")
			}
		})
	}
}

func TestStepInstallError(t *testing.T) {
	t.Log("Get install error from an installed step.")
	var normalTest = []error{
		nil,
		errors.New(""),
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				installed:  true,
				installErr: tt,
			}
			if err := s.InstallError(); err != tt && err.Error() != tt.Error() {
				t.Error("Install error should be able to get.")
			}
		})
	}

	t.Log("Get install error from a non-installed step.")
	var nonInstalledTest = []error{
		nil,
		errors.New(""),
	}
	for _, tt := range nonInstalledTest {
		t.Run("Non-installed", func(t *testing.T) {
			s := &Step{
				installErr: tt,
			}
			if err := s.InstallError(); err != ErrStepNonInstalled {
				t.Error("Install error should not be able to get.")
			}
		})
	}
}

func TestStepUninstallError(t *testing.T) {
	t.Log("Get uninstall error from an uninstalled step.")
	var normalTest = []error{
		nil,
		errors.New(""),
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				uninstalled:  true,
				uninstallErr: tt,
			}
			if err := s.UninstallError(); err != tt && err.Error() != tt.Error() {
				t.Error("Uninstall error should be able to get.")
			}
		})
	}

	t.Log("Get uninstall error from a non-uninstalled step.")
	var nonUninstalledTest = []error{
		nil,
		errors.New(""),
	}
	for _, tt := range nonUninstalledTest {
		t.Run("Non-uninstalled", func(t *testing.T) {
			s := &Step{
				uninstallErr: tt,
			}
			if err := s.UninstallError(); err != ErrStepNonUninstalled {
				t.Error("Uninstall error should not be able to get.")
			}
		})
	}
}
