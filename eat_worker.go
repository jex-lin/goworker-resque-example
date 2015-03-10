package main

import (
	"log"

	"github.com/benmanns/goworker"
)

func init() {
	goworker.Register("eat", eatWorker)
}

func eatWorker(queue string, args ...interface{}) error {
	name := args[0].(string)
	log.Println(name)
	return nil
}
