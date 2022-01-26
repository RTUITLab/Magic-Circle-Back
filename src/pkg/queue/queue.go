package queue

import (
	"errors"
	"strconv"
)

var (
	QueueIsEmpty = errors.New("Queue is empty")
)

type StringQueue []string

func NewStringQueue(slice []string) StringQueue {
	return StringQueue(slice)
}

func (s *StringQueue) Get() (string, error) {
	if len(*s) == 0 {
		return "", QueueIsEmpty
	}

	res, new := s.get()
	*s = new
	return res, nil
}

func (s StringQueue) get() (res string, newQueue []string) {
	return s[0], s[1:]
}

type IntQueue []int

func NewIntQueue(slice []int) IntQueue {
	return IntQueue(slice)
}

func (i *IntQueue) Get() (int, error) {
	if len(*i) == 0 {
		return -1, QueueIsEmpty
	}

	res, new := i.get()
	*i = new
	return res, nil
}

func (i IntQueue) get() (res int, newQueue []int) {
	return i[0], i[1:]
}

type StringQueueToIntOpts struct {
	IfNotIntElemIgnore	bool
}

type StringQueueToIntQueueFunc func(sq StringQueue) (IntQueue, error)

func StringQueueToIntQueue(opt StringQueueToIntOpts) StringQueueToIntQueueFunc {
	return func(sq StringQueue) (IntQueue, error) {
		var intQueue IntQueue
		{
			for elem, err := sq.Get(); err != QueueIsEmpty; elem, err = sq.Get() {
				if intElem, err := strconv.Atoi(elem); err != nil {
					if opt.IfNotIntElemIgnore {
						continue
					} else {
						return nil, err
					}
				} else {
					intQueue = append(intQueue, intElem)
				}
			}
		}
		return intQueue, nil
	}
}