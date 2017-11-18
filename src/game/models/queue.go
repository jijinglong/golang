package models

import (
	"container/list"
	"sync"
)

type Queue struct {
	data *list.List
	lock *sync.Mutex
}

func NewQueue() *Queue {
	q := new(Queue)
	q.data = list.New()
	q.lock = &sync.Mutex{}
	return q
}

func (q *Queue) Push(v interface{}) *list.Element {
	defer q.lock.Unlock()
	q.lock.Lock()
	return q.data.PushBack(v)
}

func (q *Queue) PopMulti(num int) []interface{} {
	defer q.lock.Unlock()
	q.lock.Lock()
	if q.data.Len() < num {
		return nil
	}

	var result []interface{}
	for i := 0; i < num; i++ {
		e := q.data.Front()
		v := q.data.Remove(e)
		e.Value = nil
		result = append(result, v)
	}
	return result
}

func (q *Queue) Remove(e *list.Element) interface{} {
	defer q.lock.Unlock()
	q.lock.Lock()
	v := q.data.Remove(e)
	e.Value = nil
	return v
}
