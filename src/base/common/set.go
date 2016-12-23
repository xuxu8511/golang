package common

import (
	"bytes"
	"fmt"
	"sync"
)

type Set struct {
	m map[interface{}]struct{}
	sync.RWMutex
}

func NewSet() *Set {
	return &Set{
		m: make(map[interface{}]struct{}),
	}
}

func (s *Set) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{}
}

func (s *Set) Remove(item interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *Set) has(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return len(s.m)
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = make(map[interface{}]struct{})
}

func (s *Set) IsEmpty() bool {
	if 0 == s.Len() {
		return true
	}
	return false
}

func (s *Set) List() []interface{} {
	s.RLock()
	defer s.RUnlock()
	list := make([]interface{}, 0, s.Len())
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func (s *Set) String() string {
	var buf bytes.Buffer
	buf.WriteString("HashSet{")
	first := true
	for key := range s.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}
