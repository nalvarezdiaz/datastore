package logger

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var Info *log.Logger
var Warning *log.Logger
var Error *log.Logger

func New(prefixes ...interface{}) {
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

	Info = log.New(os.Stdout, p+"INFO:  ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, p+"\033[1;33mWARN:  ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, p+"\033[1;31mERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}
