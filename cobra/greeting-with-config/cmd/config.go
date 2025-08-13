package cmd

import (
	"os"

	"gopkg.in/yaml.v2"
)

var cfgFile string

type Person struct {
	Name string `yaml:"name,omitempty"`
	Age  int    `yaml:"age,omitenmpty"`
}

func NewPerson(name string, age int) *Person {
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
