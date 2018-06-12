package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"gopkg.in/yaml.v2"
)

func main() {
	config, err := yaml.ReadFile("conf.yaml")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}
