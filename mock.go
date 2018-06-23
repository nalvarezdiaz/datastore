package datastore

type DsMock struct {
}

type DsMockOptions struct {
	Path string `toml:"path"`
}

func (mock *DsMock) Open(opts interface{}) {
	var options DsMockOptions
	options = opts.(DsMockOptions)
	logs.Info.Printf("open connection with file %s", options.Path)
}

func (mock *DsMock) Close() {
	logs.Info.Println("close connection")
}
