package main

import (
	ob "github.com/areThereAnyUserNamesLeft/outburst"
)

func main() {

	log := ob.NewOutBurst()
	log.Out(ob.Knot{"Scooby Doo": "Dog", "Age": 4}).Burst(3)

}
