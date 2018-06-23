package datastore

import (
	"errors"
)

type Ds interface {
	Open(interface{})
	Close()
}

const (
	mock  = 0
	redis = 1
)

func New(dsType int) (Ds, error) {
	switch dsType {
	case mock:
		initLog("ds-mock")
		return new(DsMock), nil
	case redis:
		initLog("ds-redis")
		return new(DsRedis), nil
	default:
		initLog()
		logs.Error.Println("invalid Ds type")
		return nil, errors.New("invalid Ds type")
	}
}
