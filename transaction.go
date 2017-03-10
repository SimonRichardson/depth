package depth

import (
	"sync"

	s "github.com/SimonRichardson/depth/selectors"
)

type transaction struct {
	mutex   sync.Mutex
	actions s.Sequence
	ptr     s.Key
}

func NewTransaction() s.Transaction {
	return &transaction{
		mutex:   sync.Mutex{},
		actions: NewSequence(),
		ptr:     s.Key(""),
	}
}

func (t *transaction) Do(action s.Action) {
	t.commit(action)

	if t.ptr.Empty() {
		t.ptr = action.Key()
	}

	t.actions.Add(action)
}

func (t *transaction) Undo() (s.Action, bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if t.actions.Len() == 0 || t.ptr.Empty() {
		return nil, false
	}

	if index := t.actions.FindIndex(t.ptr); index >= 0 {
		var (
			action       s.Action
			modified, ok bool
		)
		for i := t.actions.Len() - 1; i >= index; i-- {
			if action, ok = t.actions.GetAt(i); ok {
				t.revert(action)
				modified = true
			}
		}

		return action, modified
	}

	return nil, false
}

func (t *transaction) Redo() (s.Action, bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if t.actions.Len() == 0 || t.ptr.Empty() {
		return nil, false
	}

	if index := t.actions.FindIndex(t.ptr); index >= 0 {
		var (
			action       s.Action
			modified, ok bool
		)

		for i := index; i < t.actions.Len(); i++ {
			if action, ok = t.actions.GetAt(i); ok {
				t.commit(action)
				modified = true
			}
		}

		return action, modified
	}

	return nil, false
}

func (t *transaction) commit(action s.Action) {

}

func (t *transaction) revert(action s.Action) {

}
