package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Port int `yaml:"port"`
	User struct {
		Name     string `yaml:"name"`
		Age      int    `yaml:"age"`
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
	}
	ProdUrl string `yaml:"prod_url"`
}

func CheckConfig() {
	file, fileErr := ioutil.ReadFile("configs.yaml")
	if fileErr != nil {
		log.Fatalf("Не могу открыть файл: %v", fileErr)
	}

	config := Config{}

	err := yaml.Unmarshal(file, &config)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(config)
}
