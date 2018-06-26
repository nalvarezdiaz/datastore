package main

import (
	"time"

	"log"

	"github.com/nalvarezdiaz/datastore"
	"github.com/nalvarezdiaz/datastore/_examples/config"
	"github.com/nalvarezdiaz/datastore/mock"
	"github.com/nalvarezdiaz/datastore/redis"
)

func main() {
	var dsOpts interface{}

	conf, err := config.New(".conf")
	if err != nil {
		panic(err)
	}

	ds, err := datastore.New(conf.Type)
	if err != nil {
		panic(err)
	}

	switch conf.Type {
	case 0:
		dsOpts = mock.DsOptions(conf.Mock)
	case 1:
		dsOpts = redis.DsOptions(conf.Redis)
	}

	err = ds.Open(dsOpts)
	defer ds.Close()

	if err != nil {
		panic(err)
	}

	ds.Create("test1", "this is a test 1", 5)
	ds.Create("test2", struct{A string}{A: "test"}, 15)
	ds.Create("test3", "this is a test 3", 40)

	time.Sleep(10 * time.Second)

	_, err = ds.Read("test1")
	if err != nil {
		log.Println(err)
	}

	ds.Delete("test1")

	time.Sleep(10 * time.Second)
}
