package common

import (
	"bytes"
	"fmt"
	"sort"
	"sync"
)

const (
	ORDERSET_RET_NOT_FOUND int32 = -1
	ORDERSET_RET_OK        int32 = 1
	ORDERSET_INIT_SET_SIZE int32 = 100
)

type ComparatorFunc func(i, j interface{}) bool

type OrderSet struct {
	Comparator ComparatorFunc
	Values     []interface{}
	len        int
	sync.RWMutex
}

func NewOrderSet(compartor ComparatorFunc) *OrderSet {
	return &OrderSet{
		Comparator: compartor,
		Values:     make([]interface{}, 0, ORDERSET_INIT_SET_SIZE),
		len:        0,
	}
}

func (s *OrderSet) Len() int           { return s.len }
func (s *OrderSet) Swap(i, j int)      { s.Values[i], s.Values[j] = s.Values[j], s.Values[i] }
func (s *OrderSet) Less(i, j int) bool { return s.Comparator(s.Values[i], s.Values[j]) }

func (s *OrderSet) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	if index := s.has(item); ORDERSET_RET_NOT_FOUND == index {
		s.Values = append(s.Values, item)
		sort.Sort(s)
		s.len++
	}
}

func (s *OrderSet) Remove(item interface{}) int32 {
	s.Lock()
	defer s.Unlock()
	index := s.has(item)
	if ORDERSET_RET_NOT_FOUND == index {
		return ORDERSET_RET_NOT_FOUND
	}
	s.Values = append(s.Values[:index], s.Values[index+1:]...)
	s.len--
	return ORDERSET_RET_OK
}

func (s *OrderSet) has(item interface{}) int32 {
	//s.Lock()
	//defer s.Unlock()
	for index := int32(0); index < int32(s.Len()); index++ {
		if item == s.Values[index] {
			return index
		}
	}
	return ORDERSET_RET_NOT_FOUND
}

func (s *OrderSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.Values = make([]interface{}, 0, ORDERSET_INIT_SET_SIZE)
	s.len = 0
}

func (s *OrderSet) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()
	if 0 == s.Len() {
		return true
	}
	return false
}

func (s *OrderSet) List() []interface{} {
	s.RLock()
	defer s.RUnlock()
	return s.Values
}

func (s *OrderSet) String() string {
	s.Lock()
	defer s.Unlock()
	var buf bytes.Buffer
	buf.WriteString("OrderSet{")
	first := true
	for _, v := range s.Values {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", v))
	}

	buf.WriteString("}")
	return buf.String()
}
