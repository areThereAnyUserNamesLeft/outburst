package outburst

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const file = "outburst.yaml"

type confSettings struct {
	Debug      *bool  `yaml:"debug,omitempty"`
	Vulgar     *bool  `yaml:"vulgar,omitempty"`
	Emojis     *bool  `yaml:"emojis,omitempty"`
	TimeFormat string `yaml:"time_format"`
	Padding    *int   `yaml:"paddingi,omitempty"`
}

func readSettings() confSettings {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("No %s file found! Please start by saving the %s file in you project directory - (an empty %s file will supress this message and use default behaviour.\n)", file, file, file)
	}

	s := confSettings{}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("Cannot read %s file.\n", file)
	}
	err = yaml.Unmarshal(bytes, &s)
	if err != nil {
		fmt.Printf("Cannot decode %s into settings struct - please check formatting of %s.\n", file, file)
	}
	False := false
	if s.Vulgar == nil {
		s.Vulgar = &False
	}
	if s.Debug == nil {
		s.Debug = &False
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
	return s
}

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
