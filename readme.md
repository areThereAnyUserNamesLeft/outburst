# Outburst

Simple, light weight, lazy logging which gets out the way so you can program.
1. Import
```go
import ob "github.com/areThereAnyUserNamesLeft/outburst"

```
2. Initiate the logger

```go
	log := ob.NewOutBurst()
```
3. and go...
```go
	log.Out(ob.Knot{"Scooby Doo": Dog, "Age": 4}).Burst(3)
	// Or if you want an even shorter version!
	log.Info(ob.KV{"Scooby Doo": Dog, "Age": 4})
	// (Hint - Knot and KV are the same data structure)
```
where you'd normaly have
```go
	fmt.Printf("Scooby Doo": %s, Age: %v ", Dog, Age)
```
Before you dismiss `outburst` lets see what that line gets you over the fmt
package?
```
Info [2019-03-30]-[15:40:31]- Scooby Doo:Dog -  Age:4
```
Additionally, the output can be colored and have emojis to reflect the error.

For the most part this is handled by a yaml file in your projects root - see `outburst.yaml` for an example.

Extras convienience functions - 
```go
ob.ErrCheck(err, KV{"Scooby Doo": Dog, "Age": 4})

// replaces

if err != nil {
	fmt.Printf("Scooby Doo": %s, Age: %v ", Dog, Age)
}

// If the error not `nil` it will fire off the Error and the Key Value as an Error message. 
// Also the Knot/KV is an optional parameter so you can include them or not or include multiple. 
```
