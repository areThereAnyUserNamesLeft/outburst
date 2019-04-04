package outburst

import (
	"fmt"
	"github.com/fatih/color"
)

type Line struct {
	Knots       Knot
	Time        string
	Level       Level
	EmojiChoice string
	Emoji       string
	Color       *color.Color
	File        string
}

// Burst takes a map of string interface and returns a logging line
func (l Line) Burst(lvl Level) {
	str := fmt.Sprintf("%s %s %s ", l.Level.String(), l.Time, emojis[l.EmojiChoice][lvl])
	for k, v := range l.Knots {
		str = str + fmt.Sprintf(" -  %s:%v", k, v)
	}

	if (l.File != "" && lvl <= l.Level) || (l.File != "" && l.Level.String() == "Error") {
		l.writelog(str)
	}
	l.Color.Println(str)
}
