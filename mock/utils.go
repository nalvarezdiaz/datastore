package mock

import (
	"io/ioutil"
	"path/filepath"
)

func readFile(filename string) (buff []byte, err error) {
	f, _ := filepath.Abs(filename)
	return ioutil.ReadFile(f)
}

func writeFile(filename string, buff []byte) (err error) {
	f, _ := filepath.Abs(filename)
	return ioutil.WriteFile(f, buff, 0644)
}
