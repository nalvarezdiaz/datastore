package datastore

type DsRedis struct {
}

type DsRedisOptions struct {
	Host     string
	Port     int
	Pwd      string
	DbNumber int
}

func (redis *DsRedis) Open(opts interface{}) {
	var options DsRedisOptions
	options = opts.(DsRedisOptions)
	logs.Info.Printf("open connection at %s:%d with db number %d", options.Host, options.Port, options.DbNumber)
}

func (redis *DsRedis) Close() {
	logs.Info.Println("close connection")
}
