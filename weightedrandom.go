package main

import (
	"math/rand"
	"time"
)

type weightedRandomBalancer struct {
	items []interface{}
}

func NewWeightedRandomBalancer() *weightedRandomBalancer {
	return new(weightedRandomBalancer)
}

func (b *weightedRandomBalancer) Add(item interface{}, weight int) {
	for i := 0; i < weight; i++ {
		b.items = append(b.items, item)
	}
}

func (b *weightedRandomBalancer) Next() (interface{}, error) {
	length := len(b.items)
	if length == 0 {
		return nil, ErrNoAvailableItem
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(length)
	return b.items[index], nil
}
