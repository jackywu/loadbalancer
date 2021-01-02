package main

import (
	"math/rand"
	"time"
)

type randomBalancer struct {
	items []interface{}
}

func NewRandomBalancer(items []interface{}) *randomBalancer {
	b := new(randomBalancer)
	b.items = items
	return b
}

func (b *randomBalancer) Next() (interface{}, error) {
	length := len(b.items)
	if length == 0 {
		return nil, ErrNoAvailableItem
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(length)
	return b.items[index], nil
}