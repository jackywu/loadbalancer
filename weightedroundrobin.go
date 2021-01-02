/*
Weighted Round Robin
refers to: <https://github.com/lukeed/wrr/blob/master/src/index.js>
refers to: <https://github.com/lafikl/liblb/blob/master/r2/r2.go>
*/
package main

import (
	"sync"
)

type wrrBalancer struct {
	mux   sync.Mutex
	index int
	items []interface{}
}

func NewWrrBalancer() *wrrBalancer {
	return new(wrrBalancer)
}

func (b *wrrBalancer) Add(item interface{}, weight int) {
	for i := 0; i < weight; i++ {
		b.items = append(b.items, item)
	}
}

func (b *wrrBalancer) Next() (interface{}, error) {
	length := len(b.items)
	if length == 0 {
		return nil, ErrNoAvailableItem
	}
	b.mux.Lock()
	defer b.mux.Unlock()
	item := b.items[b.index]
	b.index = (b.index + 1) % length
	return item, nil
}
