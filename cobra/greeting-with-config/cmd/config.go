package cmd

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

var cfgFile string

type Person struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	Age  int    `yaml:"age,omitempty" json:"age,omitempty"`
}

func NewPerson() *Person {
	var person *Person

	bs, err := os.ReadFile(cfgFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bs, &person)
	if err != nil {
		panic(err)
	}
	return person
}

func (p *Person) DumpConfig() {
	bs, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("config.json", bs, 0644)
	if err != nil {
		panic(err)
	}
}
