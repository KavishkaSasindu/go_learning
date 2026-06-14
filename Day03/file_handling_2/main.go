package main

import (
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Protocol string `yaml:"protocol"`
	Timeout  int    `yaml:"timeout"`
	Enabled  bool   `yaml:"enabled"`
}

func loadConfig(filepath string) (*Config, error) {
	// first read file
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("No error Found")
		return nil, err
	}

	// then create a variable type of Config struct
	var cfg Config

	// then unmarshal with address of variable name (&cfg)
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// then return that address of cfg (&cfg) and nil
	return &cfg, nil
}

// use template and write new file
func writeFileNewFile(cfg ServerConfig) error {
	tmpl, err := template.ParseFiles("./templates/template.txt")
	if err != nil {
		fmt.Println("There is an error when parsing template")
	}

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer outputFile.Close()

	return tmpl.Execute(outputFile, cfg)
}

func main() {
	fmt.Println("Hello World")

	// test for read config file
	cfg, err := loadConfig("./platform/config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server host:", cfg.Server.Host)

	err = writeFileNewFile(cfg.Server)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File is generated!")
}
