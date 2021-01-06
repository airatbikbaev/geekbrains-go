package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"regexp"
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

func (c *Config) getConf() *Config {
	yamlFile, err := ioutil.ReadFile("../configs/config.yaml")

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func IsConfigCorrect() bool {
	config := Config{}
	config.getConf()

	return isProdUrlCorrect(config.ProdUrl) && isCredsCorrect(config)

}

func isProdUrlCorrect(url string) bool {
	prodUrl := regexp.MustCompile(`https://prod\..*`)

	return prodUrl.MatchString(url)
}

func isCredsCorrect(c Config) bool {
	return c.User.Login == "admin" && c.User.Password == "admin"

}
