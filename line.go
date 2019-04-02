package outburst

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

type Level int

const (
	falseVal       = false
	Trace    Level = iota
	Debug
	Info
	Warn
	Error
	Fatal
	Panic
)

func (l Level) String() string {
	return [...]string{"Trace", "Debug", "Info", "Warn", "Error", "Fatal", "Panic"}[l]
}

type Knot map[string]interface{}

type KV map[string]interface{}

type Line struct {
	Knots       Knot
	Time        string
	Level       Level
	EmojiChoice string
	Emoji       string
	Color       *color.Color
	File        string
}

// Out takes a map of string interface and returns a Line
func (o outburst) Out(knots map[string]interface{}) Line {
	lev := *o.Conf.DefaultLvl
	if !*o.Conf.Emojis {
		return Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: lev, EmojiChoice: *o.Conf.EmojiChoice, Emoji: "", Color: colors[lev.String()], File: o.Conf.LogFile}
	}
	return Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: lev, EmojiChoice: *o.Conf.EmojiChoice, Emoji: emojis[*o.Conf.EmojiChoice][lev], Color: colors[lev.String()], File: o.Conf.LogFile}
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

func (o outburst) Trace(knots map[string]interface{}) {
	l := Line{}
	if !*o.Conf.Emojis {
		l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Trace, EmojiChoice: *o.Conf.EmojiChoice, Emoji: "", Color: colors["Trace"], File: o.Conf.LogFile}
		l.Burst(Info)
		return
	}
	l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Trace, EmojiChoice: *o.Conf.EmojiChoice, Emoji: emojis[*o.Conf.EmojiChoice][*o.Conf.DefaultLvl], Color: colors["Trace"], File: o.Conf.LogFile}
	l.Burst(Info)
}
func (o outburst) Info(knots map[string]interface{}) {
	l := Line{}
	if !*o.Conf.Emojis {
		l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Info, EmojiChoice: *o.Conf.EmojiChoice, Emoji: "", Color: colors["Info"], File: o.Conf.LogFile}
		l.Burst(Info)
		return
	}
	l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Info, EmojiChoice: *o.Conf.EmojiChoice, Emoji: emojis[*o.Conf.EmojiChoice][*o.Conf.DefaultLvl], Color: colors["Info"], File: o.Conf.LogFile}
	l.Burst(Info)
}

func (o outburst) Debug(knots map[string]interface{}) {
	l := Line{}
	if !*o.Conf.Emojis {
		l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Debug, EmojiChoice: *o.Conf.EmojiChoice, Emoji: "", Color: colors["Debug"], File: o.Conf.LogFile}
		l.Burst(Debug)
		return
	}
	l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Debug, EmojiChoice: *o.Conf.EmojiChoice, Emoji: emojis[*o.Conf.EmojiChoice][*o.Conf.DefaultLvl], Color: colors["Debug"], File: o.Conf.LogFile}
	l.Burst(Debug)
}

func (o outburst) Warn(knots map[string]interface{}) {
	l := Line{}
	if !*o.Conf.Emojis {
		l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Warn, EmojiChoice: *o.Conf.EmojiChoice, Emoji: "", Color: colors["Warn"], File: o.Conf.LogFile}
		l.Burst(Warn)
		return
	}
	l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Warn, EmojiChoice: *o.Conf.EmojiChoice, Emoji: emojis[*o.Conf.EmojiChoice][*o.Conf.DefaultLvl], Color: colors["Warn"], File: o.Conf.LogFile}
	l.Burst(Warn)
}

func (o outburst) Error(knots map[string]interface{}) {
	l := Line{}
	if !*o.Conf.Emojis {
		l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Error, EmojiChoice: *o.Conf.EmojiChoice, Emoji: "", Color: colors["Error"], File: o.Conf.LogFile}
		l.Burst(Error)
		return
	}
	l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Error, EmojiChoice: *o.Conf.EmojiChoice, Emoji: emojis[*o.Conf.EmojiChoice][*o.Conf.DefaultLvl], Color: colors["Error"], File: o.Conf.LogFile}
	l.Burst(Error)
}

func (o outburst) Fatal(knots map[string]interface{}) {
	l := Line{}
	if !*o.Conf.Emojis {
		l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Fatal, EmojiChoice: *o.Conf.EmojiChoice, Emoji: "", Color: colors["Fatal"], File: o.Conf.LogFile}
		l.Burst(Fatal)
		return
	}
	l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Fatal, EmojiChoice: *o.Conf.EmojiChoice, Emoji: emojis[*o.Conf.EmojiChoice][*o.Conf.DefaultLvl], Color: colors["Fatal"], File: o.Conf.LogFile}
	l.Burst(Fatal)
}

func (o outburst) Panic(knots map[string]interface{}) {
	l := Line{}
	if !*o.Conf.Emojis {
		l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Panic, EmojiChoice: *o.Conf.EmojiChoice, Emoji: "", Color: colors["Panic"], File: o.Conf.LogFile}
		l.Burst(Panic)
		return
	}
	l = Line{Knots: knots, Time: time.Now().Format(o.Conf.TimeFormat), Level: Panic, EmojiChoice: *o.Conf.EmojiChoice, Emoji: emojis[*o.Conf.EmojiChoice][*o.Conf.DefaultLvl], Color: colors["Panic"], File: o.Conf.LogFile}
	l.Burst(Panic)
}

func (o outburst) ErrCheck(err error, knots ...map[string]interface{}) {
	if err != nil {
		o.Error(Knot{"Error": err})
		if len(knots) > 0 {
			for _, kn := range knots {
				o.Error(kn)
			}
		}

	}
}
