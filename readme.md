# Outburst

Simple logging which gets out the way and program. Include a Yaml/Json file in your projects root and initialize outburst and then just call 
```go
outburst.Log(err)
```

where you'd normaly have 
```go 
if err != nil {
	fmt.Println.fmt(err)
}
```
