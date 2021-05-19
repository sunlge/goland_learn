/*
文件内容
SiteName: seeta
SiteAddr: BeiJing
Https: true
Nginx:
    Port: 443
    LogPath:  "/var/log/nginx/nginx.log"
    Path: "/opt/nginx/"
*/

package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//Nginx nginx  配置
type Nginx struct {
	Port int `yaml:"Port"`
	LogPath string `yaml:"LogPath"`
	Path string `yaml:"Path"`
}

//Config   系统配置配置
type Config01 struct{
	Name string `yaml:"SiteName"`
	Addr string `yaml:"SiteAddr"`
	HTTPS bool `yaml:"Https"`
	SiteNginx  Nginx `yaml:"Nginx"`
}

func main() {

	var setting Config01
	config, err := ioutil.ReadFile("./jiexi_yaml.yaml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config,&setting)

	fmt.Println("setting.Name: ", setting.Name)
	fmt.Println("setting.Addr:", setting.Addr)
	fmt.Println("setting.HTTPS:", setting.HTTPS)
	fmt.Println("setting.SiteNginx.Port:", setting.SiteNginx.Port)
	fmt.Println("setting.SiteNginx.LogPath:", setting.SiteNginx.LogPath)
	fmt.Println("setting.SiteNginx.Path:", setting.SiteNginx.Path)
}