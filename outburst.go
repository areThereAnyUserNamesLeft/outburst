package outburst

import (
	"fmt"
	"time"
)

type outburst struct {
	Conf confSettings
}

func NewOutBurst() outburst {
	config := readSettings()
	return outburst{Conf: config}
}

func (o outburst) PrintSettings() {
	ob := "Outburst settings"
	if *o.Conf.Vulgar {
		ob = "Outburst's set to fecking log"
	}
	fmt.Println(ob)
	fmt.Printf("Debug      = %v \n", o.Conf.Debug)
	fmt.Printf("Vulgarity  = %v \n", o.Conf.Vulgar)
	fmt.Printf("TimeFormat = %v \n", o.Conf.TimeFormat)
	fmt.Printf("Padding    = %v \n", o.Conf.Padding)
}

// Out takes a map of string interface and returns a Line
func (o outburst) Out(knots map[string]interface{}) Line {
	lev := *o.Conf.DefaultLvl
	l := o.createLine(knots, lev)
	return l
}

func (o outburst) createLine(k Knot, lv Level) Line {
	ln := Line{}
	if !*o.Conf.Emojis {
		ln = Line{
			Knots:       k,
			Time:        time.Now().Format(o.Conf.TimeFormat),
			Level:       lv,
			EmojiChoice: *o.Conf.EmojiChoice,
			Emoji:       "",
			Color:       colors[lv.String()],
			File:        o.Conf.LogFile,
		}
		return ln
	}
	ln = Line{
		Knots:       k,
		Time:        time.Now().Format(o.Conf.TimeFormat),
		Level:       lv,
		EmojiChoice: *o.Conf.EmojiChoice,
		Emoji:       emojis[*o.Conf.EmojiChoice][*o.Conf.DefaultLvl],
		Color:       colors[lv.String()],
		File:        o.Conf.LogFile,
	}
	return ln
}

func (o outburst) Trace(knots map[string]interface{}) {
	l := o.createLine(knots, Trace)
	l.Burst(Trace)
}
func (o outburst) Info(knots map[string]interface{}) {
	l := o.createLine(knots, Info)
	l.Burst(Info)
}

func (o outburst) Debug(knots map[string]interface{}) {
	l := o.createLine(knots, Debug)
	l.Burst(Debug)
}

func (o outburst) Warn(knots map[string]interface{}) {
	l := o.createLine(knots, Warn)
	l.Burst(Warn)
}

func (o outburst) Error(knots map[string]interface{}) {
	l := o.createLine(knots, Error)
	l.Burst(Error)
}

func (o outburst) Fatal(knots map[string]interface{}) {
	l := o.createLine(knots, Fatal)
	l.Burst(Fatal)
}

func (o outburst) Panic(knots map[string]interface{}) {
	l := o.createLine(knots, Panic)
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
