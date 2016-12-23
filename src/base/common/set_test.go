package common

import (
	"testing"
)

type Cmd struct {
	cmd string
}

func TestSet(t *testing.T) {
	set := NewSet()
	t.Logf("Create a HashSet value: %v\n", set)

	set.Add(1222)
	set.Add("sadgsajkgh")
	t.Logf("set data: %v", set.List())

	set.Add(1222)
	t.Logf("set data: %v", set.List())

	set.Add(Cmd{cmd: "sdsd"})
	t.Logf("set data: %v", set.List())

	set.Add(Cmd{cmd: "sdsd"})
	t.Logf("set data: %v", set.List())
}
