package main

import (
	"log"

	"github.com/benmanns/goworker"
)

func init() {
	goworker.Register("drink", drinkWorker)
}

func drinkWorker(queue string, args ...interface{}) error {
	name := args[0].(string)
	log.Println(name)
	return nil
}
