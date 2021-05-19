package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Test struct {
	User  []string  `yaml:"user"`
	MQTT  MQ  `yaml:"mqtt"`
	Http HTTP   `yaml:"http"`
}


type Config struct {
	Test Test  `yaml:"test"`
}

type MQ struct {
	Host string  `yaml:"host"`
	Username  string  `yaml:"username"`
	Password  string  `yaml:"password"`
}


type HTTP struct {
	Port  string  `yaml:"port"`
	Host  string  `yaml:"host"`
}





//read yaml config
//注：path为yaml或yml文件的路径
func ReadYamlConfig(path string)  (*Config,error){
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil,err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	return  conf,nil
}

//test yaml
func main() {
	conf,err := ReadYamlConfig("test.yaml")
	if err != nil {
		log.Fatal(err)
	}

	byts,err := json.Marshal(conf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byts))
}