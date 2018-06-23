package mock

import (
	"encoding/json"

	"github.com/nalvarezdiaz/datastore/logger"
)

type Ds struct {
	path  string
	items Items
}

type DsOptions struct {
	Path string `toml:"path"`
}

type record struct {
	Data       string `json:"data"`
	Expiration int    `json:"expiration"`
}

type item struct {
	Key   string `json:"key"`
	Value record `json:"value"`
}

type Items []item

func (mock *Ds) Open(opts interface{}) (err error) {
	var options DsOptions

	options = opts.(DsOptions)

	logger.Info.Printf("opening connection with file %s\n", options.Path)
	bytes, err := readFile(options.Path)
	if err != nil {
		logger.Error.Printf("unable to open file %s\n", options.Path)
		return
	}

	err = json.Unmarshal(bytes, &mock.items)
	if err != nil {
		logger.Error.Printf("unable to convert data into expected structure\n")
		return
	}

	mock.path = options.Path

	logger.Info.Printf("connection opened successfully\n")
	return
}

func (mock *Ds) Close() (err error) {
	var bytes []byte

	bytes, err = json.MarshalIndent(mock.items, "", "  ")
	if err != nil {
		logger.Error.Printf("unable to save data before closing\n")
		return
	}

	err = writeFile(mock.path, bytes)
	if err != nil {
		logger.Error.Printf("unable to write storage at the specified file\n")
		return
	}

	logger.Info.Printf("connection closed successfully\n")
	return
}
