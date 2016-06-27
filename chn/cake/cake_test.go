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
	cakeshop.Work(b.N) // 237 ms
}

// adding buffers has no effect.
func BenchmarkBuffer(b *testing.B) {
	cakeshop := defaults
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 237 ms
}

// adding variability to rate of each step
// increases total time due to channel delay.
func BenchmarkVariable(b *testing.B) {
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.Work(b.N) // 253 ms
}

// adding channel buffers reduces delay
// resulting from variability.
func BenchmarkVariableBuffers(b *testing.B) {
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 242 ms
}

// making middle stage slower
// adds weight directly to the critical path
func BenchmarkSlowIcing(b *testing.B) {
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.Work(b.N) // 1.036 ms
}

// adding buffer has no effect
func BenchmarkSlowIcingBuffers(b *testing.B) {
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 1.036 ms
}

// adding more icing cooks reducing the cost of icing
// to its sequential component, following Amdahl's law.
func BenchmarkSlowIcingIcers(b *testing.B) {
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.NumIcers = 5
	cakeshop.Work(b.N) // 275 ms
}
