package depth

import s "github.com/SimonRichardson/depth/selectors"

type transaction struct {
	list  *List
	stash s.Iterator
}

func NewTransaction() s.Transaction {
	return &transaction{
		list: NewList(),
	}
}

func (t *transaction) Undo() (s.Action, bool) {
	if action, ok := t.list.Left(); ok {
		return action, action.Revert()
	}
	return nil, false
}

func (t *transaction) Redo() (s.Action, bool) {
	if action, ok := t.list.Right(); ok {
		return action, action.Commit()
	}
	return nil, false
}

func (t *transaction) Commit() bool {
	t.stash = t.list.LeftIter()
	iter := t.stash.Clone()

	for iter.HasNext() {
		action := iter.Next()
		if !action.Commit() {
			return false
		}
	}

	return true
}

func (t *transaction) Revert() bool {
	if t.stash == nil {
		return false
	}

	var (
		iter = t.stash
		list = NewList()
	)
	for iter.HasNext() {
		action := iter.Next()
		if !action.Revert() {
			return false
		}

		list.Push(action)
	}

	t.list = list
	t.stash = nil

	return true
}

type transactionReadWriter struct {
	transaction *transaction
}

func NewTransactionReadWriter(t *transaction) *transactionReadWriter {
	return &transactionReadWriter{
		transaction: t,
	}
}

func (t *transactionReadWriter) Read(p []byte) (int, error) {
	return -1, nil
}

func (t *transactionReadWriter) Write(p []byte) (int, error) {
	return -1, nil
}
