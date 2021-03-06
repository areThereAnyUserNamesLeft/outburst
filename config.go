package outburst

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const file = "outburst.yaml"

type confSettings struct {
	Debug       *bool   `yaml:"debug,omitempty"`
	DefaultLvl  *Level  `yaml:"default_level"`
	Vulgar      *bool   `yaml:"vulgar,omitempty"`
	Emojis      *bool   `yaml:"emojis,omitempty"`
	EmojiChoice *string `yaml:"emoji_choice,omitempty"`
	TimeFormat  string  `yaml:"time_format"`
	Padding     *int    `yaml:"padding,omitempty"`
	LogFile     string  `yaml:"file",omitempty`
}

func readSettings() confSettings {

	f, err := os.Open(file)

	if err != nil {
		fmt.Printf(
			`No %s file found! Please start by saving the %s file 
                         in you project directory - (an empty %s file will 
                         supress this message and use default behaviour :  - %s\n)`,
			file, file, file, err)
	}

	s := confSettings{}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("Cannot read %s file: %s\n", file, err)
	}
	err = yaml.Unmarshal(bytes, &s)
	if err != nil {
		fmt.Printf(
			`Cannot decode %s into settings struct - 
                        please check formatting of %s: %s\n`,
			file, file, err)
	}
	False := false
	if s.Vulgar == nil {
		s.Vulgar = &False
	}
	if s.Debug == nil {
		s.Debug = &False
	}
	if s.DefaultLvl == nil {
		info := Info
		s.DefaultLvl = &info
	}
	if s.Emojis == nil {
		s.Emojis = &False
	}
	if s.TimeFormat == "" {
		s.TimeFormat = "[2006-01-02]-[15:04:05]-"
	}
	four := 4
	if s.Padding == nil {
		s.Padding = &four
	}
	if s.LogFile == "" {
		s.LogFile = "~/logfile.txt"
	}
	return s
}
