package main

type Balancer interface {
	Next() interface{}
}
