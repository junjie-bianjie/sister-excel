package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

var (
	File1 string
	File2 string
)

type Yaml struct {
	Excel Excel `yaml:"excel"`
}

type Excel struct {
	File1 string `yaml:"file1"`
	File2 string `yaml:"file2"`
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

	File1 = cof.Excel.File1
	File2 = cof.Excel.File2
}
