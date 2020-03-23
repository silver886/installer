package installer

import (
	"errors"
	"testing"
)

func TestStepNew(t *testing.T) {
	t.Log("Normally create a step.")
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
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			if s, err := NewStep(tt.a, tt.b); s == nil || s.doer == nil || s.undoer == nil || err != nil {
				t.Error("Step should be able to create.")
			}
		})
	}

	t.Log("Create a step without doer or undoer.")
	var missingTest = []struct {
		a func() error
		b func() error
	}{
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
	for _, tt := range missingTest {
		t.Run("Missing", func(t *testing.T) {
			if s, err := NewStep(tt.a, tt.b); s != nil || (err != ErrStepNoDoer && err != ErrStepNoUndoer) {
				t.Error("Step should not be able to create.")
			}
		})
	}
}

func TestStepDo(t *testing.T) {
	t.Log("Normally do a step.")
	var normalTest = []struct {
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
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				doer: tt.doer,
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
		s := &Step{}
		if err := s.Do(); err != ErrStepNoDoer {
			t.Error("Step should not be able to do.")
		}
	})

	t.Log("Do an insttalled step.")
	var doneTest = []struct {
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
	for _, tt := range doneTest {
		t.Run("Done", func(t *testing.T) {
			s := &Step{
				doer: tt.doer,
				done: true,
			}
			if err := s.Do(); err != ErrStepDone {
				t.Error("Step should not be able to do.")
			}
		})
	}
}

func TestStepUndo(t *testing.T) {
	t.Log("Normally undo a step.")
	var normalTest = []struct {
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
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
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
		s := &Step{}
		if err := s.Undo(); err != ErrStepNoUndoer {
			t.Error("Step should not be able to undo.")
		}
	})

	t.Log("Undo an insttalled step.")
	var undoneTest = []struct {
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
	for _, tt := range undoneTest {
		t.Run("Undone", func(t *testing.T) {
			s := &Step{
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
				done: tt,
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
				undone: tt,
			}
			if s.Undone() != tt {
				t.Error("Undone status of step should be the same.")
			}
		})
	}
}

func TestStepDoError(t *testing.T) {
	t.Log("Get do error from an done step.")
	var normalTest = []error{
		nil,
		errors.New(""),
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				done:  true,
				doErr: tt,
			}
			if err := s.DoError(); err != tt && err.Error() != tt.Error() {
				t.Error("Do error should be able to get.")
			}
		})
	}

	t.Log("Get do error from a non-done step.")
	var nonDoneTest = []error{
		nil,
		errors.New(""),
	}
	for _, tt := range nonDoneTest {
		t.Run("Non-done", func(t *testing.T) {
			s := &Step{
				doErr: tt,
			}
			if err := s.DoError(); err != ErrStepNonDone {
				t.Error("Do error should not be able to get.")
			}
		})
	}
}

func TestStepUndoError(t *testing.T) {
	t.Log("Get undo error from an undone step.")
	var normalTest = []error{
		nil,
		errors.New(""),
	}
	for _, tt := range normalTest {
		t.Run("Normal", func(t *testing.T) {
			s := &Step{
				undone:  true,
				undoErr: tt,
			}
			if err := s.UndoError(); err != tt && err.Error() != tt.Error() {
				t.Error("Undo error should be able to get.")
			}
		})
	}

	t.Log("Get undo error from a non-undone step.")
	var nonUndoneTest = []error{
		nil,
		errors.New(""),
	}
	for _, tt := range nonUndoneTest {
		t.Run("Non-undone", func(t *testing.T) {
			s := &Step{
				undoErr: tt,
			}
			if err := s.UndoError(); err != ErrStepNonUndone {
				t.Error("Undo error should not be able to get.")
			}
		})
	}
}
