package datastore

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Logger struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

var logs *Logger

func initLog(prefixes ...interface{}) {
	var p string

	p = ""
	if len(prefixes) == 1 {
		switch prefixes[0].(type) {
		case string:
			p = "  \033[1;32m[" + prefixes[0].(string) + "]\033[0;0m "
		case int:
			p = "  \033[1;32m[" + strconv.Itoa(prefixes[0].(int)) + "]\033[0;0m "
		}
	} else if len(prefixes) > 1 {
		var pp []string
		for _, prefix := range prefixes {
			switch prefix.(type) {
			case string:
				pp = append(pp, prefix.(string))
			case int:
				pp = append(pp, strconv.Itoa(prefix.(int)))
			}
		}
		p = "  \033[1;32m[" + strings.Join(pp, ",") + "]\033[0;0m "
	}

	logs = &Logger{
		Info:    log.New(os.Stdout, p+"INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Warning: log.New(os.Stdout, p+"WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		Error:   log.New(os.Stderr, p+"ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}