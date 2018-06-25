package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type MockOptions struct {
	Path string `toml:"path"`
}

type RedisOptions struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Pwd      string `toml:"pwd"`
	DbNumber int    `toml:"db_number"`
}

type DatabaseOptions struct {
	Type  int
	Mock  MockOptions
	Redis RedisOptions
}

func readFile(filename string) (buff []byte, err error) {
	f, _ := filepath.Abs(filename)
	return ioutil.ReadFile(f)
}

func New(filename string) (config *DatabaseOptions, err error) {
	file, err := readFile(filename)
	if err != nil {
		return
	}
	_, err = toml.Decode(string(file), &config)
	return
}
