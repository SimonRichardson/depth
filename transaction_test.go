package depth

import "testing"

func Test_Transaction(t *testing.T) {
	trans := NewTransaction()
	trans.Do(MakeAddInt())
	trans.Do(MakeAddString())

	trans.Undo()
	trans.Redo()
}
