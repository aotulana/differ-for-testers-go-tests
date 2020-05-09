package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

// Config defines the properties of the config file
type Config struct {
	Host string `json:host`
	Port int `json:port`
	WorkingDirectory string `json:"-"`
	Config string
}

// Conf defines the config object
var Conf Config

// Load reads the config file
func Load() {
	environment := os.Getenv("config")

	if environment == "" {
		panic("No config specified")
	}

	dir := ""
	_, filename, _, _ := runtime.Caller(0)
	Conf.WorkingDirectory = path.Join(path.Dir(filename), "..", "..")
	dir = path.Join(path.Dir(filename), "../../config/config."+environment+".json")
	jsonFile, err := os.Open(dir)
	
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Conf)
	jsonFile.Close()

	Conf.Config = environment
}