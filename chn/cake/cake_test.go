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

// adding buffers has no effect.
func BenchmarkBuffer(b *testing.B) {
	cakeshop := defaults
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 244 ms
}

// adding variability would result in longer run time.
func BenchmarkVariable(b *testing.B) {
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceStdDev / 4

	cakeshop.Work(b.N)
}
