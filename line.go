package outburst

import (
	"fmt"
	"time"
)

type Level int

const (
	Trace    Level = 0
	Debug    Level = 1
	Info     Level = 2
	Warn     Level = 3
	Error    Level = 4
	Fatal    Level = 5
	Panic    Level = 6
	Torettes Level = 7
)

func (l Level) String() string {
	return [...]string{"Trace", "Debug", "Info", "Warn", "Error", "Fatal", "Panic", "Torettes"}[l]
}

type Knot map[string]interface{}

type Line struct {
	Knots Knot
	Time  string
	Level Level
	Emoji string
}

// Burst takes a map of string interface and returns a logging line
func (o outburst) Out(knots map[string]interface{}) Line {
	return Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Info, Emoji: ""}
}

func (l Line) Burst(lvl Level) {
	str := fmt.Sprintf("%s %s %s ", l.Level.String(), l.Time, l.Emoji)
	for k, v := range l.Knots {
		str = str + fmt.Sprintf(" -  %s:%v", k, v)
	}
	fmt.Println(str)
}
