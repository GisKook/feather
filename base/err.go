package base

import (
	"log"
)

func PrintError(err error) {
	if err != nil {
		log.Println(err)
	}
}
