package main

import (
	"fmt"
	ob "github.com/areThereAnyUserNamesLeft/outburst"
)

func main() {

	log := ob.NewOutBurst()

	log.Out(ob.Knot{"Scooby Doo": "Dog", "Age": 4}).Burst(2)

	// Lazy Methods
	log.Warn(ob.Knot{"Scooby Doo": "Dog", "Age": 4})
	log.ErrCheck(nil)
	err := fmt.Errorf("No scooby snax left!")
	log.ErrCheck(err)

}
