package main

import (
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Image  ImageConfig  `yaml:"image"`
}

type ServerConfig struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type ImageConfig struct {
	Repository string `yaml:"repository"`
	Tag        string `yaml:"tag"`
}

func loadConfigFile(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
	}

	return &config, nil
}

func generateChartvalueFile(config *Config) error {
	tmpl, err := template.ParseFiles("./template/values.yaml")
	if err != nil {
		return err
	}

	file, err := os.Create("./chart/values.yaml")
	if err != nil {
		return err
	}

	defer file.Close()

	return tmpl.Execute(file, config)

}

func main() {
	config, err := loadConfigFile("./platform/config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	err = generateChartvalueFile(config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Helm values.yaml file is generated")
}
