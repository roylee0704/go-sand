package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Shop keep tracks of manipulative states.
type Shop struct {
	Verbose bool // verbose flag
	Cakes   int  // number of cakes

	BakeTime   time.Duration //
	BakeStdDev time.Duration //
	BakeBuf    int           //

	NumIcers  int           //
	IceTime   time.Duration //
	IceStdDev time.Duration //
	IceBuf    int           //

	InscribeTime   time.Duration //
	InscribeStdDev time.Duration //
}

// baker is the first process.
func (s *Shop) baker(baked chan<- int) {
	for i := 0; i < s.Cakes; i++ {
		s.v(fmt.Sprintf("Baking cake %d\n", i))
		work(s.BakeTime, s.BakeStdDev)
		baked <- i // as id
	}
	close(baked)
}

// icer will be the second process.
func (s *Shop) icer(baked <-chan int, iced chan<- int) {
	for cake := range baked {
		s.v(fmt.Sprintf("Icing cake %d\n", cake))
		work(s.IceTime, s.IceStdDev)
		iced <- cake
	}

	// when its not sequential (numIcers>1), we aren't sure if
	// other icer has completed, hence, no close
	// close(iced)
}

// inscribe as the third and final process.
func (s *Shop) inscribe(iced <-chan int) {
	for i := 0; i < s.Cakes; i++ {
		cake := <-iced
		s.v(fmt.Sprintf("Inscribing cake %d\n", cake))
		work(s.InscribeTime, s.InscribeStdDev)
		s.v(fmt.Sprintf("Finished cake %d\n", cake))
	}
}

// work blocks goroutine for a period of time
// that is normally distributed in N(μ,σ2).
func work(μ, stddev time.Duration) {
	z := rand.NormFloat64()
	delay := μ + stddev*time.Duration(z)
	time.Sleep(delay)
}

func (s *Shop) v(msg string) {
	if s.Verbose {
		fmt.Printf(msg)
	}
}

// Work defines main algorithm for simulation
func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan int, s.BakeBuf)
		iced := make(chan int, s.IceBuf)

		go s.baker(baked) // one baker

		for i := 0; i < s.NumIcers; i++ {
			go s.icer(baked, iced) // n icers.
		}
		s.inscribe(iced) // one inscriber
	}
}
