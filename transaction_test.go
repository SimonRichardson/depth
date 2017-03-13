package depth

import "testing"

func Test_Transaction(t *testing.T) {
	trans := NewTransaction()

	trans.Undo()
	trans.Redo()
}
