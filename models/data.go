package models

import (
	"github.com/BurntSushi/toml"
	"io"
	"log"
	"os"
)

type Data struct {
	Genesis string `toml:"genesis"`
}

func NewData(file string) Data {
	var data Data
	f, err := os.Open(file)
	if err != nil {
		return data
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("failed to read file")
	}
	_, err = toml.Decode(string(b), &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
