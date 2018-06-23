package redis

import "github.com/nalvarezdiaz/datastore/logger"

type Ds struct {
}

type DsOptions struct {
	Host     string
	Port     int
	Pwd      string
	DbNumber int
}

func (redis *Ds) Open(opts interface{}) (err error) {
	var options DsOptions
	options = opts.(DsOptions)
	logger.Info.Printf("open connection at %s:%d with db number %d", options.Host, options.Port, options.DbNumber)
	return
}

func (redis *Ds) Close() (err error) {
	logger.Info.Println("close connection")
	return
}