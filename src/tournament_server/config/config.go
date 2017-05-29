package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type DBstruct struct {
	User     string
	Password string
	Database string
	Addr     string
}

type Cfg struct {
	Port uint16
	DB   DBstruct
}

func GetConf() *Cfg {
	portStr := os.Getenv("PORT")
	if portStr != "" {
		u64, err := strconv.ParseUint(portStr, 10, 32)
		if err == nil {
			return &Cfg{
				Port: uint16(u64),
				DB: DBstruct{
					User:     os.Getenv("USER"),
					Password: os.Getenv("PASSWORD"),
					Database: os.Getenv("DATABASE"),
					Addr:     os.Getenv("ADDR"),
				},
			}
		}

	}

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
