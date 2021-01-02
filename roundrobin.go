/*
	refers: <https://github.com/sbabiv/roundrobin/blob/master/roudrobin.go>
	refers: <https://github.com/hlts2/round-robin/blob/master/round_robin.go>
*/

package main

import (
	"sync"
	"sync/atomic"
)

type rrBalancer struct {
	mux   sync.Mutex
	index int
	items []interface{}
}

func NewRRBalancer(items []interface{}) *rrBalancer {
	b := new(rrBalancer)
	b.items = items
	return b
}

func (b *rrBalancer) Next() (interface{}, error) {
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

func (b *rrBalancer) NextWithoutLock() (interface{}, error) {
	length := len(b.items)
	if length == 0 {
		return nil, ErrNoAvailableItem
	}
	item := b.items[b.index]
	var tmp int32
	tmp = int32(b.index)
	b.index = int(atomic.AddInt32(&tmp, 1)) % length
	return item, nil
}