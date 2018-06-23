package config

import (
	"io/ioutil"
	"os"
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

func readFileIntoByte(filename string) []byte {
	var fmwConfigJsonBuf []byte

	f, _ := filepath.Abs(filename)
	fin, err := os.Open(f)

	if err != nil {
		panic(err)
	} else {
		fmwConfigJsonBuf, err = ioutil.ReadAll(fin)
		if err != nil {
			panic(err)
		}
	}
	fin.Close()
	return fmwConfigJsonBuf
}

func New(filename string) (config *DatabaseOptions, err error) {
	file := readFileIntoByte(filename)
	_, err = toml.Decode(string(file), &config)
	return
}
