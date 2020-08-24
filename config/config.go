package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

var (
	Cof *Yaml
)

type Yaml struct {
	Excel1 Excel `yaml:"excel1"`
	Excel2 Excel `yaml:"excel2"`
}

type Excel struct {
	File      string `yaml:"file"`
	SheetBook string `yaml:"sheetBook"`
}

func init() {
	cof := &Yaml{}
	f, err := os.Open("./application.yml")
	if err != nil {
		panic(err)
	}
	if err := yaml.NewDecoder(f).Decode(cof); err != nil {
		panic(err)
	}
	Cof = cof
}
