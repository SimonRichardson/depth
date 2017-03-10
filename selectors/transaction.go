package selectors

type Transaction interface {
	Undo() (Action, bool)
	Redo() (Action, bool)
}
