package main

import (
	"io/ioutil"
	"log"

	"github.com/chai2010/errors"
)

func loadConfig() error {
	_, err := ioutil.ReadFile("/path/to/file")
	if err != nil {
		return errors.Wrap(err, "read failed")
	}

	return nil
}

func setup() error {
	err := loadConfig()
	if err != nil {
		return errors.Wrap(err, "invalid config")
	}

	return nil
}

func main() {
	if err := setup(); err != nil {
		log.Fatal(err)
	}

	// ...
}
