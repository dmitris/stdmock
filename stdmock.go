package stdmock

import "log"

func Foobar() string {
	log.Printf("Hello I'm verbose Foobar function")
	return "hi"
}
