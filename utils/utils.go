package utils

import (
	"fmt"
	"log"
)

func HandleError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

type Destructor interface {
	Destroy() error
}

func HandleDestroy(dest Destructor) {
	if dest != nil {
		fmt.Printf("Detroy: %#v\n", dest)
		err := dest.Destroy()
		if err != nil {
			log.Panicln(err)
		}
	}
}
