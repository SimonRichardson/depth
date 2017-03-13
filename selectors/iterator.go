package selectors

type Iterator interface {
	HasNext() bool
	Next() Action
}
