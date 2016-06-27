package main

import (
	"testing"
	"time"
)

var defaults = Shop{
	Verbose: testing.Verbose(),
	Cakes:   20,

	BakeTime:     10 * time.Millisecond,
	NumIcers:     1,
	IceTime:      10 * time.Millisecond,
	InscribeTime: 10 * time.Millisecond,
}

// default benchmark with sequential concurrent processes
func Benchmark(b *testing.B) {
	// baseline: one baker, one icer, one inscriber.
	cakeshop := defaults
	cakeshop.Work(b.N) // 244 ms
}

func BenchmarkBuffer(b *testing.B) {
	// adding buffers has no effect.
	cakeshop := defaults
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 244 ms
}
