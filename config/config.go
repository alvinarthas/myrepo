package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	CONFIG     Config
	configPath = "./config/config.yml"
)

func init() {

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&CONFIG); err != nil {
		log.Fatal(err)
	}
}
