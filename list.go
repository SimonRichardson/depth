package depth

import (
	s "github.com/SimonRichardson/depth/selectors"
)

type List struct {
	left, right []s.Action
	current     s.Action
}

func NewList() *List {
	return &List{
		left:  make([]s.Action, 0),
		right: make([]s.Action, 0),
	}
}

func (l *List) Push(action s.Action) {
	if l.current != nil {
		l.left = append(l.left, l.current)
	}
	l.right = make([]s.Action, 0)
	l.current = action
}

func (l *List) Pop() (s.Action, bool) {
	l.right = make([]s.Action, 0)

	if l.current == nil {
		return nil, false
	}

	var (
		total   = len(l.left)
		current = l.current
	)

	if total < 1 {
		l.current = nil
		return l.current, l.current != nil
	}

	l.current = l.left[total-1]
	l.left = l.left[:total-1]

	return current, true
}

func (l *List) Left() (s.Action, bool) {
	if l.current == nil || len(l.right) == 0 {
		return nil, false
	}

	l.left = append(l.left, l.current)
	l.current = l.right[0]
	l.right = l.right[1:]

	return l.current, true
}

func (l *List) Right() (s.Action, bool) {
	if l.current == nil || len(l.left) == 0 {
		return nil, false
	}

	l.right = append([]s.Action{l.current}, l.right...)
	l.current = l.left[len(l.left)-1]
	l.left = l.left[:len(l.left)-1]

	return l.current, true
}

func (l *List) Len() int {
	var amount int
	if l.current != nil {
		amount = 1
	}
	return len(l.left) + amount
}

func (l *List) RemoveAll() {
	l.left = make([]s.Action, 0)
	l.right = make([]s.Action, 0)
	l.current = nil
}

func (l *List) LeftIter() s.Iterator {
	return &iterator{append(l.left, l.current)}
}

func (l *List) RightIter() s.Iterator {
	return &iterator{append([]s.Action{}, l.right...)}
}

type iterator struct {
	values []s.Action
}

func (i *iterator) HasNext() bool {
	return len(i.values) > 1
}

func (i *iterator) Next() s.Action {
	val := i.values[0]
	i.values = i.values[1:]
	return val
}

func (i *iterator) Clone() s.Iterator {
	return &iterator{append([]s.Action{}, i.values...)}
}
