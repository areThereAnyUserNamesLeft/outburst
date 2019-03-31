package outburst

import "github.com/fatih/color"

var colors = map[string]*color.Color{
	"Trace": color.New(color.FgCyan, color.Bold),
	"Debug": color.New(color.FgMagenta, color.Bold),
	"Info":  color.New(color.FgBlue, color.Bold),
	"Warn":  color.New(color.FgYellow, color.Bold),
	"Error": color.New(color.FgGreen, color.Bold),
	"Fatal": color.New(color.FgRed, color.Bold),
	"Panic": color.New(color.FgRed, color.BgWhite),
}
