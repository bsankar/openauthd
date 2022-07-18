package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type OadInitStruct struct {
	DBServerHost string `json:"dbserverhost"`
	DBServerType string `json:"dbservertype"`
	//InstallDir   string `json:"installdir"`
	Port int `json:"port"`
}

func OadInit() {
	var initdata OadInitStruct
	file, err := os.Open("./init.json")
	if err != nil {
		log.Panic("Init file missing or can't open", err.Error())
	}

	defer file.Close()

	initbytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Panic("Not able to read init file", err.Error())

	}

	json.Unmarshal(initbytes, &initdata)

	if initdata.DBServerHost == "" || initdata.DBServerType == "" || initdata.Port == 0 {
		log.Panic("Data Missing in init file")
	}

}
