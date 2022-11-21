package ch01

import (
	"testing"
)

func TestDemo3(t *testing.T) {
	m := map[int]struct {
		name string
		age  int
	}{
		1: {"user1", 10},
		2: {"user2", 20},
	}
	println(m[1].name)

}

func TestDemo3_1(t *testing.T) {
	m := make(map[string]int, 1000)
	m = map[string]int{
		"a": 1,
	}
	if v, ok := m["a"]; ok {
		println(v, ok)
	}

}

func TestDemo3_2(t *testing.T) {
	type Node struct {
		_    int
		id   int
		data *byte
		next *Node
	}

	n1 := Node{
		id:   1,
		data: nil,
	}
	n2 := Node{
		id:   2,
		data: nil,
		next: &n1,
	}
	println(n2.next.id)

}
