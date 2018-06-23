package mock

import (
	"encoding/json"

	"errors"

	"time"

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
	go mock.startGarbageCollector()

	logger.Info.Printf("connection opened successfully\n")
	return
}

func (mock *Ds) Close() (err error) {
	err = mock.write()
	if err != nil {
		return
	}

	logger.Info.Printf("connection closed successfully\n")
	return
}

func (mock *Ds) Create(key string, value string, expiration int) (err error) {

	if mock.exist(key) {
		err = errors.New("unable to create item because the key already exists")
		logger.Error.Println(err.Error())
		return
	}

	item := item{
		Key: key,
		Value: record{
			Data:       value,
			Expiration: int(time.Now().Unix()) + expiration,
		},
	}

	mock.items = append(mock.items, item)
	err = mock.write()
	if err != nil {
		return
	}

	logger.Info.Printf("item with key [%s] created successfully\n", item.Key)
	return
}

func (mock *Ds) Read(key string) (value string, err error) {
	for _, item := range mock.items {
		if item.Key == key && (item.Value.Expiration == 0 || item.Value.Expiration >= int(time.Now().Unix())) {
			return item.Value.Data, nil
		}
	}
	return "", errors.New("item has expired or does not exist")
}

func (mock *Ds) Delete(key string) (err error) {
	for i, item := range mock.items {
		if item.Key == key {
			mock.items = append(mock.items[:i], mock.items[i+1:]...)
			logger.Info.Printf("item with key [%s] has been deleted successfully\n", item.Key)
			return nil
		}
	}
	return errors.New("item does not exist")
}

func (mock *Ds) exist(key string) bool {
	for _, item := range mock.items {
		if item.Key == key {
			return true
		}
	}
	return false
}

func (mock *Ds) write() (err error) {
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

	return
}

func (mock *Ds) startGarbageCollector() {
	for {
		<-time.After(1 * time.Second)
		go mock.garbageCollector()
	}
}

func (mock *Ds) garbageCollector() {
	var tmp Items
	for i, item := range mock.items {
		if item.Value.Expiration == 0 || item.Value.Expiration >= int(time.Now().Unix()) {
			tmp = append(tmp, mock.items[i])
		} else {
			logger.Info.Printf("item with key [%s] has expired\n", item.Key)
		}
	}
	mock.items = tmp
	mock.write()
}
