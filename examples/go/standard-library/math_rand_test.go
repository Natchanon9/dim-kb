package standardlibrary

import (
	"fmt"
	"math/rand/v2"
)

func Example_mathRand() {
	fmt.Println(rand.IntN(10) < 10) // global source: auto-seeded, safe for concurrent use

	a := rand.New(rand.NewPCG(1, 2)) // fixed seed: same sequence every run (tests, simulations)
	b := rand.New(rand.NewPCG(1, 2))
	fmt.Println(a.IntN(100) == b.IntN(100), a.Float64() == b.Float64())
	// Output:
	// true
	// true true
}
