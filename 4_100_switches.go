package main

import (
	"math"
)

// LampSwitch count how many lamps is on after switched in multiplying rules.
// First iteration is multiplying of 1: 1..n
// Second iteration is multiplying of 2: 2, 4, 6, 8...n
// And so on till n
func LampSwitch(lamp int) (on int) {
	for i := 1; i <= lamp; i++ {
		// if the affected switches are multiplying of, instead of counting
		// we could check whether the modulo of the square root number of switches is 0
		// to determine is it included in multiplying of specified number
		if math.Mod(math.Sqrt(float64(i)), 1) == 0 {
			on += 1
		}
	}
	return
}
