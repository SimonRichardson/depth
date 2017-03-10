package selectors

type Sequence interface {
	Add(Action) Action
	AddAt(Action, int) Action
	GetAt(int) (Action, bool)
	GetIndex(Action) int
	Remove(Action) (Action, bool)
	RemoveAt(Action, int) (Action, bool)
	RemoveAll()
	Contains(Action) bool
	Find(Key) (Action, bool)
	FindIndex(Key) int
	Len() int
}
