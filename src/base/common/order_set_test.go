package common

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestOrderSet(t *testing.T) {
	order_set := NewOrderSet(ComparatorFunc)
	t.Logf("Create a OrderSet value: %v\n", order_set)

	for i := 0; i < 100; i++ {
		order_set.Add(strconv.Itoa(rand.Intn(100)))
	}
	t.Logf("set value: %v\n", order_set)

	remove := order_set.Values[order_set.Len()/2]
	order_set.Remove(remove)
	t.Logf("set value: %v\n", order_set)

	order_set.Clear()
	t.Logf("set value: %v\n", order_set)

	if order_set.IsEmpty() {
		t.Logf("set value: %v\n", order_set)
	}

}

//大于
func ComparatorFunc(i, j interface{}) bool {
	id, io1 := i.(int32)
	jd, jo1 := j.(int32)
	if io1 && jo1 {
		return id > jd
	}

	id2, io2 := i.(int)
	jd2, jo2 := j.(int)
	if io2 && jo2 {
		return id2 > jd2
	}

	id3, io3 := i.(string)
	jd3, jo3 := j.(string)
	if io3 && jo3 {
		return id3 > jd3
	}

	id4, io4 := i.(uint32)
	jd4, jo4 := j.(uint32)
	if io4 && jo4 {
		return id4 > jd4
	}

	id5, io5 := i.(float32)
	jd5, jo5 := j.(float32)
	if io5 && jo5 {
		return id5 > jd5
	}

	panic("not match type")
}
