package main

import (
	ob "github.com/areThereAnyUserNamesLeft/outburst"
)

func main() {

	log := ob.NewOutBurst()
	log.Warn(ob.Knot{"Scooby Doo": "Dog", "Age": 4})

	log.Out(ob.Knot{"Scooby Doo": "Dog", "Age": 4}).Burst(2)
}
