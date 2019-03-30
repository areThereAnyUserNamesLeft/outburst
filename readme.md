# Outburst

Simple, light weight, lazy logging which gets out the way and program.
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
```
where you'd normaly have
```go
	fmt.Printf("Scooby Doo": %s, Age: %v ", Dog, Age)
```
Before you dismiss `outburst` lets see what that line gets you over the fmt
package?
```
Info [2019-03-30]-[15:40:31]-   -  Scooby Doo:Dog -  Age:4

```
