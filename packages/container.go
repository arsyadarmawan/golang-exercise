package packages

import (
	"container/list"
	"fmt"
)

func Container() {
	l := list.New()
	l.PushBack(4)
	l.PushBack(3)
	l.PushFront(1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
