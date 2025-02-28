package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type App struct {
	Port         string `json:"port"`
	Prefix       string `json:"prefix"`
	Name         string `json:"name"`
	Version      string `json:"version"`
	ExpiredVideo int    `json:"expiredVideo"`
}
type Security struct {
	EnabledCors   bool   `json:"enabledCors"`
	WhitelistHost string `json:"whitelistHost"`
}

type Pexels struct {
	Host      string `json:"host"`
	SecretKey string `json:"secretKey"`
	Size      int    `json:"size"`
}

type Integration struct {
	Pexels Pexels `json:"pexels"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

type Service struct {
	Redis Redis `json:"redis"`
}

type config struct {
	App         App         `json:"app"`
	Service     Service     `json:"service"`
	Integration Integration `json:"integration"`
	Security    Security    `json:"Security"`
}

var Config config

func LoadConfig() {
	configLoc := "./config.json"
	file, err := ioutil.ReadFile(configLoc)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(file, &Config)
}
