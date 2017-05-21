package config

import (
	"encoding/json"
	"log"
	"os"
)

type Cfg struct {
	Port uint16  `json:"port"`
	DB   string `json:"db"`
}

func GetConf() *Cfg {
	file, errF := os.Open("./config.json")
	defer file.Close()
	if errF != nil {
		log.Panic(errF)
	}
	decoder := json.NewDecoder(file)
	conf := new(Cfg)
	err := decoder.Decode(&conf)
	if err != nil {
		log.Panic(err)
	}
	return conf
}
