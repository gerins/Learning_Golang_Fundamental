package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
type appConfig map[string]interface{}

func main() {
	configFilePath = os.Getenv(`Configfile`)
	appConfigContent, err := bacaFilePerBaris(configFilePath)

	if err != nil {
		panic(`Read Config File Failed`)
	}

	var config = make(appConfig)
	configReader(config, appConfigContent, "=")
	fmt.Println(config)

	dbUser := config[`dbUser`]
	fmt.Println(dbUser)
}

func configReader(configFile appConfig, config []string, separator string) {
	for _, configVal := range config {
		c := strings.Split(configVal, separator)
		key := c[0]
		value := c[1]
		configFile[key] = value
	}
}

func bacaFilePerBaris(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	var res []string
	for s.Scan() {
		res = append(res, s.Text())
	}

	err = s.Err()

	if err != nil {
		log.Fatal(err)
	}

	return res, err
}

func bacaFile(path string) {...} 
	
func tulisFile(fileName string, data []byte)  {
	err := iotil.WriteFile(filename,data, 0777)
	if err != nil{
		fmt.Println(err)
	}
}
