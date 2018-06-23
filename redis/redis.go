package redis

import (
	"strconv"
	"strings"

	"time"

	rds "github.com/go-redis/redis"
	"github.com/nalvarezdiaz/datastore/logger"
)

type Ds struct {
	Client *rds.Client
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

	redis.Client = rds.NewClient(&rds.Options{
		Addr:     strings.Join([]string{options.Host, strconv.Itoa(options.Port)}, ":"),
		Password: options.Pwd,
		DB:       options.DbNumber,
	})

	logger.Info.Printf("connection opened successfully\n")
	return
}

func (redis *Ds) Close() (err error) {
	err = redis.Client.Close()
	if err != nil {
		return
	}

	logger.Info.Printf("connection closed successfully\n")
	return
}

func (redis *Ds) Create(key string, value string, expiration int) (err error) {
	status := redis.Client.Set(key, value, time.Duration(expiration)*time.Second)
	if status.Err() != nil {
		logger.Error.Println(status.Err())
		return status.Err()
	}

	logger.Info.Printf("item with key [%s] created successfully\n", key)
	return
}

func (redis *Ds) Read(key string) (value string, err error) {
	return redis.Client.Get(key).Result()
}

func (redis *Ds) Delete(key string) (err error) {
	status := redis.Client.Del(key)
	if status.Err() != nil {
		logger.Error.Println(status.Err())
		return status.Err()
	}

	logger.Info.Printf("item with key [%s] has been deleted successfully\n", key)
	return
}
