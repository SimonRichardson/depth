package selectors

type Key string

func (k Key) Empty() bool {
	return k == ""
}

type Action interface {
	Key() Key

	Commit() bool
	Revert() bool
}
