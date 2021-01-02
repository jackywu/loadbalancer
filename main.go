package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	items := []interface{}{"hello", "world", "jacky"}
	balancer := NewRRBalancer(items)
	*/

	/*
	balancer := NewWrrBalancer()
	balancer.Add("hello", 3)
	balancer.Add("world", 2)
	balancer.Add("jacky", 1)
	 */

/*	items := []interface{}{"hello", "world", "jacky"}
	balancer := NewRandomBalancer(items)*/

	balancer := NewWeightedRandomBalancer()
	balancer.Add("hello", 3)
	balancer.Add("world", 2)
	balancer.Add("jacky", 1)

	for {
		item, err := balancer.Next()
		if err != nil {
			panic(err)
		}
		fmt.Println(item)
		time.Sleep(time.Second)
	}
}
