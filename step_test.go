package installer

import (
	"errors"
	"sync"
	"testing"
)

func TestStepNew(t *testing.T) {
	t.Log("Create a step.")
	var normalTest = []struct {
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
		{
			nil,
			func() error { return nil },
		},
		{
			nil,
			func() error { return errors.New("") },
		},
		{
			func() error { return nil },
			nil,
		},
		{
			func() error { return errors.New("") },
			nil,
		},
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			if s := NewStep(tt.a, tt.b); s == nil {
				t.Error("Step should be able to create.")
			}
		})
	}
}

func TestStepReset(t *testing.T) {
	var test = []struct {
		doer   func() error
		result error
	}{
		{
			doer:   func() error { return nil },
			result: nil,
		},
		{
			doer:   func() error { return errors.New("") },
			result: errors.New(""),
		},
	}

	t.Log("Redo a step.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				doer:  tt.doer,
			}
			s.Do()
			s.Reset()
			if err := s.Do(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Step should be able to redo.")
			}
		})
	}
}

func TestStepDo(t *testing.T) {
	var test = []struct {
		doer   func() error
		result error
	}{
		{
			doer:   func() error { return nil },
			result: nil,
		},
		{
			doer:   func() error { return errors.New("") },
			result: errors.New(""),
		},
	}

	t.Log("Normally do a step.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				doer:  tt.doer,
			}
			if err := s.Do(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Step should be able to do.")
			} else if err != nil && s.doErr.Error() != tt.result.Error() {
				t.Error("Step should be have a same error internally.")
			}
		})
	}

	t.Log("Do an emtpy step.")
	t.Run("Emtpy", func(t *testing.T) {
		s := &Step{
			mutex: &sync.Mutex{},
		}
		if err := s.Do(); err != ErrStepNoDoer {
			t.Error("Step should not be able to do.")
		}
	})

	t.Log("Do a done step.")
	for _, tt := range test {
		t.Run("Done", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				doer:  tt.doer,
				done:  true,
			}
			if err := s.Do(); err != ErrStepDone {
				t.Error("Step should not be able to do.")
			}
		})
	}
}

func TestStepUndo(t *testing.T) {
	var test = []struct {
		undoer func() error
		result error
	}{
		{
			undoer: func() error { return nil },
			result: nil,
		},
		{
			undoer: func() error { return errors.New("") },
			result: errors.New(""),
		},
	}

	t.Log("Normally undo a step.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex:  &sync.Mutex{},
				undoer: tt.undoer,
			}
			if err := s.Undo(); err != tt.result && err.Error() != tt.result.Error() {
				t.Error("Step should be able to undo.")
			} else if err != nil && s.undoErr.Error() != tt.result.Error() {
				t.Error("Step should be have a same error internally.")
			}
		})
	}

	t.Log("Undo an emtpy step.")
	t.Run("Emtpy", func(t *testing.T) {
		s := &Step{
			mutex: &sync.Mutex{},
		}
		if err := s.Undo(); err != ErrStepNoUndoer {
			t.Error("Step should not be able to undo.")
		}
	})

	t.Log("Undo an undone step.")
	for _, tt := range test {
		t.Run("Undone", func(t *testing.T) {
			s := &Step{
				mutex:  &sync.Mutex{},
				undoer: tt.undoer,
				undone: true,
			}
			if err := s.Undo(); err != ErrStepUndone {
				t.Error("Step should not be able to undo.")
			}
		})
	}
}

func TestStepDone(t *testing.T) {
	t.Log("Get done status.")
	var normalTest = []bool{
		true,
		false,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				done:  tt,
			}
			if s.Done() != tt {
				t.Error("Done status of step should be the same.")
			}
		})
	}
}

func TestStepUndone(t *testing.T) {
	t.Log("Get undone status.")
	var normalTest = []bool{
		true,
		false,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex:  &sync.Mutex{},
				undone: tt,
			}
			if s.Undone() != tt {
				t.Error("Undone status of step should be the same.")
			}
		})
	}
}

func TestStepDoneStep(t *testing.T) {
	t.Log("Get done step status.")
	var normalTest = []bool{
		true,
		false,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				done:  tt,
			}
			if (s.DoneStep() == 1) != tt {
				t.Error("Done step status should be the same.")
			}
		})
	}
}

func TestStepUndoneStep(t *testing.T) {
	t.Log("Get undone step status.")
	var normalTest = []bool{
		true,
		false,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex:  &sync.Mutex{},
				undone: tt,
			}
			if (s.UndoneStep() == 1) != tt {
				t.Error("Undone step status should be the same.")
			}
		})
	}
}

func TestStepDoneProgress(t *testing.T) {
	t.Log("Get done progress status.")
	var normalTest = []bool{
		true,
		false,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				done:  tt,
			}
			if (s.DoneProgress() == 1) != tt {
				t.Error("Done progress status of step should be the same.")
			}
			if s.DoneProgress() < 0 || s.DoneProgress() > 1 {
				t.Error("Done progress exceeds the range of 0~1.")
			}
		})
	}
}

func TestStepUndoneProgress(t *testing.T) {
	t.Log("Get undone progress status.")
	var normalTest = []bool{
		true,
		false,
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex:  &sync.Mutex{},
				undone: tt,
			}
			if (s.UndoneProgress() == 1) != tt {
				t.Error("Undone progress status of step should be the same.")
			}
			if s.UndoneProgress() < 0 || s.UndoneProgress() > 1 {
				t.Error("Undone progress exceeds the range of 0~1.")
			}
		})
	}
}

func TestStepDoError(t *testing.T) {
	var test = []error{
		nil,
		errors.New(""),
	}

	t.Log("Get do error from an done step.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				done:  true,
				doErr: tt,
			}
			if err := s.DoError(); err != tt && err.Error() != tt.Error() {
				t.Error("Do error should be able to get.")
			}
		})
	}

	t.Log("Get do error from a non-done step.")
	for _, tt := range test {
		t.Run("Non-done", func(t *testing.T) {
			s := &Step{
				mutex: &sync.Mutex{},
				doErr: tt,
			}
			if err := s.DoError(); err != ErrStepNonDone {
				t.Error("Do error should not be able to get.")
			}
		})
	}
}

func TestStepUndoError(t *testing.T) {
	var test = []error{
		nil,
		errors.New(""),
	}

	t.Log("Get undo error from an undone step.")
	for _, tt := range test {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				mutex:   &sync.Mutex{},
				undone:  true,
				undoErr: tt,
			}
			if err := s.UndoError(); err != tt && err.Error() != tt.Error() {
				t.Error("Undo error should be able to get.")
			}
		})
	}

	t.Log("Get undo error from a non-undone step.")
	for _, tt := range test {
		t.Run("Non-undone", func(t *testing.T) {
			s := &Step{
				mutex:   &sync.Mutex{},
				undoErr: tt,
			}
			if err := s.UndoError(); err != ErrStepNonUndone {
				t.Error("Undo error should not be able to get.")
			}
		})
	}
}
