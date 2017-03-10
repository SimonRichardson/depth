package depth

import (
	s "github.com/SimonRichardson/depth/selectors"
)

type sequence struct {
	actions []s.Action
}

func NewSequence() s.Sequence {
	return &sequence{actions: make([]s.Action, 0)}
}

func (q *sequence) Add(a s.Action) s.Action {
	q.actions = append(q.actions, a)
	return a
}

func (q *sequence) AddAt(a s.Action, i int) s.Action {
	q.actions = append(q.actions[:i], append([]s.Action{a}, q.actions[i:]...)...)
	return a
}

func (q *sequence) GetAt(i int) (s.Action, bool) {
	if num := len(q.actions); i >= 0 && i < num {
		return q.actions[i], true
	}
	return nil, false
}

func (q *sequence) GetIndex(a s.Action) int {
	for k, v := range q.actions {
		if v.Key() == a.Key() {
			return k
		}
	}
	return -1
}

func (q *sequence) Remove(a s.Action) (s.Action, bool) {
	if index := q.GetIndex(a); index >= 0 {
		res := q.actions[index]
		q.actions = append(q.actions[:index], q.actions[index+1:]...)
		return res, true
	}
	return nil, false
}

func (q *sequence) RemoveAt(a s.Action, i int) (s.Action, bool) {
	if num := len(q.actions); i >= 0 && i < num {
		res := q.actions[i]
		q.actions = append(q.actions[:i], q.actions[i+1:]...)
		return res, true
	}
	return nil, false
}

func (q *sequence) RemoveAll() {
	q.actions = make([]s.Action, 0)
}

func (q *sequence) Contains(a s.Action) bool {
	return q.GetIndex(a) >= 0
}

func (q *sequence) Find(k s.Key) (s.Action, bool) {
	for _, v := range q.actions {
		if v.Key() == k {
			return v, true
		}
	}
	return nil, false
}

func (q *sequence) FindIndex(key s.Key) int {
	for k, v := range q.actions {
		if v.Key() == key {
			return k
		}
	}
	return -1
}

func (q *sequence) Len() int {
	return len(q.actions)
}
