package datastore

import (
	"errors"

	"github.com/nalvarezdiaz/datastore/logger"
	"github.com/nalvarezdiaz/datastore/mock"
	"github.com/nalvarezdiaz/datastore/redis"
)

type Ds interface {
	Open(interface{}) error
	Close() error
	Create(string, interface{}, int) error
	Read(string) (interface{}, error)
	Delete(string) error
}

const (
	Mock  = 0
	Redis = 1
)

func New(dsType int) (Ds, error) {
	switch dsType {
	case Mock:
		logger.New("ds-mock")
		return new(mock.Ds), nil
	case Redis:
		logger.New("ds-redis")
		return new(redis.Ds), nil
	default:
		logger.New()
		logger.Error.Println("invalid Ds type")
		return nil, errors.New("invalid Ds type")
	}
}
